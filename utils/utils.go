package utils

import (
	"fmt"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

func AsciiToPaddedHex(str string) ([32]byte, error) {
	if len(str) > 32 {
		return [32]byte{}, fmt.Errorf("Can't pad string that is > 32 bytes")
	}

	buf := make([]byte, 32)
	copy(buf, str)

	return ethcommon.BytesToHash(buf), nil
}
