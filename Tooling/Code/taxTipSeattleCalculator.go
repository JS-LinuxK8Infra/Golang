package main

import (
	"fmt"
)

func thankYou() {
	fmt.Print("Thank you for your input. Please follow this app for additional features as they are added. \n")
}

func printCost(cost float64, calcTax func(r float64) (float64, float64)) {
	calcTipAmount, calcTaxAmount := calcTax(cost)
	fmt.Printf("The tip for your purchase in Seattle should be $%.2f\n", calcTipAmount)
	fmt.Printf("The tax on your purchase in Seattle is $%.2f\n", calcTaxAmount)
}

func modify(m *string) {
	*m = "Please enter the cost of your service, in dollars and cents: "
}

type ServiceTip struct {
	ServiceType string
	TipPercent  float64
	TaxRate     float64
}

func tipTaxCalculator(ServiceType int) func(r float64) (float64, float64) {
	tip := []ServiceTip{
		1: {ServiceType: "Restaurant", TipPercent: .18, TaxRate: .1035},
		2: {ServiceType: "Bar", TipPercent: .20, TaxRate: .1035},
		3: {ServiceType: "Delivery", TipPercent: .15, TaxRate: .1035},
		4: {ServiceType: "Taxi", TipPercent: .12, TaxRate: .1035},
		5: {ServiceType: "Haircut", TipPercent: .18, TaxRate: .1035},
		6: {ServiceType: "Manicure", TipPercent: .20, TaxRate: .1035},
		7: {ServiceType: "Pedicure", TipPercent: .20, TaxRate: .1035},
		8: {ServiceType: "Tattoo", TipPercent: .25, TaxRate: .1035},
		9: {ServiceType: "HairColor and Cut - Stylist", TipPercent: .25, TaxRate: .1035},
	}
	return func(r float64) (float64, float64) {
		calcTip := tip[ServiceType].TipPercent * r
		calcTax := tip[ServiceType].TaxRate * r
		return calcTip, calcTax
	}
}

func main() {
	request := "Please enter the cost of your service, in dollars and cents, dollar sign not needed: "
	modify(&request)

	var cost float64
	var ServiceType int
	defer thankYou()

	fmt.Print(request)
	fmt.Scanf("%f\n", &cost)
	fmt.Printf("Please enter the service you want to calculate the tip and tax for \n 1 - Resteraunt \n 2 - Bar \n 3 - Delivery \n 4 - Taxi \n 5 - Haircut \n 6 - Manicure \n 7 - Pedicure \n 8 - Tattoo \n 9 - HairColor and Cut - Stylist: ")
	fmt.Scanf("%d\n", &ServiceType)

	if ServiceType >= 1 && ServiceType < 10 {
		printCost(cost, tipTaxCalculator(ServiceType))
	} else {
		fmt.Println("Not a valid service, please select one of the provided nine choices.")
	}

}
