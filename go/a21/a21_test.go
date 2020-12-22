package a21_test

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

type dish struct {
	ingredients []string
	allergens   []string
}

func intersect(m1 ingredientSet, m2 ingredientSet) ingredientSet {
	res := make(map[string]struct{}, len(m1))
	for k1 := range m1 {
		if _, exists := m2[k1]; exists {
			res[k1] = struct{}{}
		}
	}
	return res
}

type ingredientSet map[string]struct{}

func Test_day21(t *testing.T) {
	f, err := os.Open("input")
	check(err)
	dishes := parseDishes(f)
	uniqIngredients := make(map[string]struct{})
	allergenIngredients := make(map[string]ingredientSet)
	ingredientCount := make(map[string]int)
	// For each dish
	for _, dish := range dishes {
		dishIngredients := make(ingredientSet)
		for _, ingredient := range dish.ingredients {
			ingredientCount[ingredient]++
			dishIngredients[ingredient] = struct{}{}
			uniqIngredients[ingredient] = struct{}{}
		}

		// For each allergen in the dish
		for _, allergen := range dish.allergens {
			if _, exists := allergenIngredients[allergen]; !exists {
				allergenIngredients[allergen] = dishIngredients
				continue
			}
			// If ingredients has already been added for this allergen,
			// store the intersection (shared elements) between the previous
			// and the current
			allergenIngredients[allergen] = intersect(allergenIngredients[allergen], dishIngredients)
		}
	}

	safeIngredients := make(map[string]struct{})
	for k := range uniqIngredients {
		safeIngredients[k] = struct{}{}
	}
	for _, ingredients := range allergenIngredients {
		for k := range ingredients {
			delete(safeIngredients, k)
		}
	}
	// fmt.Println(ingredientCount)
	var totalCount int
	for k := range safeIngredients {
		totalCount += ingredientCount[k]
	}
	// Part 1
	// fmt.Println(totalCount)

	// Part 2
	// fmt.Printf("%+v\n", allergenIngredients)
	result := deduceAllergens(allergenIngredients)
	fmt.Println(result)
	allergenList := make([]string, 0)
	for k := range result {
		allergenList = append(allergenList, k)
	}
	sort.Strings(allergenList)
	finalIngredientsList := make([]string, 0)
	for _, allergen := range allergenList {
		finalIngredientsList = append(finalIngredientsList, result[allergen])
	}
	fmt.Println(strings.Join(finalIngredientsList, ","))
	require.Equal(t, "hn,dgsdtj,kpksf,sjcvsr,bstzgn,kmmqmv,vkdxfj,bsfqgb", strings.Join(finalIngredientsList, ","))
}

func printMap(in map[string]ingredientSet) {
	for allergen, ingredients := range in {
		ingredientsList := make([]string, 0, len(ingredients))
		for ingredient := range ingredients {
			ingredientsList = append(ingredientsList, ingredient)
		}
		fmt.Printf("%v: { %v }\n", allergen, strings.Join(ingredientsList, ", "))
	}
}

func deduceAllergens(allergenIngredients map[string]ingredientSet) map[string]string {
	ingredientTaken := make(map[string]bool)

	// copy input
	inputCpy := make(map[string]ingredientSet)
	for k, v := range allergenIngredients {
		inputCpy[k] = make(ingredientSet)
		for kk := range v {
			inputCpy[k][kk] = struct{}{}
		}
	}
	allergenIngredients = inputCpy

	allergenIngredient := make(map[string]string)

	for {
		allOnes := true
		for allergen := range allergenIngredients {
			// If the allergen only has one ingredient, it must be the result
			if len(allergenIngredients[allergen]) == 1 {
				var ingredient string
				for ingredient = range allergenIngredients[allergen] {
				}

				// We've already added this one to the right place, continue
				if allergenIngredient[allergen] == ingredient {
					continue
				}

				// Check if ingredient has been taken
				if ingredientTaken[ingredient] {
					log.Fatalf("Ingredient %v was the only ingredient for allergen %v and it was already taken\n", ingredient, allergen)
				}

				ingredientTaken[ingredient] = true
				allergenIngredient[allergen] = ingredient

				// Remove the ingredient from all other allergens
				for otherAllergen := range allergenIngredients {
					if otherAllergen == allergen {
						continue
					}
					delete(allergenIngredients[otherAllergen], ingredient)
				}

				continue
			}
			allOnes = false
		}
		if allOnes {
			return allergenIngredient
		}
	}
}

func parseDishes(f io.Reader) []dish {
	dishes := make([]dish, 0)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		row := sc.Text()
		rowparts := strings.Split(row, "(")
		ingredients := strings.Split(strings.Trim(rowparts[0], " "), " ")
		allergens := strings.Split(
			strings.TrimRight(
				strings.TrimPrefix(rowparts[1], "contains "),
				")",
			),
			", ",
		)
		dishes = append(dishes, dish{
			ingredients: ingredients,
			allergens:   allergens,
		})
	}
	return dishes
}
