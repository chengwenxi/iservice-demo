package node

import (
	"fmt"
	"iservice/iservice/service"
	"log"
	"os"

	sdk "github.com/irisnet/irishub-sdk-go"
	"github.com/irisnet/irishub-sdk-go/types"
)

func Start(config types.ClientConfig, baseTx types.BaseTx) {
	irisClient := sdk.NewClient(config)
	irisClient.SetOutput(os.Stdout)
	serviceName := service.PriceServiceName
	baseTx.Memo = fmt.Sprintf("service invocation response for %s", serviceName)
	_, err := irisClient.Service().SubscribeSingleServiceRequest(
		serviceName, service.GetServiceCallBack(serviceName), baseTx)
	if err != nil {
		log.Printf("failed to register invocation listener, err: %s", err.Error())
		return
	}

	select {}
}
