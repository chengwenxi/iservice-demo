package service

import (
	"encoding/json"
	"math/rand"

	"github.com/irisnet/irishub-sdk-go/types"
)

var serviceMap = make(map[string]types.ServiceRespondHandler)

const (
	PriceServiceName = "price_service"
)

func init() {
	serviceMap[PriceServiceName] = priceService
}

func GetServiceCallBack(serviceName string) types.ServiceRespondHandler {
	return serviceMap[serviceName]
}

func priceService(input string) (output string, errMsg string) {
	outputBz, _ := json.Marshal(Output{Price: 100 * rand.Float32()})
	return string(outputBz), ""
}

type Output struct {
	Price float32 `json:"price"`
}
