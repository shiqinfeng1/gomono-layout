// Copyright 2024 slw 150627601@qq.com . All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pagination/v1/pagination.proto

package pagination

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on PagingRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PagingRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PagingRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PagingRequestMultiError, or
// nil if none found.
func (m *PagingRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PagingRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetFieldMask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PagingRequestValidationError{
					field:  "FieldMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PagingRequestValidationError{
					field:  "FieldMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFieldMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PagingRequestValidationError{
				field:  "FieldMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.Page != nil {
		// no validation rules for Page
	}

	if m.PageSize != nil {
		// no validation rules for PageSize
	}

	if m.Query != nil {
		// no validation rules for Query
	}

	if m.OrQuery != nil {
		// no validation rules for OrQuery
	}

	if m.NoPaging != nil {
		// no validation rules for NoPaging
	}

	if len(errors) > 0 {
		return PagingRequestMultiError(errors)
	}

	return nil
}

// PagingRequestMultiError is an error wrapping multiple validation errors
// returned by PagingRequest.ValidateAll() if the designated constraints
// aren't met.
type PagingRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PagingRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PagingRequestMultiError) AllErrors() []error { return m }

// PagingRequestValidationError is the validation error returned by
// PagingRequest.Validate if the designated constraints aren't met.
type PagingRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PagingRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PagingRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PagingRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PagingRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PagingRequestValidationError) ErrorName() string { return "PagingRequestValidationError" }

// Error satisfies the builtin error interface
func (e PagingRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPagingRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PagingRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PagingRequestValidationError{}

// Validate checks the field values on PagingResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PagingResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PagingResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PagingResponseMultiError,
// or nil if none found.
func (m *PagingResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PagingResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PagingResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PagingResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PagingResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PagingResponseMultiError(errors)
	}

	return nil
}

// PagingResponseMultiError is an error wrapping multiple validation errors
// returned by PagingResponse.ValidateAll() if the designated constraints
// aren't met.
type PagingResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PagingResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PagingResponseMultiError) AllErrors() []error { return m }

// PagingResponseValidationError is the validation error returned by
// PagingResponse.Validate if the designated constraints aren't met.
type PagingResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PagingResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PagingResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PagingResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PagingResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PagingResponseValidationError) ErrorName() string { return "PagingResponseValidationError" }

// Error satisfies the builtin error interface
func (e PagingResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPagingResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PagingResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PagingResponseValidationError{}
