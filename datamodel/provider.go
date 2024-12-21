package datamodel

type Provider struct {
	Id          int     `json:"id"`            //- идентификатор провайдера;-
	Conversion  float64 `json:"conversion"`    //- идентификатор провайдера (отношение успешных к неуспешным платежам);-
	AvgTime     float64 `json:"avg_time"`      //- среднеевремя выполнения платежа на провайдере;-
	MinSum      float64 `json:"min_sum"`       //- минимальнаясумма платежа (минимальный чек покупки);-
	MaxSum      float64 `json:"max_sum"`       //- максимальная сумма платежа (максимальный чек покупки);-
	LimitMax    float64 `json:"limit_max"`     //- максимальная общая сумма на провайдере за сутки    (припревышении значения перестают проходить платежи);-
	LimitMin    float64 `json:"limit_min"`     //- минимальнаяобщая сумма на провайдере за сутки (принедостижении значения получаем штрафы за невыработку объема);-
	LimitByCard string  `json:"limit_by_card"` //- лимитпосумме платежей в разрезе одной карты плательщика-
	Commission  float64 `json:"commission"`    //- комиссиявзимаемая провайдером;-
	Currency    string  `json:"currency"`      //- валютапровайдера.
}