package response

type NewsCategoryRes struct {
	Name string `json:"name"`
}

func (NewsCategoryRes) TableName() string {
	return "news_categories"
}
