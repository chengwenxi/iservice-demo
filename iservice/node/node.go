package node

import (
	"fmt"
	"iservice/iservice/service"
	"log"
	"os"

	"github.com/irisnet/irishub-sdk-go/client"
	sdk "github.com/irisnet/irishub-sdk-go/types"
)

func Start(config sdk.SDKConfig, baseTx sdk.BaseTx) {
	irisClient := client.NewSDKClient(config)
	irisClient.SetOutput(os.Stdout)
	serviceName := service.PriceServiceName
	baseTx.Memo = fmt.Sprintf("service invocation response for %s", serviceName)
	_, err := irisClient.Service().RegisterSingleServiceListener(
		serviceName, service.GetServiceCallBack(serviceName), baseTx)
	if err != nil {
		log.Printf("failed to register invocation listener, err: %s", err.Error())
		return
	}

	select {}
}
