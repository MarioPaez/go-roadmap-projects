package service

import (
	"encoding/json"
	"fmt"
	"github-activity/model"
	"github-activity/utils"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	URL_PATTERN = "https://api.github.com/users/<username>/events"
	REPLACE     = "<username>"
)

type GithubService interface {
	GetGithubActivity(username string) error
}

type githubService struct {
}

func NewGithubService() GithubService {
	return &githubService{}
}

func (s *githubService) GetGithubActivity(username string) error {
	url := strings.Replace(URL_PATTERN, REPLACE, username, 1)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	manageResponse(resp, username)
	return nil
}

func manageResponse(resp *http.Response, username string) error {

	code := resp.StatusCode
	switch code {
	case http.StatusOK:
		return readResponse(resp)
	case http.StatusNotFound:
		log.Printf("the username provided %q does not exist.", username)
	default:
		fmt.Printf("Status code: %d", code)
	}
	return nil
}

func readResponse(resp *http.Response) error {
	defer resp.Body.Close()
	info, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("an error occur trying to read the response.")
		return err
	}
	return getEvents(info)
}

func getEvents(info []byte) error {
	var events []model.Event
	if err := json.Unmarshal(info, &events); err != nil {
		log.Println("an error occur trying to unmarshal the content of the response.")
		return err
	}
	utils.PrettyPrint(events)
	return nil
}
