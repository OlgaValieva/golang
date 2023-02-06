package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type candy struct {
	Name  string
	Count string
	Money string
}

type CandyResponse struct {
	Change int    `json:"change"`
	Thanks string `json:"thanks"`
}

func readFlags(candy *candy) error {
	flag.StringVar(&candy.Name, "k", "", "candy type")
	flag.StringVar(&candy.Count, "c", "", "candy type")
	flag.StringVar(&candy.Money, "m", "", "candy type")
	flag.Parse()
	if candy.Name == "" || candy.Count == "" || candy.Money == "" {
		return errors.New("empty flag")
	}
	return nil
}

const (
	RootCertificatePath string = "../certs/minica.pem"
	ClientCertPath      string = "../certs/client/cert.pem"
	ClientKeyPath       string = "../certs/client/key.pem"
)

func main() {
	candy := candy{}
	if err := readFlags(&candy); err != nil {
		log.Println("error", err)
		os.Exit(1)
	}

	rootCAPool := x509.NewCertPool()
	rootCA, err := ioutil.ReadFile(RootCertificatePath)
	if err != nil {
		log.Fatalf("reading cert failed : %v", err)
	}
	rootCAPool.AppendCertsFromPEM(rootCA)

	c := http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout: 10 * time.Second,
			TLSClientConfig: &tls.Config{
				RootCAs: rootCAPool,
				// Load clients key-pair. This will be sent to server
				GetClientCertificate: func(info *tls.CertificateRequestInfo) (certificate *tls.Certificate, e error) {
					c, err := tls.LoadX509KeyPair(ClientCertPath, ClientKeyPath)
					if err != nil {
						fmt.Printf("Error loading client key pair: %v\n", err)
						return nil, err
					}
					return &c, nil
				},
			},
		},
	}

	jsonStr := fmt.Sprintf("{\"money\": %s, \"candyType\": \"%s\", \"candyCount\": %s}", candy.Money, candy.Name, candy.Count)
	req, err := http.NewRequest("POST", "https://localhost:3333/buy_candy", bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Println("error of new request")
	}
	res, err := c.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	if res.StatusCode > 299 {
		log.Println(res.Status)
		return
	}
	defer res.Body.Close()

	result, err := ioutil.ReadAll(res.Body)
	var resCandy CandyResponse
	if err := json.Unmarshal(result, &resCandy); err != nil {
		log.Println("unmarshalling body")
	}

	fmt.Println("Thank you! Your change is:", resCandy.Change)

}