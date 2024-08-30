package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
	Example of usage:
	go run . ..\dofus-data\translation\fr.json ..\dofus-data\common\Items.d2o.json ..\dofus-data\common\Recipes.d2o.json ../dofus-data/dofus-items.json
*/

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Please provide two file paths as arguments: <translation_file> <items_file> <recipes_file> <output_file>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	itemsFilePath := os.Args[2]
	recipesFilePath := os.Args[3]
	outputFilePath := os.Args[4]

	translation, err := readJSONFile(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	items, err := readItemsJSONFile(itemsFilePath)
	if err != nil {
		fmt.Printf("Error reading objects file: %v\n", err)
		os.Exit(1)
	}

	recipes, err := readRecipesJSONFile(recipesFilePath)
	if err != nil {
		fmt.Printf("Error reading recipes file: %v\n", err)
		os.Exit(1)
	}

	resultItems := createResultItems(items, recipes, translation)

	exportResultItems(outputFilePath, resultItems)
}

type Translation map[int]string

func readJSONFile(filePath string) (Translation, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var translation Translation

	err = json.Unmarshal(fileContent, &translation)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return translation, nil
}

type Item struct {
	ClassType string `json:"ClassType_"`
	ID        int    `json:"id"`
	IconID    int    `json:"iconId"`
	NameID    int    `json:"nameId"`
	RecipeIDs []int  `json:"recipeIds"`
}

type ObjectsJSON struct {
	Objects []json.RawMessage `json:"objects"`
}

func readItemsJSONFile(filePath string) (map[int]Item, error) {

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var objectsJSON ObjectsJSON
	err = json.Unmarshal(fileContent, &objectsJSON)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	result := make(map[int]Item)
	for _, rawObject := range objectsJSON.Objects {
		var item Item
		err = json.Unmarshal(rawObject, &item)
		if err != nil {
			continue
		}

		if item.ClassType == "Weapon" || item.ClassType == "Item" {
			result[item.ID] = item
		}
	}

	return result, nil
}

type Recipe struct {
	IngredientIDs []int `json:"ingredientIds"`
	Quantities    []int `json:"quantities"`
	ResultID      int   `json:"resultId"`
	ResultLevel   int   `json:"resultLevel"`
}

func readRecipesJSONFile(filePath string) (map[int]Recipe, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var recipesJSON ObjectsJSON
	err = json.Unmarshal(fileContent, &recipesJSON)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	recipeMap := make(map[int]Recipe)
	for _, rawObject := range recipesJSON.Objects {
		var recipe Recipe
		err = json.Unmarshal(rawObject, &recipe)
		if err != nil {
			continue
		}

		recipeMap[recipe.ResultID] = recipe
	}

	return recipeMap, nil
}

type Ingredient struct {
	ItemID   int `json:"itemId"`
	Quantity int `json:"quantity"`
}

// output object
type ResultItem struct {
	ItemID int          `json:"id"`
	Name   string       `json:"name"`
	IconID int          `json:"iconId"`
	Recipe []Ingredient `json:"recipe"`
}

type ResultItems = map[int]ResultItem

func createResultItems(items map[int]Item, recipes map[int]Recipe, translation map[int]string) ResultItems {
	resultItems := make(ResultItems)

	for itemID, item := range items {
		resultItems[itemID] = ResultItem{
			ItemID: itemID,
			Name:   translation[item.NameID],
			IconID: item.IconID,
			Recipe: createRecipe(recipes[itemID]),
		}
	}

	return resultItems
}

func createRecipe(recipe Recipe) []Ingredient {
	ingredients := make([]Ingredient, len(recipe.IngredientIDs))

	for i := 0; i < len(recipe.IngredientIDs); i++ {
		ingredients[i] = Ingredient{
			ItemID:   recipe.IngredientIDs[i],
			Quantity: recipe.Quantities[i],
		}
	}

	return ingredients
}

func exportResultItems(outputFilePath string, resultItems ResultItems) error {
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("error creating output file: %v", err)
	}
	defer outputFile.Close()

	encoder := json.NewEncoder(outputFile)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(resultItems)
	if err != nil {
		return fmt.Errorf("error encoding output file: %v", err)
	}

	return nil
}
