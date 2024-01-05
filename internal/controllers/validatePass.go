package controllers

import "bytes"

func validatePass(password1 []byte, password2 []byte) bool {
	if bytes.Equal(password1, password2) {
		return true
	} else {
		return false
	}
}
