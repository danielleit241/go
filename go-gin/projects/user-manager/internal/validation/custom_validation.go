package validation

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/danielleit241/internal/utils"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type CustomValidation struct {
	Tag string
	Fn  validator.Func
}

var customValidations = []CustomValidation{
	{
		Tag: "blocked_email_domain",
		Fn:  validateBlockedEmailDomain,
	},
	{
		Tag: "strong_password",
		Fn:  validateStrongPassword,
	},
}

func RegisterCustomValidations(v *validator.Validate) {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		panic("failed to get validator engine")
	}

	for _, cv := range customValidations {
		if err := v.RegisterValidation(cv.Tag, cv.Fn); err != nil {
			panic(fmt.Sprintf("failed to register validator for tag '%s': %v", cv.Tag, err))
		}
	}
}

func validateBlockedEmailDomain(fl validator.FieldLevel) bool {
	blockedDomains := map[string]bool{
		"admin.com": true,
		"test.com":  true,
	}

	email := fl.Field().String()
	emailParts := strings.Split(email, "@")

	if len(emailParts) != 2 {
		return false
	}

	domain := utils.NormalizeString(emailParts[1])

	return !blockedDomains[domain]
}

func validateStrongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	hasMinLength := len(password) >= 8
	hasUpper := strings.IndexFunc(password, unicode.IsUpper) >= 0
	hasLower := strings.IndexFunc(password, unicode.IsLower) >= 0
	hasNumber := strings.IndexFunc(password, unicode.IsDigit) >= 0
	hasSpecial := strings.IndexFunc(password, func(r rune) bool {
		return unicode.IsPunct(r) || unicode.IsSymbol(r)
	}) >= 0

	return hasMinLength && hasUpper && hasLower && hasNumber && hasSpecial
}
