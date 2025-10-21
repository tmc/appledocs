// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [LengthFormatter] class.
var (
	LengthFormatterClass     _LengthFormatterClass
	LengthFormatterClassOnce sync.Once
)

func getLengthFormatterClass() _LengthFormatterClass {
	LengthFormatterClassOnce.Do(func() {
		LengthFormatterClass = _LengthFormatterClass{objc.GetClass("NSLengthFormatter")}
	})
	return LengthFormatterClass
}

type _LengthFormatterClass struct {
	class objc.Class
}

// An interface definition for the [LengthFormatter] class.
type ILengthFormatter interface {
	IFormatter
	UnitStringFromValueUnit(value unsafe.Pointer, unit unsafe.Pointer) string
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


// Returns the unit string based on the provided value and unit.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/unitString(fromValue:unit:)
func (l_ LengthFormatter) UnitStringFromValueUnit(value unsafe.Pointer, unit unsafe.Pointer) string {
	rv := objc.Send[string](l_.ID, objc.Sel("unitStringFromValue:unit:"), value, unit)
	return rv
}

// A Boolean value that indicates whether the resulting string represents a person’s height.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/lengthformatter/isforpersonheightuse
func (l_ LengthFormatter) IsForPersonHeightUse() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isForPersonHeightUse"))
	return rv
}


// SetIsForPersonHeightUse sets the value of the isForPersonHeightUse property.
// A Boolean value that indicates whether the resulting string represents a person’s height.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/lengthformatter/isforpersonheightuse
func (l_ LengthFormatter) SetIsForPersonHeightUse(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsForPersonHeightUse:"), value)
}

// The unit style used by this formatter.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/lengthformatter/unitstyle
func (l_ LengthFormatter) UnitStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("unitStyle"))
	return rv
}


// SetUnitStyle sets the value of the unitStyle property.
// The unit style used by this formatter.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/lengthformatter/unitstyle
func (l_ LengthFormatter) SetUnitStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setUnitStyle:"), value)
}

// The number formatter used to format the numbers in length strings.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/numberFormatter
func (l_ LengthFormatter) NumberFormatter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("numberFormatter"))
	return rv
}


// SetNumberFormatter sets the value of the numberFormatter property.
// The number formatter used to format the numbers in length strings.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/numberFormatter
func (l_ LengthFormatter) SetNumberFormatter(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberFormatter:"), value)
}



