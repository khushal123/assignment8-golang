package models

type MealGenerate struct {
	Meals []struct {
		ID             int    `json:"id"`
		ImageType      string `json:"imageType"`
		Title          string `json:"title"`
		ReadyInMinutes int    `json:"readyInMinutes"`
		Servings       int    `json:"servings"`
		SourceURL      string `json:"sourceUrl"`
	} `json:"meals"`
	Nutrients struct {
		Calories      float64 `json:"calories"`
		Protein       float64 `json:"protein"`
		Fat           float64 `json:"fat"`
		Carbohydrates float64 `json:"carbohydrates"`
	} `json:"nutrients"`
}



type Recipe struct {
	Nutrition struct {
		Nutrients []struct {
			Title               string  `json:"title"`
			Amount              float64 `json:"amount"`
			Unit                string  `json:"unit"`
			PercentOfDailyNeeds float64 `json:"percentOfDailyNeeds"`
		} `json:"nutrients"`
	}`json:"nutrition"`
}
