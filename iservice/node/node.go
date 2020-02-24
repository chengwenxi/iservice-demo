package node

import (
	"log"

	"github.com/irisnet/irishub-sdk-go/client"
	sdk "github.com/irisnet/irishub-sdk-go/types"
)

const (
	TagProvider         = "provider"
	TagConsumer         = "consumer"
	TagRequestContextID = "request-context-id"
	TagRequestID        = "request-id"
	TagServiceFee       = "service-fee"
	TagRequestHeight    = "request-height"
	TagExpirationHeight = "expiration-height"
	TagSlashedCoins     = "slashed-coins"
)

func Start(config sdk.SDKConfig, baseTx sdk.BaseTx) {
	irisClient := client.New(config)
	provider := ""

	var builder = sdk.NewEventQueryBuilder()

	builder.AddCondition(TagProvider, sdk.EventValue(provider))

	_, _ = irisClient.SubscribeTx(builder, func(data sdk.EventDataTx) {
		for _, tag := range data.Result.Tags {
			var provider sdk.AccAddress
			var requestID string
			switch tag.Key {
			case TagProvider:
				var err error
				provider, err = sdk.AccAddressFromBech32(tag.Value)
				if err != nil {
					log.Printf("unknow address: %s", tag.Value)
					break
				}
			case TagRequestID:
				requestID = tag.Value
			}

			if !provider.Empty() && len(requestID) > 0 {
				RequestCallBack(irisClient, provider, requestID)
			}
		}
	})
}

func RequestCallBack(client client.Client, provider sdk.AccAddress, requestID string) {
	// TODO
	//response := MsgRespondService{
	//	RequestID: requestID,
	//	Provider:  provider,
	//	Output:    "test",
	//}

	result, err := client.BuildAndSend([]sdk.Msg{}, sdk.BaseTx{})
	if err != nil {
		log.Printf("build and send tx failed: %s", err.Error())
	}
	if result != nil {
		log.Printf("build and send tx successfully, height: %d, hash: %s",
			result.GetHeight(), result.GetHeight())
	}
}

type MsgRespondService struct {
	RequestID string         `json:"request_id"`
	Provider  sdk.AccAddress `json:"provider"`
	Output    string         `json:"output"`
	Error     string         `json:"error"`
}
