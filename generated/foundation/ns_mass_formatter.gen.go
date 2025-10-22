// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MassFormatter] class.
var (
	MassFormatterClass     _MassFormatterClass
	MassFormatterClassOnce sync.Once
)

func getMassFormatterClass() _MassFormatterClass {
	MassFormatterClassOnce.Do(func() {
		MassFormatterClass = _MassFormatterClass{objc.GetClass("NSMassFormatter")}
	})
	return MassFormatterClass
}

type _MassFormatterClass struct {
	class objc.Class
}

// An interface definition for the [MassFormatter] class.
type IMassFormatter interface {
	IFormatter
	IsForPersonMassUse() bool
	SetIsForPersonMassUse(value bool)
	NumberFormatter() NSNumberFormatter
	SetNumberFormatter(value INumberFormatter)
	UnitStyle() unsafe.Pointer
	SetUnitStyle(value unsafe.Pointer)
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


// A Boolean value that indicates whether the resulting string represents a person’s mass.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/massformatter/isforpersonmassuse
func (m_ MassFormatter) IsForPersonMassUse() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isForPersonMassUse"))
	return rv
}


// SetIsForPersonMassUse sets the value of the isForPersonMassUse property.
// A Boolean value that indicates whether the resulting string represents a person’s mass.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/massformatter/isforpersonmassuse
func (m_ MassFormatter) SetIsForPersonMassUse(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsForPersonMassUse:"), value)
}

// The number formatter used to format the numbers in a mass strings.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/massformatter/numberformatter
func (m_ MassFormatter) NumberFormatter() NSNumberFormatter {
	rv := objc.Send[NSNumberFormatter](m_.ID, objc.Sel("numberFormatter"))
	return rv
}


// SetNumberFormatter sets the value of the numberFormatter property.
// The number formatter used to format the numbers in a mass strings.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/massformatter/numberformatter
func (m_ MassFormatter) SetNumberFormatter(value INumberFormatter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberFormatter:"), value)
}

// The unit style used by this formatter.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/massformatter/unitstyle
func (m_ MassFormatter) UnitStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("unitStyle"))
	return rv
}


// SetUnitStyle sets the value of the unitStyle property.
// The unit style used by this formatter.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/massformatter/unitstyle
func (m_ MassFormatter) SetUnitStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUnitStyle:"), value)
}



