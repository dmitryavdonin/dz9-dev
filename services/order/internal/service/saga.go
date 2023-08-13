package service

import (
	"context"
	"order/internal/model"
	"order/internal/service/adapters/book"
	"order/internal/service/adapters/payment"
	"order/internal/service/adapters/store"

	"github.com/sirupsen/logrus"
)

type SagaService struct {
	storeApi   store.StoreApi
	paymentAPi payment.PaymentApi
	bookApi    book.BookApi
}

func NewSagaService(storeApi store.StoreApi, paymentAPi payment.PaymentApi, bookApi book.BookApi) *SagaService {
	return &SagaService{
		storeApi:   storeApi,
		paymentAPi: paymentAPi,
		bookApi:    bookApi,
	}
}

func (s *SagaService) CreateOrder(ctx context.Context, order model.Order) model.StatusResponse {

	logrus.Printf("Saga CreateOrder(): BEGIN order_id = %d", order.ID)

	var response = model.StatusResponse{}

	// place the order in store
	var storeOrderInfo = store.StoreOrderInfo{
		OrderId:  order.ID,
		BookId:   order.BookId,
		Quantity: order.Quantity,
	}

	logrus.Printf("Saga CreateOrder(): Try to place order in Store, order_id = %d", order.ID)
	result, err := s.storeApi.PlaceOrderInStore(ctx, storeOrderInfo)

	if err != nil {
		logrus.Errorf("Saga CreateOrder(): Cannot place order in store, order_id = %d , error = %s", order.ID, err.Error())
		response.Status = "failed"
		response.Reason = err.Error()
		return response
	}

	if result.Status == "failed" {
		logrus.Errorf("Saga CreateOrder(): Cannot place order in store, order_id = %d, status = %s, reason = %s", order.ID, result.Status, result.Reason)
		response.Status = result.Status
		response.Reason = result.Reason
		return response
	}
	// get the book price
	logrus.Printf("Saga CreateOrder(): Try to get book price, order_id = %d, book_id = %d", order.ID, order.BookId)
	bookPrice, err := s.bookApi.GetBookPrice(ctx, order.BookId)
	if err != nil {
		logrus.Errorf("Saga CreateOrder(): Cannot get book price for book_id = %d, order_id = %d, error = %s", order.BookId, order.ID, err.Error())
		response.Status = "failed"
		response.Reason = err.Error()

		// revert all saga steps
		// cancel placement the order in store
		if err = s.storeApi.CancelOrderInStore(ctx, order.ID, "Cannot get book price"); err != nil {
			logrus.Errorf("Saga CreateOrder(): Cannot cancel order in store, order_id = %d, error = %s", order.ID, err.Error())
		}

		return response
	}

	// do payment
	money := bookPrice * order.Quantity
	logrus.Printf("Saga CreateOrder(): Try to do payment money = %d, order_id = %d, book_id = %d", money, order.ID, order.BookId)
	result, err = s.paymentAPi.DoPayment(ctx, payment.PaymentInfo{OrderId: order.ID, UserId: order.UserId, Money: money})
	if err != nil {
		logrus.Errorf("Saga CreateOrder(): Cannot do payment for book_id = %d, order_id = %d, money = %d, error = %s", order.BookId, order.ID, money, err.Error())
		response.Status = "failed"
		response.Reason = err.Error()

		// revert all saga steps
		// cancel placement the order in store
		if err = s.storeApi.CancelOrderInStore(ctx, order.ID, "Payment failed"); err != nil {
			logrus.Errorf("Saga CreateOrder(): Cannot cancel order in store, order_id = %d, error = %s", order.ID, err.Error())
		}
		return response
	}

	response.Status = result.Status
	response.Reason = result.Reason

	if response.Status == "failed" {
		logrus.Printf("Saga CreateOrder(): Payment failed for book_id = %d, order_id = %d, reason = %s", order.BookId, order.ID, response.Reason)

		// revert all saga steps
		// cancel placement the order in store
		logrus.Printf("Saga CreateOrder(): Try to cancel order in store, order_id = %d", order.ID)
		if err = s.storeApi.CancelOrderInStore(ctx, order.ID, "Payment failed"); err != nil {
			logrus.Errorf("Saga CreateOrder(): Cannot cancel order in store, order_id = %d, error = %s", order.ID, err.Error())
		}
		return response
	}

	response.Status = result.Status
	response.Reason = result.Reason

	logrus.Printf("Saga CreateOrder(): END order_id = %d, status = %s, reason = %s", order.ID, response.Status, response.Reason)

	return response
}

func (s *SagaService) PlaceOrderInStore(ctx context.Context, storeOrderInfo store.StoreOrderInfo) (model.StatusResponse, error) {
	logrus.Printf("Saga PlaceOrderInStore(): BEGIN order_id = %d, book_id = %d, quantity = %d", storeOrderInfo.OrderId, storeOrderInfo.BookId, storeOrderInfo.Quantity)
	return s.storeApi.PlaceOrderInStore(ctx, storeOrderInfo)
}
