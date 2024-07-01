package schemas

type HelloResponse struct {
	ClientIP string `json:"client_ip"`
	Location string `json:"location"`
	Greeting string `json:"greeting"`
}

type ErrorResponse struct {
	ErrorMsg string `json:"error_msg"`
}

type LocationResponse struct {
	City string `json:"city"`
}

type WeatherResponse struct {
	Main Main `json:"main"`
}

type Main struct {
	Temp float64 `json:"temp"`
}
