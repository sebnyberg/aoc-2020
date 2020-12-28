package a25_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_day25(t *testing.T) {
	// cardPubKey := 5764801
	// doorPubKey := 17807724
	cardPubKey := 10212254
	doorPubKey := 12577395
	cardLoopSize := findLoopSize(7, cardPubKey)
	t.Logf("card loop size: %v", cardLoopSize)
	doorLoopSize := findLoopSize(7, doorPubKey)
	t.Logf("door loop size: %v", doorLoopSize)
	cardPrivKey := doorPubKey
	for i := 1; i < cardLoopSize; i++ {
		cardPrivKey = transform(doorPubKey, cardPrivKey)
	}
	t.Logf("card private key: %v", cardPrivKey)
	doorPrivKey := cardPubKey
	for i := 1; i < doorLoopSize; i++ {
		doorPrivKey = transform(cardPubKey, doorPrivKey)
	}
	t.Logf("door private key: %v", doorPrivKey)
	require.Equal(t, doorPrivKey, cardPrivKey)
}

func findLoopSize(subj int, pubkey int) (loopsize int) {
	val := subj
	loopsize++
	for val != pubkey {
		val = transform(subj, val)
		loopsize++
	}
	return loopsize
}

func transform(subj int, val int) int {
	val *= subj
	val %= 20201227
	return val
}
