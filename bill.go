package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func makeNewBill(name string) Bill {
	var b = Bill{
		name:  name,
		items: map[string]float64{},
		tip:   0.0,
	}
	return b
}

func (b *Bill) format() string {
	var fs = ""
	fs += "Bill Breakdown:\n"
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
	}

	var total = 0.0

	for _, v := range b.items {
		total += v
	}

	total += b.tip

	fs += fmt.Sprintf("%-25v ...$%.2f\n", "tip:", b.tip)

	fs += fmt.Sprintf("%-25v ...$%.2f\n", "total:", total)

	return fs
}

func (b *Bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *Bill) updateTip(newTip float64) {
	b.tip = newTip
}

func (b *Bill) save() {
	var data = []byte(b.format())

	var root_loc, _ = os.Getwd()

	var err = os.WriteFile(filepath.Join(root_loc, "bills", b.name+".txt"), data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill saved successfully")
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	var input, err = r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() Bill {
	var reader = bufio.NewReader(os.Stdin)
	// fmt.Print("Enter the bill name: ")
	// var name, _ = reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	var name, _ = getInput("Enter the bill name: ", reader)

	var b = makeNewBill(name)

	fmt.Printf("Created Bill - %v\n", b.name)

	return b
}

func promptOptions(b *Bill) {
	var reader = bufio.NewReader(os.Stdin)

	for {
		var opt, _ = getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
		var endPrompt = false

		switch opt {
		case "a":
			var name, _ = getInput("Enter name: ", reader)
			var price, _ = getInput("Enter price: ", reader)

			var p, err = strconv.ParseFloat(price, 64)

			if err != nil {
				fmt.Println("Invalid Number")
			} else {
				b.addItem(name, p)
				fmt.Println(name, price)
			}
		case "t":
			var tip, _ = getInput("Enter tip: ", reader)
			var t, err = strconv.ParseFloat(tip, 64)

			if err != nil {
				fmt.Println("Invalid Number")
			} else {
				b.updateTip(t)
			}
		case "s":
			fmt.Println("You choose to save the bill")
			b.save()
			endPrompt = true
		default:
			fmt.Println("That was not a valid option")
		}

		if endPrompt {
			break
		}
	}
}

func BillOps() {
	// var b = makeNewBill("surendra's bill")
	// b.addItem("pie", 3.3)
	// b.addItem("cake", 5.3)
	// b.addItem("onion soup", 6.3)
	// b.addItem("veg pie", 4.3)
	// b.addItem("toffee pudding", 5.8)
	// b.addItem("coffee", 2.0)
	// b.updateTip(2.0)
	// fmt.Println(b.format())

	var newBill = createBill()
	fmt.Println(newBill)

	promptOptions(&newBill)
}
