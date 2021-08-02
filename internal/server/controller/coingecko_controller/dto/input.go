package dto

type InputCoin struct {
	Coin 			string		`json:"coin" validate:"required"`
	Money						`json:"money" validate:"required"`
}

type InputListCoin struct {
	Coins			[]string	`json:"coins" validate:"required"`
	Money						`json:"money" validate:"required"`
}

type Money struct {
	Currency string
}


