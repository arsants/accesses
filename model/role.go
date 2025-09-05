package model

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
)

var (
	ErrRoleEnumInvalid   = errors.New("invalid role sysname")
	ErrRoleEnumUnknown   = errors.New("unknown role sysname")
	ErrRoleEnumEmpty     = errors.New("empty role sysname")
	ErrRoleEnumNotString = errors.New("role sysname must be a string")
)

type RoleEnum string

const (
	RoleCSSTAFFDESKADMIN      = "cs-staffdesk-admin"
	RoleCSSTAFFDESKMONITORING = "cs-staffdesk-monitoring"
	RoleCSSTAFFDESKTRAINER    = "cs-staffdesk-trainer"
	RoleCSSTAFFDESKS4S        = "cs-staffdesk-s4s"
	RoleCSSTAFFDESKMANAGER    = "cs-staffdesk-manager"
	RoleCSSTAFFDESKLEAD       = "cs-staffdesk-lead"
	RoleCSSTAFFDESKOPERATOR   = "cs-staffdesk-operator"
	RoleUnknown               = "unknown"
)

//nolint:gochecknoglobals
var AllRoleEnum = []RoleEnum{
	RoleCSSTAFFDESKADMIN,
	RoleCSSTAFFDESKMONITORING,
	RoleCSSTAFFDESKTRAINER,
	RoleCSSTAFFDESKS4S,
	RoleCSSTAFFDESKMANAGER,
	RoleCSSTAFFDESKLEAD,
	RoleCSSTAFFDESKOPERATOR,
}

func (e RoleEnum) IsValid() bool {
	switch e {
	case RoleCSSTAFFDESKADMIN:
		return true
	case RoleCSSTAFFDESKMONITORING:
		return true
	case RoleCSSTAFFDESKTRAINER:
		return true
	case RoleCSSTAFFDESKS4S:
		return true
	case RoleCSSTAFFDESKMANAGER:
		return true
	case RoleCSSTAFFDESKLEAD:
		return true
	case RoleCSSTAFFDESKOPERATOR:
		return true
	case RoleUnknown:
		return false
	}

	return false
}

func (e RoleEnum) String() string {
	return string(e)
}

func (e *RoleEnum) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("%w: enums must be strings", ErrRoleEnumNotString)
	}

	*e = RoleEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%w: %s", ErrRoleEnumInvalid, str)
	}

	return nil
}

func (e RoleEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

func (e *RoleEnum) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	return e.UnmarshalGQL(s)
}

func (e RoleEnum) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer

	e.MarshalGQL(&buf)

	return buf.Bytes(), nil
}

func RoleEnumFromString(s string) (RoleEnum, error) {
	if RoleEnum(s).IsValid() {
		return RoleEnum(s), nil
	}

	return RoleUnknown, fmt.Errorf("%w: %s", ErrRoleEnumInvalid, s)
}
