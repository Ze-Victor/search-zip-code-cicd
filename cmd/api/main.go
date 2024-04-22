package main

import (
	"github.com/Ze-Victor/search-zip-code/config"
	"github.com/gin-gonic/gin"
)

func main() {

	logger := config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router := gin.Default()

	initializeRoutes(router)

	err = router.Run(":8001")
	if err != nil {
		logger.Errorf("Failed to start server: %v", err)
	}

}
