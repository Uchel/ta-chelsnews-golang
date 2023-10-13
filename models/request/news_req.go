package request

type NewsReq struct {
	ID           uint
	Image        string `json:"image"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	CategoryNews string `json:"category_news"`
}
