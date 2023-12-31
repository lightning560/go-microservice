// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: job/comment/v1/comment.proto

package jobcommentv1

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

// Validate checks the field values on CommentMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CommentMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommentMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CommentMessageMultiError,
// or nil if none found.
func (m *CommentMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *CommentMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for ReplyId

	// no validation rules for Uid

	// no validation rules for OwnerUid

	if m.GetSubjectId() <= 0 {
		err := CommentMessageValidationError{
			field:  "SubjectId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for BizType

	// no validation rules for FloorId

	// no validation rules for AtUid

	// no validation rules for Content

	if len(errors) > 0 {
		return CommentMessageMultiError(errors)
	}

	return nil
}

// CommentMessageMultiError is an error wrapping multiple validation errors
// returned by CommentMessage.ValidateAll() if the designated constraints
// aren't met.
type CommentMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommentMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommentMessageMultiError) AllErrors() []error { return m }

// CommentMessageValidationError is the validation error returned by
// CommentMessage.Validate if the designated constraints aren't met.
type CommentMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommentMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommentMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommentMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommentMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommentMessageValidationError) ErrorName() string { return "CommentMessageValidationError" }

// Error satisfies the builtin error interface
func (e CommentMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommentMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommentMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommentMessageValidationError{}

// Validate checks the field values on AddLikeMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddLikeMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddLikeMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddLikeMessageMultiError,
// or nil if none found.
func (m *AddLikeMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *AddLikeMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ReplyId

	// no validation rules for Uid

	if len(errors) > 0 {
		return AddLikeMessageMultiError(errors)
	}

	return nil
}

// AddLikeMessageMultiError is an error wrapping multiple validation errors
// returned by AddLikeMessage.ValidateAll() if the designated constraints
// aren't met.
type AddLikeMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddLikeMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddLikeMessageMultiError) AllErrors() []error { return m }

// AddLikeMessageValidationError is the validation error returned by
// AddLikeMessage.Validate if the designated constraints aren't met.
type AddLikeMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddLikeMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddLikeMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddLikeMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddLikeMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddLikeMessageValidationError) ErrorName() string { return "AddLikeMessageValidationError" }

// Error satisfies the builtin error interface
func (e AddLikeMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddLikeMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddLikeMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddLikeMessageValidationError{}

// Validate checks the field values on AddReplyMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddReplyMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddReplyMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddReplyMessageMultiError, or nil if none found.
func (m *AddReplyMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *AddReplyMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetSubjectId() <= 0 {
		err := AddReplyMessageValidationError{
			field:  "SubjectId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for BizType

	// no validation rules for OwnerUid

	// no validation rules for FloorId

	// no validation rules for AtUid

	// no validation rules for Content

	// no validation rules for IsFloor

	if len(errors) > 0 {
		return AddReplyMessageMultiError(errors)
	}

	return nil
}

// AddReplyMessageMultiError is an error wrapping multiple validation errors
// returned by AddReplyMessage.ValidateAll() if the designated constraints
// aren't met.
type AddReplyMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddReplyMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddReplyMessageMultiError) AllErrors() []error { return m }

// AddReplyMessageValidationError is the validation error returned by
// AddReplyMessage.Validate if the designated constraints aren't met.
type AddReplyMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddReplyMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddReplyMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddReplyMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddReplyMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddReplyMessageValidationError) ErrorName() string { return "AddReplyMessageValidationError" }

// Error satisfies the builtin error interface
func (e AddReplyMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddReplyMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddReplyMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddReplyMessageValidationError{}
