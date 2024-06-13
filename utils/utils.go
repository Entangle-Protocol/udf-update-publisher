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

func Map[T, V any](ts []T, fn func(T) V) []V {
    result := make([]V, len(ts))
    for i, t := range ts {
        result[i] = fn(t)
    }
    return result
}
