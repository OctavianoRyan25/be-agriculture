package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/OctavianoRyan25/be-agriculture/utils/helper"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	googleOAuth "google.golang.org/api/oauth2/v2"
)

var (
	oauthConfig = &oauth2.Config{
		ClientID:     "355159605133-e1c0bul0ekhakh25c1d9rtcj9pnsdcej.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "https://be-agriculture-awh2j5ffyq-uc.a.run.app/api/v1/auth/google/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
)

func LoginGoogle(c echo.Context) error {
	url := oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func CallbackGoogle(c echo.Context) error {
	code := c.QueryParam("code")
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to exchange token: "+err.Error())
	}

	client := oauthConfig.Client(context.Background(), token)
	service, err := googleOAuth.New(client)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create oauth service: "+err.Error())
	}

	userinfo, err := service.Userinfo.V2.Me.Get().Do()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to get user info: "+err.Error())
	}

	parse, _ := strconv.Atoi(userinfo.Id)

	jwtToken, err := GenerateJWTToken(uint(parse), userinfo.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to generate JWT token: "+err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": jwtToken,
	})
}

func GenerateJWTToken(userID uint, email string) (string, error) {
	return helper.GenerateToken(userID, email, "user") // Modify as necessary to use the GenerateToken function from your helper package
}
