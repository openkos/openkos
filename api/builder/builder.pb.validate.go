// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: builder/builder.proto

package builder

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

// Validate checks the field values on Builder with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Builder) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Builder with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in BuilderMultiError, or nil if none found.
func (m *Builder) ValidateAll() error {
	return m.validate(true)
}

func (m *Builder) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Intro

	// no validation rules for ImageRepo

	// no validation rules for StackId

	// no validation rules for BuildImage

	// no validation rules for RunImage

	for idx, item := range m.GetPacks() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BuilderValidationError{
						field:  fmt.Sprintf("Packs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BuilderValidationError{
						field:  fmt.Sprintf("Packs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BuilderValidationError{
					field:  fmt.Sprintf("Packs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Status

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BuilderValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BuilderValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BuilderValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BuilderValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BuilderValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BuilderValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BuilderMultiError(errors)
	}

	return nil
}

// BuilderMultiError is an error wrapping multiple validation errors returned
// by Builder.ValidateAll() if the designated constraints aren't met.
type BuilderMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BuilderMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BuilderMultiError) AllErrors() []error { return m }

// BuilderValidationError is the validation error returned by Builder.Validate
// if the designated constraints aren't met.
type BuilderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BuilderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BuilderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BuilderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BuilderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BuilderValidationError) ErrorName() string { return "BuilderValidationError" }

// Error satisfies the builtin error interface
func (e BuilderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBuilder.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BuilderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BuilderValidationError{}

// Validate checks the field values on Pack with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Pack) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Pack with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PackMultiError, or nil if none found.
func (m *Pack) ValidateAll() error {
	return m.validate(true)
}

func (m *Pack) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Lang

	// no validation rules for Image

	for idx, item := range m.GetEnvs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PackValidationError{
						field:  fmt.Sprintf("Envs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PackValidationError{
						field:  fmt.Sprintf("Envs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PackValidationError{
					field:  fmt.Sprintf("Envs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PackMultiError(errors)
	}

	return nil
}

// PackMultiError is an error wrapping multiple validation errors returned by
// Pack.ValidateAll() if the designated constraints aren't met.
type PackMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PackMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PackMultiError) AllErrors() []error { return m }

// PackValidationError is the validation error returned by Pack.Validate if the
// designated constraints aren't met.
type PackValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PackValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PackValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PackValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PackValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PackValidationError) ErrorName() string { return "PackValidationError" }

// Error satisfies the builtin error interface
func (e PackValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPack.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PackValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PackValidationError{}

// Validate checks the field values on Env with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Env) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Env with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EnvMultiError, or nil if none found.
func (m *Env) ValidateAll() error {
	return m.validate(true)
}

func (m *Env) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Intro

	// no validation rules for DefaultValue

	if len(errors) > 0 {
		return EnvMultiError(errors)
	}

	return nil
}

// EnvMultiError is an error wrapping multiple validation errors returned by
// Env.ValidateAll() if the designated constraints aren't met.
type EnvMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EnvMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EnvMultiError) AllErrors() []error { return m }

// EnvValidationError is the validation error returned by Env.Validate if the
// designated constraints aren't met.
type EnvValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnvValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnvValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnvValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnvValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnvValidationError) ErrorName() string { return "EnvValidationError" }

// Error satisfies the builtin error interface
func (e EnvValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnv.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnvValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnvValidationError{}

// Validate checks the field values on BuilderListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *BuilderListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BuilderListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BuilderListRequestMultiError, or nil if none found.
func (m *BuilderListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *BuilderListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return BuilderListRequestMultiError(errors)
	}

	return nil
}

// BuilderListRequestMultiError is an error wrapping multiple validation errors
// returned by BuilderListRequest.ValidateAll() if the designated constraints
// aren't met.
type BuilderListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BuilderListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BuilderListRequestMultiError) AllErrors() []error { return m }

// BuilderListRequestValidationError is the validation error returned by
// BuilderListRequest.Validate if the designated constraints aren't met.
type BuilderListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BuilderListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BuilderListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BuilderListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BuilderListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BuilderListRequestValidationError) ErrorName() string {
	return "BuilderListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e BuilderListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBuilderListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BuilderListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BuilderListRequestValidationError{}

// Validate checks the field values on BuilderListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *BuilderListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BuilderListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BuilderListResponseMultiError, or nil if none found.
func (m *BuilderListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *BuilderListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BuilderListResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BuilderListResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BuilderListResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return BuilderListResponseMultiError(errors)
	}

	return nil
}

// BuilderListResponseMultiError is an error wrapping multiple validation
// errors returned by BuilderListResponse.ValidateAll() if the designated
// constraints aren't met.
type BuilderListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BuilderListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BuilderListResponseMultiError) AllErrors() []error { return m }

// BuilderListResponseValidationError is the validation error returned by
// BuilderListResponse.Validate if the designated constraints aren't met.
type BuilderListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BuilderListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BuilderListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BuilderListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BuilderListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BuilderListResponseValidationError) ErrorName() string {
	return "BuilderListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e BuilderListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBuilderListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BuilderListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BuilderListResponseValidationError{}

// Validate checks the field values on SuggestedStack with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SuggestedStack) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SuggestedStack with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SuggestedStackMultiError,
// or nil if none found.
func (m *SuggestedStack) ValidateAll() error {
	return m.validate(true)
}

func (m *SuggestedStack) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Intro

	// no validation rules for StackId

	// no validation rules for BuildImage

	// no validation rules for RunImage

	if len(errors) > 0 {
		return SuggestedStackMultiError(errors)
	}

	return nil
}

// SuggestedStackMultiError is an error wrapping multiple validation errors
// returned by SuggestedStack.ValidateAll() if the designated constraints
// aren't met.
type SuggestedStackMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SuggestedStackMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SuggestedStackMultiError) AllErrors() []error { return m }

// SuggestedStackValidationError is the validation error returned by
// SuggestedStack.Validate if the designated constraints aren't met.
type SuggestedStackValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SuggestedStackValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SuggestedStackValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SuggestedStackValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SuggestedStackValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SuggestedStackValidationError) ErrorName() string { return "SuggestedStackValidationError" }

// Error satisfies the builtin error interface
func (e SuggestedStackValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSuggestedStack.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SuggestedStackValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SuggestedStackValidationError{}

// Validate checks the field values on SuggestedStackListResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SuggestedStackListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SuggestedStackListResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SuggestedStackListResponseMultiError, or nil if none found.
func (m *SuggestedStackListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SuggestedStackListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SuggestedStackListResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SuggestedStackListResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SuggestedStackListResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SuggestedStackListResponseMultiError(errors)
	}

	return nil
}

// SuggestedStackListResponseMultiError is an error wrapping multiple
// validation errors returned by SuggestedStackListResponse.ValidateAll() if
// the designated constraints aren't met.
type SuggestedStackListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SuggestedStackListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SuggestedStackListResponseMultiError) AllErrors() []error { return m }

// SuggestedStackListResponseValidationError is the validation error returned
// by SuggestedStackListResponse.Validate if the designated constraints aren't met.
type SuggestedStackListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SuggestedStackListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SuggestedStackListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SuggestedStackListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SuggestedStackListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SuggestedStackListResponseValidationError) ErrorName() string {
	return "SuggestedStackListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SuggestedStackListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSuggestedStackListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SuggestedStackListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SuggestedStackListResponseValidationError{}
