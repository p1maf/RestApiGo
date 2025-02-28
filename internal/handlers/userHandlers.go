package handlers

import (
	"context"
	"github.com/your-username/RestApiGo/internal/userService"
	"github.com/your-username/RestApiGo/internal/web/users"
)

type UserHandler struct {
	Service *userService.UserService
}

func (u UserHandler) GetUsers(_ context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			ID:        usr.ID,
			Email:     usr.Email,
			CreatedAt: usr.CreatedAt,
			UpdatedAt: usr.UpdatedAt,
		}
		response = append(response, user)
	}

	return response, nil
}

func (u UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	user, err := u.Service.CreateUser(*request.Body)
	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

func (u UserHandler) DeleteUsersId(_ context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	err := u.Service.DeleteUserById(uint(request.Id))
	if err != nil {
		return nil, err
	}
	return users.DeleteUsersId204Response{}, err
}

func (u UserHandler) PatchUsersId(_ context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	userToUpdate := users.UpdateUser{
		Email:    request.Body.Email,
		Password: request.Body.Password,
	}

	updatedUser, err := u.Service.UpdateUserById(uint(request.Id), userToUpdate)

	if err != nil {
		return nil, err
	}

	response := users.PatchUsersId200JSONResponse{
		ID:        updatedUser.ID,
		UpdatedAt: updatedUser.UpdatedAt,
		Email:     updatedUser.Email,
		Password:  updatedUser.Password,
	}
	return response, nil
}

func UserNewHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}
