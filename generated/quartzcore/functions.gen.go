// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// QuartzCore Functions (3 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CACurrentMediaTime func() unsafe.Pointer
	_CAFrameRateRangeIsEqualToRange func(unsafe.Pointer, unsafe.Pointer) bool
	_CATransform3DMakeScale func(float64, float64, float64) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_CACurrentMediaTime, lib, "CACurrentMediaTime")
	tryRegister(&_CAFrameRateRangeIsEqualToRange, lib, "CAFrameRateRangeIsEqualToRange")
	tryRegister(&_CATransform3DMakeScale, lib, "CATransform3DMakeScale")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Returns the current absolute time, in seconds.
//
// Added in macOS 10.5.
// Returns the current absolute time, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CACurrentMediaTime()
func CACurrentMediaTime() unsafe.Pointer {
	return _CACurrentMediaTime()
}

// CAFrameRateRangeIsEqualToRange is a QuartzCore function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAFrameRateRangeIsEqualToRange
func CAFrameRateRangeIsEqualToRange(range_ unsafe.Pointer, other unsafe.Pointer) bool {
	return _CAFrameRateRangeIsEqualToRange(range_, other)
}

// Returns a transform that scales by .
//
// Added in macOS 10.5.
// Returns a transform that scales by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DMakeScale(_:_:_:)
func CATransform3DMakeScale(sx float64, sy float64, sz float64) unsafe.Pointer {
	return _CATransform3DMakeScale(sx, sy, sz)
}



