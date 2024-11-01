package app

import (
	"context"

	"github.com/bayu-aditya/ideagate/backend/internal/workerapi/usecase/handler"
	"github.com/bayu-aditya/ideagate/backend/pkg/utils/log"
	"github.com/gin-gonic/gin"
)

func Rest() {
	ctx := context.Background()
	router := gin.Default()

	usecaseHandler := handler.New()
	if err := usecaseHandler.GenerateEndpoint(ctx, router); err != nil {
		log.Fatal("generate endpoint", err)
	}

	if err := router.Run(); err != nil {
		log.Fatal("run router", err)
	}
}
