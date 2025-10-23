// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego"
	coregraphics "github.com/tmc/appledocs/generated/coregraphics"
)


// Foundation Functions (38 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_NSContainsRect func(coregraphics.CGRect, coregraphics.CGRect) bool
	_NSDecimalAdd func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalCompact func(unsafe.Pointer)
	_NSDecimalCompare func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalCopy func(unsafe.Pointer, unsafe.Pointer)
	_NSDecimalDivide func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalMultiply func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalMultiplyByPowerOf10 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalNormalize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalPower func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalRound func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_NSDecimalString func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalSubtract func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDivideRect func(coregraphics.CGRect, unsafe.Pointer, unsafe.Pointer, float64, unsafe.Pointer)
	_NSEdgeInsetsEqual func(unsafe.Pointer, unsafe.Pointer) bool
	_NSEqualPoints func(coregraphics.CGPoint, coregraphics.CGPoint) bool
	_NSEqualRects func(coregraphics.CGRect, coregraphics.CGRect) bool
	_NSEqualSizes func(coregraphics.CGSize, coregraphics.CGSize) bool
	_NSInsetRect func(coregraphics.CGRect, float64, float64) coregraphics.CGRect
	_NSIntegralRect func(coregraphics.CGRect) coregraphics.CGRect
	_NSIntegralRectWithOptions func(coregraphics.CGRect, unsafe.Pointer) coregraphics.CGRect
	_NSIntersectionRange func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSIntersectionRect func(coregraphics.CGRect, coregraphics.CGRect) coregraphics.CGRect
	_NSIntersectsRect func(coregraphics.CGRect, coregraphics.CGRect) bool
	_NSIsEmptyRect func(coregraphics.CGRect) bool
	_NSMouseInRect func(coregraphics.CGPoint, coregraphics.CGRect, bool) bool
	_NSOffsetRect func(coregraphics.CGRect, float64, float64) coregraphics.CGRect
	_NSPointFromString func(unsafe.Pointer) coregraphics.CGPoint
	_NSPointInRect func(coregraphics.CGPoint, coregraphics.CGRect) bool
	_NSRangeFromString func(unsafe.Pointer) unsafe.Pointer
	_NSRectFromString func(unsafe.Pointer) coregraphics.CGRect
	_NSSizeFromString func(unsafe.Pointer) coregraphics.CGSize
	_NSStringFromPoint func(coregraphics.CGPoint) unsafe.Pointer
	_NSStringFromRange func(unsafe.Pointer) unsafe.Pointer
	_NSStringFromRect func(coregraphics.CGRect) unsafe.Pointer
	_NSStringFromSize func(coregraphics.CGSize) unsafe.Pointer
	_NSUnionRange func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSUnionRect func(coregraphics.CGRect, coregraphics.CGRect) coregraphics.CGRect
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_NSContainsRect, lib, "NSContainsRect")
	tryRegister(&_NSDecimalAdd, lib, "NSDecimalAdd")
	tryRegister(&_NSDecimalCompact, lib, "NSDecimalCompact")
	tryRegister(&_NSDecimalCompare, lib, "NSDecimalCompare")
	tryRegister(&_NSDecimalCopy, lib, "NSDecimalCopy")
	tryRegister(&_NSDecimalDivide, lib, "NSDecimalDivide")
	tryRegister(&_NSDecimalMultiply, lib, "NSDecimalMultiply")
	tryRegister(&_NSDecimalMultiplyByPowerOf10, lib, "NSDecimalMultiplyByPowerOf10")
	tryRegister(&_NSDecimalNormalize, lib, "NSDecimalNormalize")
	tryRegister(&_NSDecimalPower, lib, "NSDecimalPower")
	tryRegister(&_NSDecimalRound, lib, "NSDecimalRound")
	tryRegister(&_NSDecimalString, lib, "NSDecimalString")
	tryRegister(&_NSDecimalSubtract, lib, "NSDecimalSubtract")
	tryRegister(&_NSDivideRect, lib, "NSDivideRect")
	tryRegister(&_NSEdgeInsetsEqual, lib, "NSEdgeInsetsEqual")
	tryRegister(&_NSEqualPoints, lib, "NSEqualPoints")
	tryRegister(&_NSEqualRects, lib, "NSEqualRects")
	tryRegister(&_NSEqualSizes, lib, "NSEqualSizes")
	tryRegister(&_NSInsetRect, lib, "NSInsetRect")
	tryRegister(&_NSIntegralRect, lib, "NSIntegralRect")
	tryRegister(&_NSIntegralRectWithOptions, lib, "NSIntegralRectWithOptions")
	tryRegister(&_NSIntersectionRange, lib, "NSIntersectionRange")
	tryRegister(&_NSIntersectionRect, lib, "NSIntersectionRect")
	tryRegister(&_NSIntersectsRect, lib, "NSIntersectsRect")
	tryRegister(&_NSIsEmptyRect, lib, "NSIsEmptyRect")
	tryRegister(&_NSMouseInRect, lib, "NSMouseInRect")
	tryRegister(&_NSOffsetRect, lib, "NSOffsetRect")
	tryRegister(&_NSPointFromString, lib, "NSPointFromString")
	tryRegister(&_NSPointInRect, lib, "NSPointInRect")
	tryRegister(&_NSRangeFromString, lib, "NSRangeFromString")
	tryRegister(&_NSRectFromString, lib, "NSRectFromString")
	tryRegister(&_NSSizeFromString, lib, "NSSizeFromString")
	tryRegister(&_NSStringFromPoint, lib, "NSStringFromPoint")
	tryRegister(&_NSStringFromRange, lib, "NSStringFromRange")
	tryRegister(&_NSStringFromRect, lib, "NSStringFromRect")
	tryRegister(&_NSStringFromSize, lib, "NSStringFromSize")
	tryRegister(&_NSUnionRange, lib, "NSUnionRange")
	tryRegister(&_NSUnionRect, lib, "NSUnionRect")
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



// Returns a Boolean value that indicates whether one rectangle completely encloses another.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether one rectangle completely encloses another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSContainsRect(_:_:)
func NSContainsRect(aRect coregraphics.CGRect, bRect coregraphics.CGRect) bool {
	return _NSContainsRect(aRect, bRect)
}

// Adds two decimal values.
//
// Added in macOS 10.0.
// Adds two decimal values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalAdd(_:_:_:_:)
func NSDecimalAdd(result unsafe.Pointer, leftOperand unsafe.Pointer, rightOperand unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalAdd(result, leftOperand, rightOperand, roundingMode)
}

// Compacts the decimal structure for efficiency.
//
// Added in macOS 10.0.
// Compacts the decimal structure for efficiency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalCompact(_:)
func NSDecimalCompact(number unsafe.Pointer) {
	_NSDecimalCompact(number)
}

// Compares two decimal values.
//
// Added in macOS 10.0.
// Compares two decimal values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalCompare(_:_:)
func NSDecimalCompare(leftOperand unsafe.Pointer, rightOperand unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalCompare(leftOperand, rightOperand)
}

// Copies the value of a decimal number.
//
// Added in macOS 10.0.
// Copies the value of a decimal number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalCopy(_:_:)
func NSDecimalCopy(destination unsafe.Pointer, source unsafe.Pointer) {
	_NSDecimalCopy(destination, source)
}

// Divides one decimal value by another.
//
// Added in macOS 10.0.
// Divides one decimal value by another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalDivide(_:_:_:_:)
func NSDecimalDivide(result unsafe.Pointer, leftOperand unsafe.Pointer, rightOperand unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalDivide(result, leftOperand, rightOperand, roundingMode)
}

// Multiplies two decimal numbers together.
//
// Added in macOS 10.0.
// Multiplies two decimal numbers together.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalMultiply(_:_:_:_:)
func NSDecimalMultiply(result unsafe.Pointer, leftOperand unsafe.Pointer, rightOperand unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalMultiply(result, leftOperand, rightOperand, roundingMode)
}

// Multiplies a decimal by the specified power of 10.
//
// Added in macOS 10.0.
// Multiplies a decimal by the specified power of 10.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalMultiplyByPowerOf10(_:_:_:_:)
func NSDecimalMultiplyByPowerOf10(result unsafe.Pointer, number unsafe.Pointer, power unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalMultiplyByPowerOf10(result, number, power, roundingMode)
}

// Normalizes the internal format of two decimal numbers to simplify later operations.
//
// Added in macOS 10.0.
// Normalizes the internal format of two decimal numbers to simplify later operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNormalize(_:_:_:)
func NSDecimalNormalize(number1 unsafe.Pointer, number2 unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalNormalize(number1, number2, roundingMode)
}

// Raises the decimal value to the specified power.
//
// Added in macOS 10.0.
// Raises the decimal value to the specified power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalPower(_:_:_:_:)
func NSDecimalPower(result unsafe.Pointer, number unsafe.Pointer, power unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalPower(result, number, power, roundingMode)
}

// Rounds off the decimal value.
//
// Added in macOS 10.0.
// Rounds off the decimal value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalRound(_:_:_:_:)
func NSDecimalRound(result unsafe.Pointer, number unsafe.Pointer, scale unsafe.Pointer, roundingMode unsafe.Pointer) {
	_NSDecimalRound(result, number, scale, roundingMode)
}

// Returns a string representation of the decimal value appropriate for the specified locale.
//
// Added in macOS 10.0.
// Returns a string representation of the decimal value appropriate for the specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalString(_:_:)
func NSDecimalString(dcm unsafe.Pointer, locale unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalString(dcm, locale)
}

// Subtracts one decimal value from another.
//
// Added in macOS 10.0.
// Subtracts one decimal value from another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalSubtract(_:_:_:_:)
func NSDecimalSubtract(result unsafe.Pointer, leftOperand unsafe.Pointer, rightOperand unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalSubtract(result, leftOperand, rightOperand, roundingMode)
}

// Divides a rectangle into two new rectangles.
//
// Added in macOS 10.0.
// Divides a rectangle into two new rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDivideRect(_:_:_:_:_:)
func NSDivideRect(inRect coregraphics.CGRect, slice unsafe.Pointer, rem unsafe.Pointer, amount float64, edge unsafe.Pointer) {
	_NSDivideRect(inRect, slice, rem, amount, edge)
}

// Returns a Boolean value that indicates whether two edge insets structures are equal.
//
// Added in macOS 10.10.
// Returns a Boolean value that indicates whether two edge insets structures are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEdgeInsetsEqual(_:_:)
func NSEdgeInsetsEqual(aInsets unsafe.Pointer, bInsets unsafe.Pointer) bool {
	return _NSEdgeInsetsEqual(aInsets, bInsets)
}

// Returns a Boolean value that indicates whether two points are equal.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether two points are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEqualPoints(_:_:)
func NSEqualPoints(aPoint coregraphics.CGPoint, bPoint coregraphics.CGPoint) bool {
	return _NSEqualPoints(aPoint, bPoint)
}

// Returns a Boolean value that indicates whether the two rectangles are equal.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether the two rectangles are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEqualRects(_:_:)
func NSEqualRects(aRect coregraphics.CGRect, bRect coregraphics.CGRect) bool {
	return _NSEqualRects(aRect, bRect)
}

// Returns a Boolean that indicates whether two size values are equal.
//
// Added in macOS 10.0.
// Returns a Boolean that indicates whether two size values are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEqualSizes(_:_:)
func NSEqualSizes(aSize coregraphics.CGSize, bSize coregraphics.CGSize) bool {
	return _NSEqualSizes(aSize, bSize)
}

// Insets a rectangle by a specified amount.
//
// Added in macOS 10.0.
// Insets a rectangle by a specified amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInsetRect(_:_:_:)
func NSInsetRect(aRect coregraphics.CGRect, dX float64, dY float64) coregraphics.CGRect {
	return _NSInsetRect(aRect, dX, dY)
}

// Adjusts the sides of a rectangle to integer values.
//
// Added in macOS 10.0.
// Adjusts the sides of a rectangle to integer values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIntegralRect(_:)
func NSIntegralRect(aRect coregraphics.CGRect) coregraphics.CGRect {
	return _NSIntegralRect(aRect)
}

// Adjusts the sides of a rectangle to integral values using the specified options.
//
// Added in macOS 10.7.
// Adjusts the sides of a rectangle to integral values using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIntegralRectWithOptions(_:_:)
func NSIntegralRectWithOptions(aRect coregraphics.CGRect, opts unsafe.Pointer) coregraphics.CGRect {
	return _NSIntegralRectWithOptions(aRect, opts)
}

// Returns the intersection of the specified ranges.
//
// Added in macOS 10.0.
// Returns the intersection of the specified ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIntersectionRange(_:_:)
func NSIntersectionRange(range1 unsafe.Pointer, range2 unsafe.Pointer) unsafe.Pointer {
	return _NSIntersectionRange(range1, range2)
}

// Calculates the intersection of two rectangles.
//
// Added in macOS 10.0.
// Calculates the intersection of two rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIntersectionRect(_:_:)
func NSIntersectionRect(aRect coregraphics.CGRect, bRect coregraphics.CGRect) coregraphics.CGRect {
	return _NSIntersectionRect(aRect, bRect)
}

// Returns a Boolean value that indicates whether two rectangles intersect.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether two rectangles intersect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIntersectsRect(_:_:)
func NSIntersectsRect(aRect coregraphics.CGRect, bRect coregraphics.CGRect) bool {
	return _NSIntersectsRect(aRect, bRect)
}

// Returns a Boolean value that indicates whether a given rectangle is empty.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether a given rectangle is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIsEmptyRect(_:)
func NSIsEmptyRect(aRect coregraphics.CGRect) bool {
	return _NSIsEmptyRect(aRect)
}

// Returns a Boolean value that indicates whether the point is in the specified rectangle.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether the point is in the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMouseInRect(_:_:_:)
func NSMouseInRect(aPoint coregraphics.CGPoint, aRect coregraphics.CGRect, flipped bool) bool {
	return _NSMouseInRect(aPoint, aRect, flipped)
}

// Offsets the rectangle by the specified amount.
//
// Added in macOS 10.0.
// Offsets the rectangle by the specified amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOffsetRect(_:_:_:)
func NSOffsetRect(aRect coregraphics.CGRect, dX float64, dY float64) coregraphics.CGRect {
	return _NSOffsetRect(aRect, dX, dY)
}

// Returns a point from a text-based representation.
//
// Added in macOS 10.0.
// Returns a point from a text-based representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointFromString(_:)
func NSPointFromString(aString unsafe.Pointer) coregraphics.CGPoint {
	return _NSPointFromString(aString)
}

// Returns a Boolean value that indicates whether a given point is in a given rectangle.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether a given point is in a given rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointInRect(_:_:)
func NSPointInRect(aPoint coregraphics.CGPoint, aRect coregraphics.CGRect) bool {
	return _NSPointInRect(aPoint, aRect)
}

// Returns a range from a textual representation.
//
// Added in macOS 10.0.
// Returns a range from a textual representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRangeFromString(_:)
func NSRangeFromString(aString unsafe.Pointer) unsafe.Pointer {
	return _NSRangeFromString(aString)
}

// Returns a rectangle from a text-based representation.
//
// Added in macOS 10.0.
// Returns a rectangle from a text-based representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRectFromString(_:)
func NSRectFromString(aString unsafe.Pointer) coregraphics.CGRect {
	return _NSRectFromString(aString)
}

// Returns an from a text-based representation.
//
// Added in macOS 10.0.
// Returns an from a text-based representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSizeFromString(_:)
func NSSizeFromString(aString unsafe.Pointer) coregraphics.CGSize {
	return _NSSizeFromString(aString)
}

// Returns a string representation of a point.
//
// Added in macOS 10.0.
// Returns a string representation of a point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromPoint(_:)
func NSStringFromPoint(aPoint coregraphics.CGPoint) unsafe.Pointer {
	return _NSStringFromPoint(aPoint)
}

// Returns a string representation of a range.
//
// Added in macOS 10.0.
// Returns a string representation of a range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromRange(_:)
func NSStringFromRange(range_ unsafe.Pointer) unsafe.Pointer {
	return _NSStringFromRange(range_)
}

// Returns a string representation of a rectangle.
//
// Added in macOS 10.0.
// Returns a string representation of a rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromRect(_:)
func NSStringFromRect(aRect coregraphics.CGRect) unsafe.Pointer {
	return _NSStringFromRect(aRect)
}

// Returns a string representation of a size.
//
// Added in macOS 10.0.
// Returns a string representation of a size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromSize(_:)
func NSStringFromSize(aSize coregraphics.CGSize) unsafe.Pointer {
	return _NSStringFromSize(aSize)
}

// Returns the union of the specified ranges.
//
// Added in macOS 10.0.
// Returns the union of the specified ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnionRange(_:_:)
func NSUnionRange(range1 unsafe.Pointer, range2 unsafe.Pointer) unsafe.Pointer {
	return _NSUnionRange(range1, range2)
}

// Calculates the union of two rectangles.
//
// Added in macOS 10.0.
// Calculates the union of two rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnionRect(_:_:)
func NSUnionRect(aRect coregraphics.CGRect, bRect coregraphics.CGRect) coregraphics.CGRect {
	return _NSUnionRect(aRect, bRect)
}



