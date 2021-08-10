package utils

import "github.com/google/go-github/github"

func GetStatusCode(resp *github.Response) int {
	if resp == nil {
		return -1
	}
	return resp.StatusCode
}
