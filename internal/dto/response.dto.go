package dto

// type Response map[string]any //agar deskriptif ,response ctx.JSON , bisa pakai gin.H

type Response struct {
	Success bool
	Data    any
	Msg     string
}
