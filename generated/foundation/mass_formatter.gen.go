// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MassFormatter] class.
var (
	massFormatterClass     _MassFormatterClass
	massFormatterClassOnce sync.Once
)

func getMassFormatterClass() _MassFormatterClass {
	massFormatterClassOnce.Do(func() {
		massFormatterClass = _MassFormatterClass{objc.GetClass("NSMassFormatter")}
	})
	return massFormatterClass
}

type _MassFormatterClass struct {
	class objc.Class
}

// An interface definition for the [MassFormatter] class.
type IMassFormatter interface {
	IFormatter
}

// A formatter that provides localized descriptions of mass and weight values.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter
type MassFormatter struct {
	Formatter
}

// MassFormatterFrom constructs a [MassFormatter] from an unsafe.Pointer.
//
// A formatter that provides localized descriptions of mass and weight values.
func MassFormatterFrom(ptr unsafe.Pointer) MassFormatter {
	return MassFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MassFormatterClass) Alloc() MassFormatter {
	rv := objc.Send[MassFormatter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MassFormatterClass) New() MassFormatter {
	rv := objc.Send[MassFormatter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MassFormatter) Init() MassFormatter {
	rv := objc.Send[MassFormatter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MassFormatter) Autorelease() MassFormatter {
	rv := objc.Send[MassFormatter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMassFormatter creates a new MassFormatter instance.
func NewMassFormatter() MassFormatter {
	return getMassFormatterClass().New()
}




