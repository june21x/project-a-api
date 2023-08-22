package util

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/validator/v10"

	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var errorTranslator ut.Translator

func InitErrorTranslator() ut.Translator {
	en := en.New()
	uni := ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")
	errorTranslator = trans

	return errorTranslator
}

func RegisterErrorTranslation(trans ut.Translator, v *validator.Validate) {
	en_translations.RegisterDefaultTranslations(v, trans)
	translateOverride(trans, v)
}

func translateOverride(trans ut.Translator, v *validator.Validate) {
	v.RegisterTranslation("gte", trans, func(ut ut.Translator) error {
		return ut.Add("gte", "{0} must be greater than or equal to {1}!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("gte", fe.Field(), fe.Param())
		return t
	})
}
