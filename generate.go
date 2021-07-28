//generate password based on passed args, return string of password

import strings

const (
	//lower and upper case latin chars, may wish to add localization here in the future
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := strings.ToUpper(lowerCase)
	Digits := "0123456789"
	//Punctation chars, yes passwords can and should include spaces ;-) 
	Punctuation := "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./' "
)