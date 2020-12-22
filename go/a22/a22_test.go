package a22_test

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"testing"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

type player struct {
	id     int          // player id (e.g. 1,2)
	cards  []int        // list of cards (top is first)
	played map[int]bool // cards which have been played
}

func Test_day22(t *testing.T) {
	f, err := os.Open("input")
	check(err)
	players := parsePlayers(f)
	fmt.Println(players)
	printDeck(players[0])
	printDeck(players[1])
	for !fight(&players[0], &players[1]) {
		fmt.Println()
		printDeck(players[0])
		printDeck(players[1])
	}
	var winner player
	if len(players[0].cards) == 0 {
		winner = players[1]
	} else {
		winner = players[0]
	}

	var score int
	for i := len(winner.cards); i > 0; i-- {
		score += i * winner.cards[len(winner.cards)-i]
	}
	fmt.Println(score)
	t.FailNow()
}

func fight(p1 *player, p2 *player) (finished bool) {
	// pop both players decks
	card1 := p1.cards[0]
	card2 := p2.cards[0]
	p1.cards = p1.cards[1:]
	p2.cards = p2.cards[1:]
	fmt.Println("Player 1 plays:", card1)
	fmt.Println("Player 2 plays:", card2)
	if card1 == card2 {
		panic("both players played the same card")
	}
	if card1 > card2 {
		fmt.Println("Player 1 wins the round!")
		p1.cards = append(p1.cards, card1, card2)
	} else {
		fmt.Println("Player 2 wins the round!")
		p2.cards = append(p2.cards, card2, card1)
	}
	return len(p1.cards) == 0 || len(p2.cards) == 0
}

func printDeck(p player) {
	fmt.Printf("Player %v's deck: %v", p.id, p.cards[0])
	for i := 1; i < len(p.cards); i++ {
		fmt.Printf(", %v", p.cards[i])
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
			id:     i + 1,
			cards:  make([]int, 0),
			played: make(map[int]bool),
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
