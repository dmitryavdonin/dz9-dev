package payment

type PaymentInfo struct {
	OrderId int
	UserId  int
	Money   int
}

type PaymentResult struct {
	Status string
	Result string
}
