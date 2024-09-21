// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/cron/client/cron_client_task.proto

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

// Validate checks the field values on CancelExecTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CancelExecTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CancelExecTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CancelExecTaskRequestMultiError, or nil if none found.
func (m *CancelExecTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CancelExecTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uuid

	if len(errors) > 0 {
		return CancelExecTaskRequestMultiError(errors)
	}

	return nil
}

// CancelExecTaskRequestMultiError is an error wrapping multiple validation
// errors returned by CancelExecTaskRequest.ValidateAll() if the designated
// constraints aren't met.
type CancelExecTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CancelExecTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CancelExecTaskRequestMultiError) AllErrors() []error { return m }

// CancelExecTaskRequestValidationError is the validation error returned by
// CancelExecTaskRequest.Validate if the designated constraints aren't met.
type CancelExecTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CancelExecTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CancelExecTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CancelExecTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CancelExecTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CancelExecTaskRequestValidationError) ErrorName() string {
	return "CancelExecTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CancelExecTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCancelExecTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CancelExecTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CancelExecTaskRequestValidationError{}

// Validate checks the field values on ExecTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ExecTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExecTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ExecTaskRequestMultiError, or nil if none found.
func (m *ExecTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ExecTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Type

	// no validation rules for Value

	// no validation rules for ExpectCode

	// no validation rules for RetryCount

	// no validation rules for RetryWaitTime

	// no validation rules for MaxExecTime

	// no validation rules for Uuid

	if len(errors) > 0 {
		return ExecTaskRequestMultiError(errors)
	}

	return nil
}

// ExecTaskRequestMultiError is an error wrapping multiple validation errors
// returned by ExecTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type ExecTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExecTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExecTaskRequestMultiError) AllErrors() []error { return m }

// ExecTaskRequestValidationError is the validation error returned by
// ExecTaskRequest.Validate if the designated constraints aren't met.
type ExecTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecTaskRequestValidationError) ErrorName() string { return "ExecTaskRequestValidationError" }

// Error satisfies the builtin error interface
func (e ExecTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecTaskRequestValidationError{}

// Validate checks the field values on ExecTaskReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ExecTaskReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExecTaskReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ExecTaskReplyMultiError, or
// nil if none found.
func (m *ExecTaskReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ExecTaskReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for Content

	// no validation rules for Time

	if len(errors) > 0 {
		return ExecTaskReplyMultiError(errors)
	}

	return nil
}

// ExecTaskReplyMultiError is an error wrapping multiple validation errors
// returned by ExecTaskReply.ValidateAll() if the designated constraints
// aren't met.
type ExecTaskReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExecTaskReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExecTaskReplyMultiError) AllErrors() []error { return m }

// ExecTaskReplyValidationError is the validation error returned by
// ExecTaskReply.Validate if the designated constraints aren't met.
type ExecTaskReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExecTaskReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExecTaskReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExecTaskReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExecTaskReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExecTaskReplyValidationError) ErrorName() string { return "ExecTaskReplyValidationError" }

// Error satisfies the builtin error interface
func (e ExecTaskReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExecTaskReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExecTaskReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExecTaskReplyValidationError{}
