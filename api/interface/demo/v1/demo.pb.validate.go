// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: interface/demo/v1/demo.proto

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

// Validate checks the field values on CreateDemoReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CreateDemoReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateDemoReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CreateDemoReqMultiError, or
// nil if none found.
func (m *CreateDemoReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateDemoReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if m.GetAge() >= 120 {
		err := CreateDemoReqValidationError{
			field:  "Age",
			reason: "value must be less than 120",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateDemoReqMultiError(errors)
	}

	return nil
}

// CreateDemoReqMultiError is an error wrapping multiple validation errors
// returned by CreateDemoReq.ValidateAll() if the designated constraints
// aren't met.
type CreateDemoReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateDemoReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateDemoReqMultiError) AllErrors() []error { return m }

// CreateDemoReqValidationError is the validation error returned by
// CreateDemoReq.Validate if the designated constraints aren't met.
type CreateDemoReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateDemoReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateDemoReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateDemoReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateDemoReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateDemoReqValidationError) ErrorName() string { return "CreateDemoReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateDemoReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateDemoReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateDemoReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateDemoReqValidationError{}

// Validate checks the field values on CreateDemoReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateDemoReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateDemoReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateDemoReplyMultiError, or nil if none found.
func (m *CreateDemoReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateDemoReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateDemoReplyMultiError(errors)
	}

	return nil
}

// CreateDemoReplyMultiError is an error wrapping multiple validation errors
// returned by CreateDemoReply.ValidateAll() if the designated constraints
// aren't met.
type CreateDemoReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateDemoReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateDemoReplyMultiError) AllErrors() []error { return m }

// CreateDemoReplyValidationError is the validation error returned by
// CreateDemoReply.Validate if the designated constraints aren't met.
type CreateDemoReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateDemoReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateDemoReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateDemoReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateDemoReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateDemoReplyValidationError) ErrorName() string { return "CreateDemoReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateDemoReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateDemoReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateDemoReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateDemoReplyValidationError{}

// Validate checks the field values on StringMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StringMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StringMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StringMessageMultiError, or
// nil if none found.
func (m *StringMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *StringMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	if len(errors) > 0 {
		return StringMessageMultiError(errors)
	}

	return nil
}

// StringMessageMultiError is an error wrapping multiple validation errors
// returned by StringMessage.ValidateAll() if the designated constraints
// aren't met.
type StringMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StringMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StringMessageMultiError) AllErrors() []error { return m }

// StringMessageValidationError is the validation error returned by
// StringMessage.Validate if the designated constraints aren't met.
type StringMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StringMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StringMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StringMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StringMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StringMessageValidationError) ErrorName() string { return "StringMessageValidationError" }

// Error satisfies the builtin error interface
func (e StringMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStringMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StringMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StringMessageValidationError{}