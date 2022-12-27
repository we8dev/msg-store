package nats

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/pokrovsky-io/msg-store/internal/entity"
	"github.com/pokrovsky-io/msg-store/internal/usecase"
	"log"
	"sync"
)

type STAN struct {
	conn    stan.Conn
	useCase usecase.Order
}

func New(sc stan.Conn, uc usecase.Order) *STAN {
	return &STAN{
		conn:    sc,
		useCase: uc,
	}
}

// TODO  Переименовать функцию
func (stn *STAN) inputHandler(msg *stan.Msg) {
	var order entity.Order
	if err := json.Unmarshal(msg.Data, &order); err != nil {
		// TODO обработать ошибку
		log.Fatal(err)
	}

	// TODO: Накапливать входящие результаты и коммитить их с опредленной частотой
	stn.useCase.Create(&order)
}

// TODO Переделать логику работы функции
// Здесь будет ждать вечно...
func (stn *STAN) Subscribe(subj string) {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		sub, err := stn.conn.Subscribe(subj, stn.inputHandler)
		if err != nil {
			// TODO: Обработать ошибку
			log.Fatal(err)
		}
		defer sub.Unsubscribe()

		select {}

	}(wg)

	wg.Wait()
}
