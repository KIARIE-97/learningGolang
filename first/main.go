package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func CreateBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("created the bill - ", b.name)

	return b
}
func promptOption(b bill) {
	reader := bufio.NewReader(os.Stdin)
    
	opt, _ := getInput("choose an option (a - add item, s - save bill, t - addtip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("whats item name", reader)
		price, _ := getInput("whats item price", reader)

		//parse price to float64
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("the price must be a number")
			promptOption(b)
		}
		b.addItem(name, p)

		fmt.Println("items added",name, price)
		promptOption(b)	
	case "t":
		tip, _ := getInput("enter tip amount in $", reader)
		
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("the tip must be a number")
			promptOption(b)
		}
		b.updateTip(t)
		fmt.Println("tip added - ", tip)
		promptOption(b)	
	case "s":
		b.save()
		fmt.Println("your bill is saved to file - ", b.name)
	default:
		fmt.Println("that was not a valid option...")
		promptOption(b)			
	}
}

func main() {
	myBill:= CreateBill()
	promptOption(myBill)
	// myBill.items()
	// myBill.tip = 10
	// myBill.items = map[string]float64{"pie": 6.00, "coffee": 20.00}
	// myBill.updateTip(10.00)
	// myBill.addItem("pie", 6.00)
	// myBill.addItem("coffee", 16.00)
	fmt.Println(myBill)
}