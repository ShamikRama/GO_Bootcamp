package pkg

import (
	pkg "Day01/pkg/dbreader"
	"fmt"
	"slices"
)

func GetName(rec pkg.Recip) (names []string) {
	for _, value := range rec.Cakes {
		names = append(names, value.Name)
	}
	return names
}

func FindAddedCakes(oldNames, newNames []string) []string {
	added := make([]string, 0, len(newNames))
	for _, value := range newNames {
		if !slices.Contains(oldNames, value) {
			added = append(added, value)
		}
	}
	return added
}

func FindRemovedCakes(oldNames, newNames []string) []string {
	removed := make([]string, 0, len(oldNames))
	for _, value := range oldNames {
		if !slices.Contains(newNames, value) {
			removed = append(removed, value)
		}
	}
	return removed
}

func CompareCakes(oldcake, newcake pkg.Recip) {
	oldNames := GetName(oldcake)
	newNames := GetName(newcake)

	addedCakes := FindAddedCakes(oldNames, newNames)
	removedCakes := FindRemovedCakes(oldNames, newNames)

	for _, value := range addedCakes {
		fmt.Printf("ADDED cake \"%s\"\n", value)
	}

	for _, value := range removedCakes {
		fmt.Printf("REMOVED cake \"%s\"\n", value)
	}
}

func FindTimeChanges(oldcake pkg.Cake, newcake pkg.Cake) {
	if oldcake.Time != newcake.Time {
		fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", newcake, newcake.Time, oldcake.Time)
	}
}

func FindCakeByName(name string, rec pkg.Recip) (cake pkg.Cake) {
	for _, value := range rec.Cakes {
		if value.Name == name {
			cake = value
		}
	}
	return cake
}
