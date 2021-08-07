package services

type GithubService struct {
}

func GetGithubService() *GithubService {
	return &GithubService{}
}
func (gs *GithubService) SayHello() string {
	return "hello world"
}
