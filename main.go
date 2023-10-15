// Download the helper library from https://www.twilio.com/docs/go/install
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/joho/godotenv"
	"github.com/petini96/fut-message/delivery"
	p "github.com/petini96/fut-message/player"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Erro ao carregar as variáveis de ambiente:", err)
	}

	twilioConnector := delivery.NewTwilioConnector()
	filePath := "storage/json/serie_a.json"

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	var players []p.Player
	if err := json.Unmarshal(data, &players); err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}

	sort.Slice(players, func(i, j int) bool {
		return players[i].Name < players[j].Name
	})

	for _, player := range players {
		// if player.PresenceLevel <= 5 {
		// 	println(player.Name, "(", player.PresenceLevel, ")", ": REBAIXADO PARA 2º DIVISÃO.")
		// }
		// if player.PresenceLevel > 5 && player.PresenceLevel <= 7 {
		// 	println(player.Name, "(", player.PresenceLevel, ")", ": EM OBSERVAÇÃO.")
		// }
		// if player.PresenceLevel > 7 {
		// 	fmt.Println(player.Name, "(", player.PresenceLevel, ")", ": JOGADOR REGULAR.")
		// }

		// for _, position := range player.Position {
		// 	if position == "goalkeeper" {
		// 		fmt.Println(player.Name)
		// 		fmt.Println()
		// 	}
		// }
		playerSMS := p.NewPlayerSMS(&twilioConnector)
		playerSMS.NotificateGame(player)

	}

}
