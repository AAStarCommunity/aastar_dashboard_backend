package oauth

import (
	"aastar_dashboard_back/model"
	"aastar_dashboard_back/repository"
	"aastar_dashboard_back/rpc_server/middlewares"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	clientId     string
	clientSecret string
)

func SetGithubOAuthAppInfo(clientIdPtr, clientSecretPrt *string) {
	clientId = *clientIdPtr
	clientSecret = *clientSecretPrt
}

type githubAccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

/*
	{
	  "login": "fanhousanbu",
	  "id": 6702542,
	  "node_id": "MDQ6VXNlcjY3MDI1NDI=",
	  "avatar_url": "https://avatars.githubusercontent.com/u/6702542?v=4",
	  "gravatar_id": "",
	  "url": "https://api.github.com/users/fanhousanbu",
	  "html_url": "https://github.com/fanhousanbu",
	  "followers_url": "https://api.github.com/users/fanhousanbu/followers",
	  "following_url": "https://api.github.com/users/fanhousanbu/following{/other_user}",
	  "gists_url": "https://api.github.com/users/fanhousanbu/gists{/gist_id}",
	  "starred_url": "https://api.github.com/users/fanhousanbu/starred{/owner}{/repo}",
	  "subscriptions_url": "https://api.github.com/users/fanhousanbu/subscriptions",
	  "organizations_url": "https://api.github.com/users/fanhousanbu/orgs",
	  "repos_url": "https://api.github.com/users/fanhousanbu/repos",
	  "events_url": "https://api.github.com/users/fanhousanbu/events{/privacy}",
	  "received_events_url": "https://api.github.com/users/fanhousanbu/received_events",
	  "type": "User",
	  "site_admin": false,
	  "name": "徐超越",
	  "company": null,
	  "blog": "https://www.gotocloud.vip",
	  "location": "杭州",
	  "email": "go-to-cloud@qq.com",
	  "hireable": null,
	  "bio": null,
	  "twitter_username": null,
	  "public_repos": 17,
	  "public_gists": 0,
	  "followers": 10,
	  "following": 4,
	  "created_at": "2014-02-17T08:06:15Z",
	  "updated_at": "2024-05-10T05:15:11Z"
	}
*/
type githubUserInfo struct {
	Id        int    `json:"id"`
	Login     string `json:"login"`
	AvatarUrl string `json:"avatar_url"`
	Url       string `json:"url"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}

func getGithubAccessToken(token string) (*githubAccessToken, error) {
	req, err := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token?"+
			"client_id="+clientId+"&"+
			"client_secret="+clientSecret+"&"+
			"code="+token,
		nil)
	if err != nil {
		fmt.Println("GithubAccessToken Request failed:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Post Failed:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("StatusCode:", resp.StatusCode)
		return nil, err
	}
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll error:", err)
		return nil, err
	}

	ret := &githubAccessToken{}
	items := strings.Split(string(responseBody), "&")
	accessToken := strings.Split(items[0], "=")[1]
	scope := strings.Split(string(items[1]), "=")[1]
	tokenType := strings.Split(string(items[2]), "=")[1]
	ret.AccessToken = accessToken
	ret.Scope = scope
	ret.TokenType = tokenType
	return ret, err
}

func getGithubUserInfo(accessToken *githubAccessToken) (*githubUserInfo, error) {
	req, err := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil)
	if err != nil {
		fmt.Println("GithubAccessToken Request failed:", err)
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", accessToken.TokenType+" "+accessToken.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Get Failed:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("StatusCode:", resp.StatusCode)
		return nil, err
	}
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll error:", err)
		return nil, err
	}

	githubUserInfo := &githubUserInfo{}
	logrus.Debug("GithubUserInfo:", string(responseBody))
	err = json.Unmarshal(responseBody, githubUserInfo)
	return githubUserInfo, err
}

// GithubOAuthLogin
// @Tags OAuth
// @Description Github OAuth Login
// @Accept json
// @Product json
// @Param code query string true "Github OAuth Code"
// @Router /oauth/github [get]
// @Success 200
func GithubOAuthLogin(ctx *gin.Context) {
	token := ctx.Query("code")

	if accessToken, err := getGithubAccessToken(token); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	} else {
		if githubUser, err := getGithubUserInfo(accessToken); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		} else {
			user, err := repository.FindUserByGitHubId(githubUser.Id)
			if err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					ctx.JSON(http.StatusBadRequest, err)
					return
				}
				user = &model.User{}
				//New User By GitHub
				user.GithubId = githubUser.Id
				user.GithubAvatarUrl = githubUser.AvatarUrl
				user.GithubLogin = githubUser.Login
				user.GithubName = githubUser.Name
				user.Email = githubUser.Email
				user.UserId = uuid.New().String()
				err := repository.CreateUser(user)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, err)
					return
				}
			} else {
				if user.GithubName != githubUser.Name {
					user.GithubName = githubUser.Name
				}
				if user.GithubAvatarUrl != githubUser.AvatarUrl {
					user.GithubAvatarUrl = githubUser.AvatarUrl
				}
			}

			_ = repository.UpdateUserLatestLoginTime(user)
			ctx.Set("user_id", user.UserId)
			middlewares.GinJwtMiddleware().LoginHandler(ctx)

			// https://github.com/AAStarCommunity/EthPaymaster_BackService/blob/cedeb46d0cac7dae88ba52117f6fb057e37ad217/rpc_server/api/auth.go#L17
		}
	}
}
func GithubOAuthBind(ctx *gin.Context) {

}
