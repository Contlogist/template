package response

type Base struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error" nullable:"true"`
}
