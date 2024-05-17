package oauth

import "aastar_dashboard_back/config"

func Init() {
	githubClientId := config.GetSystemConfigByKey("OAuth.Github.ClientId")
	githubClientSecret := config.GetSystemConfigByKey("OAuth.Github.ClientSecret")
	SetGithubOAuthAppInfo(&githubClientId, &githubClientSecret)

}
