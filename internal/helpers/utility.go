package helpers

import (
	"fmt"
	"math/rand"
	"time"

	tbTypes "github.com/tigerbeetledb/tigerbeetle-go/pkg/types"
)

func Uint128(value string) tbTypes.Uint128 {
	x, err := tbTypes.HexStringToUint128(value)
	if err != nil {
		panic(err)
	}
	return x
}

// GetRandID actually only returns a 64-bit integer but that's fine for out purposes here
func GetRandID() tbTypes.Uint128 {
	rand.Seed(time.Now().UnixNano())

	return Uint128(fmt.Sprintf("%d", rand.Uint64()))
}
