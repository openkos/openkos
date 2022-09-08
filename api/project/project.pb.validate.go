// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: project/project.proto

package project

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

// Validate checks the field values on Project with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Project) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Project with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ProjectMultiError, or nil if none found.
func (m *Project) ValidateAll() error {
	return m.validate(true)
}

func (m *Project) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if l := utf8.RuneCountInString(m.GetName()); l < 3 || l > 32 {
		err := ProjectValidationError{
			field:  "Name",
			reason: "value length must be between 3 and 32 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetIntro()); l < 5 || l > 150 {
		err := ProjectValidationError{
			field:  "Intro",
			reason: "value length must be between 5 and 150 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetGit()); l < 10 || l > 512 {
		err := ProjectValidationError{
			field:  "Git",
			reason: "value length must be between 10 and 512 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetImages() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProjectValidationError{
						field:  fmt.Sprintf("Images[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProjectValidationError{
						field:  fmt.Sprintf("Images[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProjectValidationError{
					field:  fmt.Sprintf("Images[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetPorts() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProjectValidationError{
						field:  fmt.Sprintf("Ports[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProjectValidationError{
						field:  fmt.Sprintf("Ports[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProjectValidationError{
					field:  fmt.Sprintf("Ports[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetConfigs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProjectValidationError{
						field:  fmt.Sprintf("Configs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProjectValidationError{
						field:  fmt.Sprintf("Configs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProjectValidationError{
					field:  fmt.Sprintf("Configs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetMiddlewares() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ProjectValidationError{
						field:  fmt.Sprintf("Middlewares[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ProjectValidationError{
						field:  fmt.Sprintf("Middlewares[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProjectValidationError{
					field:  fmt.Sprintf("Middlewares[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetLivenessProbe()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProjectValidationError{
					field:  "LivenessProbe",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProjectValidationError{
					field:  "LivenessProbe",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLivenessProbe()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectValidationError{
				field:  "LivenessProbe",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetReadinessProbe()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProjectValidationError{
					field:  "ReadinessProbe",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProjectValidationError{
					field:  "ReadinessProbe",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetReadinessProbe()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectValidationError{
				field:  "ReadinessProbe",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProjectValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProjectValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedBy()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProjectValidationError{
					field:  "CreatedBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProjectValidationError{
					field:  "CreatedBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedBy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectValidationError{
				field:  "CreatedBy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProjectValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProjectValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedBy()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProjectValidationError{
					field:  "UpdatedBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProjectValidationError{
					field:  "UpdatedBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedBy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectValidationError{
				field:  "UpdatedBy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ProjectMultiError(errors)
	}

	return nil
}

// ProjectMultiError is an error wrapping multiple validation errors returned
// by Project.ValidateAll() if the designated constraints aren't met.
type ProjectMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProjectMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProjectMultiError) AllErrors() []error { return m }

// ProjectValidationError is the validation error returned by Project.Validate
// if the designated constraints aren't met.
type ProjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectValidationError) ErrorName() string { return "ProjectValidationError" }

// Error satisfies the builtin error interface
func (e ProjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectValidationError{}

// Validate checks the field values on HealthProbe with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HealthProbe) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthProbe with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HealthProbeMultiError, or
// nil if none found.
func (m *HealthProbe) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthProbe) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Port

	// no validation rules for Path

	// no validation rules for PeriodSeconds

	// no validation rules for TimeoutSeconds

	// no validation rules for SuccessThreshold

	// no validation rules for FailureThreshold

	// no validation rules for InitialDelaySeconds

	if len(errors) > 0 {
		return HealthProbeMultiError(errors)
	}

	return nil
}

// HealthProbeMultiError is an error wrapping multiple validation errors
// returned by HealthProbe.ValidateAll() if the designated constraints aren't met.
type HealthProbeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthProbeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthProbeMultiError) AllErrors() []error { return m }

// HealthProbeValidationError is the validation error returned by
// HealthProbe.Validate if the designated constraints aren't met.
type HealthProbeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthProbeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthProbeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthProbeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthProbeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthProbeValidationError) ErrorName() string { return "HealthProbeValidationError" }

// Error satisfies the builtin error interface
func (e HealthProbeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthProbe.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthProbeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthProbeValidationError{}

// Validate checks the field values on Image with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Image) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Image with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ImageMultiError, or nil if none found.
func (m *Image) ValidateAll() error {
	return m.validate(true)
}

func (m *Image) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Latest

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ImageValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ImageValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ImageValidationError{
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
				errors = append(errors, ImageValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ImageValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ImageValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ImageMultiError(errors)
	}

	return nil
}

// ImageMultiError is an error wrapping multiple validation errors returned by
// Image.ValidateAll() if the designated constraints aren't met.
type ImageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ImageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ImageMultiError) AllErrors() []error { return m }

// ImageValidationError is the validation error returned by Image.Validate if
// the designated constraints aren't met.
type ImageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImageValidationError) ErrorName() string { return "ImageValidationError" }

// Error satisfies the builtin error interface
func (e ImageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImageValidationError{}

// Validate checks the field values on Port with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Port) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Port with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PortMultiError, or nil if none found.
func (m *Port) ValidateAll() error {
	return m.validate(true)
}

func (m *Port) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if val := m.GetPort(); val <= 0 || val > 65535 {
		err := PortValidationError{
			field:  "Port",
			reason: "value must be inside range (0, 65535]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _Port_Protocol_InLookup[m.GetProtocol()]; !ok {
		err := PortValidationError{
			field:  "Protocol",
			reason: "value must be in list [TCP UDP STCP]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _Port_AppProtocol_InLookup[m.GetAppProtocol()]; !ok {
		err := PortValidationError{
			field:  "AppProtocol",
			reason: "value must be in list [http http2 tcp]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PortMultiError(errors)
	}

	return nil
}

// PortMultiError is an error wrapping multiple validation errors returned by
// Port.ValidateAll() if the designated constraints aren't met.
type PortMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PortMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PortMultiError) AllErrors() []error { return m }

// PortValidationError is the validation error returned by Port.Validate if the
// designated constraints aren't met.
type PortValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PortValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PortValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PortValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PortValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PortValidationError) ErrorName() string { return "PortValidationError" }

// Error satisfies the builtin error interface
func (e PortValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPort.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PortValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PortValidationError{}

var _Port_Protocol_InLookup = map[string]struct{}{
	"TCP":  {},
	"UDP":  {},
	"STCP": {},
}

var _Port_AppProtocol_InLookup = map[string]struct{}{
	"http":  {},
	"http2": {},
	"tcp":   {},
}

// Validate checks the field values on Configuration with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Configuration) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Configuration with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ConfigurationMultiError, or
// nil if none found.
func (m *Configuration) ValidateAll() error {
	return m.validate(true)
}

func (m *Configuration) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetFilename()); l < 2 || l > 16 {
		err := ConfigurationValidationError{
			field:  "Filename",
			reason: "value length must be between 2 and 16 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Content

	if l := utf8.RuneCountInString(m.GetMountPath()); l < 2 || l > 255 {
		err := ConfigurationValidationError{
			field:  "MountPath",
			reason: "value length must be between 2 and 255 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Level

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfigurationValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfigurationValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigurationValidationError{
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
				errors = append(errors, ConfigurationValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfigurationValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigurationValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ConfigurationMultiError(errors)
	}

	return nil
}

// ConfigurationMultiError is an error wrapping multiple validation errors
// returned by Configuration.ValidateAll() if the designated constraints
// aren't met.
type ConfigurationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfigurationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfigurationMultiError) AllErrors() []error { return m }

// ConfigurationValidationError is the validation error returned by
// Configuration.Validate if the designated constraints aren't met.
type ConfigurationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigurationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigurationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigurationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigurationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigurationValidationError) ErrorName() string { return "ConfigurationValidationError" }

// Error satisfies the builtin error interface
func (e ConfigurationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfiguration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigurationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigurationValidationError{}

// Validate checks the field values on Middleware with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Middleware) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Middleware with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MiddlewareMultiError, or
// nil if none found.
func (m *Middleware) ValidateAll() error {
	return m.validate(true)
}

func (m *Middleware) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if _, ok := _Middleware_Type_InLookup[m.GetType()]; !ok {
		err := MiddlewareValidationError{
			field:  "Type",
			reason: "value must be in list [MySQL MongoDB Redis Memcached Kafka RabbitMQ iNotify]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	{
		sorted_keys := make([]string, len(m.GetProperties()))
		i := 0
		for key := range m.GetProperties() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetProperties()[key]
			_ = val

			// no validation rules for Properties[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, MiddlewareValidationError{
							field:  fmt.Sprintf("Properties[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, MiddlewareValidationError{
							field:  fmt.Sprintf("Properties[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return MiddlewareValidationError{
						field:  fmt.Sprintf("Properties[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return MiddlewareMultiError(errors)
	}

	return nil
}

// MiddlewareMultiError is an error wrapping multiple validation errors
// returned by Middleware.ValidateAll() if the designated constraints aren't met.
type MiddlewareMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MiddlewareMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MiddlewareMultiError) AllErrors() []error { return m }

// MiddlewareValidationError is the validation error returned by
// Middleware.Validate if the designated constraints aren't met.
type MiddlewareValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MiddlewareValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MiddlewareValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MiddlewareValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MiddlewareValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MiddlewareValidationError) ErrorName() string { return "MiddlewareValidationError" }

// Error satisfies the builtin error interface
func (e MiddlewareValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMiddleware.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MiddlewareValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MiddlewareValidationError{}

var _Middleware_Type_InLookup = map[string]struct{}{
	"MySQL":     {},
	"MongoDB":   {},
	"Redis":     {},
	"Memcached": {},
	"Kafka":     {},
	"RabbitMQ":  {},
	"iNotify":   {},
}

// Validate checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListRequestMultiError, or
// nil if none found.
func (m *ListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return ListRequestMultiError(errors)
	}

	return nil
}

// ListRequestMultiError is an error wrapping multiple validation errors
// returned by ListRequest.ValidateAll() if the designated constraints aren't met.
type ListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRequestMultiError) AllErrors() []error { return m }

// ListRequestValidationError is the validation error returned by
// ListRequest.Validate if the designated constraints aren't met.
type ListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRequestValidationError) ErrorName() string { return "ListRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRequestValidationError{}

// Validate checks the field values on ListResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListResponseMultiError, or
// nil if none found.
func (m *ListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListResponse) validate(all bool) error {
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
					errors = append(errors, ListResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListResponseMultiError(errors)
	}

	return nil
}

// ListResponseMultiError is an error wrapping multiple validation errors
// returned by ListResponse.ValidateAll() if the designated constraints aren't met.
type ListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListResponseMultiError) AllErrors() []error { return m }

// ListResponseValidationError is the validation error returned by
// ListResponse.Validate if the designated constraints aren't met.
type ListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResponseValidationError) ErrorName() string { return "ListResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResponseValidationError{}
