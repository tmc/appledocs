// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ISO8601DateFormatter] class.
var iSO8601DateFormatterClass = _ISO8601DateFormatterClass{objc.GetClass("NSISO8601DateFormatter")}

type _ISO8601DateFormatterClass struct {
	class objc.Class
}

// An interface definition for the [ISO8601DateFormatter] class.
type IISO8601DateFormatter interface {
	IFormatter
}

// A formatter that converts between dates and their ISO 8601 string representations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter

type ISO8601DateFormatter struct {
	Formatter
}

// ISO8601DateFormatterFrom constructs a [ISO8601DateFormatter] from an unsafe.Pointer.
//
// A formatter that converts between dates and their ISO 8601 string representations.
func ISO8601DateFormatterFrom(ptr unsafe.Pointer) ISO8601DateFormatter {
	return ISO8601DateFormatter{
		Formatter: FormatterFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (ic _ISO8601DateFormatterClass) Alloc() ISO8601DateFormatter {
	rv := objc.Send[ISO8601DateFormatter](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ic _ISO8601DateFormatterClass) New() ISO8601DateFormatter {
	rv := objc.Send[ISO8601DateFormatter](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ISO8601DateFormatter) Init() ISO8601DateFormatter {
	rv := objc.Send[ISO8601DateFormatter](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ISO8601DateFormatter) Autorelease() ISO8601DateFormatter {
	rv := objc.Send[ISO8601DateFormatter](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewISO8601DateFormatter creates a new ISO8601DateFormatter instance.
func NewISO8601DateFormatter() ISO8601DateFormatter {
	return iSO8601DateFormatterClass.New()
}




