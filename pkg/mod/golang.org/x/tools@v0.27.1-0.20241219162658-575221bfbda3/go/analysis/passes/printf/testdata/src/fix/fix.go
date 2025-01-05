// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains tests of the printf checker's suggested fixes.

package fix

import (
	"fmt"
	"log"
	"os"
)

func nonConstantFormat(s string) { // #60529
	fmt.Printf(s) // want `non-constant format string in call to fmt.Printf`
	fmt.Printf(s, "arg")
	fmt.Fprintf(os.Stderr, s) // want `non-constant format string in call to fmt.Fprintf`
	log.Printf(s)             // want `non-constant format string in call to log.Printf`
}
