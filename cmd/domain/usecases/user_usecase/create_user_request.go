package user_usecase

type CreateUserRequest struct {
	Name     *string `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}

type GetUserByIDResponse struct {
	ID        string  `json:"id"`
	Name      *string `json:"name"`
	Email     string  `json:"email"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
