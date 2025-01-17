package request

type ProcListRequest struct {
	BUYNO string `form:"BUYNO" json:"BUYNO" binding:"required"` // Lấy từ form-data hoặc JSON, bắt buộc
	GSBH  string `form:"GSBH" json:"GSBH" binding:"required"`   // Lấy từ form-data hoặc JSON, bắt buộc
}
