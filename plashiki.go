package plashikigo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// GetAnimeID ...
func GetAnimeID(id string) (map[int]Episode, error) {
	resp, err := http.Get("https://plashiki.su/api/anime/v2/" + id)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	anime := AnimeIDResponse{}
	if err := json.Unmarshal(body, &anime); err != nil {
		return nil, err
	}
	if anime.OK {
		return anime.Result, nil
	}
	return nil, errors.New("Request error: " + anime.Reason)
}

// GetAnimeEpisode ...
func GetAnimeEpisode(id string, num string) (map[int]Episode, error) {
	resp, err := http.Get("https://plashiki.su/api/anime/v2/" + id + "/episodes/" + num)
	if err != nil {
		return map[int]Episode{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return map[int]Episode{}, err
	}
	anime := AnimeEpisodeResponse{}
	if err := json.Unmarshal(body, &anime); err != nil {
		return map[int]Episode{}, err
	}
	if anime.OK {
		return anime.Result, nil
	}
	return map[int]Episode{}, errors.New("Request error: " + anime.Reason)
}

// GetAnimeEpisodesNumber ...
func GetAnimeEpisodesNumber(id string) ([]int, error) {
	resp, err := http.Get("https://plashiki.su/api/anime/v2/" + id + "/episodes")
	if err != nil {
		return []int{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []int{}, err
	}
	anime := AnimeEpisodesResponse{}
	if err := json.Unmarshal(body, &anime); err != nil {
		return []int{}, err
	}
	if anime.OK {
		return anime.Result, nil
	}
	return []int{}, errors.New("Request error: " + anime.Reason)
}

// GetAnimeQuery ...
func GetAnimeQuery(args map[string]string) ([]Translation, error) {
	baseURL, err := url.Parse("https://plashiki.su")
	if err != nil {
		return nil, err
	}
	baseURL.Path += "/api/anime/query"
	params := url.Values{}
	if val, ok := args["anime"]; ok {
		params.Add("anime", val)
	}
	if val, ok := args["searchLimit"]; ok {
		params.Add("searchLimit", val)
	}
	if val, ok := args["episode"]; ok {
		params.Add("episode", val)
	}
	if val, ok := args["kind"]; ok {
		params.Add("kind", val)
	}
	if val, ok := args["lang"]; ok {
		params.Add("lang", val)
	}
	if val, ok := args["quality"]; ok {
		params.Add("quality", val)
	}
	if val, ok := args["author"]; ok {
		params.Add("author", val)
	}
	if val, ok := args["uploader"]; ok {
		params.Add("uploader", val)
	}
	if val, ok := args["player"]; ok {
		params.Add("player", val)
	}
	baseURL.RawQuery = params.Encode()
	resp, err := http.Get(baseURL.String())
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	anime := AnimeQueryResponse{}
	if err := json.Unmarshal(body, &anime); err != nil {
		return nil, err
	}
	if anime.OK {
		return anime.Result, nil
	}
	return nil, errors.New("Request error: " + anime.Reason)
}
