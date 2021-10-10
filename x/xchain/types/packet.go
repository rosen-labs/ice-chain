package types

import (
	"fmt"
	"strings"
)

func GetRevAddressAndUID(r string) (string, string, error) {
	s := strings.Split(r, "___")
	if len(s) != 2 {
		return "", "", fmt.Errorf("error wrong format reciever")
	}
	return s[0], s[1], nil
}
