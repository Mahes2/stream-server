package api

type Request struct {
	Id   string `json:"id"`
	Hash string `json:"hash"`
}

type Response struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}
