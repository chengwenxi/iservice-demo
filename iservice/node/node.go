package node

import (
	"fmt"
	"iservice/iservice/service"
	"log"

	"github.com/irisnet/irishub-sdk-go/client"
	sdk "github.com/irisnet/irishub-sdk-go/types"
)

func Start(config sdk.SDKConfig, baseTx sdk.BaseTx) {
	irisClient := client.New(config)
	serviceName := service.PriceServiceName
	baseTx.Memo = fmt.Sprintf("service invocation response for %s", serviceName)
	err := irisClient.Service.RegisterSingleInvocationListener(
		serviceName, service.GetServiceCallBack(serviceName), baseTx)
	if err != nil {
		log.Printf("failed to register invocation listener, err: %s", err.Error())
		return
	}

	select {}
}
