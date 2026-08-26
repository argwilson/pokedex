package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string, cache Cache) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cacheDat, ok := cache.Get(url)
	fmt.Printf("From Cache: %t\n", ok)
	if ok {
		cacheLocationsResp := RespShallowLocations{}
		err := json.Unmarshal(cacheDat, &cacheLocationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return cacheLocationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}
	cache.Add(url, dat)

	return locationsResp, nil
}
