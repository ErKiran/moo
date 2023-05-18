# moo

## setting

### Edit benchNum in main.go (recommend: 10000)

already set 1 for human input/estimating

### fill Estimate func type

以下の関数を実装し、 `main.go` から呼び出す

```go
func(difficulty int) game.Estimate
```

`game.Estimate` の func は1度のゲーム(answerが同じ間)の中で何度も call される

`game.Question` を予想した答えと共に実行すると、これの `Hit` と `Blow` が返却され質問した回数がカウントされる

実装した `func(difficulty int) game.Estimate` は `goroutine` セーフである必要がある

> ベンチマークのため、並列で呼び出される

`game` package は自由に利用して良い

## run benchmark

```bash
$ go run main.go
```

- example.)

```
~~~
avg. spent: 883.156µs avg. estimates count: 5.5789
```

> 理論値不明 (5.3くらい?)

for using an example algorithm (too idiot)

`sample/dummy.go`

```go
func EstimateWithRandom(difficulty int) game.Estimate {
	return func(fn game.Question) (res []int) {
		r := game.GetMooNum(difficulty)
		fn(r)
		return r
	}
}
```

`main.go`

```go
	for n := 0; n < benchNum; n++ {
		queue <- sample.EstimateHuman(difficulty)
	}
```



# Changes (KIRAN ADHIKARI)

Benchmarks with `EstimateWithOptimization` function

``` sh 
avg. spent: 233.091µs avg. estimates count: 5.5686
```

``` sh 
avg. spent: 210.879µs avg. estimates count: 5.5648
```


**Summary for the changes and approach taken to optimized the guess game**

Cracking `Fisher-Yates` Shuffle

### EstimateWithOptimization:

* This function generates an automated estimate function for a game with a given difficulty.
* It starts by generating all possible candidates using the GenerateCandidates function.
* It calculates the maximum number of candidates using the FactorialDivision function.
* This function returns an anonymous function that takes a game.Question and returns an estimate.

In each invocation of the anonymous function:
1. It checks if there are no more candidates left and returns nil if so.

2. It selects the first candidate from the list of candidates and removes it.
3. It obtains the hit and blow values by calling the question function with the selected estimate.
4. It filters the candidates based on the hit and blow values, keeping only the ones that match.
5. The filtered candidates are stored in the filteredCandidates slice.
6. The candidates are updated to contain only the filtered candidates.
The selected estimate is returned.


### FactorialDivision
* This function calculates the division of two factorials: n! / (n-k)!.
* It performs the computation by multiplying numbers from n down to n-k.
* This function checks for invalid inputs (n < k or n <= 0 or k <= 0) and returns 0 in such cases.

### GenerateCandidates
* This function generates all possible combinations of unique digits for a given difficulty level.
* It initializes an empty slice to store the candidates.
* It uses a recursive function, generateCandidates, to generate the candidates.
* The generateCandidates function takes two parameters: the current digit being processed and the current candidate being formed.
* If the current digit is equal to the desired difficulty level, a new candidate is created by copying the current candidate and appended to the list of candidates.
* Otherwise, it iterates through digits from 0 to 9 and checks if the digit is not used in the current candidate. If not used, it marks the digit as used, updates the current candidate with the digit, and recursively calls generateCandidates with the next digit.
* After the recursive call, the digit is marked as not used to allow it to be used in other candidates.
* The function starts the recursive process with the initial values of 0 for the current digit and an empty candidate.
* Finally, it returns the list of generated candidates.