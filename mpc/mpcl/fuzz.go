//
// Copyright (c) 2021 Markku Rossi
//
// All rights reserved.
//

package mpcl

import (
	"github.com/markkurossi/mpc/compiler"
	"github.com/markkurossi/mpc/compiler/utils"
)

func Fuzz(data []byte) int {
	circ, _, err := compiler.New(utils.NewParams()).Compile(string(data))
	if err != nil {
		if circ != nil {
			panic("circ != nil on error")
		}
		return 0
	}
	return 1
}
