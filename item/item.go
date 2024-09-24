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