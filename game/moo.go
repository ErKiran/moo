package game

import (
	"fmt"
)

var (
	nums = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// DebugMode makes log output enable
	DebugMode = true
)

type (
	// Question returns hits, blow by guess []int
	Question func(guess []int) (hits, blow int)
	// Estimate returns a next guess
	Estimate func(q Question) (guess []int)
	// Game means one game field
	Game struct {
		difficulty int
		answer     []int
	}
)

// NewGame returns a new game field
func NewGame(d int) *Game {
	if d < 1 || d > 9 {
		fmt.Println(d, "is invalid moo digit, difficulty set to", d)
		d = 4
	}
	return &Game{
		difficulty: d,
		answer:     GetMooNum(d),
	}
}

// GetDifficulty returns digits
func (x *Game) GetDifficulty() int {
	return x.difficulty
}

// GetAnswer returns answer
func (x *Game) GetAnswer() []int {
	return x.answer
}

// GetQuestion returns a question func
func (x *Game) GetQuestion(count *int) Question {
	*count = 0
	return func(g []int) (h, b int) {
		*count++
		h = x.GetHit(g)
		b = x.GetBlow(g)
		if DebugMode {
			fmt.Println(g, ": hits:", h, "blow:", b)
		}
		return
	}
}

// GetHit returns hit count in this game
func (x *Game) GetHit(g []int) int {
	return GetHit(g, x.answer)
}

// GetBlow returns blow count in this game
func (x *Game) GetBlow(g []int) int {
	return GetBlow(g, x.answer)
}

// Equals returns bool which guess = answer
func (x *Game) Equals(g []int) bool {
	return Equals(g, x.answer)
}

// GetHit returns hit
func GetHit(guess []int, answer []int) int {
	count := 0
	if len(guess) != len(answer) {
		return 0
	}
	for i, v := range answer {
		if guess[i] == v {
			count++
		}
	}
	return count
}

// GetBlow returns blow
func GetBlow(guess []int, answer []int) int {
	count := 0
	if len(guess) != len(answer) {
		return 0
	}
	for i, g := range guess {
		for j, a := range answer {
			if g == a && i != j {
				count++
			}
		}
	}
	return count
}

// FactorialDivision calculates the division of two factorials: n! / (n-k)!
// It computes the result by multiplying numbers from n down to n-k.
func FactorialDivision(n, k int) int {
	if n < k || n <= 0 || k <= 0 {
		return 0
	}

	result := 1
	for i := 0; i < k; i++ {
		result *= n - i
	}

	return result
}

// GenerateCandidates generates all possible combinations of unique digits for a given difficulty level.
func GenerateCandidates(difficulty int) [][]int {
	candidates := make([][]int, 0)
	used := make([]bool, 10)

	var generateCandidates func(int, []int)
	generateCandidates = func(digit int, candidate []int) {
		if digit == difficulty {
			newCandidate := make([]int, difficulty)
			copy(newCandidate, candidate)
			candidates = append(candidates, newCandidate)
			return
		}

		for i := 0; i < 10; i++ {
			if !used[i] {
				used[i] = true
				candidate[digit] = i
				generateCandidates(digit+1, candidate)
				used[i] = false
			}
		}
	}

	generateCandidates(0, make([]int, difficulty))
	return candidates
}
