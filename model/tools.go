package model

const (
	URLErr    = "your subscription or external configuration is invalid"
	GenSubErr = "generate sing-box configuration error"
)

type Response struct {
	Status int    `json:"Status"`
	Error  string `json:"ErrorMsg"`
}
