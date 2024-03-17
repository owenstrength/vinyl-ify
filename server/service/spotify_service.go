package service

import (
	"Server/config"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func GetUserID(authJSON string) (string, error) {
	// Make GET request to Spotify API

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	if err != nil {
		return "", err
	}

	authBody := strings.Split(authJSON, ",")

	accessToken := authBody[0][14:len(authBody[0])]

	fmt.Println("ACCESS TOKEN: ", accessToken)

	// Set authorization header
	req.Header.Set("Authorization", "Bearer "+accessToken)
	fmt.Println("set header")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	fmt.Println("RESPONSE: ", resp)
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get user ID: %s", resp.Status)
	}

	res, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return "", err
	}
	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println(string(res))
	return string(res), nil

}

func GetAuthToken(code string) (string, error) {
	cfg := config.LoadConfig()

	form := url.Values{}
	form.Add("grant_type", "authorization_code")
	form.Add("code", code)
	form.Add("redirect_uri", "http://localhost:8000/callback")

	fmt.Println("Encoded form data:", form.Encode())

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}

	// Set authorization header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	auth := base64.StdEncoding.EncodeToString([]byte(cfg.SpotifyClientID + ":" + cfg.SpotifyClientSecret))
	req.Header.Set("Authorization", "Basic "+auth)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get user ID: %s", resp.Status)
	}

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetUserTopArtists(authJSON string) (string, error) {
	// Make GET request to Spotify API
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/top/artists?limit=30", nil)
	if err != nil {
		return "", err
	}

	authBody := strings.Split(authJSON, ",")

	accessToken := authBody[0][14:len(authBody[0])]

	// Set authorization header
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get user ID: %s", resp.Status)
	}

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println(string(body))
	return string(body), nil
}
