package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const (
	apiHerokuConfigCall = "https://api.heroku.com/addons/%s/config"
)

type DBHerokuUrl struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func getHerokuConfigURLByPostgresID() (*pgx.ConnConfig, error) {
	request, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(apiHerokuConfigCall, os.Getenv("HEROKU_POSTGRES_ID")),
		nil,
	)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Accept", "application/vnd.heroku+json; version=3")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("HEROKU_API_KEY")))
	c := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := c.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return parseHerokuGetConnResponse(resp)
}

func parseHerokuGetConnResponse(resp *http.Response) (*pgx.ConnConfig, error) {
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var urlData []DBHerokuUrl
	err = json.Unmarshal(bodyBytes, &urlData)
	if err != nil {
		return nil, err
	}
	if len(urlData) == 0 {
		return nil, errors.New("empty database list")
	}

	return pgx.ParseConfig(urlData[0].Value)
}
