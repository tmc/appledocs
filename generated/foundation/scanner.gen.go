// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Scanner] class.
var ScannerClass = _ScannerClass{objc.GetClass("NSScanner")}

type _ScannerClass struct {
	class objc.Class
}

type Scanner struct {
	objc.ID
}

func ScannerFrom(ptr unsafe.Pointer) Scanner {
	return Scanner{
		ID: objc.ID(ptr),
	}
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


