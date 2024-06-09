package ninjarmm

import (
	"context"
	"fmt"

	"github.com/sourcegraph/conc/pool"
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

func (qc *QueryClient[T]) Do(path string, query map[string]string) ([]T, error) {
	req1 := qc.base.httpClient.R().
		SetQueryParams(query).
		SetResult(&QueryResult[T]{}).
		SetError(&types.NinjaRMMAPIError{})

	res1, err := req1.Get(path)
	if err != nil {
		return nil, err
	}
	err = check.ForError(res1)
	if err != nil {
		return nil, err
	}
	data1, ok := res1.Result().(*QueryResult[T])
	if !ok {
		return nil, fmt.Errorf("failed to parse response: %s", string(res1.Body()))
	}
	final := make([]T, 0)
	errs := make([]error, 0)
	if data1.Cursor != nil {
		batchCount := float64(data1.Cursor.Count / int32(qc.batchSize))
		p := pool.NewWithResults[[]T]().WithContext(context.Background())
		for i := 0; i <= int(batchCount); i++ {
			p.Go(func(ctx context.Context) ([]T, error) {
				q := query
				q["cursor"] = data1.Cursor.Name
				q["pageSize"] = fmt.Sprint(qc.batchSize)
				req := qc.base.httpClient.R().
					SetResult(&QueryResult[T]{}).
					SetError(&types.NinjaRMMAPIError{}).
					SetQueryParams(q)
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
					return nil, fmt.Errorf("failed to parse response: %s", string(res1.Body()))
				}
				return data.Results, nil
			})
		}
		results, err := p.Wait()
		if err != nil {
			return nil, err
		}
		for _, result := range results {
			final = append(final, result...)
		}
	} else {
		final = append(final, data1.Results...)
	}
	for _, err := range errs {
		return nil, err
	}
	return final, nil
}
