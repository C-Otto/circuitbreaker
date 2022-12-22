// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: circuitbreaker.proto

package circuitbreakerrpc

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on GetInfoRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetInfoRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// GetInfoRequestValidationError is the validation error returned by
// GetInfoRequest.Validate if the designated constraints aren't met.
type GetInfoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetInfoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetInfoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetInfoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetInfoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetInfoRequestValidationError) ErrorName() string { return "GetInfoRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetInfoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetInfoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetInfoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetInfoRequestValidationError{}

// Validate checks the field values on GetInfoResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetInfoResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ConnectedNode

	return nil
}

// GetInfoResponseValidationError is the validation error returned by
// GetInfoResponse.Validate if the designated constraints aren't met.
type GetInfoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetInfoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetInfoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetInfoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetInfoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetInfoResponseValidationError) ErrorName() string { return "GetInfoResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetInfoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetInfoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetInfoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetInfoResponseValidationError{}

// Validate checks the field values on UpdateLimitRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateLimitRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Node

	// no validation rules for MaxHourlyRate

	// no validation rules for MaxPending

	return nil
}

// UpdateLimitRequestValidationError is the validation error returned by
// UpdateLimitRequest.Validate if the designated constraints aren't met.
type UpdateLimitRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateLimitRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateLimitRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateLimitRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateLimitRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateLimitRequestValidationError) ErrorName() string {
	return "UpdateLimitRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateLimitRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateLimitRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateLimitRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateLimitRequestValidationError{}

// Validate checks the field values on UpdateLimitResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateLimitResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateLimitResponseValidationError is the validation error returned by
// UpdateLimitResponse.Validate if the designated constraints aren't met.
type UpdateLimitResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateLimitResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateLimitResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateLimitResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateLimitResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateLimitResponseValidationError) ErrorName() string {
	return "UpdateLimitResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateLimitResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateLimitResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateLimitResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateLimitResponseValidationError{}

// Validate checks the field values on ListLimitsRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListLimitsRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListLimitsRequestValidationError is the validation error returned by
// ListLimitsRequest.Validate if the designated constraints aren't met.
type ListLimitsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListLimitsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListLimitsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListLimitsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListLimitsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListLimitsRequestValidationError) ErrorName() string {
	return "ListLimitsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListLimitsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListLimitsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListLimitsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListLimitsRequestValidationError{}

// Validate checks the field values on ListLimitsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListLimitsResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetGlobalLimit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListLimitsResponseValidationError{
				field:  "GlobalLimit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetLimits() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListLimitsResponseValidationError{
					field:  fmt.Sprintf("Limits[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListLimitsResponseValidationError is the validation error returned by
// ListLimitsResponse.Validate if the designated constraints aren't met.
type ListLimitsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListLimitsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListLimitsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListLimitsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListLimitsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListLimitsResponseValidationError) ErrorName() string {
	return "ListLimitsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListLimitsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListLimitsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListLimitsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListLimitsResponseValidationError{}

// Validate checks the field values on NodeLimit with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *NodeLimit) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Node

	if v, ok := interface{}(m.GetLimit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeLimitValidationError{
				field:  "Limit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetCounters() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeLimitValidationError{
					field:  fmt.Sprintf("Counters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// NodeLimitValidationError is the validation error returned by
// NodeLimit.Validate if the designated constraints aren't met.
type NodeLimitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeLimitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeLimitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeLimitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeLimitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeLimitValidationError) ErrorName() string { return "NodeLimitValidationError" }

// Error satisfies the builtin error interface
func (e NodeLimitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeLimit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeLimitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeLimitValidationError{}

// Validate checks the field values on Limit with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Limit) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for MaxHourlyRate

	// no validation rules for MaxPending

	return nil
}

// LimitValidationError is the validation error returned by Limit.Validate if
// the designated constraints aren't met.
type LimitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LimitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LimitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LimitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LimitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LimitValidationError) ErrorName() string { return "LimitValidationError" }

// Error satisfies the builtin error interface
func (e LimitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLimit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LimitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LimitValidationError{}

// Validate checks the field values on Counter with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Counter) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Total

	// no validation rules for Successes

	return nil
}

// CounterValidationError is the validation error returned by Counter.Validate
// if the designated constraints aren't met.
type CounterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CounterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CounterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CounterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CounterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CounterValidationError) ErrorName() string { return "CounterValidationError" }

// Error satisfies the builtin error interface
func (e CounterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCounter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CounterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CounterValidationError{}
