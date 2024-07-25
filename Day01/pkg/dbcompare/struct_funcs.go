package pkg

import (
	pkg "Day01/pkg/dbreader"
	"fmt"
	"slices"
)

// find cakes name
func FindCakeName(rec pkg.Recip) (names []string) {
	for _, value := range rec.Cakes {
		names = append(names, value.Name)
	}
	return names
}

// find cake by their name
func FindCakeByName(name string, rec pkg.Recip) (cake pkg.Cake) {
	for _, value := range rec.Cakes {
		if name == value.Name {
			cake = value
		}
	}
	return cake
}

// find new cakes
func FindAddedCakes(oldNames, newNames []string) []string {
	added := make([]string, 0, len(newNames))
	for _, value := range newNames {
		if !slices.Contains(oldNames, value) {
			added = append(added, value)
		}
	}
	return added
}

// find removed cakes
func FindRemovedCakes(oldNames, newNames []string) []string {
	removed := make([]string, 0, len(oldNames))
	for _, value := range oldNames {
		if !slices.Contains(newNames, value) {
			removed = append(removed, value)
		}
	}
	return removed
}

// compare cakes by methods removed and added cakes
func CompareCakes(oldcake, newcake pkg.Recip) {
	oldNames := FindCakeName(oldcake)
	newNames := FindCakeName(newcake)

	addedCakes := FindAddedCakes(oldNames, newNames)
	removedCakes := FindRemovedCakes(oldNames, newNames)

	for _, value := range addedCakes {
		fmt.Printf("ADDED cake \"%s\"\n", value)
	}

	for _, value := range removedCakes {
		fmt.Printf("REMOVED cake \"%s\"\n", value)
	}
}

// find time changes
func FindTimeChanges(oldcake pkg.Cake, newcake pkg.Cake) {
	if oldcake.Time != newcake.Time {
		fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", newcake.Name, newcake.Time, oldcake.Time)
	}
}

// find ingredient name
func FindIngredients(cake pkg.Cake) (names []string) {
	for _, value := range cake.Ingredients {
		names = append(names, value.Name)
	}
	return names
}

// find ingredient by their name
func FindIngredientByName(name string, cake []pkg.Ingredient) (ingredient pkg.Ingredient) {
	for _, value := range cake {
		if name == value.Name {
			ingredient = value
		}
	}
	return ingredient
}

// find added ingredient
func FindAddedIngredient(oldnames, newnames []string) []string {
	added := make([]string, 0, len(newnames))
	for _, value := range newnames {
		if !slices.Contains(oldnames, value) {
			added = append(added, value)
		}
	}
	return added
}

// find removed ingredient
func FindRemovedIngredient(oldnames, newnames []string) []string {
	removed := make([]string, 0, len(oldnames))
	for _, value := range oldnames {
		if !slices.Contains(newnames, value) {
			removed = append(removed, value)
		}
	}
	return removed
}

// compare counts
func CompareCounts(newingredient, oldingredient pkg.Ingredient, newcakename string) {
	zero := ""
	if newingredient.Count != zero && oldingredient.Count == zero {
		fmt.Printf("ADDED unit count for ingredient \"%s\" for cake \"%s\"\n", newingredient.Name, newcakename)
	} else if newingredient.Count == zero && oldingredient.Count != zero {
		fmt.Printf("REMOVED unit count for ingredient \"%s\" for cake \"%s\"\n", newingredient.Name, newcakename)
	} else if newingredient.Count != oldingredient.Count {
		fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",
			newingredient.Name, newcakename, newingredient.Count, oldingredient.Count)
	}
}

// compare units
func CompareUnits(newingredient, oldingredient pkg.Ingredient, newcakename string) {
	zero := ""
	if newingredient.Unit != zero && oldingredient.Unit == zero {
		fmt.Printf("ADDED unit for ingredient \"%s\" for cake \"%s\"\n", newingredient.Name, newcakename)
	} else if newingredient.Unit == zero && oldingredient.Unit != zero {
		fmt.Printf("REMOVED unit for ingredient \"%s\" for cake \"%s\"\n", newingredient.Name, newcakename)
	} else if newingredient.Unit != oldingredient.Unit {
		fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",
			newingredient.Name, newcakename, newingredient.Unit, oldingredient.Unit)
	}
}

func IngredientsChanges(newCake pkg.Cake, oldCake pkg.Cake) {
	// Find ingredient names for both cakes
	newIngredientsNames := FindIngredients(newCake)
	oldIngredientsNames := FindIngredients(oldCake)

	// Find added and removed ingredients
	addedIngredients := FindAddedIngredient(oldIngredientsNames, newIngredientsNames)
	removedIngredients := FindRemovedIngredient(oldIngredientsNames, newIngredientsNames)

	// Print added ingredients
	for _, ingredientName := range addedIngredients {
		fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", ingredientName, newCake.Name)
	}

	// Print removed ingredients
	for _, ingredientName := range removedIngredients {
		fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", ingredientName, newCake.Name)
	}

	// Compare details for not new and not removed ingredients
	for _, newIngredientName := range newIngredientsNames {
		if slices.Contains(oldIngredientsNames, newIngredientName) {
			newIngredient := FindIngredientByName(newIngredientName, newCake.Ingredients)
			oldIngredient := FindIngredientByName(newIngredientName, oldCake.Ingredients)
			CompareCounts(newIngredient, oldIngredient, newCake.Name)
			CompareUnits(newIngredient, oldIngredient, newCake.Name)
		}
	}
}

func Compare(oldCakes pkg.Recip, newCakes pkg.Recip) {
	oldNames := FindCakeName(oldCakes)

	CompareCakes(oldCakes, newCakes)
	// details changes in not new and not removed cakes
	for _, newCake := range newCakes.Cakes {
		if slices.Contains(oldNames, newCake.Name) {
			oldCake := FindCakeByName(newCake.Name, oldCakes)
			FindTimeChanges(newCake, oldCake)
			IngredientsChanges(newCake, oldCake)
		}
	}
}
