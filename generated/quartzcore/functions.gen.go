// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/ebitengine/purego"
	coregraphics "github.com/tmc/appledocs/generated/coregraphics"
)


// QuartzCore Functions (16 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CACurrentMediaTime func() unsafe.Pointer
	_CAFrameRateRangeIsEqualToRange func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CAFrameRateRangeMake func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CATransform3DConcat func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CATransform3DEqualToTransform func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CATransform3DGetAffineTransform func(unsafe.Pointer) coregraphics.CGAffineTransform
	_CATransform3DInvert func(unsafe.Pointer) unsafe.Pointer
	_CATransform3DIsAffine func(unsafe.Pointer) unsafe.Pointer
	_CATransform3DIsIdentity func(unsafe.Pointer) unsafe.Pointer
	_CATransform3DMakeAffineTransform func(coregraphics.CGAffineTransform) unsafe.Pointer
	_CATransform3DMakeRotation func(float64, float64, float64, float64) unsafe.Pointer
	_CATransform3DMakeScale func(float64, float64, float64) unsafe.Pointer
	_CATransform3DMakeTranslation func(float64, float64, float64) unsafe.Pointer
	_CATransform3DRotate func(unsafe.Pointer, float64, float64, float64, float64) unsafe.Pointer
	_CATransform3DScale func(unsafe.Pointer, float64, float64, float64) unsafe.Pointer
	_CATransform3DTranslate func(unsafe.Pointer, float64, float64, float64) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_CACurrentMediaTime, lib, "CACurrentMediaTime")
	tryRegister(&_CAFrameRateRangeIsEqualToRange, lib, "CAFrameRateRangeIsEqualToRange")
	tryRegister(&_CAFrameRateRangeMake, lib, "CAFrameRateRangeMake")
	tryRegister(&_CATransform3DConcat, lib, "CATransform3DConcat")
	tryRegister(&_CATransform3DEqualToTransform, lib, "CATransform3DEqualToTransform")
	tryRegister(&_CATransform3DGetAffineTransform, lib, "CATransform3DGetAffineTransform")
	tryRegister(&_CATransform3DInvert, lib, "CATransform3DInvert")
	tryRegister(&_CATransform3DIsAffine, lib, "CATransform3DIsAffine")
	tryRegister(&_CATransform3DIsIdentity, lib, "CATransform3DIsIdentity")
	tryRegister(&_CATransform3DMakeAffineTransform, lib, "CATransform3DMakeAffineTransform")
	tryRegister(&_CATransform3DMakeRotation, lib, "CATransform3DMakeRotation")
	tryRegister(&_CATransform3DMakeScale, lib, "CATransform3DMakeScale")
	tryRegister(&_CATransform3DMakeTranslation, lib, "CATransform3DMakeTranslation")
	tryRegister(&_CATransform3DRotate, lib, "CATransform3DRotate")
	tryRegister(&_CATransform3DScale, lib, "CATransform3DScale")
	tryRegister(&_CATransform3DTranslate, lib, "CATransform3DTranslate")
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



// Returns the current absolute time, in seconds. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CACurrentMediaTime()
func CACurrentMediaTime() unsafe.Pointer {
	return _CACurrentMediaTime()
	}


// CAFrameRateRangeIsEqualToRange is a QuartzCore function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAFrameRateRangeIsEqualToRange
func CAFrameRateRangeIsEqualToRange(range_ unsafe.Pointer, other unsafe.Pointer) unsafe.Pointer {
	return _CAFrameRateRangeIsEqualToRange(range_, other)
	}


// CAFrameRateRangeMake is a QuartzCore function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAFrameRateRangeMake
func CAFrameRateRangeMake(minimum unsafe.Pointer, maximum unsafe.Pointer, preferred unsafe.Pointer) unsafe.Pointer {
	return _CAFrameRateRangeMake(minimum, maximum, preferred)
	}


// Concatenates to and returns the result: . [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DConcat(_:_:)
func CATransform3DConcat(a unsafe.Pointer, b unsafe.Pointer) unsafe.Pointer {
	return _CATransform3DConcat(a, b)
	}


// Returns a Boolean value that indicates whether the two transforms are exactly equal. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DEqualToTransform(_:_:)
func CATransform3DEqualToTransform(a unsafe.Pointer, b unsafe.Pointer) unsafe.Pointer {
	return _CATransform3DEqualToTransform(a, b)
	}


// Returns the affine transform represented by . [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DGetAffineTransform(_:)
func CATransform3DGetAffineTransform(t unsafe.Pointer) coregraphics.CGAffineTransform {
	return _CATransform3DGetAffineTransform(t)
	}


// Inverts and returns the result. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DInvert(_:)
func CATransform3DInvert(t unsafe.Pointer) unsafe.Pointer {
	return _CATransform3DInvert(t)
	}


// Returns a Boolean value that indicates whether a transform can be exactly represented by an affine transform. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DIsAffine(_:)
func CATransform3DIsAffine(t unsafe.Pointer) unsafe.Pointer {
	return _CATransform3DIsAffine(t)
	}


// Returns a Boolean value that indicates whether the transform is the identity transform. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DIsIdentity(_:)
func CATransform3DIsIdentity(t unsafe.Pointer) unsafe.Pointer {
	return _CATransform3DIsIdentity(t)
	}


// Returns a transform with the same effect as affine transform . [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DMakeAffineTransform(_:)
func CATransform3DMakeAffineTransform(m coregraphics.CGAffineTransform) unsafe.Pointer {
	return _CATransform3DMakeAffineTransform(m)
	}


// Returns a transform that rotates by radians about the vector . [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DMakeRotation(_:_:_:_:)
func CATransform3DMakeRotation(angle float64, x float64, y float64, z float64) unsafe.Pointer {
	return _CATransform3DMakeRotation(angle, x, y, z)
	}


// Returns a transform that scales by . [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DMakeScale(_:_:_:)
func CATransform3DMakeScale(sx float64, sy float64, sz float64) unsafe.Pointer {
	return _CATransform3DMakeScale(sx, sy, sz)
	}


// Returns a transform that translates by . [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DMakeTranslation(_:_:_:)
func CATransform3DMakeTranslation(tx float64, ty float64, tz float64) unsafe.Pointer {
	return _CATransform3DMakeTranslation(tx, ty, tz)
	}


// Rotates by radians about the vector and returns the result. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DRotate(_:_:_:_:_:)
func CATransform3DRotate(t unsafe.Pointer, angle float64, x float64, y float64, z float64) unsafe.Pointer {
	return _CATransform3DRotate(t, angle, x, y, z)
	}


// Scales by and returns the result: . [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DScale(_:_:_:_:)
func CATransform3DScale(t unsafe.Pointer, sx float64, sy float64, sz float64) unsafe.Pointer {
	return _CATransform3DScale(t, sx, sy, sz)
	}


// Translates by and returns the result: . [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DTranslate(_:_:_:_:)
func CATransform3DTranslate(t unsafe.Pointer, tx float64, ty float64, tz float64) unsafe.Pointer {
	return _CATransform3DTranslate(t, tx, ty, tz)
	}




