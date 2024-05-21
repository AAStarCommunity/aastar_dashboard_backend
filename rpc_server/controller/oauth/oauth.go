package oauth

import "aastar_dashboard_back/config"

func Init() {
	githubClientId := config.GetSystemConfigByKey("OAuth.Github.ClientId")
	githubClientSecret := config.GetSystemConfigByKey("OAuth.Github.ClientSecret")
	if githubClientId == "" || githubClientSecret == "" {
		panic("githubClientId or githubClientSecret is empty")
	}
	SetGithubOAuthAppInfo(&githubClientId, &githubClientSecret)

}
