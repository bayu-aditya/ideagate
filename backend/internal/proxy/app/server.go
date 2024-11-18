package app

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var (
	addr       = ":8080"
	wsUpgrader = websocket.Upgrader{}
)

func Server() {
	http.HandleFunc("/proxy", websocketProxyHandler)
	http.HandleFunc("/", healthCheckHandler)

	logrus.Infof("Proxy server running on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		logrus.Fatal(err)
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func websocketProxyHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Error(err)
		return
	}
	defer conn.Close()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			now := time.Now().Second()
			conn.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(now)))
		}
	}()

	for {
		msgType, message, err := conn.ReadMessage()
		if err != nil {
			logrus.Error("read message", err)
			break
		}

		logrus.Info("get message", msgType, string(message))

		if err = conn.WriteMessage(msgType, message); err != nil {
			logrus.Error("write message", err)
			break
		}
	}
}
