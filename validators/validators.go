package validators

import "regexp"

// Regex to validate CPF or CNPJ
func IsCpfOrCnpj(e string) bool {
	regex, _ := regexp.Compile(`([0-9]{2}[\.]?[0-9]{3}[\.]?[0-9]{3}[\/]?[0-9]{4}[-]?[0-9]{2})|([0-9]{3}[\.]?[0-9]{3}[\.]?[0-9]{3}[-]?[0-9]{2})`)
	return regex.MatchString(e)
}

// Regex to validate if is 0 or 1
func IsZeroOrOne(e string) bool {
	regex, _ := regexp.Compile("[0-1]{1}")
	return regex.MatchString(e)
}

// Regex to validate fomart Date
func IsDate(e string) bool {
	regex, _ := regexp.Compile(`\d{4}-\d{2}-\d{2}`)
	return regex.MatchString(e)
}

// Regex to validate money in Brazilian Real
func IsMonetaryReal(e string) bool {
	regex, _ := regexp.Compile(`^[1-9]\d{0,2}(\.\d{3})*,\d{2}$`)
	return regex.MatchString(e)
}
