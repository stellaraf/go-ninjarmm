package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

func isGenericError(token interface{}) bool {
	reflection := reflect.ValueOf(token)
	field := reflection.FieldByName("Error")
	return field.IsValid()
}

func isRequestError(data interface{}) bool {
	reflection := reflect.ValueOf(data)
	resultCode := reflection.FieldByName("ResultCode")
	errorMessage := reflection.FieldByName("ErrorMessage")
	incidentID := reflection.FieldByName("IncidentID")
	return resultCode.IsValid() && errorMessage.IsValid() && incidentID.IsValid()
}

func isApiError(data interface{}) bool {
	reflection := reflect.ValueOf(data)
	errorDescription := reflection.FieldByName("ErrorDescription")
	return errorDescription.IsValid()
}

func getNinjaRMMError(data interface{}) string {
	if isRequestError(data) {
		return data.(ninjaRMMRequestError).ErrorMessage
	}
	if isApiError(data) {
		return data.(ninaRMMAPIError).ErrorDescription
	}
	if isGenericError(data) {
		return data.(ninjaRMMBaseError).Error
	}
	return fmt.Sprintf("%s", data)
}

func isArray(in interface{}) bool {
	return reflect.TypeOf(in).Kind() == reflect.Slice
}

func isString(in interface{}) bool {
	return reflect.TypeOf(in).Kind() == reflect.String
}

func arrayContains[T comparable](arr []T, item T) bool {
	for _, element := range arr {
		if element == item {
			return true
		}
	}
	return false
}

func LoadEnv() (env Environment, err error) {
	err = godotenv.Load()
	if err != nil {
		return
	}
	clientId := os.Getenv("NINJARMM_CLIENT_ID")
	clientSecret := os.Getenv("NINJARMM_CLIENT_SECRET")
	encryptionPassphrase := os.Getenv("NINJARMM_ENCRYPTION_PASSPHRASE")
	baseURL := os.Getenv("NINJARMM_BASE_URL")
	env = Environment{
		ClientID:             clientId,
		ClientSecret:         clientSecret,
		EncryptionPassphrase: encryptionPassphrase,
		BaseURL:              baseURL,
	}
	return
}

func checkForError(response *resty.Response) (err error) {
	var possibleError interface{}
	body := response.Body()
	err = json.Unmarshal(body, &possibleError)
	if err != nil {
		return
	}
	var errorDetail interface{} = "unknown"

	if isString(possibleError) {
		errorDetail = possibleError.(string)
		err = fmt.Errorf("request failed with error '%s'", errorDetail)
		return
	}
	if !isArray(possibleError) {
		data := possibleError.(map[string]interface{})

	loop:
		for key := range data {
			switch key {
			case "error_description":
				errorDetail = data[key]
				break loop
			case "error_message":
				errorDetail = data[key]
				break loop
			case "error":
				errorDetail = data[key]
				break loop
			}
		}
	}
	if errorDetail == "unknown" {
		return nil
	}
	err = fmt.Errorf("request failed with %d error '%s'", response.StatusCode(), errorDetail)
	return
}
