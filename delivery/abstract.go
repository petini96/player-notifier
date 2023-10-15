package delivery

type DeliverySMS interface {
	sendSMS(message string) error
}
