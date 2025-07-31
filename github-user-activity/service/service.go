package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github-activity/model"
	"github-activity/utils"
	"io"
	"net/http"
	"strings"
)

const (
	URL_PATTERN = "https://api.github.com/users/<username>/events"
	REPLACE     = "<username>"
)

func GetGithubActivity(username string) error {
	url := strings.Replace(URL_PATTERN, REPLACE, username, 1) //Interesante entender como lo hace por detrás

	resp, err := http.Get(url)
	if err != nil {
		return errors.New(err.Error()) //TODO: Handle errors by us
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		info, err := io.ReadAll(resp.Body)
		if err != nil {
			return err //Cual es la diferencia con  errors.New(err.Error())?
		}
		var events []model.Event
		fmt.Println("Se ha realizado correctamente la petición!")
		if err := json.Unmarshal(info, &events); err != nil {
			return err
		}
		utils.PrettyPrint(events)
	}
	return nil
}

// Output:
// - Pushed 3 commits to kamranahmedse/developer-roadmap
// - Opened a new issue in kamranahmedse/developer-roadmap
// - Starred kamranahmedse/developer-roadmap
