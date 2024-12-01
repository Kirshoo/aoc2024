package utils

import (
	"net/http"
	"fmt"
	"aoc24/pkg/config"
)

var (
	cookieVariable = "SESSION_COOKIE"
)

type Input struct {
	Year int
	Day int
}

func NewInput(year, day int) *Input {
	return &Input{Year: year, Day: day}
}

func (in *Input) Get() (*http.Response, error) {
	// Get session cookie stored in .env. It will be used in a headers
	cookie, ok := config.Look(cookieVariable)
	if !ok {
		return nil, fmt.Errorf("can't find %s inside pkg/config/.env", cookieVariable)
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", in.Year, in.Day)
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("error creating a request:")
		return nil, err
	}

	client := &http.Client{}
	// Very important to specify, that its a session and not just random string! 
	req.Header.Add("Cookie", fmt.Sprintf("session=%s", cookie))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error getting an input:")
		return nil, err
	}

	return res, nil
}
