package router

import "C"
import (
	"fmt"
)
import "net/http"
import "server/models"
import "encoding/json"

func CandyHandler(w http.ResponseWriter, r *http.Request) {
	var req models.CandyRequest
	var resp models.CandyResponseOk
	defer r.Body.Close()

	if r.Method != "POST" {
		return
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		var res models.CandyResponseFailed
		res.Error = models.ErrInputData.Error()
		w.WriteHeader(http.StatusBadRequest)
		bytes, err := json.Marshal(res)
		if err != nil {
			fmt.Printf("failed to marshal json: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(bytes)
		return
	}

	err = ValidateInputData(req)
	if err != nil {
		var res models.CandyResponseFailed
		res.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		bytes, err := json.Marshal(res)
		if err != nil {
			fmt.Printf("failed to marshal json: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(bytes)
		return
	}

	change, err := CalculateChange(req)
	if err != nil {
		var res models.CandyResponseFailed
		res.Error = err.Error()
		w.WriteHeader(http.StatusPaymentRequired)
		bytes, err := json.Marshal(res)
		if err != nil {
			fmt.Printf("failed to marshal json: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(bytes)
		return
	}

	resp.Thanks = C.GoString(C.ask_cow(C.CString("Thank you")))

	resp.Change = change

	bytes, err := json.Marshal(resp)
	if err != nil {
		fmt.Printf("failed to marshal json: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	w.WriteHeader(http.StatusCreated)
}

func ValidateInputData(req models.CandyRequest) error {
	if req.CandyCount < 0 || req.Money < 0 {
		return models.ErrInputData
	}

	if _, exists := models.CandyPrice[req.CandyType]; !exists {
		return models.ErrInputData
	}
	return nil
}

func CalculateChange(req models.CandyRequest) (int, error) {

	price := models.CandyPrice[req.CandyType]

	if req.Money < req.CandyCount*price {
		return 0, models.ErrNoMoney
	}

	return req.Money - (req.CandyCount * price), nil
}