package contents

type contentDefault struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	SubTitle    string `json:"sub_tilte"`
	Body        string `json:"body"`
	SubBody     string `json:"sub_body"`
	Pictures	[]Pictures `json:"pictures"`
	Status      bool   `json:"status"`
	CreatedDate string `json:"created_date"`
}

type Content contentDefault {
	User 		string `json:"user"`
	Section		string `json:"section"`
	Category 	string `json:"category"`
}
