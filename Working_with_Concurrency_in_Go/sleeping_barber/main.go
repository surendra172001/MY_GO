// package main

// import (
// 	"fmt"
// 	"time"
// )

// var seatCnt = 10
// var cutDuration = 1000 * time.Millisecond
// var shopDuration = 10 * time.Second

// // var arrivalRate = 100

// func main() {

// 	// fmt.Println("Welcome to Barber shop")
// 	fmt.Println("----------------------")

// 	// Create a channel for the clients
// 	shopClosing := make(chan bool)
// 	clientsChan := make(chan string, seatCnt)
// 	doneChan := make(chan bool)

// 	// Create the structure for the barbershop
// 	shop := BarberShop{
// 		seatCnt:     seatCnt,
// 		cutDuration: cutDuration,
// 		clientsChan: clientsChan,
// 		barberCnt:   0,
// 		barberDone:  doneChan,
// 	}

// 	// Add barbers to the barbershop
// 	shop.AddBarber("Frank")

// 	// Create a routine to open the barbershop
// 	shop.OpenShop()

// 	// create a routine that will close the barbershop after shop duration

// 	go func() {
// 		<-time.After(shopDuration)
// 		shopClosing <- true
// 		shop.CloseShop()
// 	}()

// 	// Create a routine that will send clients to the barbershop

// 	// go func() {
// 	// 	randMillis := r.Int() % (2 * arrivalRate)
// 	// 	for {
// 	// 		select {
// 	// 		case shop.clientsChan <- time.After(randMillis * time.Millisecond):
// 	// 		}
// 	// 	}
// 	// }()

// 	close(shopClosing)
// }

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// --------------- problem description -------------
// There could be a barber shop with one or more barbers and one or more seats
// If a customer arrives at the and it is closed the customer goes home
// If a customer arrives and shop is open then customer checks if empty chairs are available or not
// If empty chairs are not present the customer goes home
// If empty chairs are present the customer checks if any barber is available
// If barber is available then customers wakes up the barber for cutting its hair
// If barber is not available then customer occupies one chair and waits for its turn
// When the shop is closing then any new customer will not be taken in
// However the barber will cut hairs of all the waiting customers in the shop and then it will go home

// variables
var seatingCapacity = 2
var cutDuration = 1000 * time.Millisecond
var arrivalRate = 100

// welcome message
// create some channels
// create data structure for barber shop
// create barber shop
// add barbers to the shop
// run a routine for the shop
// add clients to the shop
// close the barber shop after some time
// print the ending message
func main() {
	fmt.Print(`
░██████╗██╗░░░░░███████╗███████╗██████╗░██╗███╗░░██╗░██████╗░  ██████╗░░█████╗░██████╗░██████╗░███████╗██████╗░
██╔════╝██║░░░░░██╔════╝██╔════╝██╔══██╗██║████╗░██║██╔════╝░  ██╔══██╗██╔══██╗██╔══██╗██╔══██╗██╔════╝██╔══██╗
╚█████╗░██║░░░░░█████╗░░█████╗░░██████╔╝██║██╔██╗██║██║░░██╗░  ██████╦╝███████║██████╔╝██████╦╝█████╗░░██████╔╝
░╚═══██╗██║░░░░░██╔══╝░░██╔══╝░░██╔═══╝░██║██║╚████║██║░░╚██╗  ██╔══██╗██╔══██║██╔══██╗██╔══██╗██╔══╝░░██╔══██╗
██████╔╝███████╗███████╗███████╗██║░░░░░██║██║░╚███║╚██████╔╝  ██████╦╝██║░░██║██║░░██║██████╦╝███████╗██║░░██║
╚═════╝░╚══════╝╚══════╝╚══════╝╚═╝░░░░░╚═╝╚═╝░░╚══╝░╚═════╝░  ╚═════╝░╚═╝░░╚═╝╚═╝░░╚═╝╚═════╝░╚══════╝╚═╝░░╚═╝

██████╗░██████╗░░█████╗░██████╗░██╗░░░░░███████╗███╗░░░███╗
██╔══██╗██╔══██╗██╔══██╗██╔══██╗██║░░░░░██╔════╝████╗░████║
██████╔╝██████╔╝██║░░██║██████╦╝██║░░░░░█████╗░░██╔████╔██║
██╔═══╝░██╔══██╗██║░░██║██╔══██╗██║░░░░░██╔══╝░░██║╚██╔╝██║
██║░░░░░██║░░██║╚█████╔╝██████╦╝███████╗███████╗██║░╚═╝░██║
╚═╝░░░░░╚═╝░░╚═╝░╚════╝░╚═════╝░╚══════╝╚══════╝╚═╝░░░░░╚═╝
`)
	color.Yellow("Sleeping barber problem")
	color.Yellow("-----------------------")

	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		NumberOfBarbers: 0,
		HairCutDuration: cutDuration,
		ClientChan:      clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	shop.AddBarber("Frank")
	shop.AddBarber("Gabriel")
	shop.AddBarber("Susan")
	shop.AddBarber("Kelly")
	shop.AddBarber("Pat")

	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(10 * time.Second)
		shopClosing <- true
		shop.CloseShopForDay()
		closed <- true
	}()

	go func() {
		for i := 0; ; i++ {
			randomMilliSeconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Duration(randomMilliSeconds) * time.Millisecond):
				shop.AddClient(fmt.Sprintf("client%d", i))
			}
		}
	}()

	<-closed
}
