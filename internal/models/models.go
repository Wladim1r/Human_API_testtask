package models

// @Description Detailed information about a person
// @Example {"id":1,"name":"Vladimir","surname":"Sokolov","patronymic":"Dmitrievich","age":18,"gender":"male","nationality":"RU"}
type Human struct {
	ID          uint   `json:"id"                   gorm:"primaryKey" example:"1"`
	Name        string `json:"name"                                   example:"Vladimir"`
	Surname     string `json:"surname"                                example:"Sokolov"`
	Patronymic  string `json:"patronymic,omitempty"                   example:"Dmitrievich"`
	Age         uint   `json:"age"                                    example:"18"`
	Gender      string `json:"gender"                                 example:"male"`
	Nationality string `json:"nationality"                            example:"RU"`
}

// @Description Required information for creating a new user
// @Example {"name":"Vladimir","surname":"Sokolov","patronymic":"Dmitrievich"}
type PostRequest struct {
	Name       string `json:"name"                 example:"Vladimir"`
	Surname    string `json:"surname"              example:"Sokolov"`
	Patronymic string `json:"patronymic,omitempty" example:"Dmitrievich"`
}

// @Description Probability of nationality and Country ID for a given name
// @Example {"country_id":"RU","probability":0.7223171}
type CountryProbability struct {
	CountryID   string  `json:"country_id"  example:"RU"`
	Probability float64 `json:"probability" example:"0.7223171"`
}

// @Description Contains nationality prediction results for a given name from external API
// @Example {"count":5,"name":"Vladimir","country":[{"country_id":"RU","probability":0.7223171}]}
type NationalizeResponse struct {
	Count   int                  `json:"count"   example:"5"`
	Name    string               `json:"name"    example:"Vladimir"`
	Country []CountryProbability `json:"country"`
}

// @Description Default error response
// @Example {"error":"User not found"}
type ErrorResponse struct {
	Error string `json:"error" example:"error description"`
}

// @Description Default successfully response
// @Example {"message":"User created successfully"}
type SuccessResponse struct {
	Message string `json:"message" example:"message description"`
}
