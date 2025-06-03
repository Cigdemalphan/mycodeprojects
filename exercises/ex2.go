package main

import "fmt"

func main(){

var temperature int 
var humidity int

fmt.Println("What's the temperature?")
fmt.Scanln(&temperature)
fmt.Println("What's the humidity?")
fmt.Scanln(&humidity)

if (temperature <10){
	if (humidity>80){
		fmt.Println("Wear a waterproof coat")
	}else{
		fmt.Println("Wear a warm coat.")
	}
}else if (temperature >=10 && temperature <=20){
	if (humidity>80){
		fmt.Println("Wear a waterproof jacket.")
	}else{
		fmt.Println("Wear a jacket.")
	}
	}else {
	if (humidity>80){
		fmt.Println("Wear a light clothes and carry an umbrella")
		}else{
			fmt.Println("Wear a light clothes.")
		}
	}
}
   