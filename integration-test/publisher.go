package main

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
	"sync"
	"time"
)

// TODO: Вынести в конфиг
const (
	clusterID = "test-cluster"
	clientID  = "test-client-2"
	subject   = "foo"
)

// TODO: Добавить отправку файлов
// TODO: Добавить парсинг JSON
// TODO: Создать моковую модель данных для отправки
// TODO Добавить библиотеку для мокирования

func main() {
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	wg := &sync.WaitGroup{}

	wg.Add(1)

	go Publish(sc, wg)

	wg.Wait()
}

var tm, _ = time.Parse(time.RFC3339, "2021-11-26T06:22:19Z")
var sampleData = &Order{
	OrderUid:    "b563feb7b2b84b6test",
	TrackNumber: "WBILMTESTTRACK",
	Entry:       "WBIL",
	Delivery: Delivery{
		Name:    "Test Testov",
		Phone:   "+9720000000",
		Zip:     "2639809",
		City:    "Kiryat Mozkin",
		Address: "Ploshad Mira 15",
		Region:  "Kraiot",
		Email:   "test@gmail.com",
	},
	Payment: Payment{
		Transaction:  "b563feb7b2b84b6test",
		RequestId:    "",
		Currency:     "USD",
		Provider:     "wbpay",
		Amount:       1817,
		PaymentDt:    1637907727,
		Bank:         "alpha",
		DeliveryCost: 1500,
		GoodsTotal:   317,
		CustomFee:    0,
	},
	Items: []Item{
		{
			ChrtId:      9934930,
			TrackNumber: "WBILMTESTTRACK",
			Price:       453,
			Rid:         "ab4219087a764ae0btest",
			Name:        "Mascaras",
			Sale:        30,
			Size:        "0",
			TotalPrice:  317,
			NmId:        2389212,
			Brand:       "Vivienne Sabo",
			Status:      202,
		},
	},
	Locale:            "en",
	InternalSignature: "",
	CustomerId:        "test",
	DeliveryService:   "meest",
	ShardKey:          "9",
	SmId:              99,
	DateCreated:       tm,
	OofShard:          "1",
}

func Publish(sc stan.Conn, wg *sync.WaitGroup) {
	defer wg.Done()

	jsonData, err := json.Marshal(sampleData)
	if err != nil {
		log.Fatal(err)
	}

	for {
		if err := sc.Publish(subject, jsonData); err != nil {
			log.Fatal(err)
		}

		time.Sleep(time.Second)
	}
}
