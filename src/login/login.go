package login

// Login data struct that represents the user credentials.
type Login struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}
