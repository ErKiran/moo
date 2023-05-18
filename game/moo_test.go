package game

import (
	"reflect"
	"testing"
)

func TestGetHit(t *testing.T) {
	{
		g := []int{1, 2, 3, 4}
		a := []int{8, 2, 4, 3}
		expected := 1
		res := GetHit(g, a)
		if expected != res {
			t.Error(g, "and", a, "should hits", expected, "result:", res)
		}
	}
	{
		g := []int{1, 2, 3, 4}
		a := []int{1, 2, 3, 4}
		expected := 4
		res := GetHit(g, a)
		if expected != res {
			t.Error(g, "and", a, "should hits", expected, "result:", res)
		}
	}
	{
		g := []int{1, 2, 3, 4}
		a := []int{5, 6, 7, 9}
		expected := 0
		res := GetHit(g, a)
		if expected != res {
			t.Error(g, "and", a, "should hits", expected, "result:", res)
		}
	}
}

func TestFactorialDivision(t *testing.T) {
	testCases := []struct {
		n        int
		k        int
		expected int
	}{
		// Input: n = 5, k = 3
		// Explanation: n! = 5! = 5 * 4 * 3 * 2 * 1 = 120, (n-k)! = (5-3)! = 2! = 2 * 1 = 2. The expected result is 120 / 2 = 60.
		{5, 3, 60},
		// Input: n = 10, k = 5
		// Explanation: n! = 10! = 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1 = 3,628,800, (n-k)! = (10-5)! = 5! = 5 * 4 * 3 * 2 * 1 = 120. The expected result is 3,628,800 / 120 = 30,240.
		{10, 5, 30240}, // 10! / 5! = 10 * 9 * 8 * 7 * 6 = 30240
		// Input: n = 6, k = 6
		// Explanation: n! = 6! = 6*5*4*3*2*1 = 720, (n-k)! = (6-6)! = 0! = 1. The expected result is 720/1 = 720
		{6, 6, 720},
		{8, 10, 0}, // Invalid input: n < k, should return 0
		{0, 5, 0},  // Invalid input: n <= 0, should return 0
		{5, 0, 0},  // Invalid input: k <= 0, should return 0
	}

	for _, tc := range testCases {
		result := FactorialDivision(tc.n, tc.k)
		if result != tc.expected {
			t.Errorf("FactorialDivision(%d, %d) returned %d, expected %d", tc.n, tc.k, result, tc.expected)
		}
	}
}

func TestGenerateCandidates(t *testing.T) {
	testCases := []struct {
		difficulty int
		expected   [][]int
	}{
		{1, [][]int{{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}}, // Difficulty 1, all single-digit numbers
		{2, [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}, {0, 7}, {0, 8}, {0, 9}, {1, 0}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7}, {1, 8}, {1, 9}, {2, 0}, {2, 1}, {2, 3}, {2, 4}, {2, 5}, {2, 6}, {2, 7}, {2, 8}, {2, 9}, {3, 0}, {3, 1}, {3, 2}, {3, 4}, {3, 5}, {3, 6}, {3, 7}, {3, 8}, {3, 9}, {4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 5}, {4, 6}, {4, 7}, {4, 8}, {4, 9}, {5, 0}, {5, 1}, {5, 2}, {5, 3}, {5, 4}, {5, 6}, {5, 7}, {5, 8}, {5, 9}, {6, 0}, {6, 1}, {6, 2}, {6, 3}, {6, 4}, {6, 5}, {6, 7}, {6, 8}, {6, 9}, {7, 0}, {7, 1}, {7, 2}, {7, 3}, {7, 4}, {7, 5}, {7, 6}, {7, 8}, {7, 9}, {8, 0}, {8, 1}, {8, 2}, {8, 3}, {8, 4}, {8, 5}, {8, 6}, {8, 7}, {8, 9}, {9, 0}, {9, 1}, {9, 2}, {9, 3}, {9, 4}, {9, 5}, {9, 6}, {9, 7}, {9, 8}}}, // Difficulty 2, all two-digit numbers
		// skips others cause it will be toooooooooooooooo long
	}

	for _, tc := range testCases {
		result := GenerateCandidates(tc.difficulty)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("GenerateCandidates(%d) returned %v, expected %v", tc.difficulty, result, tc.expected)
		}
	}
}

func TestGetBlow(t *testing.T) {
	{
		g := []int{1, 2, 3, 4}
		a := []int{8, 2, 4, 3}
		expected := 2
		res := GetBlow(g, a)
		if expected != res {
			t.Error(g, "and", a, "should blows", expected, "result:", res)
		}
	}
	{
		g := []int{1, 2, 3, 4}
		a := []int{1, 2, 3, 4}
		expected := 0
		res := GetBlow(g, a)
		if expected != res {
			t.Error(g, "and", a, "should blows", expected, "result:", res)
		}
	}
	{
		g := []int{1, 2, 3, 4}
		a := []int{4, 3, 2, 1}
		expected := 4
		res := GetBlow(g, a)
		if expected != res {
			t.Error(g, "and", a, "should blows", expected, "result:", res)
		}
	}
}
