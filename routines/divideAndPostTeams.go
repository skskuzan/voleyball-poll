package routines

import (
	"encoding/csv"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"volleyball_bot/helpers"
	"volleyball_bot/models"
)

func DivideAndPostTeams() {
	participants := getParticipantsWithStats()

	// Implement your team balancing algorithm here
	teams := balanceTeams(participants)

	msg:=teamsMessage(teams)

	// Send teams to the group
	helpers.SendMessage(chatID, msg)
}

func teamsMessage(teams [][]models.Player) string{
	msg:=""
	for i,team:=range teams{
		msg += fmt.Sprint("Команда ", i+1,":\n")
		for _,player:=range team{
			msg +=  fmt.Sprint(player.Name,"\n"
		}
		msg += "\n"
	}

	return msg
}

func getParticipantsWithStats() []models.Player {
	file, err := os.Open("players.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var players []models.Player

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		pass, _ := strconv.Atoi(record[3])
		shoot, _ := strconv.Atoi(record[4])
		def, _ := strconv.Atoi(record[5])

		// Map CSV fields to Player struct
		player := models.Player{
			Name:         record[0],
			TelegramTag:  record[1],
			Position:     record[2],
			PassLevel:    pass,
			ShootLevel:   shoot,
			DefenceLevel: def,
		}

		// Check if the player is in the participants list
		if isParticipant(player.TelegramTag) {
			players = append(players, player)
		}
	}

	return players
}

func balanceTeams(players []models.Player) [][]models.Player {
	// Example: Simple split, implement your own logic
	sort.Slice(players, func(i, j int) bool {
		return players[i].OverallSkill() > players[j].OverallSkill()
	})

	teamA := players[:len(players)/2]
	teamB := players[len(players)/2:]

	return [][]models.Player{teamA, teamB}
}
