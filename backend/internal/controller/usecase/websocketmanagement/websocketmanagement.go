package websocketmanagement

import (
	"encoding/json"
	"log"
	"time"

	entitywebsocket "github.com/bayu-aditya/ideagate/backend/internal/shared/domain/entity/websocket"
	"github.com/gorilla/websocket"
)

type IWebsocketManagementUsecase interface {
	WorkerSubscriber()
	WorkerPublisher()
	Close() error
}

func NewWebsocketManagement(wsConn *websocket.Conn) IWebsocketManagementUsecase {
	return &websocketManagement{
		conn: wsConn,
	}
}

type websocketManagement struct {
	conn *websocket.Conn
}

func (w *websocketManagement) WorkerSubscriber() {
	for {
		_, message, err := w.conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			continue
		}

		var eventRequest entitywebsocket.Event
		if err = json.Unmarshal(message, &eventRequest); err != nil {
			log.Println("unmarshal event request:", err)
			continue
		}

		log.Printf("recv: %s", eventRequest)

		// TODO switch usecase
		//usecaseProcess := w.usecaseSwitcher(eventType)

		// TODO handle by usecase
		//result, err := usecaseProcess.Process(data)
		//if err != nil {
		//	// TODO handle error
		//	log.Println("error:", err)
		//	// TODO construct error into event response and send error to server
		//}

		// construct event response
		eventResponse := entitywebsocket.Event{
			Id:        eventRequest.Id,
			ProjectId: eventRequest.ProjectId,
			Data:      "hupla data response",
		}

		eventResponseJson, err := json.Marshal(eventResponse)
		if err != nil {
			log.Println("Marshal event response:", err)
			continue
		}

		if err = w.conn.WriteMessage(websocket.TextMessage, eventResponseJson); err != nil {
			log.Println("write:", err)
			continue
		}
	}
}

func (w *websocketManagement) WorkerPublisher() {
	for {
		time.Sleep(2 * time.Second)
		if err := w.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
			log.Println("write:", err)
		}
	}
}

func (w *websocketManagement) Close() error {
	// Cleanly close the connection by sending a close message and then waiting for the server to close the connection.
	err := w.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return err
	}

	if err = w.conn.Close(); err != nil {
		return err
	}

	return nil
}
