package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessages = []string{
	"¡Buen Trabajo!",
	"¡Bien Hecho!",
	"Deberias Comprar un Boleto de Loteria",
}

var loseMessages = []string{
	"¡Qué Lastima!",
	"¡Intentalo de Nuevo!",
	"Hoy simplemente no es tu día.",
}

var drawMessages = []string{
	"Las grandes mentes piensan igual.",
	"Oh no. Intentalo de nuevo.",
	"Nadie gana, pero puedes intentarlo de nuevo",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var ComputerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		ComputerChoice = "La computadora Eligio PIEDRA"
	case PAPER:
		computerChoiceInt = PAPER
		ComputerChoice = "La computadora Eligio PAPEL"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		ComputerChoice = "La computadora Eligio TIJERA"
	}

	messageInt := rand.Intn(3)

	var message string

	if playerValue == computerValue {
		roundResult = "Es un empate"
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "El jugador Gana"
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "La computadora Gana"
		message = loseMessages[messageInt]
	}

	return Round{
		Message:           message,
		ComputerChoice:    ComputerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
