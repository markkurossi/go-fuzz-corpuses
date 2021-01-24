//
// Copyright (c) 2021 Markku Rossi
//
// All rights reserved.
//

package mpcl

import (
	"bytes"
	"os"

	"github.com/markkurossi/iql"
)

func Fuzz(data []byte) int {
	client := iql.NewClient(os.Stdout)
	err := client.Parse(bytes.NewReader(data), "data")
	if err != nil {
		return 0
	}
	return 1
}
