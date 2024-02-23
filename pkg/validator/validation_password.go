package validator

import "github.com/go-playground/validator/v10"

// PasswordValidation はパスワードのバリデーションを行います。
// 8文字以上、128文字以下、大文字、小文字、数字、記号を含む必要があります。
// 記号は次のものを含みます。
// !"#$%&'()*+,-./:;<=>?@[\]^_`{|}~
func PasswordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 {
		return false
	}

	if len(password) > 128 {
		return false
	}

	if !containsUpper(password) {
		return false
	}

	if !containsLower(password) {
		return false
	}

	if !containsNumber(password) {
		return false
	}

	if !containsSymbol(password) {
		return false
	}

	return true
}

func containsNumber(s string) bool {
	for _, r := range s {
		if '0' <= r && r <= '9' {
			return true
		}
	}
	return false
}

func containsUpper(s string) bool {
	for _, r := range s {
		if 'A' <= r && r <= 'Z' {
			return true
		}
	}
	return false
}

func containsLower(s string) bool {
	for _, r := range s {
		if 'a' <= r && r <= 'z' {
			return true
		}
	}
	return false
}

func containsSymbol(s string) bool {
	for _, r := range s {
		if '!' <= r && r <= '/' {
			return true
		}
		if ':' <= r && r <= '@' {
			return true
		}
		if '[' <= r && r <= '`' {
			return true
		}
		if '{' <= r && r <= '~' {
			return true
		}
	}
	return false
}
