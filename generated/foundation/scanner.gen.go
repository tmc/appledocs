// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Scanner] class.
var scannerClass = _ScannerClass{objc.GetClass("NSScanner")}

type _ScannerClass struct {
	class objc.Class
}

// An interface definition for the [Scanner] class.
type IScanner interface {
	objectivec.IObject
	ScanDouble(result unsafe.Pointer) bool
	ScanFloat(result unsafe.Pointer) bool
	ScanInt(result unsafe.Pointer) bool
}

// A string parser that scans for substrings or characters in a character set, and for numeric values from decimal, hexadecimal, and floating-point representations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner

type Scanner struct {
	objectivec.Object
}

// ScannerFrom constructs a [Scanner] from an unsafe.Pointer.
//
// A string parser that scans for substrings or characters in a character set, and for numeric values from decimal, hexadecimal, and floating-point representations.
func ScannerFrom(ptr unsafe.Pointer) Scanner {
	return Scanner{objectivec.Object{objc.ID(ptr)}}
}

// Scans for a double value, returning a found value by reference. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/scanDouble(_:)
func (s_ Scanner) ScanDouble(result unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("scanDouble:"), result)
	return rv
}
// Scans for a float value, returning a found value by reference. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/scanFloat(_:)
func (s_ Scanner) ScanFloat(result unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("scanFloat:"), result)
	return rv
}
// Scans for an int value from a decimal representation, returning a found value by reference. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/scanInt32(_:)
func (s_ Scanner) ScanInt(result unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("scanInt:"), result)
	return rv
}


