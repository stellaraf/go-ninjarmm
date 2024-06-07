package ninjarmm

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

type statement [3]string

func (smt statement) String() string {
	return fmt.Sprintf("%s %s %s", smt[0], smt[1], smt[2])
}

type deviceFilter struct {
	items   []statement
	offline bool
	online  bool
	group   int
}

type Operator string

type ApprovalStatus string

func (s ApprovalStatus) String() string {
	return string(s)
}

func (op Operator) String() string {
	return string(op)
}

const EQ Operator = "eq"
const NEQ Operator = "neq"
const IN Operator = "in"
const NOTIN Operator = "notin"
const BEFORE Operator = "before"
const AFTER Operator = "after"
const and Operator = " AND "

const APPROVED ApprovalStatus = "APPROVED"
const PENDING ApprovalStatus = "PENDING"

func NewDeviceFilter() *deviceFilter {
	return &deviceFilter{items: []statement{}}
}

func stringifyValue(value []any) string {
	if len(value) == 1 {
		return fmt.Sprint(value[0])
	}
	vals := make([]string, 0, len(value))
	for _, val := range value {
		vals = append(vals, fmt.Sprint(val))
	}
	csv := strings.Join(vals, ",")
	return fmt.Sprintf("(%s)", csv)
}

func genericize[T any](s []T) []any {
	out := make([]any, 0, len(s))
	for _, v := range s {
		out = append(out, v)
	}
	return out
}

func (df *deviceFilter) String() string {
	values := make([]string, 0, len(df.items)+1)
	for _, smt := range df.items {
		values = append(values, smt.String())
	}
	if df.online && !df.offline {
		values = append(values, "online")
	}
	if df.offline && !df.online {
		values = append(values, "offline")
	}
	if df.group != 0 {
		values = append(values, fmt.Sprintf("group %d", df.group))
	}
	return strings.Join(values, and.String())
}

func (df *deviceFilter) Encode() string {
	return url.PathEscape(df.String())
}

func (df *deviceFilter) add(item string, op Operator, value ...any) *deviceFilter {
	if len(value) == 0 {
		panic("no values provided")
	}
	if (op == EQ || op == NEQ) && len(value) > 1 {
		panic(fmt.Sprintf("%s may only be used with one value", op))
	}
	smt := statement{item, op.String(), stringifyValue(value)}
	df.items = append(df.items, smt)
	return df
}

func (df *deviceFilter) Org(op Operator, value ...int) *deviceFilter {
	return df.add("org", op, genericize(value)...)
}

func (df *deviceFilter) Loc(op Operator, value ...int) *deviceFilter {
	return df.add("loc", op, genericize(value)...)
}

func (df *deviceFilter) Role(op Operator, value ...int) *deviceFilter {
	return df.add("role", op, genericize(value)...)
}

func (df *deviceFilter) ID(op Operator, value ...int) *deviceFilter {
	return df.add("id", op, genericize(value)...)
}

func (df *deviceFilter) Class(op Operator, value ...string) *deviceFilter {
	return df.add("class", op, genericize(value)...)
}

func (df *deviceFilter) Status(op Operator, value ApprovalStatus) *deviceFilter {
	if op != EQ && op != NEQ {
		panic(fmt.Errorf("status operator must be EQ or NEQ"))
	}
	return df.add("status", op, value)
}

func (df *deviceFilter) Offline() *deviceFilter {
	df.offline = true
	return df
}

func (df *deviceFilter) Online() *deviceFilter {
	df.online = true
	return df
}

func (df *deviceFilter) Created(op Operator, value time.Time) *deviceFilter {
	if op != EQ && op != BEFORE && op != AFTER {
		panic(fmt.Errorf("created operator must be EQ, BEFORE, or AFTER"))
	}
	return df.add("created", op, value.Format("2006-01-02"))
}

func (df *deviceFilter) Group(id int) *deviceFilter {
	df.group = id
	return df
}
