package utils
import "fmt"
import "item"
import "buyer"



func Command_line () int{
	var menu int
	for {
		menu = 0 // 첫 메뉴

		fmt.Println("(1) 구매")
		fmt.Println("(2) 잔여 수량 확인")
		fmt.Println("(3) 잔여 마일리지 확인")
		fmt.Println("(4) 배송 상태 확인")
		fmt.Println("(5) 장바구니 확인")
		fmt.Println("(6) 프로그램 종료")
		fmt.Println()
		fmt.Print("실행할 기능을 입력하시오 : ")	

		fmt.Scanln(&menu)
		if menu < 1 || menu > 6 {
			fmt.Println()
			fmt.Println("잘못된 명령어 입니다. 다시 입력하세요.")
			fmt.Println()
			continue
		} else {
			break
		}
	}
	return menu
}

func PrintItemInfo (arr []item.Item) {
	fmt.Println()
	for i:=0 ; i < len(arr) ; i ++ {
		fmt.Printf("%s: %d개\n",arr[i].Name,arr[i].Amount)	
	}
	fmt.Println()
	fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
	fmt.Println()
	fmt.Scanln()
}

func PrintItemInfo2 (arr []item.Item ,man *buyer.Buyer , noo *int, deliverystart chan *map[string]int) {
	for{
		fmt.Println()
		for i:=0 ; i < len(arr) ; i ++ {
			fmt.Printf("(%d) %s: %d원 [잔여수량 %d개]\n", i + 1 , arr[i].Name,arr[i].Price,arr[i].Amount)	
		}
		fmt.Println()
		
		var choice int
		fmt.Print("구매할 물품을 선택하세요 (1~5) : ")
		fmt.Scanln(&choice)
		
		switch {
			case choice == 1:
				buying(arr, man , choice , noo , deliverystart)
				return
			case choice == 2:
				buying(arr, man , choice, noo, deliverystart)
				return
			case choice == 3:
				buying(arr, man , choice , noo , deliverystart)
				return
			case choice == 4:
				buying(arr, man , choice , noo , deliverystart)
				return
			case choice == 5:
				buying(arr, man , choice , noo , deliverystart)
				return
			default:
				fmt.Println()
				fmt.Println("유효한 물품을 선택하세요. (1~6)")
				fmt.Println()
				continue
		}
	}
	
}

func buying(arr []item.Item ,man *buyer.Buyer , choice int , noo *int , deliverystart chan *map[string]int) {
	for {
		var num int
		fmt.Print("구매수량 : ")
		fmt.Scanln(&num)
		fmt.Println()
		for {
			var a int
			fmt.Print("(1)바로구매 , (2)장바구니에 담기 : ")
			fmt.Scanln(&a)
			fmt.Println()
			
			if a == 1 {
				buy(arr, man , choice , num, noo , deliverystart) 
				return	
			} else if a == 2 {
				goToShoppingBag(arr, man , choice , num)
				return
			} else {
				fmt.Println("잘못된 입력입니다. 다시입력하세요")
				continue
			}
		}

	}
}

func goToShoppingBag(arr []item.Item ,man *buyer.Buyer , choice int , num int ) {

	if arr[choice -1].Amount < man.ShoppingBucket[arr[choice-1].Name] + num {
		fmt.Println()
		fmt.Println("물품의 잔여 수량을 초과했습니다")
		fmt.Println()
		return
	}
	
	var check bool = false
	for i := range man.ShoppingBucket {
		if i == arr[choice -1].Name {
			check = true
		}
	}
	
	if check == true {
		man.ShoppingBucket[arr[choice-1].Name] += num
	} else {
		man.ShoppingBucket[arr[choice-1].Name] = num
	}
	
	fmt.Println("상품이 장바구니에 추가되었습니다.")
	fmt.Println()
}

func buy(arr []item.Item ,man *buyer.Buyer , choice int , num int, noo *int , deliverystart chan *map[string]int) {
	defer func() {
		if r:= recover(); r!=nil{
			fmt.Println(r)
			buying(arr, man , choice, noo , deliverystart)
		}
	}()
	
	if *noo >= 5 {
		fmt.Println()
		fmt.Println("배송 한도를 초과했습니다. 배송이 완료되면 주문하세요.")
		fmt.Println()
		return
	}
	
	
	
	if arr[choice-1].CanManBuyIt(man,num) == true {
		arr[choice - 1].Amount -= num
		man.Point -= arr[choice - 1].Price * num
		fmt.Println("상품 주문이 접수되었습니다.")
		
		tmp := map[string]int{}
		tmp[arr[choice-1].Name] = num
		deliverystart <- &tmp //배송시작
		*noo ++
	} else {
		//주문불가 에러처리
		if num < 0 {
			panic("올바른 수량을 입력하세요.")
		}else {
			panic("주문이 불가능합니다.")
		}
	}
}

func RequiredPoint(arr []item.Item ,man *buyer.Buyer) bool {
	//장바구니 담은 모든 상품의 가격과 보유 포인트 비교
	bucketpoint:=0
	for idx , val := range man.ShoppingBucket {
		for i:= 0 ; i < len(arr) ; i ++ {
			if arr[i].Name == idx {
				bucketpoint += arr[i].Price * val
			}
		}
	}
	
	if man.Point < bucketpoint {
		fmt.Println()
		fmt.Printf("필요 마일리지 : %d\n", bucketpoint)
		fmt.Printf("보유 마일리지 : %d\n", man.Point)
		fmt.Println()
		fmt.Printf("마일리지가 %d점 부족합니다.\n", bucketpoint-man.Point)
		return false
	}
	return true
}

func ExcessAmount(arr []item.Item ,man *buyer.Buyer) bool {
	for idx , val := range man.ShoppingBucket {
		for i := 0 ; i < len(arr) ; i ++ {
			if arr[i].Name == idx {
				if arr[i].Amount < val {
					fmt.Println()
					fmt.Printf("%s 잔여수량: %d, %d개 수량부족.\n", arr[i].Name , arr[i].Amount , val - arr[i].Amount)
					fmt.Println()
					return false
				}
			}
		}
	}
	return true
}

func BuyInBucket(arr []item.Item ,man *buyer.Buyer ,noo *int , deliverystart chan *map[string]int) {
	// //구매
	if *noo >= 5 {
		fmt.Println()
		fmt.Println("배송 한도를 초과했습니다. 배송이 완료되면 주문하세요.")
		fmt.Println()
		return
	}
	for idx , val := range man.ShoppingBucket {
		for i:= 0 ; i < len(arr) ; i ++ {
			if arr[i].Name == idx {
				arr[i].Amount -= val
				man.Point -= arr[i].Price * val
			}
		}
	}
	deliverystart <- &man.ShoppingBucket
	//man.ShoppingBucket = map[string]int{} 여기서 쇼핑백을 초기화하면 deliveryStatus 에서 사용불가 (포인터)
	fmt.Println()
	fmt.Println("상품 주문이 접수되었습니다.")
	fmt.Println()
	*noo ++
	 //배송시작
}
