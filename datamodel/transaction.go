package datamodel

type Transaction struct {
	//EventTimeRes //- время платежа;-
	Amount    float64 `jaon:"amount"`     //- сумма платежа;-
	Cur       string  `jsonL:"cur"`       //- валюта платежа;-
	Payment   string  `json:"payment"`    //- идентификатор платежа;-
	CardToken string  `json:"card_token"` //- токен однозначно коррелирующий с номером карты плательщика
}