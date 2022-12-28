package cache

import (
	"errors"
	"github.com/pokrovsky-io/msg-store/internal/entity"
	"strconv"
	"time"
)

var (
	ErrOrderNotFound = errors.New("order not found")
)

type Cache struct {
	// TODO Указатели или сами сущности?
	orders []*entity.Order
}

func New() *Cache {
	// TODO Установить capacity в зависимости от кол-ва элементов

	preset := make([]*entity.Order, 0, 10)

	for i := 0; i <= 10; i++ {
		d := sampleData
		d.OrderUid = strconv.Itoa(i)
		preset = append(preset, d)
	}

	return &Cache{
		// TODO Хранить в массиве, а не в мапе
		//orders: make([]*entity.Order, 0, 100),
		orders: preset,
	}
}

func (c *Cache) Create(order *entity.Order) {
	id := len(c.orders) + 1

	c.orders[id] = order
}

//func (c *Cache) Get(id int) (*entity.Order, error) {
//	order, ok := c.orders[id]
//	if !ok {
//		return nil, ErrOrderNotFound
//	}
//
//	return order, nil
//}

var tm, _ = time.Parse(time.RFC3339, "2021-11-26T06:22:19Z")
var sampleData = &entity.Order{
	OrderUid:    "b563feb7b2b84b6test",
	TrackNumber: "WBILMTESTTRACK",
	Entry:       "WBIL",
	Delivery: entity.Delivery{
		Name:    "Test Testov",
		Phone:   "+9720000000",
		Zip:     "2639809",
		City:    "Kiryat Mozkin",
		Address: "Ploshad Mira 15",
		Region:  "Kraiot",
		Email:   "test@gmail.com",
	},
	Payment: entity.Payment{
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
	Items: []entity.Item{
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

func (c *Cache) Get(id int) (*entity.Order, error) {
	if id < len(c.orders) {
		return c.orders[id], nil
	}

	return nil, ErrOrderNotFound
}

//func (c *Cache) SaveOrders(orders ...*entity.Order) {
//	startId := len(c.orders) + 1
//
//	for i, order := range orders {
//		c.orders[startId+i] = order
//	}
//}

//
//func (c *Cache) RemoveOrders(ids ...int) {
//	for _, id := range ids {
//		delete(c.orders, id)
//	}
//}
//
//func (c *Cache) ClearStorage() {
//	c.orders = make(map[int]*entity.Order)
//}
