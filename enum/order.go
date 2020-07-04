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
	OrderSideSell OrderSide = 2
)

const (
	OrderTypeLimit OrderType = 1
	OrderTypeMarket OrderType = 2
)

func OrderSideWithInt(i int32) OrderSide {
	if i == 1 {
		return OrderSideBuy
	}

	return OrderSideSell
}