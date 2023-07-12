package ninjarmm

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/go-resty/resty/v2"
	"github.com/stellaraf/go-utils"
	"github.com/stellaraf/go-utils/environment"
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

func loadEnv() (env environmentT, err error) {
	err = environment.Load(&env)
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

	if utils.IsString(possibleError) {
		errorDetail = possibleError.(string)
		err = fmt.Errorf("request failed with error '%s'", errorDetail)
		return
	}
	if !utils.IsSlice(possibleError) {
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
