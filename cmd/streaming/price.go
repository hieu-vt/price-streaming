package streaming

import (
	"context"
	"encoding/json"
	goservice "github.com/200Lab-Education/go-sdk"
	"log"
	"math/rand"
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

func RandomSymbol() string {
	rand.Seed(time.Now().UnixNano())

	// Define the alphabet
	alphabet := "ABC"

	// Shuffle the alphabet
	runes := []rune(alphabet)
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})

	// Create a string with the first three characters
	str := string(runes[:3])
	return "ASX:" + str
}

func RandomQuote() string {
	// Generate random numbers for the quote
	symbol := RandomSymbol()

	openPrice := rand.Float64() * 100
	closePrice := rand.Float64() * 100
	lastPrice := rand.Float64() * 100
	quoteVol := rand.Intn(1000000)

	// Create the Quote data structure
	quote := Quote{
		OpenPrice:  openPrice,
		ClosePrice: closePrice,
		LastPrice:  lastPrice,
		QuoteVol:   quoteVol,
	}

	mess := Message{
		Symbol: symbol,
		Data:   quote,
	}

	result, _ := json.Marshal(mess)

	return string(result)
}

func StreamingPrice(sc goservice.ServiceContext) error {
	pb := sc.MustGet(common.PluginNatsPubsub).(nats.NatsPubsub)
	ctx := context.Background()

	go func() {
		time.Sleep(time.Second * 5)

		for {
			time.Sleep(time.Millisecond)
			data := RandomQuote()
			log.Println("Message: ", data)
			pb.Publish(ctx, common.StreamingPriceChannel, data)
		}
	}()

	go func() {
		time.Sleep(time.Second * 5)

		for {
			time.Sleep(time.Millisecond)
			data := RandomQuote()
			log.Println("Message: ", data)
			pb.Publish(ctx, common.StreamingPriceChannel, data)
		}
	}()

	go func() {
		time.Sleep(time.Second * 5)

		for {
			time.Sleep(time.Millisecond)
			data := RandomQuote()
			log.Println("Message: ", data)
			pb.Publish(ctx, common.StreamingPriceChannel, data)
		}
	}()

	go func() {
		time.Sleep(time.Second * 5)

		for {
			time.Sleep(time.Millisecond)
			data := RandomQuote()
			log.Println("Message: ", data)
			pb.Publish(ctx, common.StreamingPriceChannel, data)
		}
	}()

	go func() {
		time.Sleep(time.Second * 5)

		for {
			time.Sleep(time.Millisecond)
			data := RandomQuote()
			log.Println("Message: ", data)
			pb.Publish(ctx, common.StreamingPriceChannel, data)
		}
	}()

	go func() {
		time.Sleep(time.Second * 5)

		for {
			time.Sleep(time.Millisecond)
			data := RandomQuote()
			log.Println("Message: ", data)
			pb.Publish(ctx, common.StreamingPriceChannel, data)
		}
	}()

	go func() {
		time.Sleep(time.Second * 5)

		for {
			time.Sleep(time.Millisecond)
			data := RandomQuote()
			log.Println("Message: ", data)
			pb.Publish(ctx, common.StreamingPriceChannel, data)
		}
	}()

	go func() {
		time.Sleep(time.Second * 5)

		for {
			time.Sleep(time.Millisecond)
			data := RandomQuote()
			log.Println("Message: ", data)
			pb.Publish(ctx, common.StreamingPriceChannel, data)
		}
	}()

	go func() {
		time.Sleep(time.Second * 5)

		for {
			time.Sleep(time.Millisecond)
			data := RandomQuote()
			log.Println("Message: ", data)
			pb.Publish(ctx, common.StreamingPriceChannel, data)
		}
	}()

	go func() {
		time.Sleep(time.Second * 5)

		for {
			time.Sleep(time.Millisecond)
			data := RandomQuote()
			log.Println("Message: ", data)
			pb.Publish(ctx, common.StreamingPriceChannel, data)
		}
	}()

	go func() {
		time.Sleep(time.Second * 5)

		for {
			time.Sleep(time.Millisecond)
			data := RandomQuote()
			log.Println("Message: ", data)
			pb.Publish(ctx, common.StreamingPriceChannel, data)
		}
	}()

	return nil
}
