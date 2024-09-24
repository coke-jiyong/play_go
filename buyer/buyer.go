package buyer
import "fmt"

type Buyer struct {
	Point int
	ShoppingBucket map[string]int //물품:수량
}

func NewBuyer() *Buyer {
	man:= new(Buyer)
	man.Point = 1000000
	man.ShoppingBucket = map[string]int{}
	return man
}

func (b Buyer) PrintUserInfo() {
	fmt.Printf("현재 잔여 마일리지는 %d점입니다.\n", b.Point)
	fmt.Print("엔터를 입력하면 메뉴 화면으로 돌아갑니다.")
	fmt.Scanln()
}

func (b Buyer) PrintShoppingBucket() (int, error) {
	if len(b.ShoppingBucket) == 0 {
		fmt.Println()
		fmt.Println("장바구니가 비었습니다.")
		fmt.Println()
		return 0 , fmt.Errorf("empty")
	}
	fmt.Println()
	fmt.Println("[장바구니 목록]")
	for i , v := range b.ShoppingBucket {
		fmt.Printf("%s : %d개\n", i ,v)
	}
	fmt.Println()
	fmt.Println("엔터를 입력하면 장바구니 메뉴 화면으로 갑니다.")
	fmt.Println()
	fmt.Scanln()
	
	var bucketmenu int
	for {
		fmt.Println("(1) 장바구니 상품 주문")
		fmt.Println("(2) 장바구니 초기화")
		fmt.Println("(3) 메뉴로 돌아가기")
		fmt.Print("실행할 기능을 입력하시오 :")
		fmt.Scanln(&bucketmenu)
		fmt.Println()
		
		if bucketmenu != 1 && bucketmenu != 2 && bucketmenu != 3{
			fmt.Println()
			fmt.Println("잘못된 명령어. 다시 입력하세요.")
			fmt.Println()
			continue
		} else {
			return bucketmenu , nil
		}
	}
}

func (b *Buyer) ClearShoppingBucket() {
	b.ShoppingBucket = map[string]int{}
	fmt.Println()
	fmt.Println("장바구니 초기화 완료.",len(b.ShoppingBucket))
	fmt.Println()
}