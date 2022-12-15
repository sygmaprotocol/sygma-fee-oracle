// The Licensed Work is (c) 2022 Sygma
// SPDX-License-Identifier: BUSL-1.1

package util

import (
	"regexp"
)

var checkIntegerRegex = regexp.MustCompile(`^[0-9]+$`)

func CheckInteger(integerString string) bool {
	return checkIntegerRegex.MatchString(integerString)
}
