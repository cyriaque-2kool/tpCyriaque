package main

import "fmt"

func monkeymoney(amount int, coins []int) []int {
	n := len(coins)
	chosen := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for amount >= coins[i] {
			amount -= coins[i]
			chosen[i]++
		}
	}
	if amount != 0 {
		return nil
	}
	return chosen
}

func main() {
	coins := []int{1, 3, 9}
	amount := 69
	fmt.Println(monkeymoney(amount, coins))
}
