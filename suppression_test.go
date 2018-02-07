package main

import (
	"testing"
)

func Test_isMd5(t *testing.T) {

	tests := []string{
		"4A8A08F09D37B73795649038408B5F33",
		"8277E0910D750195B448797616E091AD",
		"E1671797C52E15F763380B45E841EC32",
		"brice@gmail.com",
		"gmail.com",
		"*.tld",
	}

	for _, v := range tests {
		hashedTrimmed := forceMd5(v)
		if !md5Regex.MatchString(hashedTrimmed) {
			t.Error(hashedTrimmed + " is not Md5")
		}
	}

}

func Test_main(t *testing.T) {
	t.Error("")
}
