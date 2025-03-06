package validator

import (
	"app/utils"
	"errors"
)

type Password struct{}

var specialChars = []rune{
	'!', '@', '#', '$', '%', '^', '&', '*', '(', ')',
	'-', '_', '=', '+', '[', ']', '{', '}', '\\', '|',
	';', ':', '\'', '"', ',', '<', '.', '>', '/', '?', '`', '~',
}

func (p Password) Validate(passwd string) []error {
	var (
		numeric bool
		special bool
		upper   bool
		errs    = make([]error, 0, 4)
	)

	for _, c := range passwd {
		if c >= 'A' && c <= 'Z' {
			upper = true
		} else if c >= '0' && c <= '9' {
			numeric = true
		} else if utils.In(specialChars, c) {
			special = true
		}

		if upper && numeric && special {
			break
		}
	}

	if !numeric {
		errs = append(errs, errors.New("missing numeric character"))
	}
	if !upper {
		errs = append(errs, errors.New("missing upper case character"))
	}
	if !special {
		errs = append(errs, errors.New("missing special character"))

	}
	if len(passwd) < 8 {
		errs = append(errs, errors.New("password needs to have at least 8 characters"))
	}

	return errs
}
