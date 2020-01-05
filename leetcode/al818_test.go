package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test818RacecarSample(t *testing.T) {
	//racecarTest(t, 3, "AA")
	//
	//racecarTest(t, 4, "AARRA")
	//
	//racecarTest(t, 5, "AARARAA")
	//racecarTest(t, 6, "AAARA")
	racecarTest(t, 20, "AAAARARAAARA")

	printRacecar(20)
}

func racecarTest(t *testing.T, target int, expected string) {
	result := racecarCmds(target)
	assert.Equal(t, expected, result, "%v-->%v != %v", target, result, expected)
}

func printRacecar(target int) {
	racecarCmds(target)
}