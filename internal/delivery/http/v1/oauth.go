package v1

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

var (
	oauthConfig  *oauth2.Config
)

// state: random1111111111
// code: 4/0AY0e-g........
func getUserInfo(state, code, urlForGettingUserInfo string) ([]byte, error) {
	if state != viper.GetString("oauth.state") {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := oauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}
	// token.AccessToken: a0AfH6.......
	response, err := http.Get(urlForGettingUserInfo + url.QueryEscape(token.AccessToken))
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}

func oAuthCallbackHandler(c echo.Context, urlForGettingUserInfo string) error {
	content, err := getUserInfo(
		c.FormValue("state"),
		c.FormValue("code"),
		urlForGettingUserInfo,
	)
	if err != nil {
		fmt.Println(err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	//fmt.Fprintf(w, "Content: %s\n", content)
	return c.String(http.StatusOK, string(content))
}

func (h *Handler) oAuthStart(c echo.Context) error {
	var htmlIndex = `<html>
<body>
	<a href="/api/v1/oauth/google">Google Log In</a> <br/>
	<a href="/api/v1/oauth/fb">Facebook Log In</a> <br/>
</body>
</html>`

	return c.HTML(http.StatusOK, htmlIndex)
}


func (h *Handler) fBLogin(c echo.Context) error {
	hostAndPort := viper.GetString("http.protocol") + "://" + viper.GetString("http.host") + ":" +
		viper.GetString("http.port")
	oauthConfig = &oauth2.Config{
		RedirectURL:  hostAndPort + "/api/v1/oauth/fb/callback",
		ClientID:     os.Getenv("FB_CLIENT_ID"),
		ClientSecret: os.Getenv("FB_CLIENT_SECRET"),
		Scopes:       []string{"public_profile", "email"},
		Endpoint:     facebook.Endpoint,
	}

	urlToRedirect := oauthConfig.AuthCodeURL(viper.GetString("oauth.state"))
	return c.Redirect(http.StatusTemporaryRedirect, urlToRedirect)
}

func (h *Handler) fBLoginCallback(c echo.Context) error {
	return oAuthCallbackHandler(
		c,
		"https://graph.facebook.com/me?access_token=",
	)
}

func (h *Handler) googleLogin(c echo.Context) error {
	hostAndPost := viper.GetString("http.protocol") + "://" + viper.GetString("http.host") + ":" +
		viper.GetString("http.port")
	oauthConfig = &oauth2.Config{
		RedirectURL:  hostAndPost + "/api/v1/oauth/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	redirectUrl := oauthConfig.AuthCodeURL(viper.GetString("oauth.state"))
	// redirectUrl example: https://accounts.google.com/o/oauth2/auth?client_id=5555555&state=random11111
	return c.Redirect(http.StatusTemporaryRedirect, redirectUrl)
}

func (h *Handler) googleLoginCallback(c echo.Context) error {
	return oAuthCallbackHandler(
		c,
		"https://www.googleapis.com/oauth2/v2/userinfo?access_token=",
	)
}