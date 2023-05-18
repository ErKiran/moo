package sample

import (
	"fmt"
	"os"
	"strings"

	"github.com/speecan/moo/game"
)

// EstimateHuman is played by human
func EstimateHuman(difficulty int) game.Estimate {
	return func(fn game.Question) (res []int) {
		var input string
		fmt.Print("?: ")
		fmt.Fscanln(os.Stdin, &input)
		guess := game.Str2Int(strings.Split(input, ""))
		fn(guess)
		return guess
	}
}

// EstimateWithRandom is idiot algo.
// returns estimate number with simply random
func EstimateWithRandom(difficulty int) game.Estimate {
	return func(fn game.Question) (res []int) {
		r := game.GetMooNum(difficulty)
		fn(r)
		return r
	}
}

// EstimateWithRandom2 is idiot algo.
// exclude duplicate queries
func EstimateWithRandom2(difficulty int) game.Estimate {
	query := make([][]int, 0)
	isDuplicated := func(i []int) bool {
		for _, v := range query {
			if game.Equals(v, i) {
				return true
			}
		}
		return false
	}
	return func(fn game.Question) (res []int) {
		var r []int
		for {
			r = game.GetMooNum(difficulty)
			if !isDuplicated(r) {
				break
			}
		}
		fn(r)
		query = append(query, r)
		return r
	}
}

// EstimateWithOptimization generates an automated estimate function for a game with a given difficulty.
// The function uses a combination of candidate elimination and feedback from the game to make educated guesses.
func EstimateWithOptimization(difficulty int) game.Estimate {
	candidates := game.GenerateCandidates(difficulty)
	maxCandidates := game.FactorialDivision(10, difficulty)
	filteredCandidates := make([][]int, maxCandidates)

	return func(question game.Question) []int {
		if len(candidates) == 0 {
			fmt.Println("No more candidates.")
			return nil
		}

		estimate := candidates[0]
		candidates = candidates[1:]

		hit, blow := question(estimate)
		filteredIndex := 0

		for i := 0; i < len(candidates); i++ {
			if game.GetHit(candidates[i], estimate) == hit && game.GetBlow(candidates[i], estimate) == blow {
				filteredCandidates[filteredIndex] = candidates[i]
				filteredIndex++
			}
		}

		candidates = filteredCandidates[:filteredIndex]
		return estimate
	}
}
