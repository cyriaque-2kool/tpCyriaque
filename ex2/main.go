package main

import "fmt"

func Ft_missing(num []int) int {
	n := len(nums)
	sumAttendu := n * (n + 1) / 2
	Vraisum := 0
	for _, value := range nums {
		Vraisum += num
		return sumAttendu - Vraisum

	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 8}
	missing := Ft_missing(nums)
	fmt.Println("Le nombre manquant est le : ", missing)

}

