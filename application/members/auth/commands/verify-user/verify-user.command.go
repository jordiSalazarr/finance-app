package auth_useCases

type VerifyUserCommand struct {
	Email string
	Code  string
}
