package customvalidation

import (
	"errors"
	"unicode"

	validation "github.com/go-ozzo/ozzo-validation"
)

/*
	== GROUPS
	sometimes rules are used in multiple places,
	so we can group them to make it maintanable
	refer: https://github.com/go-ozzo/ozzo-validation#rule-groups
*/

var PasswordRequiredLength = []validation.Rule{
	validation.Required,
	validation.Length(8, 0),
}
var PasswordValidation = append(PasswordRequiredLength, validation.By(CheckPassword))

// == END OF GROUPS



/*	== CUSTOM RULES
	a place to store custom rules for validation purposes.
	refer: https://github.com/go-ozzo/ozzo-validation#creating-custom-rules
*/

func CheckPassword(pass interface{}) error {
	password, ok := pass.(string)
	if !ok {
		return errors.New("must be a string")
	}
	
	isLower := false
	isUpper := false
	isNumeric := false
	// isSymbol := false
	errs := map[string]string{
		"isLower" : "lowercase (a-z)",
		"isUpper" : "uppercase (A-Z)",
		"isNumeric" : "number (0-9)",
		// "isSymbol" : "symbol (.-,:; etc)",
	}
	
	for _, val := range password {
		if !isLower && unicode.IsLower(val) {
			isLower = true
			delete(errs, "isLower")
		}
		if !isUpper && unicode.IsUpper(val){
			isUpper = true
			delete(errs, "isUpper")
		}
		if !isNumeric && unicode.IsNumber(val) {
			isNumeric = true
			delete(errs, "isNumeric")
		}
		// if !isSymbol && unicode.IsSymbol(val) {
		// 	isSymbol = true
		// 	delete(errs, "isNumeric")
		// }
		if 
			isLower && 
			isUpper && 
			isNumeric &&
			// isSymbol &&
		true {
			return nil
		}
	}
	// parse to error
	var str string = "Must contain at least one: "
	for _, val := range errs{
		str = str + val + ", "
	}
	return errors.New(str[0:len(str)-2])
}

func CheckName(str interface{}) error{
	name, ok := str.(string)
	if !ok {
		return errors.New("must be string")
	}

	err := errors.New("must be a-z/A-Z/space/./-/'")
	for _, val := range name {
		if unicode.IsLower(val) || unicode.IsUpper(val){
			continue
		} else if val == '.' || val == '\'' || val == '-' || val == ' '{
			continue
		} else {
			return err
		}
	}
	return nil
}

// == END OF CUSTOM RULES