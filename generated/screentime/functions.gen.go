// Code generated from Apple documentation for ScreenTime. DO NOT EDIT.

package screentime

import (
	"github.com/ebitengine/purego"
)


// ScreenTime Functions (0 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

func init() {
	// Framework has no exported C functions, only classes/protocols
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	_ = lib // Suppress unused variable warning
}





