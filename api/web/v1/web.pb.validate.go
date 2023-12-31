// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: web/v1/web.proto

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

// Validate checks the field values on CreateWarehouseRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateWarehouseRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateWarehouseRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateWarehouseRequestMultiError, or nil if none found.
func (m *CreateWarehouseRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateWarehouseRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return CreateWarehouseRequestMultiError(errors)
	}

	return nil
}

// CreateWarehouseRequestMultiError is an error wrapping multiple validation
// errors returned by CreateWarehouseRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateWarehouseRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateWarehouseRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateWarehouseRequestMultiError) AllErrors() []error { return m }

// CreateWarehouseRequestValidationError is the validation error returned by
// CreateWarehouseRequest.Validate if the designated constraints aren't met.
type CreateWarehouseRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateWarehouseRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateWarehouseRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateWarehouseRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateWarehouseRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateWarehouseRequestValidationError) ErrorName() string {
	return "CreateWarehouseRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateWarehouseRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateWarehouseRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateWarehouseRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateWarehouseRequestValidationError{}

// Validate checks the field values on CreateWarehouseResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateWarehouseResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateWarehouseResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateWarehouseResponseMultiError, or nil if none found.
func (m *CreateWarehouseResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateWarehouseResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateWarehouseResponseMultiError(errors)
	}

	return nil
}

// CreateWarehouseResponseMultiError is an error wrapping multiple validation
// errors returned by CreateWarehouseResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateWarehouseResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateWarehouseResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateWarehouseResponseMultiError) AllErrors() []error { return m }

// CreateWarehouseResponseValidationError is the validation error returned by
// CreateWarehouseResponse.Validate if the designated constraints aren't met.
type CreateWarehouseResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateWarehouseResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateWarehouseResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateWarehouseResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateWarehouseResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateWarehouseResponseValidationError) ErrorName() string {
	return "CreateWarehouseResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateWarehouseResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateWarehouseResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateWarehouseResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateWarehouseResponseValidationError{}

// Validate checks the field values on NodeType with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NodeType) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NodeType with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NodeTypeMultiError, or nil
// if none found.
func (m *NodeType) ValidateAll() error {
	return m.validate(true)
}

func (m *NodeType) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return NodeTypeMultiError(errors)
	}

	return nil
}

// NodeTypeMultiError is an error wrapping multiple validation errors returned
// by NodeType.ValidateAll() if the designated constraints aren't met.
type NodeTypeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NodeTypeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NodeTypeMultiError) AllErrors() []error { return m }

// NodeTypeValidationError is the validation error returned by
// NodeType.Validate if the designated constraints aren't met.
type NodeTypeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeTypeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeTypeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeTypeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeTypeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeTypeValidationError) ErrorName() string { return "NodeTypeValidationError" }

// Error satisfies the builtin error interface
func (e NodeTypeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeType.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeTypeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeTypeValidationError{}
