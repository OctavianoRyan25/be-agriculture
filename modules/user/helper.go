package user

import (
	"math/rand"
)

func MapUserRequestToUser(userRequest *UserRequest) *User {
	return &User{
		Name:      userRequest.Name,
		Email:     userRequest.Email,
		Password:  userRequest.Password,
		Url_Image: userRequest.Url_Image,
	}
}

func MapUserToResponse(user *User) *UserResponse {
	return &UserResponse{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Is_Active:  user.Is_Active,
		OTP:        user.OTP,
		Url_Image:  user.Url_Image,
		Created_at: user.Created_at,
	}
}

func MapLoginRequestToUser(loginRequest *LoginRequest) *User {
	return &User{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}
}

func RandomOTP() string {
	var code [6]byte
	_, err := rand.Read(code[:])
	if err != nil {
		panic(err) // atau tangani error sesuai kebutuhan
	}
	for i := range code {
		code[i] = (code[i] % 10) + '0' // Ubah ke angka 0-9
	}
	return string(code[:])
}
