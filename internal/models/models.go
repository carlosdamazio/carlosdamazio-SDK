package models

type (
	Book struct {
		ID   string `json:"_id"`
		Name string `json:"name"`
	}

	BookChapter struct {
		ID          string `json:"_id"`
		ChapterName string `json:"chapterName"`
	}

	Movie struct {
		ID                         string  `json:"_id"`
		Name                       string  `json:"name"`
		RuntimeInMinutes           float64 `json:"runtimeInMinutes"`
		BudgetInMillions           float64 `json:"budgetInMillions"`
		BoxOfficeRevenueInMillions float64 `json:"BoxOfficeRevenueInMillions"`
		AcademyAwardNominations    float64 `json:"academyAwardNominations"`
		RottenTomatoesScore        float64 `json:"rottenTomatoesScore"`
	}

	Quote struct {
		ID        string `json:"_id"`
		Dialog    string `json:"dialog"`
		Movie     string `json:"movie"`
		Character string `json:"character"`
	}

	Character struct {
		ID      string `json:"_id"`
		Height  string `json:"height"`
		Race    string `json:"race"`
		Gender  string `json:"gender"`
		Birth   string `json:"birth"`
		Spouse  string `json:"spouse"`
		Death   string `json:"death"`
		Realm   string `json:"realm"`
		Hair    string `json:"hair"`
		Name    string `json:"name"`
		WikiUrl string `json:"wikiUrl"`
	}

	Chapter struct {
		ID          string `json:"_id"`
		ChapterName string `json:"chapterName"`
		Book        string `json:"book"`
	}
)
