package routes

type Holiday struct {
	holiday_id 		int `json:"holiday_id"`
	date 			string `json:"date"`
	remark 			string `json:"remark"`
	created_date 	string `json: "created_date"`
	created_by 		string `json: "created_by"`
}