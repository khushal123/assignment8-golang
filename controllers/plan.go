package controllers

import (
	"assignment8/models"
	"assignment8/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var plannerURl string = "https://api.spoonacular.com/mealplanner/generate?timeFrame=%v&targetCalories=%v&diet=%v&apiKey=%v"

var recipeUrl string = "https://api.spoonacular.com/recipes/informationBulk?includeNutrition=true&apiKey=%v&ids=%v"

const SPOON_API_KEY = "7184f398eb9c4c279cc8f535d33e4970"

func Plan(w http.ResponseWriter, r *http.Request) {
	dayOfWeek := r.FormValue("day")
	calories := r.FormValue("calories")
	diet := r.FormValue("diet")
	apiUrl := fmt.Sprintf(plannerURl, dayOfWeek, calories, diet, SPOON_API_KEY)
	resp, err := http.Get(apiUrl)
	if err != nil {
		utils.Response(w, "400", []byte(err.Error()))
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.Response(w, "400", []byte(err.Error()))
		return
	}
	if resp.StatusCode != 200 {
		utils.Response(w, "400", body)
		return
	}
	// fmt.Fprintln(body)
	var meal models.MealGenerate
	json.Unmarshal(body, &meal)

	var IDS []int
	for i := 0; i < len(meal.Meals); i++ {
		fmt.Println(meal.Meals[i])
		IDS = append(IDS, meal.Meals[i].ID)
	}
	ids := utils.ArrayToString(IDS, ",")
	apiUrl = fmt.Sprintf(recipeUrl, SPOON_API_KEY, ids)
	recipeResp, err := http.Get(apiUrl)
	if err != nil {
		utils.Response(w, "400", []byte(err.Error()))
		return
	}
	defer recipeResp.Body.Close()
	body, err = ioutil.ReadAll(recipeResp.Body)
	if err != nil {
		utils.Response(w, "400", []byte(err.Error()))
		return
	}
	if resp.StatusCode != 200 {
		utils.Response(w, "400", body)
		return
	}

	var recipies models.Recipies
	json.Unmarshal(body, &recipies)
	suggestions := make(map[string]string)
	for i := 0; i < len(recipies); i++ {
		nutrition := recipies[i].Nutrition
		nutrients := nutrition.Nutrients
		servings := meal.Meals[i].Servings
		key := fmt.Sprintf("%v %v", "Meal", i)
		suggestions[key] = fmt.Sprintf("%v %v - %v %v", servings, meal.Meals[i].Title, nutrients[0].Amount, nutrients[0].Title)
	}
	finalResponse := models.FinalResponse{Day: dayOfWeek, Suggestions: suggestions}
	response, err := json.Marshal(finalResponse)
	fmt.Println(finalResponse)

	utils.Response(w, "200", response)
}
