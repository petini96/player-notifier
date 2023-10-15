package player

import (
	"github.com/petini96/fut-message/delivery"
)

type PlayerSMS struct {
	*delivery.TwilioConnector
}

func NewPlayerSMS(t *delivery.TwilioConnector) *PlayerSMS {
	return &PlayerSMS{
		TwilioConnector: t,
	}
}

func (p PlayerSMS) RememberFut(player Player) {
	message := "Boa tarde " + player.Nick + " meu 10! O mosntro da SÉRIE (" + player.Serie + ") da quadra fênix! "
	message += "Segundona braba. Se puder, coloque o nome na lista pra ajudar na organização."
	message += " Tamo junto! Boa semana. "
	// message := "Fala " + player.Nick + ", o mosntro da SÉRIE (" + player.Serie + ") da quadra fênix! "
	// message += "Só de boa meu jogador ? "
	// message += " Aqui é do Futebol de Quinta(Fênix)... "
	// message += " Só passando pra te lembrar que hoje tem fut!!! "
	// message += "Coloque o nome na lista para garantir sua presença. Tamo junto! "

	// if player.Away {
	// 	message += "Faz tempo que você não joga um fut... Bora nessa quinta pow!"
	// }

	p.TwilioConnector.SendSMS(message, player.Number)
}

func (p PlayerSMS) NotificateGame(player Player) {
	message := "JOGO DO SÃO PAULO CANCELADO!!! Calma, calma, é brincadeira mano " + player.Nick + ", o monstro sagrado da quadra fenix! O melhor jogador da SÉRIE (" + player.Serie + ")! "
	message += "Bora jogar bola, se puder, coloque o nome na lista pra ajudar na organização."
	message += " Tamo junto! "

	p.TwilioConnector.SendSMS(message, player.Number)
}
