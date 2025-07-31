package lasagna

const NoodleGramsPerLayer int = 50
const SauceLitersPerLayer float64 = 0.2

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTimePerLayer int) int {
	if prepTimePerLayer == 0 {
		prepTimePerLayer = 2
	}

	return len(layers) * prepTimePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var gramsOfNoodles int = 0
	var litersOfSauce float64 = 0

	for _, layer := range(layers) {
		switch layer {
		case "sauce":
			litersOfSauce += SauceLitersPerLayer
		case "noodles":
			gramsOfNoodles += NoodleGramsPerLayer
		}
	}

	return gramsOfNoodles, litersOfSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaleFactor float64 = float64(portions) / 2.0
	var scaledQuantities = make([]float64, len(quantities))

	for i,_ := range(quantities) {
		scaledQuantities[i] = scaleFactor * quantities[i]
	}

	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
