package dto

type User struct {
	// key datatype struct_tag = untuk memudahkan aplikasi membaca data slelain json bisa tipe yg lain example 'form'
	Name string `json:"nama"`
	Age  int8   `json:"umur"`
}

type Account struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
