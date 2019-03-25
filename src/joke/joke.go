package joke

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//make sure your Procfile contains "web: main" and located in root app directory
type Joke struct {
	ID   uint32 `json:"id"`
	Joke string `json:"joke"`
}

type JokeResponse struct {
	Type  string `json:"type"`
	Value Joke   `json:"value"`
}

func GetJoke(url string) string {
	c := http.Client{}
	resp, err := c.Get(url)
	if err != nil {
		return "Jokes API not responding"
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	joke := JokeResponse{}

	err = json.Unmarshal(body, &joke)
	if err != nil {
		return "Joke error"
	}
	return joke.Value.Joke
}
