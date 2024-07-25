package pkg

import (
	pkg "Day01/pkg/dbreader"
	"fmt"
	"slices"
)

func GetName(rec pkg.Recip) (names []string) {
	for _, cake := range rec.Cakes {
		names = append(names, cake.Name)
	}
	return names
}

func FindAddedCakes(oldNames, newNames []string) []string {
	added := make([]string, 0, len(newNames))
	for _, newName := range newNames {
		if !slices.Contains(oldNames, newName) {
			added = append(added, newName)
		}
	}
	return added
}

func FindRemovedCakes(oldNames, newNames []string) []string {
	removed := make([]string, 0, len(oldNames))
	for _, oldName := range oldNames {
		if !slices.Contains(newNames, oldName) {
			removed = append(removed, oldName)
		}
	}
	return removed
}

func CompareCakes(oldcake, newcake pkg.Recip) {
	oldNames := GetName(oldcake)
	newNames := GetName(newcake)

	addedCakes := FindAddedCakes(oldNames, newNames)
	removedCakes := FindRemovedCakes(oldNames, newNames)

	for _, cake := range addedCakes {
		fmt.Printf("ADDED cake \"%s\"\n", cake)
	}

	for _, cake := range removedCakes {
		fmt.Printf("REMOVED cake \"%s\"\n", cake)
	}
}

func FindTimeChanges(oldcake pkg.Cake, newcake pkg.Cake) {
	if oldcake.Time != newcake.Time {
		fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", newcake, newcake.Time, oldcake.Time)
	}
}
