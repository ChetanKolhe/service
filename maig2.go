package main

import (
	"fmt"
	"time"
)

// Solution

// user 1   --> solution ---> fetch ---> solution datastructure -->
//  --> cached the solution
//  user 1 --> it fetched from cached
//  It should get deleted after some time , 15 minute

type Solution struct {
	name string
}

func getSolution() []Solution {

	solutionList := []Solution{
		Solution{name: "Solution_1"},
		Solution{name: "Solution_2"},
		Solution{name: "Solution_3"},
		Solution{name: "Solution_4"},
		Solution{name: "Solution_5"},
	}

	return solutionList
}

type SolutionFetcher struct {
	solutions []Solution
	cache     map[string]Solution
}

func clearSolution(solutionName string, solutionFetcher *SolutionFetcher) {
	time.Sleep(2 * time.Second)
	delete(solutionFetcher.cache, solutionName)
}

func (sl *SolutionFetcher) getSolution(solutonName string) (Solution, error) {

	// if solution found in cache
	value, ok := sl.cache[solutonName]
	if ok {
		fmt.Println("fetch from cache")
		return value, nil
	}

	// if solution not found seach the solution and add to cache

	for _, solution := range sl.solutions {

		if solution.name == solutonName {

			fmt.Println("fetch from db")
			sl.cache[solution.name] = solution
			go clearSolution(solutonName, sl)
			return solution, nil
		}
	}

	return Solution{}, fmt.Errorf("Solution not found")
}

func main() {

	solutionFet := SolutionFetcher{
		solutions: getSolution(),
		cache:     map[string]Solution{},
	}

	fmt.Println(solutionFet.getSolution("Solution_1"))
	fmt.Println(solutionFet.getSolution("Solution_1"))
	fmt.Println(solutionFet.getSolution("Solution_1"))
	fmt.Println(solutionFet.getSolution("Solution_1"))
	time.Sleep(3 * time.Second)
	fmt.Println(solutionFet.getSolution("Solution_1"))

}
