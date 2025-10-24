// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

/* debug [functions.gen.go]: Generating 0 functions for CoreLocation */
import (
	"github.com/ebitengine/purego"
)

// CoreLocation Functions (0 total)
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
