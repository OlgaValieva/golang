package main
 
import (
	"log"
	"net/http"
	"server/router"
)

const ServerPort = "localhost:3333"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/buy_candy", router.CandyHandler)
	log.Fatal(http.ListenAndServeTLS(ServerPort, "../certs/server/cert.pem", "../certs/server/key.pem", mux))
}