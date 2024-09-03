package extends

import (
	"fmt"
	"math/big"
)

type NiceBigInt struct {
	big.Int
}

func (n NiceBigInt) MarshalJSON() ([]byte, error) {
	return []byte(n.String()), nil
}

func (n *NiceBigInt) UnmarshalJSON(p []byte) error {
	if string(p) == "null" {
		return nil
	}
	var z big.Int
	_, ok := z.SetString(string(p), 10)
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", p)
	}
	n.Int = z
	return nil
}
