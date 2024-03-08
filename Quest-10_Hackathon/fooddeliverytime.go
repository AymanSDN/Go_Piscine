package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var meal food
	if order == "burger" {
		meal.preptime = 15
	} else if order == "chips" {
		meal.preptime = 10
	} else if order == "nuggets" {
		meal.preptime = 12
	} else {
		return 404
	}
	return meal.preptime
}
