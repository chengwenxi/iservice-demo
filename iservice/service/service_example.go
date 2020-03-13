package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"iservice/iservice/market"

	"github.com/irisnet/irishub-sdk-go/rpc"
)

var serviceMap = make(map[string]rpc.ServiceRespondHandler)

const (
	PriceServiceName = "price_service"
)

func init() {
	serviceMap[PriceServiceName] = priceService
}

func GetServiceCallBack(serviceName string) rpc.ServiceRespondHandler {
	return serviceMap[serviceName]
}

func priceService(input string) (output string, errMsg string) {
	var request Input
	err := json.Unmarshal([]byte(input), &request)
	if err != nil {
		return "", errors.New(fmt.Sprintf("can not parse input json string : %s", err.Error())).Error()
	}

	// get price from public market
	mk := market.GetMarket()
	price, errMsg := mk.GetPrice(request.Base, request.Quote)

	if len(errMsg) > 0 {
		return "", errMsg
	}

	outputBz, _ := json.Marshal(Output{Price: price})
	return string(outputBz), errMsg
}

type Input struct {
	Base  string `json:"base"`
	Quote string `json:"quote"`
}

type Output struct {
	Price float64 `json:"price"`
}
