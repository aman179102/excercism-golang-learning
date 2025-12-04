package lasagna

// PreparationTime calculates total preparation time based on layers count and avg time.
func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		avgTime = 2
	}
	return len(layers) * avgTime
}

// Quantities returns total noodles (grams) and sauce (liters) based on layers.
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// AddSecretIngredient replaces "?" in myList with last item of friendsList.
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe scales ingredients for desired number of portions.
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaled := make([]float64, len(quantities))
	for i, qty := range quantities {
		scaled[i] = qty * float64(portions) / 2
	}
	return scaled
}
