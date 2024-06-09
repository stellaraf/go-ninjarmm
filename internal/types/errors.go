package types

import (
	"fmt"
	"strings"
)

type NinjaRMMBaseError struct {
	Error string `json:"error"`
}

type NinjaRMMAPIError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorCode        int    `json:"error_code,omitempty"`
}

type NinjaRMMRequestError struct {
	ResultCode   string `json:"resultCode"`
	ErrorMessage string `json:"errorMessage"`
	IncidentID   string `json:"incidentId"`
}

type NinjaRMMPutError struct {
	ErrorMessage struct {
		Code   string         `json:"code"`
		Params map[string]any `json:"params"`
	} `json:"errorMessage"`
}

func (e *NinjaRMMPutError) GetErrorMessage(deviceID int) string {
	paramsPairs := []string{}
	for k, v := range e.ErrorMessage.Params {
		param := fmt.Sprintf("%s: %v", k, v)
		paramsPairs = append(paramsPairs, param)
	}
	params := strings.Join(paramsPairs, ", ")
	msg := strings.ReplaceAll(e.ErrorMessage.Code, "_", "")
	errMsg := fmt.Sprintf("failed to set maintenance for device '%d' due to error '%s'", deviceID, msg)
	if len(params) > 0 {
		errMsg += fmt.Sprintf(" (details: %s)", params)
	}
	return errMsg
}
