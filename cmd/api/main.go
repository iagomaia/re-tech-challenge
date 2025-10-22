package main

import (
	"log"
	"net/http"
	"os"

	"github.com/iagomaia/re-tech-challenge/internal/infra/server"
	"github.com/iagomaia/re-tech-challenge/internal/infra/utils"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	handler := server.GetServerRoutes()

	logger := utils.GetLogger()
	port := utils.GetStringValueOrDefault(os.Getenv("SV_PORT"), "3000")
	logger.Info().Msgf("Starting server on port %v", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, handler))
}
