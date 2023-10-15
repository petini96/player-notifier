package delivery

import (
	"fmt"
	"log"
	"os"
	"time"

	api "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/twilio/twilio-go"
)

var client = twilio.NewRestClient()
var twilioFrom = os.Getenv("TWILIO_FROM")

type TwilioConnector struct {
	*api.CreateMessageParams
	*twilio.RestClient
}

func NewTwilioConnector() TwilioConnector {
	twilioConnector := &TwilioConnector{
		CreateMessageParams: &api.CreateMessageParams{},
		RestClient:          client,
	}
	log.Fatal(twilioFrom)
	twilioConnector.setFrom(twilioFrom)
	return *twilioConnector
}

func (t TwilioConnector) SetBody(message string) {
	t.CreateMessageParams.SetBody(message)
}

func (t TwilioConnector) setFrom(number string) {
	t.CreateMessageParams.SetFrom(number)
}

func (t TwilioConnector) SetTo(number string) {
	t.CreateMessageParams.SetTo(number)
}

func (t TwilioConnector) SendSMS(message string, number string) error {
	t.SetBody(message)
	t.SetTo(number)
	resp, err := t.RestClient.Api.CreateMessage(t.CreateMessageParams)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
			return err
		} else {
			fmt.Println(resp.Sid)
			return err
		}
	}
	time.Sleep(1 * time.Second)

	return nil
}
