package main

import (
	"fmt"
	"os"
)

//struct
type bill struct{
	name string
	items map[string]float64
	tip float64
}
//make new bills
func newBill(name string) bill{
	b:= bill{
		name: name,
		items: map[string]float64{},
		tip: 0,
	}
	return b
}

//receiver function- we pass a copy of bill
func (b *bill) format() string {
 fs := "bill breakdown: \n"
 var total float64 = 0

 //list items
 for k, v := range b.items {
	fs += fmt.Sprintf("%v ....$%0.2f \n", k+ ":", v)
	total += v
 }
 //tip
 fs += fmt.Sprintf("%v ...$%0.2f \n", "tip:", b.tip)

 //total
 fs += fmt.Sprintf("%v ...$%0.2f", "total:", total+b.tip)

 return  fs
}

//
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
} 

func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("bill saved to file")
}