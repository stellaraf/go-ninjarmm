package ninjarmm

import (
	"fmt"

	"github.com/stellaraf/go-ninjarmm/internal/check"
	"github.com/stellaraf/go-ninjarmm/internal/types"
)

type QueryClient[T any] struct {
	base      *Client
	batchSize int
}

func NewQueryClient[T any](client *Client, batchSize int) *QueryClient[T] {
	return &QueryClient[T]{
		base:      client,
		batchSize: batchSize,
	}
}

func (qc *QueryClient[T]) Do(path string, query map[string]string) (*QueryResult[T], error) {
	req := qc.base.httpClient.R().
		SetResult(&QueryResult[T]{}).
		SetError(&types.NinjaRMMAPIError{})

	if query != nil {
		req.SetQueryParams(query)
	}

	res, err := req.Get(path)
	if err != nil {
		return nil, err
	}
	err = check.ForError(res)
	if err != nil {
		return nil, err
	}
	data, ok := res.Result().(*QueryResult[T])
	if !ok {
		return nil, fmt.Errorf("failed to parse response: %s", string(res.Body()))
	}
	return data, nil
}
