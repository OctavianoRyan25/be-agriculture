package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	googleOAuth "google.golang.org/api/oauth2/v2"
)

var (
	oauthConfigforAndro = &oauth2.Config{
		ClientID:     "355159605133-claot9pia5m9o2b37g1klcb6ai4f784l.apps.googleusercontent.com",
		ClientSecret: "",
		RedirectURL:  "https://be-agriculture-awh2j5ffyq-uc.a.run.app/api/v1/auth-andro/google/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
)

func LoginGoogleforAndro(c echo.Context) error {
	url := oauthConfigforAndro.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func CallbackGoogleforAndro(c echo.Context) error {
	code := c.QueryParam("code")
	token, err := oauthConfigforAndro.Exchange(context.Background(), code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to exchange token: "+err.Error())
	}

	client := oauthConfigforAndro.Client(context.Background(), token)
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
