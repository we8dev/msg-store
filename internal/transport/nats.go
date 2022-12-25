package transport

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"github.com/pokrovsky-io/msg-store/config"
	"github.com/pokrovsky-io/msg-store/internal/model"
	"github.com/pokrovsky-io/msg-store/internal/usecase"
	"log"
	"sync"
)

type STAN struct {
	config  config.NATS
	usecase usecase.UseCase
}

func NewSTAN(cfg config.NATS) *STAN {
	fmt.Println(cfg)

	return &STAN{
		config:  cfg,
		usecase: usecase.NewUscase(),
	}
}

// TODO  Переименовать функцию
func (stn *STAN) inputHandler(msg *stan.Msg) {
	var order model.Order
	if err := json.Unmarshal(msg.Data, &order); err != nil {
		// TODO обработать ошибку
		log.Fatal(err)
	}

	//stn.usecase.Create()

	fmt.Println(order)
}

// TODO Переделать логику работы функции
// Здесь будет ждать вечно...
func (stn *STAN) Subscribe() *stan.Conn {
	sc, err := stan.Connect(stn.config.ClusterID, stn.config.ClientID)
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func(sc stan.Conn, wg *sync.WaitGroup) {
		defer wg.Done()

		sub, err := sc.Subscribe(stn.config.Subject, stn.inputHandler)
		if err != nil {
			log.Fatal(err)
		}
		defer sub.Unsubscribe()

		select {}

	}(sc, wg)

	wg.Wait()

	return &sc
}
