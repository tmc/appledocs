// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RelativeDateTimeFormatter] class.
var (
	relativeDateTimeFormatterClass     _RelativeDateTimeFormatterClass
	relativeDateTimeFormatterClassOnce sync.Once
)

func getRelativeDateTimeFormatterClass() _RelativeDateTimeFormatterClass {
	relativeDateTimeFormatterClassOnce.Do(func() {
		relativeDateTimeFormatterClass = _RelativeDateTimeFormatterClass{objc.GetClass("NSRelativeDateTimeFormatter")}
	})
	return relativeDateTimeFormatterClass
}

type _RelativeDateTimeFormatterClass struct {
	class objc.Class
}

// An interface definition for the [RelativeDateTimeFormatter] class.
type IRelativeDateTimeFormatter interface {
	IFormatter
}

// A formatter that creates locale-aware string representations of a relative date or time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter
type RelativeDateTimeFormatter struct {
	Formatter
}

// RelativeDateTimeFormatterFrom constructs a [RelativeDateTimeFormatter] from an unsafe.Pointer.
//
// A formatter that creates locale-aware string representations of a relative date or time.
func RelativeDateTimeFormatterFrom(ptr unsafe.Pointer) RelativeDateTimeFormatter {
	return RelativeDateTimeFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RelativeDateTimeFormatterClass) Alloc() RelativeDateTimeFormatter {
	rv := objc.Send[RelativeDateTimeFormatter](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RelativeDateTimeFormatterClass) New() RelativeDateTimeFormatter {
	rv := objc.Send[RelativeDateTimeFormatter](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RelativeDateTimeFormatter) Init() RelativeDateTimeFormatter {
	rv := objc.Send[RelativeDateTimeFormatter](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RelativeDateTimeFormatter) Autorelease() RelativeDateTimeFormatter {
	rv := objc.Send[RelativeDateTimeFormatter](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRelativeDateTimeFormatter creates a new RelativeDateTimeFormatter instance.
func NewRelativeDateTimeFormatter() RelativeDateTimeFormatter {
	return getRelativeDateTimeFormatterClass().New()
}




