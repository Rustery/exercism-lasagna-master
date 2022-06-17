package lasagna

// The estimate for the total preparation time based on the number of layers
func PreparationTime(layers []string, time int) int {
	if time <= 0 {
		time = 2
	}
	return len(layers) * time
}

// The quantity of noodles and sauce needed to make your meal
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, v := range layers {
		switch v {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// Modify the list of your ingredients directly with friend's secret ingredient
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// Calculate the amounts for different numbers of portions
func ScaleRecipe(quantities []float64, portions int) (amount []float64) {
	for _, v := range quantities {
		amount = append(amount, v*float64(portions)/2)
	}
	return amount
}
