package crypto_utils

import (
	"fmt"
	"testing"
)

func TestGetSha256(t *testing.T) {
	fmt.Println(GetSha256("123"))
}