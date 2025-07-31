package service

import (
	"fmt"
	"strings"
)

const (
	URL_PATTERN = "https://api.github.com/users/<username>/events"
	REPLACE     = "<username>"
)

func GetGithubActivity(username string) error {
	url := strings.Replace(URL_PATTERN, REPLACE, username, 1) //Interesante entender como lo hace por detrás
	fmt.Println(url)

	return nil
}
