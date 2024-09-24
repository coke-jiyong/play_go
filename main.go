package main
 
import "fmt"
import "buyer"
import "item"
import "utils"
 
func main() {
	
	items:= make([]item.Item ,5)
	buyer1:= buyer.NewBuyer()
	
	items[0] = item.Item{"텀블러", 10000, 30}
	items[1] = item.Item{"롱패딩", 500000, 20}
	items[2] = item.Item{"투미 백팩", 400000, 20}
	items[3] = item.Item{"나이키 운동화", 150000, 50}
	items[4] = item.Item{"빼빼로", 1200, 500}
	
	
	for{
		result:=utils.Command_line()
		
		switch {
			case result == 1 : 
			utils.PrintItemInfo2(items,buyer1)
		case result == 2 :
			utils.PrintItemInfo(items)
		case result == 3 :
			buyer1.PrintUserInfo()
		case result == 5 :
			a , err := buyer1.PrintShoppingBucket()
			if err != nil {
				continue
			}
			if a == 1 {
				if utils.RequiredPoint(items, buyer1)==true && utils.ExcessAmount(items, buyer1)==true {
					utils.BuyInBucket(items,buyer1)
				}
			} else if a == 2 {
				buyer1.ClearShoppingBucket()
			} else {
				fmt.Println()
				fmt.Println("주문불가.")
				fmt.Println()
			}
			
		case result == 6 :
			return
		}
	}

	
}



