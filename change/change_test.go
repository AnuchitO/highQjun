package change

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var successCases = []struct {
	name   string
	coins  []int
	target int
	want   []int
}{
	{
		"single coin change",
		[]int{1, 5, 10, 25, 100},
		25,
		[]int{25},
	},
	{
		"multiple coin change",
		[]int{1, 5, 10, 25, 100},
		15,
		[]int{5, 10},
	},
	{
		"change with Lilliputian Coins",
		[]int{1, 4, 15, 20, 50},
		23,
		[]int{4, 4, 15},
	},
	{
		"change with Lower Elbonia Coins",
		[]int{1, 5, 10, 21, 25},
		63,
		[]int{21, 21, 21},
	},
	{
		"large target values",
		[]int{1, 2, 5, 10, 20, 50, 100},
		999,
		[]int{2, 2, 5, 20, 20, 50, 100, 100, 100, 100, 100, 100, 100, 100, 100},
	},
	{
		"possible change without unit coins available",
		[]int{2, 5, 10, 20, 50},
		21,
		[]int{2, 2, 2, 5, 10},
	},
	{
		"another possible change without unit coins available",
		[]int{4, 5},
		27,
		[]int{4, 4, 4, 5, 5, 5},
	},
	{
		"no coins make 0 change",
		[]int{1, 5, 10, 21, 25},
		0,
		[]int{},
	},
}

func TestChangeSuccess(t *testing.T) {
	for _, tc := range successCases {
		t.Run(tc.name, func(t *testing.T) {
			got, _ := Change(tc.coins, tc.target)

			assert.Equal(t, tc.want, got, "it should be equal")
		})
	}
}

var failureCases = []struct {
	name   string
	coins  []int
	target int
	err    error
}{
	{
		"no coins make 0 change",
		[]int{5, 10},
		3,
		ErrNoChange,
	},
}

func TestChangeFailure(t *testing.T) {
	for _, tc := range failureCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := Change(tc.coins, tc.target)

			assert.Equal(t, tc.err, err, "it should be equal")
		})
	}
}
