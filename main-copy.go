// Download the helper library from https://www.twilio.com/docs/go/install
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	api "github.com/twilio/twilio-go/rest/api/v2010"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
)

type Player struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Number string `json:"number"`
	Nick   string `json:"nick"`
	Serie  string `json:"serie"`
	Away   bool   `json:"away"`
}

var client = twilio.NewRestClient()
var twilioFrom = os.Getenv("TWILIO_FROM")

func sendSMS(player Player) {
	message := "Fala " + player.Nick + ", o mosntro da SÉRIE (" + player.Serie + ") da quadra fênix! "
	message += "Só de boa meu jogador ? "
	message += " Aqui é do Futebol de Quinta(Fênix)... "
	message += " Só passando pra te lembrar que hoje tem fut!!! "
	message += "Coloque o nome na lista para garantir sua presença. Tamo junto! "

	if player.Away {
		message += "Faz tempo que você não joga um fut... Bora nessa quinta pow!"
	}
	params := &api.CreateMessageParams{}
	params.SetBody(message)
	params.SetFrom(twilioFrom)
	params.SetTo(player.Number)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}
	time.Sleep(1 * time.Second)
}
func main2() {

	//twilioTo := os.Getenv("TWILIO_TO")

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Erro ao carregar as variáveis de ambiente:", err)
	}

	// Caminho do arquivo "serie_a.json" (altere para o caminho correto no seu sistema)
	filePath := "storage/json/serie_a.json"

	// Lê o conteúdo do arquivo
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	// Cria uma variável para armazenar o slice de players
	var players []Player

	// Decodifica o JSON em um slice de players
	if err := json.Unmarshal(data, &players); err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}

	//time.Sleep(2 * time.Second)
	// Navega pelo slice de players e exibe os detalhes de cada estudante
	// for _, player := range players {

	// 	sendSMS(player)
	// 	// fmt.Println(player.Name)
	// 	// fmt.Println(player.Number)
	// 	fmt.Println()

	// }

}
