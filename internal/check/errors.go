package check

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/go-resty/resty/v2"
	"github.com/stellaraf/go-ninjarmm/internal/types"
	"github.com/stellaraf/go-utils"
)

func IsGenericError(token interface{}) bool {
	reflection := reflect.ValueOf(token)
	if reflection.Kind() == reflect.Ptr {
		return reflection.Elem().FieldByName("Error").IsValid()
	}
	return reflection.FieldByName("Error").IsValid()
}

func IsRequestError(data interface{}) bool {

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

func IsApiError(data interface{}) bool {
	reflection := reflect.ValueOf(data)
	if reflection.Kind() == reflect.Ptr {
		return reflection.Elem().FieldByName("ErrorDescription").IsValid()
	}
	return reflection.FieldByName("ErrorDescription").IsValid()
}

func ForError(response *resty.Response) (err error) {
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
			case "error_message", "errorMessage":
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

func GetNinjaRMMError(data interface{}) string {
	if IsRequestError(data) {
		return data.(types.NinjaRMMRequestError).ErrorMessage
	}
	if IsApiError(data) {
		return data.(types.NinjaRMMAPIError).ErrorDescription
	}
	if IsGenericError(data) {
		return data.(types.NinjaRMMBaseError).Error
	}
	return fmt.Sprintf("%s", data)
}
