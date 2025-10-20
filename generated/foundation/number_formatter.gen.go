// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NumberFormatter] class.
var (
	numberFormatterClass     _NumberFormatterClass
	numberFormatterClassOnce sync.Once
)

func getNumberFormatterClass() _NumberFormatterClass {
	numberFormatterClassOnce.Do(func() {
		numberFormatterClass = _NumberFormatterClass{objc.GetClass("NSNumberFormatter")}
	})
	return numberFormatterClass
}

type _NumberFormatterClass struct {
	class objc.Class
}

// An interface definition for the [NumberFormatter] class.
type INumberFormatter interface {
	IFormatter
}

// A formatter that converts between numeric values and their textual representations.
//
// Instances of format the textual representation of cells that contain objects and convert textual representations of numeric values into objects. The representation encompasses integers, floats, and doubles; floats and doubles can be formatted to a specified decimal position. objects can also impose ranges on the numeric values cells can accept.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter
type NumberFormatter struct {
	Formatter
}

// NumberFormatterFrom constructs a [NumberFormatter] from an unsafe.Pointer.
//
// A formatter that converts between numeric values and their textual representations.
func NumberFormatterFrom(ptr unsafe.Pointer) NumberFormatter {
	return NumberFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NumberFormatterClass) Alloc() NumberFormatter {
	rv := objc.Send[NumberFormatter](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NumberFormatterClass) New() NumberFormatter {
	rv := objc.Send[NumberFormatter](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NumberFormatter) Init() NumberFormatter {
	rv := objc.Send[NumberFormatter](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NumberFormatter) Autorelease() NumberFormatter {
	rv := objc.Send[NumberFormatter](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNumberFormatter creates a new NumberFormatter instance.
func NewNumberFormatter() NumberFormatter {
	return getNumberFormatterClass().New()
}




