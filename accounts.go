package upbit

type Accounts interface {
	// order 메서드는 주문을 하되 Config.Timeout 만큼이 지나가면 주문을 자동으로 취소한다.
	// 매수/매도에 둘다 사용한다.
	Order(b *Bot, coin *Coin, side string, volume, price float64) (bool, error)
	// 내부에 있는 upbit.API 에서의 접근을 위해 accounts 를 반환해야 한다.
	Accounts() ([]map[string]interface{}, error)
}