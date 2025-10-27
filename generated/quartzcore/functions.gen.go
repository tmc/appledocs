// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore


import (
	"unsafe"

	"github.com/ebitengine/purego"
	corefoundation "github.com/tmc/appledocs/generated/corefoundation"
)


// QuartzCore Functions (16 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CACurrentMediaTime func() TimeInterval
	_CAFrameRateRangeIsEqualToRange func(FrameRateRange, FrameRateRange) bool
	_CAFrameRateRangeMake func(float32, float32, float32) FrameRateRange
	_CATransform3DConcat func(Transform3D, Transform3D) Transform3D
	_CATransform3DEqualToTransform func(Transform3D, Transform3D) bool
	_CATransform3DGetAffineTransform func(Transform3D) corefoundation.CGAffineTransform
	_CATransform3DInvert func(Transform3D) Transform3D
	_CATransform3DIsAffine func(Transform3D) bool
	_CATransform3DIsIdentity func(Transform3D) bool
	_CATransform3DMakeAffineTransform func(corefoundation.CGAffineTransform) Transform3D
	_CATransform3DMakeRotation func(float64, float64, float64, float64) Transform3D
	_CATransform3DMakeScale func(float64, float64, float64) Transform3D
	_CATransform3DMakeTranslation func(float64, float64, float64) Transform3D
	_CATransform3DRotate func(Transform3D, float64, float64, float64, float64) Transform3D
	_CATransform3DScale func(Transform3D, float64, float64, float64) Transform3D
	_CATransform3DTranslate func(Transform3D, float64, float64, float64) Transform3D
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



// Returns the current absolute time, in seconds.
//
// Added in macOS 10.5.
// Returns the current absolute time, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CACurrentMediaTime()
func CACurrentMediaTime() TimeInterval {
	return _CACurrentMediaTime()
}

// CAFrameRateRangeIsEqualToRange is a QuartzCore function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAFrameRateRangeIsEqualToRange
func CAFrameRateRangeIsEqualToRange(range_ FrameRateRange, other FrameRateRange) bool {
	return _CAFrameRateRangeIsEqualToRange(range_, other)
}

// CAFrameRateRangeMake is a QuartzCore function.
//
// Added in macOS 12.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAFrameRateRangeMake
func CAFrameRateRangeMake(minimum float32, maximum float32, preferred float32) FrameRateRange {
	return _CAFrameRateRangeMake(minimum, maximum, preferred)
}

// Concatenates to and returns the result: .
//
// Added in macOS 10.5.
// Concatenates to and returns the result: .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DConcat(_:_:)
func CATransform3DConcat(a Transform3D, b Transform3D) Transform3D {
	return _CATransform3DConcat(a, b)
}

// Returns a Boolean value that indicates whether the two transforms are exactly equal.
//
// Added in macOS 10.5.
// Returns a Boolean value that indicates whether the two transforms are exactly equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DEqualToTransform(_:_:)
func CATransform3DEqualToTransform(a Transform3D, b Transform3D) bool {
	return _CATransform3DEqualToTransform(a, b)
}

// Returns the affine transform represented by .
//
// Added in macOS 10.5.
// Returns the affine transform represented by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DGetAffineTransform(_:)
func CATransform3DGetAffineTransform(t Transform3D) corefoundation.CGAffineTransform {
	return _CATransform3DGetAffineTransform(t)
}

// Inverts and returns the result.
//
// Added in macOS 10.5.
// Inverts and returns the result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DInvert(_:)
func CATransform3DInvert(t Transform3D) Transform3D {
	return _CATransform3DInvert(t)
}

// Returns a Boolean value that indicates whether a transform can be exactly represented by an affine transform.
//
// Added in macOS 10.5.
// Returns a Boolean value that indicates whether a transform can be exactly represented by an affine transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DIsAffine(_:)
func CATransform3DIsAffine(t Transform3D) bool {
	return _CATransform3DIsAffine(t)
}

// Returns a Boolean value that indicates whether the transform is the identity transform.
//
// Added in macOS 10.5.
// Returns a Boolean value that indicates whether the transform is the identity transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DIsIdentity(_:)
func CATransform3DIsIdentity(t Transform3D) bool {
	return _CATransform3DIsIdentity(t)
}

// Returns a transform with the same effect as affine transform .
//
// Added in macOS 10.5.
// Returns a transform with the same effect as affine transform .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DMakeAffineTransform(_:)
func CATransform3DMakeAffineTransform(m corefoundation.CGAffineTransform) Transform3D {
	return _CATransform3DMakeAffineTransform(m)
}

// Returns a transform that rotates by radians about the vector .
//
// Added in macOS 10.5.
// Returns a transform that rotates by radians about the vector .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DMakeRotation(_:_:_:_:)
func CATransform3DMakeRotation(angle float64, x float64, y float64, z float64) Transform3D {
	return _CATransform3DMakeRotation(angle, x, y, z)
}

// Returns a transform that scales by .
//
// Added in macOS 10.5.
// Returns a transform that scales by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DMakeScale(_:_:_:)
func CATransform3DMakeScale(sx float64, sy float64, sz float64) Transform3D {
	return _CATransform3DMakeScale(sx, sy, sz)
}

// Returns a transform that translates by .
//
// Added in macOS 10.5.
// Returns a transform that translates by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DMakeTranslation(_:_:_:)
func CATransform3DMakeTranslation(tx float64, ty float64, tz float64) Transform3D {
	return _CATransform3DMakeTranslation(tx, ty, tz)
}

// Rotates by radians about the vector and returns the result.
//
// Added in macOS 10.5.
// Rotates by radians about the vector and returns the result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DRotate(_:_:_:_:_:)
func CATransform3DRotate(t Transform3D, angle float64, x float64, y float64, z float64) Transform3D {
	return _CATransform3DRotate(t, angle, x, y, z)
}

// Scales by and returns the result: .
//
// Added in macOS 10.5.
// Scales by and returns the result: .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DScale(_:_:_:_:)
func CATransform3DScale(t Transform3D, sx float64, sy float64, sz float64) Transform3D {
	return _CATransform3DScale(t, sx, sy, sz)
}

// Translates by and returns the result: .
//
// Added in macOS 10.5.
// Translates by and returns the result: .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransform3DTranslate(_:_:_:_:)
func CATransform3DTranslate(t Transform3D, tx float64, ty float64, tz float64) Transform3D {
	return _CATransform3DTranslate(t, tx, ty, tz)
}




