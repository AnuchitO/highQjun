package change

import "errors"

type State struct {
	remaining int
	coins     []int
	change    []int
}

var ErrNoChange = errors.New("could not make change")

func Change(coins []int, target int) (result []int, err error) {
	queue := []State{{remaining: target, coins: coins, change: []int{}}}
	var state State
	for 0 < len(queue) {
		state, queue = queue[0], queue[1:]
		if state.remaining == 0 &&
			(len(result) == 0 || len(state.change) < len(result)) {
			result = state.change
		} else if 0 < state.remaining && 0 < len(state.coins) &&
			(len(result) == 0 || len(state.change) < len(result)) {
			coin := state.coins[len(state.coins)-1]
			queue = append(queue, State{
				state.remaining - coin,
				state.coins,
				append([]int{coin}, state.change...),
			})
			queue = append(queue, State{
				state.remaining,
				state.coins[:len(state.coins)-1],
				state.change,
			})
		}
	}
	if len(result) == 0 && target != 0 {
		err = ErrNoChange
	}
	return
}

func Sum(coins []int) (result int) {
	for _, coin := range coins {
		result += coin
	}
	return
}
