//
// Copyright (c) 2021 Markku Rossi
//
// All rights reserved.
//

package mpcl

import (
	"bytes"

	"github.com/markkurossi/mpc/circuit"
)

func Fuzz(data []byte) int {
	circ, err := circuit.ParseMPCLC(bytes.NewReader(data))
	if err != nil {
		if circ != nil {
			panic("circ != nil on error")
		}
		return 0
	}
	return 1
}
