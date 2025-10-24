// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSLengthFormatter */


/* debug [class_header]: Header for NSLengthFormatter */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LengthFormatter */
// An interface definition for the [LengthFormatter] class.
type ILengthFormatter interface {
	IFormatter
	
/* debug [class_interface_properties]: Properties for LengthFormatter */
	// properties:
	ForPersonHeightUse() bool
	SetForPersonHeightUse(value bool)
	NumberFormatter() INumberFormatter
	SetNumberFormatter(value INumberFormatter)
	UnitStyle() FormattingUnitStyle
	SetUnitStyle(value FormattingUnitStyle)
	IsForPersonHeightUse() bool
	SetIsForPersonHeightUse(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LengthFormatter */
	// methods:
	GetObjectValueForStringErrorDescription(obj objectivec.IObject, string_ IString, error_ IString) bool
	StringFromMeters(numberInMeters float64) IString
	StringFromValueUnit(value float64, unit LengthFormatterUnit) IString
	UnitStringFromMetersUsedUnit(numberInMeters float64, unitp LengthFormatterUnit) IString
	UnitStringFromValueUnit(value float64, unit LengthFormatterUnit) IString
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LengthFormatter */
// Alloc allocates a new instance without initialization.
func (lc _LengthFormatterClass) Alloc() LengthFormatter {
	rv := objc.Send[LengthFormatter](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LengthFormatter */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LengthFormatter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LengthFormatter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LengthFormatter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LengthFormatter */

// This method is not supported for the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/getObjectValue(_:for:errorDescription:)
func (l_ LengthFormatter) GetObjectValueForStringErrorDescription(obj objectivec.IObject, string_ IString, error_ IString) bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("getObjectValue:forString:errorDescription:"), obj, string_, error_)
	return rv
}/* debug [instance_methods/method]: GetObjectValueForStringErrorDescription */


// Returns a length string for the provided value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/string(fromMeters:)
func (l_ LengthFormatter) StringFromMeters(numberInMeters float64) IString {
	rv := objc.Send[String](l_.ID, objc.Sel("stringFromMeters:"), numberInMeters)
	return rv
}/* debug [instance_methods/method]: StringFromMeters */


// Returns a properly formatted length string for the given value and unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/string(fromValue:unit:)
func (l_ LengthFormatter) StringFromValueUnit(value float64, unit LengthFormatterUnit) IString {
	rv := objc.Send[String](l_.ID, objc.Sel("stringFromValue:unit:"), value, unit)
	return rv
}/* debug [instance_methods/method]: StringFromValueUnit */


// Returns the unit string for the provided value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/unitString(fromMeters:usedUnit:)
func (l_ LengthFormatter) UnitStringFromMetersUsedUnit(numberInMeters float64, unitp LengthFormatterUnit) IString {
	rv := objc.Send[String](l_.ID, objc.Sel("unitStringFromMeters:usedUnit:"), numberInMeters, unitp)
	return rv
}/* debug [instance_methods/method]: UnitStringFromMetersUsedUnit */


// Returns the unit string based on the provided value and unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/unitString(fromValue:unit:)
func (l_ LengthFormatter) UnitStringFromValueUnit(value float64, unit LengthFormatterUnit) IString {
	rv := objc.Send[String](l_.ID, objc.Sel("unitStringFromValue:unit:"), value, unit)
	return rv
}/* debug [instance_methods/method]: UnitStringFromValueUnit */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LengthFormatter */

// A Boolean value that indicates whether the resulting string represents a person’s height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/isForPersonHeightUse
func (l_ LengthFormatter) ForPersonHeightUse() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("forPersonHeightUse"))
	return rv
}/* debug [instance_properties/getter]: forPersonHeightUse */


// A Boolean value that indicates whether the resulting string represents a person’s height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/isForPersonHeightUse
func (l_ LengthFormatter) SetForPersonHeightUse(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setForPersonHeightUse:"), value)
}/* debug [instance_properties/setter]: forPersonHeightUse */


// The number formatter used to format the numbers in length strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/numberFormatter
func (l_ LengthFormatter) NumberFormatter() INumberFormatter {
	rv := objc.Send[NumberFormatter](l_.ID, objc.Sel("numberFormatter"))
	return rv
}/* debug [instance_properties/getter]: numberFormatter */


// The number formatter used to format the numbers in length strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/numberFormatter
func (l_ LengthFormatter) SetNumberFormatter(value INumberFormatter) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNumberFormatter:"), value)
}/* debug [instance_properties/setter]: numberFormatter */


// The unit style used by this formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/unitStyle
func (l_ LengthFormatter) UnitStyle() FormattingUnitStyle {
	rv := objc.Send[FormattingUnitStyle](l_.ID, objc.Sel("unitStyle"))
	return rv
}/* debug [instance_properties/getter]: unitStyle */


// The unit style used by this formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/LengthFormatter/unitStyle
func (l_ LengthFormatter) SetUnitStyle(value FormattingUnitStyle) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setUnitStyle:"), value)
}/* debug [instance_properties/setter]: unitStyle */


// A Boolean value that indicates whether the resulting string represents a person’s height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/lengthformatter/isforpersonheightuse
func (l_ LengthFormatter) IsForPersonHeightUse() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isForPersonHeightUse"))
	return rv
}/* debug [instance_properties/getter]: isForPersonHeightUse */


// A Boolean value that indicates whether the resulting string represents a person’s height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/lengthformatter/isforpersonheightuse
func (l_ LengthFormatter) SetIsForPersonHeightUse(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsForPersonHeightUse:"), value)
}/* debug [instance_properties/setter]: isForPersonHeightUse */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSLengthFormatter */



