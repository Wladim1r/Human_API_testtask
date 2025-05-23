package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Wladim1r/testtask/internal/models"
)

func IsPositive(num string) (int, error) {
	if num == "" {
		return 0, nil
	}

	numInt, err := strconv.Atoi(num)
	if err != nil {
		return 0, fmt.Errorf("could not convert string into int %w", err)
	}

	if numInt <= 0 {
		return 0, fmt.Errorf("number must be positive")
	}

	return numInt, nil
}

func ParseResponse(url string) (map[string]interface{}, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("could not read URL %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read body request %w", err)
	}

	var result map[string]interface{}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("could not parse into JSON %w", err)
	}

	return result, nil
}

func ParseResponseNationalize(url string) (models.NationalizeResponse, error) {
	response, err := http.Get(url)
	if err != nil {
		return models.NationalizeResponse{}, fmt.Errorf("could not read URL %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return models.NationalizeResponse{}, fmt.Errorf("could not read body request %w", err)
	}

	var result models.NationalizeResponse

	if err := json.Unmarshal(body, &result); err != nil {
		return models.NationalizeResponse{}, fmt.Errorf("could not parse into JSON %w", err)
	}

	return result, nil
}
