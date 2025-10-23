// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	GetObjectValueForStringErrorDescription(obj objectivec.IObject, string_ string, error_ string) bool
	StringFromMeters(numberInMeters float64) String
	StringFromValueUnit(value float64, unit NSLengthFormatterUnit) String
	UnitStringFromMetersUsedUnit(numberInMeters float64, unitp NSLengthFormatterUnit) String
	UnitStringFromValueUnit(value float64, unit NSLengthFormatterUnit) String
	ForPersonHeightUse() bool
	SetForPersonHeightUse(value bool)
	NumberFormatter() NSNumberFormatter
	SetNumberFormatter(value INumberFormatter)
	UnitStyle() FormattingUnitStyle
	SetUnitStyle(value FormattingUnitStyle)
	IsForPersonHeightUse() bool
	SetIsForPersonHeightUse(value bool)
}

// A formatter that provides localized descriptions of linear distances, such as length and height measurements.


// A formatter that provides localized descriptions of linear distances, such as length and height measurements.
//
// [Full Topic]
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



// This method is not supported for the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/getObjectValue(_:for:errorDescription:)
func (l_ LengthFormatter) GetObjectValueForStringErrorDescription(obj objectivec.IObject, string_ string, error_ string) bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("getObjectValue:forString:errorDescription:"), obj, objc.String(string_), objc.String(error_))
	return rv
}


// Returns a length string for the provided value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/string(fromMeters:)
func (l_ LengthFormatter) StringFromMeters(numberInMeters float64) String {
	rv := objc.Send[String](l_.ID, objc.Sel("stringFromMeters:"), numberInMeters)
	return rv
}


// Returns a properly formatted length string for the given value and unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/string(fromValue:unit:)
func (l_ LengthFormatter) StringFromValueUnit(value float64, unit NSLengthFormatterUnit) String {
	rv := objc.Send[String](l_.ID, objc.Sel("stringFromValue:unit:"), value, unit)
	return rv
}


// Returns the unit string for the provided value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/unitString(fromMeters:usedUnit:)
func (l_ LengthFormatter) UnitStringFromMetersUsedUnit(numberInMeters float64, unitp NSLengthFormatterUnit) String {
	rv := objc.Send[String](l_.ID, objc.Sel("unitStringFromMeters:usedUnit:"), numberInMeters, unitp)
	return rv
}


// Returns the unit string based on the provided value and unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/unitString(fromValue:unit:)
func (l_ LengthFormatter) UnitStringFromValueUnit(value float64, unit NSLengthFormatterUnit) String {
	rv := objc.Send[String](l_.ID, objc.Sel("unitStringFromValue:unit:"), value, unit)
	return rv
}


// A Boolean value that indicates whether the resulting string represents a person’s height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/isForPersonHeightUse
func (l_ LengthFormatter) ForPersonHeightUse() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("forPersonHeightUse"))
	return rv
}


// A Boolean value that indicates whether the resulting string represents a person’s height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/isForPersonHeightUse
func (l_ LengthFormatter) SetForPersonHeightUse(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setForPersonHeightUse:"), value)
}


// The number formatter used to format the numbers in length strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/numberFormatter
func (l_ LengthFormatter) NumberFormatter() NSNumberFormatter {
	rv := objc.Send[NSNumberFormatter](l_.ID, objc.Sel("numberFormatter"))
	return rv
}


// The number formatter used to format the numbers in length strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/numberFormatter
func (l_ LengthFormatter) SetNumberFormatter(value INumberFormatter) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberFormatter:"), value)
}


// The unit style used by this formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/unitStyle
func (l_ LengthFormatter) UnitStyle() FormattingUnitStyle {
	rv := objc.Send[FormattingUnitStyle](l_.ID, objc.Sel("unitStyle"))
	return rv
}


// The unit style used by this formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/unitStyle
func (l_ LengthFormatter) SetUnitStyle(value FormattingUnitStyle) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setUnitStyle:"), value)
}


// A Boolean value that indicates whether the resulting string represents a person’s height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/lengthformatter/isforpersonheightuse
func (l_ LengthFormatter) IsForPersonHeightUse() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isForPersonHeightUse"))
	return rv
}


// A Boolean value that indicates whether the resulting string represents a person’s height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/lengthformatter/isforpersonheightuse
func (l_ LengthFormatter) SetIsForPersonHeightUse(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsForPersonHeightUse:"), value)
}



