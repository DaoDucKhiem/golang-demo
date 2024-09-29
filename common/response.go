package common

type Response struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
	Total  interface{} `json:"total,omitempty"`
}
