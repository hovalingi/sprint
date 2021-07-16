package main

import "fmt"

func sayHello() {
 	fmt.Println("Hello world")
}
func main()  {
	for i := 0; i < 10; i++ {
	//	fmt.Print(i)
		sayHello()
	}

	click := 0
	distanceCount := 0
	sprintZone := 110

	for distanceCount <= sprintZone {
		click++ // click = click + 1 // click+=1
		distanceCount+=10
		fmt.Println("distanceCount =", distanceCount)
		fmt.Println("click =", click)
	}
}


	// БегунРобот ---> За одно нажатие кнопки 10 метров
	// СпринЗона = 110 метров


