package enum

type OrderAction uint8
type OrderSide uint8
type OrderType uint8

const (
	OrderActionCreate OrderAction = 1
	OrderActionCancel OrderAction = 2
)

const (
	OrderSideBuy OrderSide = 1
	OrderSideSell OrderAction = 2
)

const (
	OrderTypeLimit OrderType = 1
	OrderTypeMarket OrderAction = 2
)

