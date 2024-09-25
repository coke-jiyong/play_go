package item

import _"fmt"
import "buyer"
type Item struct {
	Name string
	Price int
	Amount int
}

func (i Item)CanManBuyIt(man *buyer.Buyer , num int) bool {
	if man.Point < i.Price * num || num > i.Amount  || num < 0{
		return false
	}
	return true
}

type Delivery struct {
	Status string
	Onedelivery map[string]int // 한번에 배송하는 물품의 뜻으로 생각
}

func NewDelivery() Delivery {
	d := Delivery{}
	d.Onedelivery = map[string]int{}
	return d
}