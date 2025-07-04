// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/discover.proto

package v1

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

// Validate checks the field values on SearchRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SearchRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SearchRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SearchRequestMultiError, or
// nil if none found.
func (m *SearchRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SearchRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Query

	// no validation rules for Page

	// no validation rules for PageSize

	if len(errors) > 0 {
		return SearchRequestMultiError(errors)
	}

	return nil
}

// SearchRequestMultiError is an error wrapping multiple validation errors
// returned by SearchRequest.ValidateAll() if the designated constraints
// aren't met.
type SearchRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SearchRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SearchRequestMultiError) AllErrors() []error { return m }

// SearchRequestValidationError is the validation error returned by
// SearchRequest.Validate if the designated constraints aren't met.
type SearchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchRequestValidationError) ErrorName() string { return "SearchRequestValidationError" }

// Error satisfies the builtin error interface
func (e SearchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchRequestValidationError{}

// Validate checks the field values on SearchResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SearchResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SearchResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SearchResponseMultiError,
// or nil if none found.
func (m *SearchResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SearchResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetCategories() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SearchResponseValidationError{
						field:  fmt.Sprintf("Categories[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SearchResponseValidationError{
						field:  fmt.Sprintf("Categories[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SearchResponseValidationError{
					field:  fmt.Sprintf("Categories[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetPrograms() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SearchResponseValidationError{
						field:  fmt.Sprintf("Programs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SearchResponseValidationError{
						field:  fmt.Sprintf("Programs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SearchResponseValidationError{
					field:  fmt.Sprintf("Programs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetEpisodes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SearchResponseValidationError{
						field:  fmt.Sprintf("Episodes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SearchResponseValidationError{
						field:  fmt.Sprintf("Episodes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SearchResponseValidationError{
					field:  fmt.Sprintf("Episodes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for TotalCount

	// no validation rules for Page

	// no validation rules for PageSize

	// no validation rules for TotalPages

	if len(errors) > 0 {
		return SearchResponseMultiError(errors)
	}

	return nil
}

// SearchResponseMultiError is an error wrapping multiple validation errors
// returned by SearchResponse.ValidateAll() if the designated constraints
// aren't met.
type SearchResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SearchResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SearchResponseMultiError) AllErrors() []error { return m }

// SearchResponseValidationError is the validation error returned by
// SearchResponse.Validate if the designated constraints aren't met.
type SearchResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchResponseValidationError) ErrorName() string { return "SearchResponseValidationError" }

// Error satisfies the builtin error interface
func (e SearchResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchResponseValidationError{}

// Validate checks the field values on FeaturedRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *FeaturedRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FeaturedRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FeaturedRequestMultiError, or nil if none found.
func (m *FeaturedRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *FeaturedRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return FeaturedRequestMultiError(errors)
	}

	return nil
}

// FeaturedRequestMultiError is an error wrapping multiple validation errors
// returned by FeaturedRequest.ValidateAll() if the designated constraints
// aren't met.
type FeaturedRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FeaturedRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FeaturedRequestMultiError) AllErrors() []error { return m }

// FeaturedRequestValidationError is the validation error returned by
// FeaturedRequest.Validate if the designated constraints aren't met.
type FeaturedRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FeaturedRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FeaturedRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FeaturedRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FeaturedRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FeaturedRequestValidationError) ErrorName() string { return "FeaturedRequestValidationError" }

// Error satisfies the builtin error interface
func (e FeaturedRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeaturedRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FeaturedRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FeaturedRequestValidationError{}

// Validate checks the field values on FeaturedResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *FeaturedResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FeaturedResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FeaturedResponseMultiError, or nil if none found.
func (m *FeaturedResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *FeaturedResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetPrograms() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FeaturedResponseValidationError{
						field:  fmt.Sprintf("Programs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FeaturedResponseValidationError{
						field:  fmt.Sprintf("Programs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FeaturedResponseValidationError{
					field:  fmt.Sprintf("Programs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return FeaturedResponseMultiError(errors)
	}

	return nil
}

// FeaturedResponseMultiError is an error wrapping multiple validation errors
// returned by FeaturedResponse.ValidateAll() if the designated constraints
// aren't met.
type FeaturedResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FeaturedResponseMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FeaturedResponseMultiError) AllErrors() []error { return m }

// FeaturedResponseValidationError is the validation error returned by
// FeaturedResponse.Validate if the designated constraints aren't met.
type FeaturedResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FeaturedResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FeaturedResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FeaturedResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FeaturedResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FeaturedResponseValidationError) ErrorName() string { return "FeaturedResponseValidationError" }

// Error satisfies the builtin error interface
func (e FeaturedResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeaturedResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FeaturedResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FeaturedResponseValidationError{}
