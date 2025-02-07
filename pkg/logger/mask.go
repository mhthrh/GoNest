package logger

import (
	"encoding/json"
	. "github.com/ggwhite/go-masker"
	"go.uber.org/zap"
)

type Level string

var (
	NameLevel       Level = "NameLevel"
	PasswordLevel   Level = "PasswordLevel"
	AddressLevel    Level = "AddressLevel"
	EmailLevel      Level = "EmailLevel"
	MobileLevel     Level = "MobileLevel"
	TelephoneLevel  Level = "TelephoneLevel"
	IDLevel         Level = "IDLevel"
	CreditCardLevel Level = "CreditCardLevel"
)

func MaskingStr(v string, maskLevel Level) string {
	var maskedStr string
	switch maskLevel {
	case NameLevel:
		maskedStr = Name(v)
	case AddressLevel:
		maskedStr = Address(v)
	case EmailLevel:
		maskedStr = Email(v)
	case MobileLevel:
		maskedStr = Mobile(v)
	case TelephoneLevel:
		maskedStr = Telephone(v)
	case IDLevel:
		maskedStr = ID(v)
	case CreditCardLevel:
		maskedStr = CreditCard(v)
	default:
		maskedStr = Password(v)
	}

	return maskedStr
}

func MaskingStruct(v interface{}) string {
	output, err := Struct(v)
	if err != nil {
		zap.S().Errorf("MaskingStruct Error: %v", err)
		return ""
	}
	bytes, err := json.Marshal(output)
	if err != nil {
		zap.S().Errorf("MaskingStruct Error: %v", err)
		return ""
	}
	return string(bytes)
}
