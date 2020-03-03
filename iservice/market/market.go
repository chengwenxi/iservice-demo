package market

import (
	"iservice/iservice/market/bian"
	"iservice/iservice/market/huobi"
)

type Market interface {
	GetPrice(base string, quote string) (price float64, error string)
}

type EmptyMarket struct {
}

func (market EmptyMarket) GetPrice(base string, quote string) (price float64, error string) {
	return 0, ""
}

var MarketType string

func GetMarket() Market {
	switch MarketType {
	case "huobi":
		return huobi.HuobiMarket{}
	case "bian":
		return bian.BianMarket{}
	default:
		return EmptyMarket{}
	}
}
