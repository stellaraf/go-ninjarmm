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

// NewDeviceFilter creates a device filter object, which provides a native method to create a
// device filter.
//
// https://app.ninjaone.com/apidocs-beta/core-resources/articles/devices/device-filters
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

// String returns a non-escaped and non-encoded string representation of the device filter.
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

// Encode returns a URL-encoded string representation of the device filter, which is what NinjaRMM
// expects to receive.
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

// Org adds an organization filter.
func (df *deviceFilter) Org(op Operator, value ...int) *deviceFilter {
	return df.add("org", op, genericize(value)...)
}

// Loc adds a location filter.
func (df *deviceFilter) Loc(op Operator, value ...int) *deviceFilter {
	return df.add("loc", op, genericize(value)...)
}

// Role adds a role filter.
func (df *deviceFilter) Role(op Operator, value ...int) *deviceFilter {
	return df.add("role", op, genericize(value)...)
}

// ID adds a device ID filter.
func (df *deviceFilter) ID(op Operator, value ...int) *deviceFilter {
	return df.add("id", op, genericize(value)...)
}

// Class adds a Node Class filter.
func (df *deviceFilter) Class(op Operator, value ...string) *deviceFilter {
	return df.add("class", op, genericize(value)...)
}

// Status adds a device approval status filter. Only EQ or NEQ operators may be used.
func (df *deviceFilter) Status(op Operator, value ApprovalStatus) *deviceFilter {
	if op != EQ && op != NEQ {
		panic(fmt.Errorf("status operator must be EQ or NEQ"))
	}
	return df.add("status", op, value)
}

// Offline adds a device offline filter.
func (df *deviceFilter) Offline() *deviceFilter {
	df.offline = true
	return df
}

// Online adds a device online filter.
func (df *deviceFilter) Online() *deviceFilter {
	df.online = true
	return df
}

// Created adds a device creation date filter. Only BEFORE and AFTER operators may be used.
func (df *deviceFilter) Created(op Operator, value time.Time) *deviceFilter {
	if op != EQ && op != BEFORE && op != AFTER {
		panic(fmt.Errorf("created operator must be EQ, BEFORE, or AFTER"))
	}
	return df.add("created", op, value.Format("2006-01-02"))
}

// Group adds a group membership filter.
func (df *deviceFilter) Group(id int) *deviceFilter {
	df.group = id
	return df
}
