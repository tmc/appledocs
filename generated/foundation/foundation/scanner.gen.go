// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Scanner] class.
var ScannerClass objc.Class

func init() {
	ScannerClass = objc.GetClass("NSScanner")
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
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Scanner/scanDouble(_:)
func (s_ Scanner) ScanDouble(result unsafe.Pointer) bool {
	sel := objc.RegisterName("scanDouble:")
	ret := s_.ID.Send(sel, result)
	return ret != 0
}
// Scans for a float value, returning a found value by reference. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Scanner/scanFloat(_:)
func (s_ Scanner) ScanFloat(result unsafe.Pointer) bool {
	sel := objc.RegisterName("scanFloat:")
	ret := s_.ID.Send(sel, result)
	return ret != 0
}
// Scans for an int value from a decimal representation, returning a found value by reference. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/Scanner/scanInt32(_:)
func (s_ Scanner) ScanInt(result unsafe.Pointer) bool {
	sel := objc.RegisterName("scanInt:")
	ret := s_.ID.Send(sel, result)
	return ret != 0
}


