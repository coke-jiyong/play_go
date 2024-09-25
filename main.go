package main
 
import "fmt"
import "buyer"
import "item"
import "utils"
import "time"
func main() {
	
	items:= make([]item.Item ,5)
	buyer1:= buyer.NewBuyer()
	
	items[0] = item.Item{"텀블러", 10000, 30}
	items[1] = item.Item{"롱패딩", 500000, 20}
	items[2] = item.Item{"투미 백팩", 400000, 20}
	items[3] = item.Item{"나이키 운동화", 150000, 50}
	items[4] = item.Item{"빼빼로", 1200, 500}
	
	noo := 0 //number of ordered
	
	deliverylist := make([]item.Delivery, 5) 
	for i := 0; i < 5; i++ { 
		deliverylist[i] = item.NewDelivery()
	}
	deliverystart := make(chan *map[string]int)
	
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond) //고루틴 순서대로 실행되도록 약간 딜레이
		go deliveryStatus(deliverystart, i , deliverylist , &noo)
	}
	for{
		result:=utils.Command_line()
		
		switch {
			case result == 1 : 
			utils.PrintItemInfo2(items,buyer1, &noo , deliverystart)
		case result == 2 :
			utils.PrintItemInfo(items)
		case result == 3 :
			buyer1.PrintUserInfo()
		case result == 4 :
			if noo == 0 {
				fmt.Println()
				fmt.Println("배송중인 상품이 없습니다.")
				fmt.Println()
			} else {
				printdeliveryStatus(deliverylist)
			}
		case result == 5 :
			a , err := buyer1.PrintShoppingBucket()
			if err != nil {
				continue
			}
			if a == 1 {
				if utils.RequiredPoint(items, buyer1)==true && utils.ExcessAmount(items, buyer1)==true {
					utils.BuyInBucket(items,buyer1, &noo, deliverystart)
				}
			} else if a == 2 {
				buyer1.ClearShoppingBucket()
			} else if a == 3{
				break
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


//구매목록을 바로구매 , 장바구니구매 -> deliveryStatus 로 가져와야함. 
func deliveryStatus(deliverystart chan *map[string]int, i int, deliverylist []item.Delivery, noo *int) {
	for{
		a , ok := <-deliverystart
		if ok { 
			//deliverylist[i].Onedelivery = a
			for idx ,val := range *a {
				deliverylist[i].Onedelivery[idx]=val
			}
			
			deliverylist[i].Status = "주문접수"
			time.Sleep(time.Second*10)

			deliverylist[i].Status = "배송중"
			time.Sleep(time.Second*20)

			deliverylist[i].Status = "배송완료"
			time.Sleep(time.Second*10)
			
			deliverylist[i].Status = "" 
			*a = map[string]int{}
			*noo--
		}
	}	
}

func printdeliveryStatus(deliverylist []item.Delivery){
	fmt.Println()
	for i:= 0 ; i < len(deliverylist) ; i ++ {
		if len(deliverylist[i].Onedelivery) != 0 {
			for idx ,val := range deliverylist[i].Onedelivery {
				if deliverylist[i].Status != "" {
					fmt.Printf("%s, %d개/ " , idx , val)	
				}
			}
			fmt.Printf("배송상황: %s\n" , deliverylist[i].Status)	
		}

	}
	fmt.Println()
	fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
	fmt.Println()
	fmt.Scanln()
}