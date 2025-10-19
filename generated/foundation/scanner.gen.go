// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Scanner] class.
var (
	scannerClass     _ScannerClass
	scannerClassOnce sync.Once
)

func getScannerClass() _ScannerClass {
	scannerClassOnce.Do(func() {
		scannerClass = _ScannerClass{objc.GetClass("NSScanner")}
	})
	return scannerClass
}

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

// A string parser that scans for substrings or characters in a character set, and for numeric values from decimal, hexadecimal, and floating-point representations.
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

// Alloc allocates a new instance without initialization.
func (sc _ScannerClass) Alloc() Scanner {
	rv := objc.Send[Scanner](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScannerClass) New() Scanner {
	rv := objc.Send[Scanner](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Scanner) Init() Scanner {
	rv := objc.Send[Scanner](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Scanner) Autorelease() Scanner {
	rv := objc.Send[Scanner](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScanner creates a new Scanner instance.
func NewScanner() Scanner {
	return getScannerClass().New()
}


// Scans for a double value, returning a found value by reference.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/scanDouble(_:)
func (s_ Scanner) ScanDouble(result unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("scanDouble:"), result)
	return rv
}
// Scans for a float value, returning a found value by reference.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/scanFloat(_:)
func (s_ Scanner) ScanFloat(result unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("scanFloat:"), result)
	return rv
}
// Scans for an int value from a decimal representation, returning a found value by reference.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/scanInt32(_:)
func (s_ Scanner) ScanInt(result unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("scanInt:"), result)
	return rv
}


