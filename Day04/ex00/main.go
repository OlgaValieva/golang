package main
 
import (
	"fmt"
	"log"
	"net/http"
	"server/router"
)

func main() {
	log.Printf("Beginning")
	http.HandleFunc("/buy_candy", router.CandyHandler)
	http.ListenAndServe("localhost:3333", nil)
	fmt.Println("Hello")
}