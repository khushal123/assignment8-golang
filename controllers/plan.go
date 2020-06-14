package controllers

import (
	"assignment8/models"
	"assignment8/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var url string = "https://api.spoonacular.com/mealplanner/generate?timeFrame=%v&targetCalories=%v&diet=%v&apiKey=%v"

const SPOON_API_KEY = "7184f398eb9c4c279cc8f535d33e4970"

func Plan(w http.ResponseWriter, r *http.Request) {
	dayOfWeek := r.FormValue("dayOfWeek")
	calories := r.FormValue("calories")
	diet := r.FormValue("diet")

	apiUrl := fmt.Sprintf(url, dayOfWeek, calories, diet, SPOON_API_KEY)
	fmt.Println(apiUrl)
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
	var mealGenerate models.MealGenerate
	json.Unmarshal(body, &mealGenerate)
	fmt.Printf("API Response as struct %+v\n", mealGenerate)

	utils.Response(w, "200", body)
}
