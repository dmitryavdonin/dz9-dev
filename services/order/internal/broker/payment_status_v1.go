package broker

import (
	"encoding/json"
	"order/internal/service"

	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type PaymentStatusEvent struct {
	Data struct {
		OrderID       int64  `json:"order_id"`
		PaymentStatus string `json:"payment_status"`
		Error         string `json:"error"`
	} `json:"data"`
}

type PaymentStatusHandler struct {
	service *service.Order
}

func BuildPaymentStatusHandler(service *service.Order) PaymentStatusHandler {
	return PaymentStatusHandler{service: service}
}

func (gch PaymentStatusHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		gce := PaymentStatusEvent{}
		err := json.Unmarshal(msg.Value, &gce)
		if err != nil {
			logrus.Errorf("Event hasn't been handled.")
			session.MarkMessage(msg, "")
			continue
		}

		if gce.Data.PaymentStatus == "Success" {
			logrus.Infof("SUCCESS! Payment done")
			//service.Order
		} else {
			logrus.Warnf("FAILED! Payment failed: %s", gce.Data.Error)
		}

		// TODO: update order status

		// _, err = gch.db.Exec(context.Background(), `UPDATE orders SET status_id = 2 WHERE id = $1`, gce.Data.OrderID)
		// if err != nil {
		// 	log.Error().Err(err).Msg("Event hasn't been inserted.")
		// }

		session.MarkMessage(msg, "")
	}

	return nil
}

func (PaymentStatusHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (PaymentStatusHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}
