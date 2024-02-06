package jwt

type Payload struct {
	ID            int `json:"id"`
	UserAccessBit int `json:"uab"`
	CompanyID     int `json:"cid"`
	DateStart     int `json:"nbf"`
	DateEnd       int `json:"exp"`
}
