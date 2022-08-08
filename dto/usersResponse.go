package dto

import "invertory/domain"

type UsersResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Token string `json:"token"`
}

func FormatUsersResponse(users domain.Users, token string) UsersResponse {
	UsersResponse := UsersResponse{
		ID:    users.ID,
		Email: users.Email,
		Role:  users.Role,
		Token: token,
	}
	return UsersResponse
}
