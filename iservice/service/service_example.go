package service

import "github.com/irisnet/irishub-sdk-go/types"

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
	return "", ""
}
