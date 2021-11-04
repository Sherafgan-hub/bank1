///package types

// Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д.).
//type Money int64

// Currency представляет код валюты
//type Currency string

// Status представляет собой статус платежа
//type Status string

//Предопределённые статусы платежей.
/*const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusIprogress Status = "INPROGRESS"
)*/

// Category представляет собой категорию, в которой был совершён платёж (авто, аптеки, рестораны и т.д.).
/*type Category string

// Коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// PAN представляет номер карты
type PAN string

// Card представляет информацию о платёжной карте
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

// Payment представляет информацию о платеже.
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status	 Status
}

type PaymentSource struct {
	Type    string
	Number  PAN
	Balance Money
}*/
//Payment NotFound