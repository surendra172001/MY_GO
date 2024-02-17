// package main

// import (
// 	"time"

// 	"github.com/fatih/color"
// )

// type BarberShop struct {
// 	seatCnt     int
// 	cutDuration time.Duration
// 	clientsChan chan string
// 	barberCnt   int
// 	barberDone  chan bool
// 	open        bool
// }

// func (shop *BarberShop) AddBarber(barber string) {
// 	shop.barberCnt++
// 	color.Yellow("Barber %s enters the shop\n", barber)
// 	go func() {
// 		sleeping := false
// 		for {
// 			if len(shop.clientsChan) == 0 {
// 				color.Cyan("There is no one in wating area so barber %s is taking nap\n", barber)
// 				sleeping = true
// 			}

// 			client, ok := <-shop.clientsChan

// 			if ok {
// 				if sleeping {
// 					color.Cyan("%s wakes up %s from the nap", client, barber)
// 					sleeping = false
// 				}
// 				shop.DoHairCut(barber, client)
// 			} else {
// 				shop.SendBarberHome(barber)
// 			}

// 		}
// 	}()
// }

// func (shop *BarberShop) DoHairCut(barber string, client string) {
// 	color.Green("%s is cutting the hair of %s", barber, client)
// 	time.Sleep(shop.cutDuration)
// 	color.Green("%s is done cutting hair of %s", barber, client)
// }

// func (shop *BarberShop) SendBarberHome(barber string) {
// 	color.Yellow("%s is going home", barber)
// 	shop.barberDone <- true
// }

// func (shop *BarberShop) OpenShop() {
// 	color.Yellow("Shop is open for the day")
// 	shop.open = true
// }

//	func (shop *BarberShop) CloseShop() {
//		color.Yellow("Shop is closing for the day")
//		close(shop.clientsChan)
//		close(shop.barberDone)
//	}
package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	NumberOfBarbers int
	HairCutDuration time.Duration
	BarbersDoneChan chan bool
	ClientChan      chan string
	Open            bool
}

// checks if any client is there
// if no client barber goes to sleep
// if client arrives and barber is sleeping the client wakes him up
// then barber cuts the hair of the client
// then barber checks if shop is closed, if it is closed then barber goes home
func (shop *BarberShop) AddBarber(barber string) {
	shop.NumberOfBarbers++
	color.Cyan("%s has arrived and goes to seating area to do haircut\n", barber)
	go func() {
		isSleeping := false
		for {
			if len(shop.ClientChan) == 0 {
				isSleeping = true
				color.Yellow("There is nothing to do so %s takes a nap", barber)
			}

			client, shopOpen := <-shop.ClientChan
			if shopOpen {
				if isSleeping {
					isSleeping = false
					color.Yellow("%s wakes up %s\n", client, barber)
				}
				shop.CutHair(barber, client)
			} else {
				shop.SendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *BarberShop) CutHair(barber, client string) {
	color.Green("%s is cutting %s's hair", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is finished cutting %s's hair", barber, client)
}

func (shop *BarberShop) SendBarberHome(barber string) {
	color.Cyan("Shop is closed so %s is going home", barber)
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) CloseShopForDay() {
	color.Cyan("Closing shop for the day")
	close(shop.ClientChan)

	for i := 0; i < shop.NumberOfBarbers; i++ {
		<-shop.BarbersDoneChan
	}

	close(shop.BarbersDoneChan)
	color.Green("------------------------------------------------------")
	color.Green("The shop is closed for the day now. Everyone goes home")
}

func (shop *BarberShop) AddClient(client string) {
	color.Yellow("%s arrives", client)

	select {
	case shop.ClientChan <- client:
		if shop.Open {
			color.Green("%s takes a seat in the waiting area", client)
		} else {
			color.Red("shop is already closed, so %s leaves", client)
			return
		}
	default:
		color.Red("Waiting Area is full so %s leaves", client)
	}

}
