// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [LengthFormatter] class.
var (
	lengthFormatterClass     _LengthFormatterClass
	lengthFormatterClassOnce sync.Once
)

func getLengthFormatterClass() _LengthFormatterClass {
	lengthFormatterClassOnce.Do(func() {
		lengthFormatterClass = _LengthFormatterClass{objc.GetClass("NSLengthFormatter")}
	})
	return lengthFormatterClass
}

type _LengthFormatterClass struct {
	class objc.Class
}

// An interface definition for the [LengthFormatter] class.
type ILengthFormatter interface {
	IFormatter
}

// A formatter that provides localized descriptions of linear distances, such as length and height measurements.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter
type LengthFormatter struct {
	Formatter
}

// LengthFormatterFrom constructs a [LengthFormatter] from an unsafe.Pointer.
//
// A formatter that provides localized descriptions of linear distances, such as length and height measurements.
func LengthFormatterFrom(ptr unsafe.Pointer) LengthFormatter {
	return LengthFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (lc _LengthFormatterClass) Alloc() LengthFormatter {
	rv := objc.Send[LengthFormatter](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LengthFormatterClass) New() LengthFormatter {
	rv := objc.Send[LengthFormatter](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LengthFormatter) Init() LengthFormatter {
	rv := objc.Send[LengthFormatter](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LengthFormatter) Autorelease() LengthFormatter {
	rv := objc.Send[LengthFormatter](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLengthFormatter creates a new LengthFormatter instance.
func NewLengthFormatter() LengthFormatter {
	return getLengthFormatterClass().New()
}




