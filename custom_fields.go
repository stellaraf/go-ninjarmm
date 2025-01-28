package ninjarmm

import "encoding/json"

// ParseCustomFields parses a custom fields map to a predefined struct type.
func ParseCustomFields[T any](cf map[string]any) (*T, error) {
	b, err := json.Marshal(&cf)
	if err != nil {
		return nil, err
	}
	var t *T
	err = json.Unmarshal(b, &t)
	if err != nil {
		return nil, err
	}
	return t, nil
}
