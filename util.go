package ninjarmm

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stellaraf/go-utils"
	"github.com/stellaraf/go-utils/environment"
)

func isGenericError(token interface{}) bool {
	reflection := reflect.ValueOf(token)
	if reflection.Kind() == reflect.Ptr {
		return reflection.Elem().FieldByName("Error").IsValid()
	}
	return reflection.FieldByName("Error").IsValid()
}

func isRequestError(data interface{}) bool {

	reflection := reflect.ValueOf(data)
	if reflection.Kind() == reflect.Ptr {
		elem := reflection.Elem()
		resultCode := elem.FieldByName("ResultCode").IsValid()
		errorMessage := elem.FieldByName("ErrorMessage").IsValid()
		incidentID := elem.FieldByName("IncidentID").IsValid()
		return resultCode && errorMessage && incidentID

	}
	resultCode := reflection.FieldByName("ResultCode").IsValid()
	errorMessage := reflection.FieldByName("ErrorMessage").IsValid()
	incidentID := reflection.FieldByName("IncidentID").IsValid()
	return resultCode && errorMessage && incidentID
}

func isApiError(data interface{}) bool {
	reflection := reflect.ValueOf(data)
	if reflection.Kind() == reflect.Ptr {
		return reflection.Elem().FieldByName("ErrorDescription").IsValid()
	}
	return reflection.FieldByName("ErrorDescription").IsValid()
}

func getNinjaRMMError(data interface{}) string {
	if isRequestError(data) {
		return data.(NinjaRMMRequestError).ErrorMessage
	}
	if isApiError(data) {
		return data.(NinjaRMMAPIError).ErrorDescription
	}
	if isGenericError(data) {
		return data.(NinjaRMMBaseError).Error
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

func timeToFractional(t time.Time) float64 {
	f := float64(t.UnixNano()) / 1e9
	return f
}
