package a22_test

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

type player struct {
	id    int   // player id (e.g. 1,2)
	cards []int // list of cards, from top to bottom of pile
}

func (p player) Copy() player {
	new := player{
		id:    p.id,
		cards: make([]int, len(p.cards)),
	}
	copy(new.cards, p.cards)
	return new
}

var ngames int

func Test_day22(t *testing.T) {
	f, err := os.Open("input")
	check(err)
	players := parsePlayers(f)

	winnerID := playGame(&players[0], &players[1])
	fmt.Println("Winner:", winnerID)
	fmt.Println("Score:", score(players[winnerID-1].cards))
	require.Equal(t, 33266, score(players[winnerID-1].cards))
}

func score(cards []int) (res int) {
	for i := len(cards); i > 0; i-- {
		res += i * cards[len(cards)-i]
	}
	return res
}

func playGame(p1 *player, p2 *player) int {
	ngames++
	// gameID := ngames
	// fmt.Printf("=== Game %v ===\n\n", gameID)

	gameHistory := make(map[string]struct{})
	round := 0
	for {
		round++

		// fmt.Printf("-- Round %v -- (Game %v) --\n", round, gameID)
		// printDeck(*p1)
		// printDeck(*p2)

		deckID := makeDeckID(p1.cards, p2.cards)

		// If there was a previous round in the game with the same cards
		// and order, player 1 wins the game immediately
		if _, exists := gameHistory[deckID]; exists {
			return p1.id
		}

		// Mark current game as "played"
		gameHistory[deckID] = struct{}{}

		// Draw one card from each deck
		card1 := p1.cards[0]
		// fmt.Printf("Player %v plays: %v\n", 1, card1)
		card2 := p2.cards[0]
		// fmt.Printf("Player %v plays: %v\n", 2, card2)
		p1.cards = p1.cards[1:]
		p2.cards = p2.cards[1:]

		// The winner will be determined either via a sub-game,
		// or a regular round
		var winnerID int

		// SUB-GAME
		// If both players have at least as many cards remaining in their deck
		// as the value of the card they just drew,
		// determine the winner through recursive combat
		if len(p1.cards) >= card1 && len(p2.cards) >= card2 {
			// fmt.Println("Playing a sub-game to determine the winner...")
			// fmt.Println()

			// Copy players and their decks
			newP1 := p1.Copy()
			newP2 := p2.Copy()
			// printDeck(*p1)
			// printDeck(*p2)
			// printDeck(newP1)
			// printDeck(newP2)

			// The quantity of cards copied to the sub-game is determined
			// by the card that was drawn for each player
			newP1.cards = newP1.cards[:min(card1, len(newP1.cards))]
			newP2.cards = newP2.cards[:min(card2, len(newP2.cards))]

			winnerID = playGame(&newP1, &newP2)
			// fmt.Printf("...anyway, back to game %v.\n", gameID)
		} else {
			// Regular game
			if card1 == card2 {
				log.Fatalln("both players had the same card value")
			}
			if card1 > card2 {
				winnerID = 1
			} else {
				winnerID = 2
			}
		}

		// fmt.Printf("Player %v wins round %v of game %v!\n", winnerID, round, gameID)
		switch winnerID {
		case 1:
			p1.cards = append(p1.cards, card1, card2)
		case 2:
			p2.cards = append(p2.cards, card2, card1)
		}

		// Check if game is over
		if len(p1.cards) == 0 || len(p2.cards) == 0 {
			// fmt.Printf("The winner of game %v is player %v!\n\n", gameID, winnerID)
			return winnerID
		}

		// fmt.Println()
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func makeDeckID(p1cards []int, p2cards []int) string {
	var sb strings.Builder
	for _, card := range p1cards {
		sb.WriteString(strconv.Itoa(card))
		sb.WriteRune(',')
	}
	sb.WriteRune('|')
	for _, card := range p2cards {
		sb.WriteString(strconv.Itoa(card))
		sb.WriteRune(',')
	}
	return sb.String()
}

func printDeck(p player) {
	fmt.Printf("Player %v's deck: ", p.id)
	for i := 0; i < len(p.cards); i++ {
		fmt.Printf("%v", p.cards[i])
		if i < len(p.cards)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("\n")
}

func parsePlayers(f io.Reader) []player {
	sc := bufio.NewScanner(f)

	players := make([]player, 0)

	for i := 0; ; i++ {
		// Parse player header
		if !sc.Scan() {
			return players
		}
		curPlayer := player{
			id:    i + 1,
			cards: make([]int, 0),
		}
		for sc.Scan() {
			row := sc.Text()
			if row == "" {
				break
			}
			n, err := strconv.Atoi(row)
			check(err)
			curPlayer.cards = append(curPlayer.cards, n)
		}
		players = append(players, curPlayer)
	}
}
