package lib


type GithubClientInterface interface {
	GetRepository(owner string, name string) error
}

type GithubClient struct {}

func (githubClient *GithubClient) GetRepository(owner string, name string) error {
	return nil
}
