package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int	
}

type pairs struct {
	pairItems []pair
}

func (p *pairs) Sort() []pair {
	
	for i:= 1; i < len(p.pairItems); i++ {
		key := p.pairItems[i].count
		for j := i - 1; j >= 0 && key < p.pairItems[j].count; j-- {
			p.pairItems[j], p.pairItems[j+1] = p.pairItems[j+1], p.pairItems[j]
		}
		
	}
	return p.pairItems
}

func (p *pairs) AddItem(pairItem pair) {
	p.pairItems = append(p.pairItems, pairItem)
}

func MostAppearItem(items []string) []pair {

	// Tidak menggunakan map karna urutan key pada map tidak pasti

	var pairedItems = pairs{}

	for i := 0; i < len(items); i++ {
		count := 1
		for j := 0; j < len(items); j++{
			if i == j {
				continue
			}
			if items[j] == ""{
				continue
			}
			if items[j] == items[i]{
				count++
				items[j] = ""
			}
		}
		if items[i] != ""{
			pairedItems.AddItem(pair{items[i],count})
		}
	}

		
	return pairedItems.Sort()
}


func main() {
	
	fmt.Println(MostAppearItem([]string{"js","js","golang","ruby","ruby","js","js"}))
	fmt.Println(MostAppearItem([]string{"A","B","B","C","A","A","B","A","D","D"}))
	fmt.Println(MostAppearItem([]string{"football","basketball","tenis"}))
}