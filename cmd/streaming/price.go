package streaming

import (
	"context"
	"encoding/json"
	"fmt"
	goservice "github.com/200Lab-Education/go-sdk"
	"io/ioutil"
	"math/rand"
	"net/http"
	"price-streaming/common"
	"price-streaming/plugin/nats"
	"time"
)

type Quote struct {
	OpenPrice  float64
	ClosePrice float64
	LastPrice  float64
	QuoteVol   int
}

type Message struct {
	Symbol string
	Data   Quote
}

type Data struct {
	Symbol string
}

type Response struct {
	Data []Data
}

type PriceResponse struct {
	Symbol        string  `json:"symbol"`
	TradePrice    float64 `json:"trade_price"`
	Volume        int     `json:"volume"`
	ChangePoint   int     `json:"change_point"`
	ChangePercent int     `json:"change_percent"`
	Bid           float64 `json:"bid"`
	BidQty        int     `json:"bid_qty"`
	Ask           float64 `json:"ask"`
	AskQty        int     `json:"ask_qty"`
	Open          float64 `json:"open"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	PreClose      float64 `json:"pre_close"`
	Close         float64 `json:"close"`
	MarketStatus  string  `json:"market_status"`
}

func GetSymbolDetail(url string) []string {
	var listSymbol []string
	//var responseObject Response
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiIrODQ5NjYxODcyMzIiLCJzdWIiOiIrODQ5NjYxODcyMzIiLCJkZXZpY2VfaWQiOiJhYmMiLCJ1c2VybmFtZSI6InVzZXIyaHNtIiwicGFydHlfaWQiOiIxMDAwMDIiLCJsb2dpbl90eXBlIjoiU1NPIiwidXNlcl9sb2dpbl90eXBlIjoidXNlcl9tb2JpbGUiLCJleHAiOjIwMDc5NTMzMDguOTUzLCJpYXQiOjE2OTI1OTMzMDh9.YVXuTzLYZpV6c2-9A7KuPDxb-Asjckr2blpSRcjKWdgt4Goq2-jF7blL-sPEMkcV4kpa_EgRXW-g_2zgphyzQgfuNTRcvvS89Y5P9Fv-62bmK--Nv24Fsa9WoAaapw9GEmrMl03EWoCzJLII7zGXS7sij8WudAs59rz8BN-Hs8FK_XZLXiKqvcpF6CeHt4MF_YXbG7tYm8dnAc45rFJZP6aGId5Zz0q0u0hpI0r7MP7Fq6irncaj6hfqhXSgR1qymFVjiwtZ2mPmpS50q64PL-jTUeDKeyc5XpzmCK-wYMJ_eQIT353gtmtNsMOYN2iNKqu7mIDsxi871Nn46mISSUs9fIdHkQei_3z8j5yMoLqBHtfA4kC35JYJTlwbInaKbSC3tQo9fnmVc9uf4gVSxil_ntkuvEz6AX4f4Gc1jRFwlfpckH0PqBb1vzhsAUU0QQRFxmZdijCie6hyDkd_1ZKMv4oDlV77pYWDknFeU2voi9TErWnEyU0pyLKYSshd9Nn-_xhqne-Un_njQ79aHyHe6IMIbHyZOGX8SUJRBvkRQ4Kg4T3GRQPFFu447ykh00_Ohe1EVP32Nl0i4MrPLaJiRqdd5R0ASfWhPnS752Xv3Y20pMElZ5JwjHju7AQN2jkr_hN3KcMHo1ToXouhotGbYEadp_yubh6jfujdLkc")
	//json.Unmarshal(responseData, &responseObject)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Print(err.Error())
	}

	var symbols Response

	if err := json.Unmarshal([]byte(body), &symbols); err != nil {
		fmt.Println("Error unmarshaling Json: ", err)
	}

	for _, v := range symbols.Data {
		listSymbol = append(listSymbol, v.Symbol)
	}

	return listSymbol
}

func randomPriceResponse(symbol string) PriceResponse {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	return PriceResponse{
		Symbol:        symbol,
		TradePrice:    rand.Float64() * 100.0, // Example range, modify as needed
		Volume:        rand.Intn(1000000),     // Example range, modify as needed
		ChangePoint:   rand.Intn(100),         // Example range, modify as needed
		ChangePercent: rand.Intn(100),         // Example range, modify as needed
		Bid:           rand.Float64() * 100.0, // Example range, modify as needed
		BidQty:        rand.Intn(10000),       // Example range, modify as needed
		Ask:           rand.Float64() * 100.0, // Example range, modify as needed
		AskQty:        rand.Intn(10000),       // Example range, modify as needed
		Open:          rand.Float64() * 100.0, // Example range, modify as needed
		High:          rand.Float64() * 100.0, // Example range, modify as needed
		Low:           rand.Float64() * 100.0, // Example range, modify as needed
		PreClose:      rand.Float64() * 100.0, // Example range, modify as needed
		Close:         rand.Float64() * 100.0, // Example range, modify as needed
		MarketStatus:  "main_session",         // Example, modify as needed
	}
}

func sendDataWith(ctx context.Context, pb nats.NatsPubsub, listSymbol []string) {
	go func() {
		time.Sleep(time.Second * 2)

		for {
			time.Sleep(time.Millisecond)
			for _, v := range listSymbol {
				priceOfSymbol := randomPriceResponse(v)

				body, _ := json.Marshal(priceOfSymbol)
				err := pb.Publish(ctx, common.StreamingPriceChannel, string(body))

				if err != nil {
					fmt.Println("Send pubsub error")
					return
				}
			}

		}
	}()
}

func StreamingPrice(sc goservice.ServiceContext) error {
	pb := sc.MustGet(common.PluginNatsPubsub).(nats.NatsPubsub)
	ctx := context.Background()

	const url = "https://cgsi-api.equix.app/v1/live/symbol?search=A"

	listSymbols := GetSymbolDetail(url)

	sendDataWith(ctx, pb, listSymbols)

	return nil
}
