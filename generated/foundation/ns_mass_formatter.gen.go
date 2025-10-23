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
	// properties:
	UnitStyle() FormattingUnitStyle /* not a class type */
	SetUnitStyle(value FormattingUnitStyle /* not a class type */)
	IsForPersonMassUse() bool /* primitive/slice/pointer */
	SetIsForPersonMassUse(value bool /* primitive/slice/pointer */)
	NumberFormatter() INumberFormatter
	SetNumberFormatter(value INumberFormatter)
	// methods:
}

// A formatter that provides localized descriptions of mass and weight values.


// A formatter that provides localized descriptions of mass and weight values.
//
// [Full Topic]
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



// The unit style used by this formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter/unitStyle
func (m_ MassFormatter) UnitStyle() FormattingUnitStyle /* not a class type */ {
	rv := objc.Send[FormattingUnitStyle](m_.ID, objc.Sel("unitStyle"))
	return rv
}


// The unit style used by this formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/MassFormatter/unitStyle
func (m_ MassFormatter) SetUnitStyle(value FormattingUnitStyle /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUnitStyle:"), value)
}


// A Boolean value that indicates whether the resulting string represents a person’s mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/massformatter/isforpersonmassuse
func (m_ MassFormatter) IsForPersonMassUse() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("isForPersonMassUse"))
	return rv
}


// A Boolean value that indicates whether the resulting string represents a person’s mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/massformatter/isforpersonmassuse
func (m_ MassFormatter) SetIsForPersonMassUse(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsForPersonMassUse:"), value)
}


// The number formatter used to format the numbers in a mass strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/massformatter/numberformatter
func (m_ MassFormatter) NumberFormatter() INumberFormatter {
	rv := objc.Send[NumberFormatter](m_.ID, objc.Sel("numberFormatter"))
	return rv
}


// The number formatter used to format the numbers in a mass strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/massformatter/numberformatter
func (m_ MassFormatter) SetNumberFormatter(value INumberFormatter) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberFormatter:"), value)
}



