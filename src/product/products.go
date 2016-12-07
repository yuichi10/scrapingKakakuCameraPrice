package product

type Product struct {
	Category           string `json:"category"`              //カテゴリー
	Maker              string `json:"maker"`                 //maker
	Name               string `json:"name"`                  //製品名
	ReleaseDate        string `json:"release_date"`          //発売日
	SecondHandMinPrice string `json:"second_hand_min_price"` //中古最低価格
	SecondHandMaxPrice string `json:"second_hand_max_price"` //中古最高価格
	LowestNewPrice     string `json:"lowest_new_price"`      //新品最低価格
}

type DslrCamera struct {
	Category           string `json:"category"`              //カテゴリー
	Maker              string `json:"maker"`                 //maker
	Name               string `json:"name"`                  //製品名
	ReleaseDate        string `json:"release_date"`          //発売日
	SecondHandMinPrice string `json:"second_hand_min_price"` //中古最低価格
	SecondHandMaxPrice string `json:"second_hand_max_price"` //中古最高価格
	LowestNewPrice     string `json:"lowest_new_price"`      //新品最低価格
}

type Lens struct {
	Category           string `json:"category"`              //カテゴリー
	Maker              string `json:"maker"`                 //maker
	Name               string `json:"name"`                  //製品名
	ReleaseDate        string `json:"release_date"`          //発売日
	SecondHandMinPrice string `json:"second_hand_min_price"` //中古最低価格
	SecondHandMaxPrice string `json:"second_hand_max_price"` //中古最高価格
	LowestNewPrice     string `json:"lowest_new_price"`      //新品最低価格
}

type VideoCamera struct {
	Category           string `json:"category"`              //カテゴリー
	Maker              string `json:"maker"`                 //maker
	Name               string `json:"name"`                  //製品名
	ReleaseDate        string `json:"release_date"`          //発売日
	SecondHandMinPrice string `json:"second_hand_min_price"` //中古最低価格
	SecondHandMaxPrice string `json:"second_hand_max_price"` //中古最高価格
	LowestNewPrice     string `json:"lowest_new_price"`      //新品最低価格
}
