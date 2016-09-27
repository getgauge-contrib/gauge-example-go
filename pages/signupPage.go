package pages

type SignupPage struct {
	Page Page
}

var (
	SignupPageURL = URL + "signup"

	UsernameID        = "user_username"
	UsernameEmailID   = "user_email"
	UserPasswordID    = "user_password"
	UserPasswordConID = "user_password_confirmation"
	UserCommitName    = "commit"
)
