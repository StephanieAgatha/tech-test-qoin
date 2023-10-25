package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Dice struct {
	value int
}

type Player struct {
	Name   string
	Dice   []*Dice
	Points int
}

func (d *Dice) roll() *Dice {
	d.value = rand.Intn(6) + 1
	return d
}

func (p *Player) play() {
	for _, die := range p.Dice {
		die.roll()
	}
}

func removeIndices(dice []*Dice, indices []int) []*Dice {
	newDice := []*Dice{}
	for i, die := range dice {
		if !contains(indices, i) {
			newDice = append(newDice, die)
		}
	}
	return newDice
}

func contains(arr []int, elem int) bool {
	for _, e := range arr {
		if e == elem {
			return true
		}
	}
	return false
}

func countActivePlayers(players []Player) int {
	count := 0
	for _, player := range players {
		if len(player.Dice) > 0 {
			count++
		}
	}
	return count
}

func findWinner(players []Player) Player {
	var winner Player
	for i, player := range players {
		if i == 0 || player.Points > winner.Points {
			winner = player
		}
	}
	return winner
}

func playGame(round int, players []Player) {
	fmt.Printf("==================\nGiliran %d lempar dadu:\n", round)

	// roll dadu secara bersamaann
	for i := range players {
		player := &players[i]
		player.play() //start a game
		fmt.Printf("%s (%d): ", player.Name, player.Points)
		var diceVals []string
		for _, die := range player.Dice {
			diceVals = append(diceVals, fmt.Sprintf("%d", die.value))
		}
		fmt.Printf("%s\n", strings.Join(diceVals, " "))
	}

	roundOver := true
	shouldPrintEval := true

	for i := range players {
		player := &players[i]

		if len(player.Dice) > 0 {
			roundOver = false

			onesIndices := []int{}
			sixIndices := []int{}

			// dadu angka 1 dan 6
			for j, die := range player.Dice {
				if die.value == 1 {
					onesIndices = append(onesIndices, j)
				} else if die.value == 6 {
					sixIndices = append(sixIndices, j)
				}
			}

			// move dadu angka 1
			for _, index := range onesIndices {
				nextIndex := (i + 1) % len(players)
				players[nextIndex].Dice = append(players[nextIndex].Dice, player.Dice[index])
			}
			player.Dice = removeIndices(player.Dice, onesIndices)

			// Hapus dadu angka 6
			player.Dice = removeIndices(player.Dice, sixIndices)

			// poin+1 untuk setiap dadu angka 6
			player.Points += len(sixIndices)

			// print dadu angka 1 setelah evaluasi
			if len(onesIndices) > 0 {
				if shouldPrintEval {
					fmt.Printf("Setelah evaluasi:\n")
					shouldPrintEval = false
				}
				var onesVals []string
				for range onesIndices {
					onesVals = append(onesVals, "1")
				}
				diceVals := []string{}
				for _, die := range player.Dice {
					diceVals = append(diceVals, fmt.Sprintf("%d", die.value))
				}
				fmt.Printf("%s (%d): %s %s\n", player.Name, player.Points, strings.Join(diceVals, " "), strings.Join(onesVals, " "))
			}
		}
	}

	if roundOver || countActivePlayers(players) == 1 {
		// pindahkan angka 1 ke pemain sebelahnya setelah evaluasii
		for i := range players {
			player := &players[i]
			if len(player.Dice) > 0 {
				nextIndex := (i + 1) % len(players)
				if len(players[nextIndex].Dice) > 0 {
					onesIndices := []int{}
					for j, die := range player.Dice {
						if die.value == 1 {
							onesIndices = append(onesIndices, j)
						}
					}
					if len(onesIndices) > 0 {
						players[nextIndex].Dice = append(players[nextIndex].Dice, player.Dice[onesIndices[0]])
						player.Dice = removeIndices(player.Dice, onesIndices[:1])
					}
				}
			}
		}
	}

	if roundOver || countActivePlayers(players) == 1 {
		return
	}

	playGame(round+1, players)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var numPlayers, numDice int
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scan(&numPlayers)
	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scan(&numDice)

	players := make([]Player, numPlayers)
	for i := range players {
		players[i].Name = fmt.Sprintf("Pemain #%d", i+1)
		players[i].Dice = make([]*Dice, numDice)
		for j := range players[i].Dice {
			players[i].Dice[j] = &Dice{}
		}
	}

	playGame(1, players)

	winner := findWinner(players)
	fmt.Printf("Game berakhir karena hanya %s yang memiliki dadu.\n", winner.Name)
	fmt.Printf("Game dimenangkan oleh %s karena memiliki poin lebih banyak dari pemain lainnya.\n", winner.Name)
}
