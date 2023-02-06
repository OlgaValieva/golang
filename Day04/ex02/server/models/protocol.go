package models

import "errors"

var ErrInputData = errors.New("400: some error in input data")
var ErrNoMoney = errors.New("402: not enough money")

var CandyPrice = map[string]int {
	"CE": 10,
	"AA": 15,
	"NT": 17,
	"DE": 21,
	"YR": 23,
}

type CandyRequest struct {
	Money int `json:"money"`
	CandyType string `json:"candyType"`
	CandyCount int `json:"candyCount"`
}

type CandyResponseOk struct {
	Change int `json:"change"`
	Thanks string `json:"thanks"`
}

type CandyResponseFailed struct {
	Error string `json:"error"`
}
