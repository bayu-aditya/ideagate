package app

import (
	"encoding/json"
	"net/http"

	"github.com/bayu-aditya/ideagate/backend/internal/proxy/usecase"
	entitypubsub "github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/pubsub"
	entitywebsocket "github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/websocket"
	"github.com/bayu-aditya/ideagate/backend/internal/shared/ports/pubsub"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var (
	addr       = ":8080"
	wsUpgrader = websocket.Upgrader{}
)

func RunRestHttp(pubSub pubsub.IPubSubAdapter, usecaseWebsocketClientManagement usecase.IWebsocketClientManagement) {
	http.HandleFunc("/", healthCheckHandler)
	http.HandleFunc("/event/send", sendEventHandler(pubSub))
	http.HandleFunc("/event/ws", websocketProxyHandler(usecaseWebsocketClientManagement))

	logrus.Infof("Proxy server running on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		logrus.Fatal(err)
	}
}

func healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func sendEventHandler(pubSub pubsub.IPubSubAdapter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		ctx := r.Context()

		requestBody := struct {
			ProjectId string
			Data      any
		}{}

		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if requestBody.ProjectId == "" {
			http.Error(w, "Project ID is required", http.StatusBadRequest)
			return
		}

		// construct event
		event := entitywebsocket.Event{
			Id:        uuid.NewString(),
			ProjectId: requestBody.ProjectId,
			Data:      requestBody.Data,
		}

		eventBytes, err := json.Marshal(event)
		if err != nil {
			http.Error(w, "Failed to marshal event", http.StatusInternalServerError)
			return
		}

		if err = pubSub.Publish(ctx, entitypubsub.TopicEventRequest, eventBytes); err != nil {
			http.Error(w, "Failed to publish event", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func websocketProxyHandler(usecaseWebsocketClientManagement usecase.IWebsocketClientManagement) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		projectID := r.URL.Query().Get("project_id")
		if projectID == "" {
			http.Error(w, "Project ID is required", http.StatusBadRequest)
			return
		}

		connectionID := r.Header.Get("connection_id")
		if connectionID == "" {
			http.Error(w, "Connection ID is required", http.StatusBadRequest)
			return
		}

		conn, err := wsUpgrader.Upgrade(w, r, nil)
		if err != nil {
			logrus.Error(err)
			return
		}
		defer conn.Close()

		websocketClient := usecaseWebsocketClientManagement.NewClient(ctx, conn, projectID, connectionID)
		go websocketClient.WorkerReceiveEvent()
		<-websocketClient.WaitConnectionClose()
	}
}
