package model

//go:generate reform

//reform:account
type Account struct {
	Id       int    `reform:"id,pk"`
	Email    string `reform:"email"`
	Role     string `reform:"role"`
	Password string `reform:"password"`
}
