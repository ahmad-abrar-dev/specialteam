package special

type Specialteam struct {
	ID      int    `json:"id" form:"id" gorm:"primaryKey"`
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Phone   string `json:"phone" form:"phone"`
	Address string `json:"address" form:"address"`
	Image   string `json:"image" form:"image"`
}

type Response struct {
	ErrorCode int         `json:"error_code" form:"error_code"`
	Message   string      `json:"message" form:"message"`
	Data      interface{} `json:"data"`
}
