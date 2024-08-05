package types

import (
	"encoding/json"
	"fmt"

	"go.stellar.af/go-utils/mmap"
)

type Error struct {
	Message     string         `json:"message"`
	Code        string         `json:"code"`
	Description string         `json:"description"`
	Details     map[string]any `json:"details"`
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) Unwrap() error {
	return fmt.Errorf(e.Message)
}

func (e Error) Is(err error) bool {
	return err.Error() == e.Message
}

func (e *Error) UnmarshalJSON(b []byte) error {
	var data map[string]any
	err := json.Unmarshal(b, &data)
	if err != nil {
		bs := string(b)
		if bs == "" {
			return err
		}
		e.Message = bs
		return nil
	}
	details := make(map[string]any)
	if msg, ok := mmap.AssertValue[string](data, "error"); ok {
		e.Message = msg
	} else if msg, ok := mmap.AssertValue[string](data, "errorMessage"); ok {
		e.Message = msg
	}
	if desc, ok := mmap.AssertValue[string](data, "error_description"); ok {
		e.Description = desc
	}
	if code, ok := mmap.AssertValue[float64](data, "error_code"); ok {
		e.Code = fmt.Sprint(code)
	}
	if code, ok := mmap.AssertValue[string](data, "resultCode"); ok {
		if e.Message == "" {
			e.Message = code
		} else {
			e.Code = code
		}
	}
	if id, ok := mmap.AssertValue[string](data, "incidentId"); ok {
		details["incidentId"] = id
	}
	if errObj, ok := mmap.AssertValue[map[string]any](data, "errorMessage"); ok {
		if code, ok := mmap.AssertValue[string](errObj, "code"); ok {
			if e.Message == "" {
				e.Message = code
			} else {
				e.Code = code
			}
		}
		if params, ok := mmap.AssertValue[map[string]any](errObj, "params"); ok {
			for k, v := range params {
				details[k] = v
			}
		}
	}
	e.Details = details
	if e.Message == "" {
		e.Message = "Unknown Error"
	}
	return nil
}
