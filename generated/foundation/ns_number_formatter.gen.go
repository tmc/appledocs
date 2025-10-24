// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSNumberFormatter */


/* debug [class_header]: Header for NSNumberFormatter */
// The class instance for the [NumberFormatter] class.
var (
	NumberFormatterClass     _NumberFormatterClass
	NumberFormatterClassOnce sync.Once
)

func getNumberFormatterClass() _NumberFormatterClass {
	NumberFormatterClassOnce.Do(func() {
		NumberFormatterClass = _NumberFormatterClass{objc.GetClass("NSNumberFormatter")}
	})
	return NumberFormatterClass
}

type _NumberFormatterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NumberFormatter */
// An interface definition for the [NumberFormatter] class.
type INumberFormatter interface {
	IFormatter
	
/* debug [class_interface_properties]: Properties for NumberFormatter */
	// properties:
	AllowsFloats() bool
	SetAllowsFloats(value bool)
	AlwaysShowsDecimalSeparator() bool
	SetAlwaysShowsDecimalSeparator(value bool)
	AttributedStringForNil() IAttributedString
	SetAttributedStringForNil(value IAttributedString)
	AttributedStringForNotANumber() IAttributedString
	SetAttributedStringForNotANumber(value IAttributedString)
	AttributedStringForZero() IAttributedString
	SetAttributedStringForZero(value IAttributedString)
	CurrencyCode() IString
	SetCurrencyCode(value IString)
	CurrencyDecimalSeparator() IString
	SetCurrencyDecimalSeparator(value IString)
	CurrencyGroupingSeparator() IString
	SetCurrencyGroupingSeparator(value IString)
	CurrencySymbol() IString
	SetCurrencySymbol(value IString)
	DecimalSeparator() IString
	SetDecimalSeparator(value IString)
	ExponentSymbol() IString
	SetExponentSymbol(value IString)
	Format() IString
	SetFormat(value IString)
	FormatWidth() uint
	SetFormatWidth(value uint)
	FormatterBehavior() NumberFormatterBehavior
	SetFormatterBehavior(value NumberFormatterBehavior)
	FormattingContext() FormattingContext
	SetFormattingContext(value FormattingContext)
	GeneratesDecimalNumbers() bool
	SetGeneratesDecimalNumbers(value bool)
	GroupingSeparator() IString
	SetGroupingSeparator(value IString)
	GroupingSize() uint
	SetGroupingSize(value uint)
	HasThousandSeparators() bool
	SetHasThousandSeparators(value bool)
	InternationalCurrencySymbol() IString
	SetInternationalCurrencySymbol(value IString)
	Lenient() bool
	SetLenient(value bool)
	PartialStringValidationEnabled() bool
	SetPartialStringValidationEnabled(value bool)
	Locale() ILocale
	SetLocale(value ILocale)
	LocalizesFormat() bool
	SetLocalizesFormat(value bool)
	Maximum() INumber
	SetMaximum(value INumber)
	MaximumFractionDigits() uint
	SetMaximumFractionDigits(value uint)
	MaximumIntegerDigits() uint
	SetMaximumIntegerDigits(value uint)
	MaximumSignificantDigits() uint
	SetMaximumSignificantDigits(value uint)
	Minimum() INumber
	SetMinimum(value INumber)
	MinimumFractionDigits() uint
	SetMinimumFractionDigits(value uint)
	MinimumGroupingDigits() int
	SetMinimumGroupingDigits(value int)
	MinimumIntegerDigits() uint
	SetMinimumIntegerDigits(value uint)
	MinimumSignificantDigits() uint
	SetMinimumSignificantDigits(value uint)
	MinusSign() IString
	SetMinusSign(value IString)
	Multiplier() INumber
	SetMultiplier(value INumber)
	NegativeFormat() IString
	SetNegativeFormat(value IString)
	NegativeInfinitySymbol() IString
	SetNegativeInfinitySymbol(value IString)
	NegativePrefix() IString
	SetNegativePrefix(value IString)
	NegativeSuffix() IString
	SetNegativeSuffix(value IString)
	NilSymbol() IString
	SetNilSymbol(value IString)
	NotANumberSymbol() IString
	SetNotANumberSymbol(value IString)
	NumberStyle() NumberFormatterStyle
	SetNumberStyle(value NumberFormatterStyle)
	PaddingCharacter() IString
	SetPaddingCharacter(value IString)
	PaddingPosition() NumberFormatterPadPosition
	SetPaddingPosition(value NumberFormatterPadPosition)
	PerMillSymbol() IString
	SetPerMillSymbol(value IString)
	PercentSymbol() IString
	SetPercentSymbol(value IString)
	PlusSign() IString
	SetPlusSign(value IString)
	PositiveFormat() IString
	SetPositiveFormat(value IString)
	PositiveInfinitySymbol() IString
	SetPositiveInfinitySymbol(value IString)
	PositivePrefix() IString
	SetPositivePrefix(value IString)
	PositiveSuffix() IString
	SetPositiveSuffix(value IString)
	RoundingBehavior() IDecimalNumberHandler
	SetRoundingBehavior(value IDecimalNumberHandler)
	RoundingIncrement() INumber
	SetRoundingIncrement(value INumber)
	RoundingMode() NumberFormatterRoundingMode
	SetRoundingMode(value NumberFormatterRoundingMode)
	SecondaryGroupingSize() uint
	SetSecondaryGroupingSize(value uint)
	TextAttributesForNegativeInfinity() IDictionary
	SetTextAttributesForNegativeInfinity(value IDictionary)
	TextAttributesForNegativeValues() IDictionary
	SetTextAttributesForNegativeValues(value IDictionary)
	TextAttributesForNil() IDictionary
	SetTextAttributesForNil(value IDictionary)
	TextAttributesForNotANumber() IDictionary
	SetTextAttributesForNotANumber(value IDictionary)
	TextAttributesForPositiveInfinity() IDictionary
	SetTextAttributesForPositiveInfinity(value IDictionary)
	TextAttributesForPositiveValues() IDictionary
	SetTextAttributesForPositiveValues(value IDictionary)
	TextAttributesForZero() IDictionary
	SetTextAttributesForZero(value IDictionary)
	ThousandSeparator() IString
	SetThousandSeparator(value IString)
	UsesGroupingSeparator() bool
	SetUsesGroupingSeparator(value bool)
	UsesSignificantDigits() bool
	SetUsesSignificantDigits(value bool)
	ZeroSymbol() IString
	SetZeroSymbol(value IString)
	IsLenient() bool
	SetIsLenient(value bool)
	IsPartialStringValidationEnabled() bool
	SetIsPartialStringValidationEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NumberFormatter */
	// methods:
	GetObjectValueForStringRangeError(obj objectivec.IObject, string_ IString, rangep objc.IObject /* cross-framework: Range */, error_ IError) bool
	NumberFromString(string_ IString) INumber
	StringFromNumber(number INumber) IString
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NumberFormatter */
// Alloc allocates a new instance without initialization.
func (nc _NumberFormatterClass) Alloc() NumberFormatter {
	rv := objc.Send[NumberFormatter](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NumberFormatter */
// A formatter that converts between numeric values and their textual representations.
//
// Instances of format the textual representation of cells that contain objects and convert textual representations of numeric values into objects. The representation encompasses integers, floats, and doubles; floats and doubles can be formatted to a specified decimal position. objects can also impose ranges on the numeric values cells can accept.


// A formatter that converts between numeric values and their textual representations.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NumberFormatter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NumberFormatter */

// Returns an constant that indicates default formatter behavior for new instances of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/defaultFormatterBehavior()
func (nc _NumberFormatterClass) DefaultFormatterBehavior() NumberFormatterBehavior {
	rv := objc.Send[NumberFormatterBehavior](objc.ID(nc.class), objc.Sel("defaultFormatterBehavior"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultFormatterBehavior) */


// Returns a localized number string with the specified style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/localizedString(from:number:)
func (nc _NumberFormatterClass) LocalizedStringFromNumberNumberStyle(num INumber, nstyle NumberFormatterStyle) IString {
	rv := objc.Send[String](objc.ID(nc.class), objc.Sel("localizedStringFromNumber:numberStyle:"), num, nstyle)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LocalizedStringFromNumberNumberStyle) */


// Sets the default formatter behavior for new instances of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/setDefaultFormatterBehavior(_:)
func (nc _NumberFormatterClass) SetDefaultFormatterBehavior(behavior NumberFormatterBehavior) {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("setDefaultFormatterBehavior:"), behavior)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetDefaultFormatterBehavior) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NumberFormatter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NumberFormatter */

// Returns by reference a cell-content object after creating it from a range of characters in a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/getObjectValue(_:for:range:)
func (n_ NumberFormatter) GetObjectValueForStringRangeError(obj objectivec.IObject, string_ IString, rangep objc.IObject /* cross-framework: Range */, error_ IError) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("getObjectValue:forString:range:error:"), obj, string_, rangep, error_)
	return rv
}/* debug [instance_methods/method]: GetObjectValueForStringRangeError */


// Returns an object created by parsing a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/number(from:)
func (n_ NumberFormatter) NumberFromString(string_ IString) INumber {
	rv := objc.Send[Number](n_.ID, objc.Sel("numberFromString:"), string_)
	return rv
}/* debug [instance_methods/method]: NumberFromString */


// Returns a string containing the formatted value of the provided number object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/string(from:)
func (n_ NumberFormatter) StringFromNumber(number INumber) IString {
	rv := objc.Send[String](n_.ID, objc.Sel("stringFromNumber:"), number)
	return rv
}/* debug [instance_methods/method]: StringFromNumber */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NumberFormatter */

// Determines whether the receiver allows as input floating-point values (that is, values that include the period character [ ]).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/allowsFloats
func (n_ NumberFormatter) AllowsFloats() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("allowsFloats"))
	return rv
}/* debug [instance_properties/getter]: allowsFloats */


// Determines whether the receiver allows as input floating-point values (that is, values that include the period character [ ]).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/allowsFloats
func (n_ NumberFormatter) SetAllowsFloats(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAllowsFloats:"), value)
}/* debug [instance_properties/setter]: allowsFloats */


// Determines whether the receiver always shows the decimal separator, even for integer numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/alwaysShowsDecimalSeparator
func (n_ NumberFormatter) AlwaysShowsDecimalSeparator() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("alwaysShowsDecimalSeparator"))
	return rv
}/* debug [instance_properties/getter]: alwaysShowsDecimalSeparator */


// Determines whether the receiver always shows the decimal separator, even for integer numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/alwaysShowsDecimalSeparator
func (n_ NumberFormatter) SetAlwaysShowsDecimalSeparator(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAlwaysShowsDecimalSeparator:"), value)
}/* debug [instance_properties/setter]: alwaysShowsDecimalSeparator */


// The attributed string the receiver uses to display values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForNil
func (n_ NumberFormatter) AttributedStringForNil() IAttributedString {
	rv := objc.Send[AttributedString](n_.ID, objc.Sel("attributedStringForNil"))
	return rv
}/* debug [instance_properties/getter]: attributedStringForNil */


// The attributed string the receiver uses to display values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForNil
func (n_ NumberFormatter) SetAttributedStringForNil(value IAttributedString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAttributedStringForNil:"), value)
}/* debug [instance_properties/setter]: attributedStringForNil */


// The attributed string the receiver uses to display “not a number” values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForNotANumber
func (n_ NumberFormatter) AttributedStringForNotANumber() IAttributedString {
	rv := objc.Send[AttributedString](n_.ID, objc.Sel("attributedStringForNotANumber"))
	return rv
}/* debug [instance_properties/getter]: attributedStringForNotANumber */


// The attributed string the receiver uses to display “not a number” values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForNotANumber
func (n_ NumberFormatter) SetAttributedStringForNotANumber(value IAttributedString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAttributedStringForNotANumber:"), value)
}/* debug [instance_properties/setter]: attributedStringForNotANumber */


// The attributed string that the receiver uses to display zero values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForZero
func (n_ NumberFormatter) AttributedStringForZero() IAttributedString {
	rv := objc.Send[AttributedString](n_.ID, objc.Sel("attributedStringForZero"))
	return rv
}/* debug [instance_properties/getter]: attributedStringForZero */


// The attributed string that the receiver uses to display zero values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForZero
func (n_ NumberFormatter) SetAttributedStringForZero(value IAttributedString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAttributedStringForZero:"), value)
}/* debug [instance_properties/setter]: attributedStringForZero */


// The receiver’s currency code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyCode
func (n_ NumberFormatter) CurrencyCode() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("currencyCode"))
	return rv
}/* debug [instance_properties/getter]: currencyCode */


// The receiver’s currency code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyCode
func (n_ NumberFormatter) SetCurrencyCode(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCurrencyCode:"), value)
}/* debug [instance_properties/setter]: currencyCode */


// The string used by the receiver as a currency decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyDecimalSeparator
func (n_ NumberFormatter) CurrencyDecimalSeparator() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("currencyDecimalSeparator"))
	return rv
}/* debug [instance_properties/getter]: currencyDecimalSeparator */


// The string used by the receiver as a currency decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyDecimalSeparator
func (n_ NumberFormatter) SetCurrencyDecimalSeparator(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCurrencyDecimalSeparator:"), value)
}/* debug [instance_properties/setter]: currencyDecimalSeparator */


// The currency grouping separator for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyGroupingSeparator
func (n_ NumberFormatter) CurrencyGroupingSeparator() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("currencyGroupingSeparator"))
	return rv
}/* debug [instance_properties/getter]: currencyGroupingSeparator */


// The currency grouping separator for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyGroupingSeparator
func (n_ NumberFormatter) SetCurrencyGroupingSeparator(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCurrencyGroupingSeparator:"), value)
}/* debug [instance_properties/setter]: currencyGroupingSeparator */


// The string used by the receiver as a local currency symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencySymbol
func (n_ NumberFormatter) CurrencySymbol() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("currencySymbol"))
	return rv
}/* debug [instance_properties/getter]: currencySymbol */


// The string used by the receiver as a local currency symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencySymbol
func (n_ NumberFormatter) SetCurrencySymbol(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCurrencySymbol:"), value)
}/* debug [instance_properties/setter]: currencySymbol */


// The character the receiver uses as a decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/decimalSeparator
func (n_ NumberFormatter) DecimalSeparator() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("decimalSeparator"))
	return rv
}/* debug [instance_properties/getter]: decimalSeparator */


// The character the receiver uses as a decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/decimalSeparator
func (n_ NumberFormatter) SetDecimalSeparator(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDecimalSeparator:"), value)
}/* debug [instance_properties/setter]: decimalSeparator */


// The string used to represent an exponent symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/exponentSymbol
func (n_ NumberFormatter) ExponentSymbol() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("exponentSymbol"))
	return rv
}/* debug [instance_properties/getter]: exponentSymbol */


// The string used to represent an exponent symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/exponentSymbol
func (n_ NumberFormatter) SetExponentSymbol(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExponentSymbol:"), value)
}/* debug [instance_properties/setter]: exponentSymbol */


// The receiver’s format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/format
func (n_ NumberFormatter) Format() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("format"))
	return rv
}/* debug [instance_properties/getter]: format */


// The receiver’s format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/format
func (n_ NumberFormatter) SetFormat(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFormat:"), value)
}/* debug [instance_properties/setter]: format */


// The format width used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/formatWidth
func (n_ NumberFormatter) FormatWidth() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("formatWidth"))
	return rv
}/* debug [instance_properties/getter]: formatWidth */


// The format width used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/formatWidth
func (n_ NumberFormatter) SetFormatWidth(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFormatWidth:"), value)
}/* debug [instance_properties/setter]: formatWidth */


// The formatter behavior of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/formatterBehavior
func (n_ NumberFormatter) FormatterBehavior() NumberFormatterBehavior {
	rv := objc.Send[NumberFormatterBehavior](n_.ID, objc.Sel("formatterBehavior"))
	return rv
}/* debug [instance_properties/getter]: formatterBehavior */


// The formatter behavior of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/formatterBehavior
func (n_ NumberFormatter) SetFormatterBehavior(value NumberFormatterBehavior) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFormatterBehavior:"), value)
}/* debug [instance_properties/setter]: formatterBehavior */


// The capitalization formatting context used when formatting a number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/formattingContext
func (n_ NumberFormatter) FormattingContext() FormattingContext {
	rv := objc.Send[FormattingContext](n_.ID, objc.Sel("formattingContext"))
	return rv
}/* debug [instance_properties/getter]: formattingContext */


// The capitalization formatting context used when formatting a number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/formattingContext
func (n_ NumberFormatter) SetFormattingContext(value FormattingContext) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFormattingContext:"), value)
}/* debug [instance_properties/setter]: formattingContext */


// Determines whether the receiver creates instances of when it converts strings to number objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/generatesDecimalNumbers
func (n_ NumberFormatter) GeneratesDecimalNumbers() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("generatesDecimalNumbers"))
	return rv
}/* debug [instance_properties/getter]: generatesDecimalNumbers */


// Determines whether the receiver creates instances of when it converts strings to number objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/generatesDecimalNumbers
func (n_ NumberFormatter) SetGeneratesDecimalNumbers(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setGeneratesDecimalNumbers:"), value)
}/* debug [instance_properties/setter]: generatesDecimalNumbers */


// The string used by the receiver for a grouping separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/groupingSeparator
func (n_ NumberFormatter) GroupingSeparator() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("groupingSeparator"))
	return rv
}/* debug [instance_properties/getter]: groupingSeparator */


// The string used by the receiver for a grouping separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/groupingSeparator
func (n_ NumberFormatter) SetGroupingSeparator(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setGroupingSeparator:"), value)
}/* debug [instance_properties/setter]: groupingSeparator */


// The grouping size of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/groupingSize
func (n_ NumberFormatter) GroupingSize() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("groupingSize"))
	return rv
}/* debug [instance_properties/getter]: groupingSize */


// The grouping size of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/groupingSize
func (n_ NumberFormatter) SetGroupingSize(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setGroupingSize:"), value)
}/* debug [instance_properties/setter]: groupingSize */


// Determines whether the receiver uses thousand separators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/hasThousandSeparators
func (n_ NumberFormatter) HasThousandSeparators() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("hasThousandSeparators"))
	return rv
}/* debug [instance_properties/getter]: hasThousandSeparators */


// Determines whether the receiver uses thousand separators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/hasThousandSeparators
func (n_ NumberFormatter) SetHasThousandSeparators(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHasThousandSeparators:"), value)
}/* debug [instance_properties/setter]: hasThousandSeparators */


// The international currency symbol used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/internationalCurrencySymbol
func (n_ NumberFormatter) InternationalCurrencySymbol() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("internationalCurrencySymbol"))
	return rv
}/* debug [instance_properties/getter]: internationalCurrencySymbol */


// The international currency symbol used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/internationalCurrencySymbol
func (n_ NumberFormatter) SetInternationalCurrencySymbol(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setInternationalCurrencySymbol:"), value)
}/* debug [instance_properties/setter]: internationalCurrencySymbol */


// Determines whether the receiver will use heuristics to guess at the number which is intended by a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/isLenient
func (n_ NumberFormatter) Lenient() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("lenient"))
	return rv
}/* debug [instance_properties/getter]: lenient */


// Determines whether the receiver will use heuristics to guess at the number which is intended by a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/isLenient
func (n_ NumberFormatter) SetLenient(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLenient:"), value)
}/* debug [instance_properties/setter]: lenient */


// Determines whether partial string validation is enabled for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/isPartialStringValidationEnabled
func (n_ NumberFormatter) PartialStringValidationEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("partialStringValidationEnabled"))
	return rv
}/* debug [instance_properties/getter]: partialStringValidationEnabled */


// Determines whether partial string validation is enabled for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/isPartialStringValidationEnabled
func (n_ NumberFormatter) SetPartialStringValidationEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPartialStringValidationEnabled:"), value)
}/* debug [instance_properties/setter]: partialStringValidationEnabled */


// The locale of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/locale
func (n_ NumberFormatter) Locale() ILocale {
	rv := objc.Send[Locale](n_.ID, objc.Sel("locale"))
	return rv
}/* debug [instance_properties/getter]: locale */


// The locale of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/locale
func (n_ NumberFormatter) SetLocale(value ILocale) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocale:"), value)
}/* debug [instance_properties/setter]: locale */


// Determines whether the dollar sign character ( ), decimal separator character ( ), and thousand separator character ( ) are converted to appropriately localized characters as specified by the user’s localization preference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/localizesFormat
func (n_ NumberFormatter) LocalizesFormat() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("localizesFormat"))
	return rv
}/* debug [instance_properties/getter]: localizesFormat */


// Determines whether the dollar sign character ( ), decimal separator character ( ), and thousand separator character ( ) are converted to appropriately localized characters as specified by the user’s localization preference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/localizesFormat
func (n_ NumberFormatter) SetLocalizesFormat(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocalizesFormat:"), value)
}/* debug [instance_properties/setter]: localizesFormat */


// The highest number allowed as input by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximum
func (n_ NumberFormatter) Maximum() INumber {
	rv := objc.Send[Number](n_.ID, objc.Sel("maximum"))
	return rv
}/* debug [instance_properties/getter]: maximum */


// The highest number allowed as input by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximum
func (n_ NumberFormatter) SetMaximum(value INumber) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMaximum:"), value)
}/* debug [instance_properties/setter]: maximum */


// The maximum number of digits after the decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumFractionDigits
func (n_ NumberFormatter) MaximumFractionDigits() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("maximumFractionDigits"))
	return rv
}/* debug [instance_properties/getter]: maximumFractionDigits */


// The maximum number of digits after the decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumFractionDigits
func (n_ NumberFormatter) SetMaximumFractionDigits(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMaximumFractionDigits:"), value)
}/* debug [instance_properties/setter]: maximumFractionDigits */


// The maximum number of digits before the decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumIntegerDigits
func (n_ NumberFormatter) MaximumIntegerDigits() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("maximumIntegerDigits"))
	return rv
}/* debug [instance_properties/getter]: maximumIntegerDigits */


// The maximum number of digits before the decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumIntegerDigits
func (n_ NumberFormatter) SetMaximumIntegerDigits(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMaximumIntegerDigits:"), value)
}/* debug [instance_properties/setter]: maximumIntegerDigits */


// The maximum number of significant digits for the number formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumSignificantDigits
func (n_ NumberFormatter) MaximumSignificantDigits() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("maximumSignificantDigits"))
	return rv
}/* debug [instance_properties/getter]: maximumSignificantDigits */


// The maximum number of significant digits for the number formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumSignificantDigits
func (n_ NumberFormatter) SetMaximumSignificantDigits(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMaximumSignificantDigits:"), value)
}/* debug [instance_properties/setter]: maximumSignificantDigits */


// The lowest number allowed as input by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimum
func (n_ NumberFormatter) Minimum() INumber {
	rv := objc.Send[Number](n_.ID, objc.Sel("minimum"))
	return rv
}/* debug [instance_properties/getter]: minimum */


// The lowest number allowed as input by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimum
func (n_ NumberFormatter) SetMinimum(value INumber) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimum:"), value)
}/* debug [instance_properties/setter]: minimum */


// The minimum number of digits after the decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumFractionDigits
func (n_ NumberFormatter) MinimumFractionDigits() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("minimumFractionDigits"))
	return rv
}/* debug [instance_properties/getter]: minimumFractionDigits */


// The minimum number of digits after the decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumFractionDigits
func (n_ NumberFormatter) SetMinimumFractionDigits(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimumFractionDigits:"), value)
}/* debug [instance_properties/setter]: minimumFractionDigits */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumGroupingDigits
func (n_ NumberFormatter) MinimumGroupingDigits() int {
	rv := objc.Send[int](n_.ID, objc.Sel("minimumGroupingDigits"))
	return rv
}/* debug [instance_properties/getter]: minimumGroupingDigits */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumGroupingDigits
func (n_ NumberFormatter) SetMinimumGroupingDigits(value int) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimumGroupingDigits:"), value)
}/* debug [instance_properties/setter]: minimumGroupingDigits */


// The minimum number of digits before the decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumIntegerDigits
func (n_ NumberFormatter) MinimumIntegerDigits() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("minimumIntegerDigits"))
	return rv
}/* debug [instance_properties/getter]: minimumIntegerDigits */


// The minimum number of digits before the decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumIntegerDigits
func (n_ NumberFormatter) SetMinimumIntegerDigits(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimumIntegerDigits:"), value)
}/* debug [instance_properties/setter]: minimumIntegerDigits */


// The minimum number of significant digits for the number formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumSignificantDigits
func (n_ NumberFormatter) MinimumSignificantDigits() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("minimumSignificantDigits"))
	return rv
}/* debug [instance_properties/getter]: minimumSignificantDigits */


// The minimum number of significant digits for the number formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumSignificantDigits
func (n_ NumberFormatter) SetMinimumSignificantDigits(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimumSignificantDigits:"), value)
}/* debug [instance_properties/setter]: minimumSignificantDigits */


// The string used to represent a minus sign.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minusSign
func (n_ NumberFormatter) MinusSign() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("minusSign"))
	return rv
}/* debug [instance_properties/getter]: minusSign */


// The string used to represent a minus sign.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minusSign
func (n_ NumberFormatter) SetMinusSign(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinusSign:"), value)
}/* debug [instance_properties/setter]: minusSign */


// The multiplier of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/multiplier
func (n_ NumberFormatter) Multiplier() INumber {
	rv := objc.Send[Number](n_.ID, objc.Sel("multiplier"))
	return rv
}/* debug [instance_properties/getter]: multiplier */


// The multiplier of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/multiplier
func (n_ NumberFormatter) SetMultiplier(value INumber) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMultiplier:"), value)
}/* debug [instance_properties/setter]: multiplier */


// The format the receiver uses to display negative values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeFormat
func (n_ NumberFormatter) NegativeFormat() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("negativeFormat"))
	return rv
}/* debug [instance_properties/getter]: negativeFormat */


// The format the receiver uses to display negative values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeFormat
func (n_ NumberFormatter) SetNegativeFormat(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNegativeFormat:"), value)
}/* debug [instance_properties/setter]: negativeFormat */


// The string used to represent a negative infinity symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeInfinitySymbol
func (n_ NumberFormatter) NegativeInfinitySymbol() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("negativeInfinitySymbol"))
	return rv
}/* debug [instance_properties/getter]: negativeInfinitySymbol */


// The string used to represent a negative infinity symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeInfinitySymbol
func (n_ NumberFormatter) SetNegativeInfinitySymbol(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNegativeInfinitySymbol:"), value)
}/* debug [instance_properties/setter]: negativeInfinitySymbol */


// The string the receiver uses as a prefix for negative values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativePrefix
func (n_ NumberFormatter) NegativePrefix() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("negativePrefix"))
	return rv
}/* debug [instance_properties/getter]: negativePrefix */


// The string the receiver uses as a prefix for negative values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativePrefix
func (n_ NumberFormatter) SetNegativePrefix(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNegativePrefix:"), value)
}/* debug [instance_properties/setter]: negativePrefix */


// The string the receiver uses as a suffix for negative values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeSuffix
func (n_ NumberFormatter) NegativeSuffix() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("negativeSuffix"))
	return rv
}/* debug [instance_properties/getter]: negativeSuffix */


// The string the receiver uses as a suffix for negative values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeSuffix
func (n_ NumberFormatter) SetNegativeSuffix(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNegativeSuffix:"), value)
}/* debug [instance_properties/setter]: negativeSuffix */


// The string used to represent a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/nilSymbol
func (n_ NumberFormatter) NilSymbol() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("nilSymbol"))
	return rv
}/* debug [instance_properties/getter]: nilSymbol */


// The string used to represent a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/nilSymbol
func (n_ NumberFormatter) SetNilSymbol(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNilSymbol:"), value)
}/* debug [instance_properties/setter]: nilSymbol */


// The string used to represent a NaN (“not a number”) value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/notANumberSymbol
func (n_ NumberFormatter) NotANumberSymbol() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("notANumberSymbol"))
	return rv
}/* debug [instance_properties/getter]: notANumberSymbol */


// The string used to represent a NaN (“not a number”) value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/notANumberSymbol
func (n_ NumberFormatter) SetNotANumberSymbol(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNotANumberSymbol:"), value)
}/* debug [instance_properties/setter]: notANumberSymbol */


// The number style used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/numberStyle
func (n_ NumberFormatter) NumberStyle() NumberFormatterStyle {
	rv := objc.Send[NumberFormatterStyle](n_.ID, objc.Sel("numberStyle"))
	return rv
}/* debug [instance_properties/getter]: numberStyle */


// The number style used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/numberStyle
func (n_ NumberFormatter) SetNumberStyle(value NumberFormatterStyle) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNumberStyle:"), value)
}/* debug [instance_properties/setter]: numberStyle */


// The string that the receiver uses to pad numbers in the formatted string representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/paddingCharacter
func (n_ NumberFormatter) PaddingCharacter() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("paddingCharacter"))
	return rv
}/* debug [instance_properties/getter]: paddingCharacter */


// The string that the receiver uses to pad numbers in the formatted string representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/paddingCharacter
func (n_ NumberFormatter) SetPaddingCharacter(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPaddingCharacter:"), value)
}/* debug [instance_properties/setter]: paddingCharacter */


// The padding position used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/paddingPosition
func (n_ NumberFormatter) PaddingPosition() NumberFormatterPadPosition {
	rv := objc.Send[NumberFormatterPadPosition](n_.ID, objc.Sel("paddingPosition"))
	return rv
}/* debug [instance_properties/getter]: paddingPosition */


// The padding position used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/paddingPosition
func (n_ NumberFormatter) SetPaddingPosition(value NumberFormatterPadPosition) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPaddingPosition:"), value)
}/* debug [instance_properties/setter]: paddingPosition */


// The string used to represent a per-mill (per-thousand) symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/perMillSymbol
func (n_ NumberFormatter) PerMillSymbol() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("perMillSymbol"))
	return rv
}/* debug [instance_properties/getter]: perMillSymbol */


// The string used to represent a per-mill (per-thousand) symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/perMillSymbol
func (n_ NumberFormatter) SetPerMillSymbol(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPerMillSymbol:"), value)
}/* debug [instance_properties/setter]: perMillSymbol */


// The string used to represent a percent symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/percentSymbol
func (n_ NumberFormatter) PercentSymbol() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("percentSymbol"))
	return rv
}/* debug [instance_properties/getter]: percentSymbol */


// The string used to represent a percent symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/percentSymbol
func (n_ NumberFormatter) SetPercentSymbol(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPercentSymbol:"), value)
}/* debug [instance_properties/setter]: percentSymbol */


// The string used to represent a plus sign.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/plusSign
func (n_ NumberFormatter) PlusSign() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("plusSign"))
	return rv
}/* debug [instance_properties/getter]: plusSign */


// The string used to represent a plus sign.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/plusSign
func (n_ NumberFormatter) SetPlusSign(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPlusSign:"), value)
}/* debug [instance_properties/setter]: plusSign */


// The format the receiver uses to display positive values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveFormat
func (n_ NumberFormatter) PositiveFormat() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("positiveFormat"))
	return rv
}/* debug [instance_properties/getter]: positiveFormat */


// The format the receiver uses to display positive values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveFormat
func (n_ NumberFormatter) SetPositiveFormat(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPositiveFormat:"), value)
}/* debug [instance_properties/setter]: positiveFormat */


// The string used to represent a positive infinity symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveInfinitySymbol
func (n_ NumberFormatter) PositiveInfinitySymbol() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("positiveInfinitySymbol"))
	return rv
}/* debug [instance_properties/getter]: positiveInfinitySymbol */


// The string used to represent a positive infinity symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveInfinitySymbol
func (n_ NumberFormatter) SetPositiveInfinitySymbol(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPositiveInfinitySymbol:"), value)
}/* debug [instance_properties/setter]: positiveInfinitySymbol */


// The string the receiver uses as the prefix for positive values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positivePrefix
func (n_ NumberFormatter) PositivePrefix() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("positivePrefix"))
	return rv
}/* debug [instance_properties/getter]: positivePrefix */


// The string the receiver uses as the prefix for positive values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positivePrefix
func (n_ NumberFormatter) SetPositivePrefix(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPositivePrefix:"), value)
}/* debug [instance_properties/setter]: positivePrefix */


// The string the receiver uses as the suffix for positive values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveSuffix
func (n_ NumberFormatter) PositiveSuffix() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("positiveSuffix"))
	return rv
}/* debug [instance_properties/getter]: positiveSuffix */


// The string the receiver uses as the suffix for positive values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveSuffix
func (n_ NumberFormatter) SetPositiveSuffix(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPositiveSuffix:"), value)
}/* debug [instance_properties/setter]: positiveSuffix */


// The rounding behavior used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingBehavior
func (n_ NumberFormatter) RoundingBehavior() IDecimalNumberHandler {
	rv := objc.Send[DecimalNumberHandler](n_.ID, objc.Sel("roundingBehavior"))
	return rv
}/* debug [instance_properties/getter]: roundingBehavior */


// The rounding behavior used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingBehavior
func (n_ NumberFormatter) SetRoundingBehavior(value IDecimalNumberHandler) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRoundingBehavior:"), value)
}/* debug [instance_properties/setter]: roundingBehavior */


// The rounding increment used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingIncrement
func (n_ NumberFormatter) RoundingIncrement() INumber {
	rv := objc.Send[Number](n_.ID, objc.Sel("roundingIncrement"))
	return rv
}/* debug [instance_properties/getter]: roundingIncrement */


// The rounding increment used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingIncrement
func (n_ NumberFormatter) SetRoundingIncrement(value INumber) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRoundingIncrement:"), value)
}/* debug [instance_properties/setter]: roundingIncrement */


// The rounding mode used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingMode-swift.property
func (n_ NumberFormatter) RoundingMode() NumberFormatterRoundingMode {
	rv := objc.Send[NumberFormatterRoundingMode](n_.ID, objc.Sel("roundingMode"))
	return rv
}/* debug [instance_properties/getter]: roundingMode */


// The rounding mode used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingMode-swift.property
func (n_ NumberFormatter) SetRoundingMode(value NumberFormatterRoundingMode) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRoundingMode:"), value)
}/* debug [instance_properties/setter]: roundingMode */


// The secondary grouping size of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/secondaryGroupingSize
func (n_ NumberFormatter) SecondaryGroupingSize() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("secondaryGroupingSize"))
	return rv
}/* debug [instance_properties/getter]: secondaryGroupingSize */


// The secondary grouping size of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/secondaryGroupingSize
func (n_ NumberFormatter) SetSecondaryGroupingSize(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSecondaryGroupingSize:"), value)
}/* debug [instance_properties/setter]: secondaryGroupingSize */


// The text attributes used to display the negative infinity symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNegativeInfinity
func (n_ NumberFormatter) TextAttributesForNegativeInfinity() IDictionary {
	rv := objc.Send[Dictionary](n_.ID, objc.Sel("textAttributesForNegativeInfinity"))
	return rv
}/* debug [instance_properties/getter]: textAttributesForNegativeInfinity */


// The text attributes used to display the negative infinity symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNegativeInfinity
func (n_ NumberFormatter) SetTextAttributesForNegativeInfinity(value IDictionary) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTextAttributesForNegativeInfinity:"), value)
}/* debug [instance_properties/setter]: textAttributesForNegativeInfinity */


// The text attributes to be used in displaying negative values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNegativeValues
func (n_ NumberFormatter) TextAttributesForNegativeValues() IDictionary {
	rv := objc.Send[Dictionary](n_.ID, objc.Sel("textAttributesForNegativeValues"))
	return rv
}/* debug [instance_properties/getter]: textAttributesForNegativeValues */


// The text attributes to be used in displaying negative values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNegativeValues
func (n_ NumberFormatter) SetTextAttributesForNegativeValues(value IDictionary) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTextAttributesForNegativeValues:"), value)
}/* debug [instance_properties/setter]: textAttributesForNegativeValues */


// The text attributes used to display the symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNil
func (n_ NumberFormatter) TextAttributesForNil() IDictionary {
	rv := objc.Send[Dictionary](n_.ID, objc.Sel("textAttributesForNil"))
	return rv
}/* debug [instance_properties/getter]: textAttributesForNil */


// The text attributes used to display the symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNil
func (n_ NumberFormatter) SetTextAttributesForNil(value IDictionary) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTextAttributesForNil:"), value)
}/* debug [instance_properties/setter]: textAttributesForNil */


// The text attributes used to display the NaN (“not a number”) string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNotANumber
func (n_ NumberFormatter) TextAttributesForNotANumber() IDictionary {
	rv := objc.Send[Dictionary](n_.ID, objc.Sel("textAttributesForNotANumber"))
	return rv
}/* debug [instance_properties/getter]: textAttributesForNotANumber */


// The text attributes used to display the NaN (“not a number”) string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNotANumber
func (n_ NumberFormatter) SetTextAttributesForNotANumber(value IDictionary) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTextAttributesForNotANumber:"), value)
}/* debug [instance_properties/setter]: textAttributesForNotANumber */


// The text attributes used to display the positive infinity symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForPositiveInfinity
func (n_ NumberFormatter) TextAttributesForPositiveInfinity() IDictionary {
	rv := objc.Send[Dictionary](n_.ID, objc.Sel("textAttributesForPositiveInfinity"))
	return rv
}/* debug [instance_properties/getter]: textAttributesForPositiveInfinity */


// The text attributes used to display the positive infinity symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForPositiveInfinity
func (n_ NumberFormatter) SetTextAttributesForPositiveInfinity(value IDictionary) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTextAttributesForPositiveInfinity:"), value)
}/* debug [instance_properties/setter]: textAttributesForPositiveInfinity */


// The text attributes to be used in displaying positive values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForPositiveValues
func (n_ NumberFormatter) TextAttributesForPositiveValues() IDictionary {
	rv := objc.Send[Dictionary](n_.ID, objc.Sel("textAttributesForPositiveValues"))
	return rv
}/* debug [instance_properties/getter]: textAttributesForPositiveValues */


// The text attributes to be used in displaying positive values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForPositiveValues
func (n_ NumberFormatter) SetTextAttributesForPositiveValues(value IDictionary) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTextAttributesForPositiveValues:"), value)
}/* debug [instance_properties/setter]: textAttributesForPositiveValues */


// The text attributes used to display a zero value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForZero
func (n_ NumberFormatter) TextAttributesForZero() IDictionary {
	rv := objc.Send[Dictionary](n_.ID, objc.Sel("textAttributesForZero"))
	return rv
}/* debug [instance_properties/getter]: textAttributesForZero */


// The text attributes used to display a zero value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForZero
func (n_ NumberFormatter) SetTextAttributesForZero(value IDictionary) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTextAttributesForZero:"), value)
}/* debug [instance_properties/setter]: textAttributesForZero */


// The character the receiver uses as a thousand separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/thousandSeparator
func (n_ NumberFormatter) ThousandSeparator() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("thousandSeparator"))
	return rv
}/* debug [instance_properties/getter]: thousandSeparator */


// The character the receiver uses as a thousand separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/thousandSeparator
func (n_ NumberFormatter) SetThousandSeparator(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setThousandSeparator:"), value)
}/* debug [instance_properties/setter]: thousandSeparator */


// Determines whether the receiver displays the group separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/usesGroupingSeparator
func (n_ NumberFormatter) UsesGroupingSeparator() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("usesGroupingSeparator"))
	return rv
}/* debug [instance_properties/getter]: usesGroupingSeparator */


// Determines whether the receiver displays the group separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/usesGroupingSeparator
func (n_ NumberFormatter) SetUsesGroupingSeparator(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUsesGroupingSeparator:"), value)
}/* debug [instance_properties/setter]: usesGroupingSeparator */


// A Boolean value indicating whether the formatter uses minimum and maximum significant digits when formatting numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/usesSignificantDigits
func (n_ NumberFormatter) UsesSignificantDigits() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("usesSignificantDigits"))
	return rv
}/* debug [instance_properties/getter]: usesSignificantDigits */


// A Boolean value indicating whether the formatter uses minimum and maximum significant digits when formatting numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/usesSignificantDigits
func (n_ NumberFormatter) SetUsesSignificantDigits(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUsesSignificantDigits:"), value)
}/* debug [instance_properties/setter]: usesSignificantDigits */


// The string used to represent a zero value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/zeroSymbol
func (n_ NumberFormatter) ZeroSymbol() IString {
	rv := objc.Send[String](n_.ID, objc.Sel("zeroSymbol"))
	return rv
}/* debug [instance_properties/getter]: zeroSymbol */


// The string used to represent a zero value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/zeroSymbol
func (n_ NumberFormatter) SetZeroSymbol(value IString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setZeroSymbol:"), value)
}/* debug [instance_properties/setter]: zeroSymbol */


// Determines whether the receiver will use heuristics to guess at the number which is intended by a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/numberformatter/islenient
func (n_ NumberFormatter) IsLenient() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isLenient"))
	return rv
}/* debug [instance_properties/getter]: isLenient */


// Determines whether the receiver will use heuristics to guess at the number which is intended by a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/numberformatter/islenient
func (n_ NumberFormatter) SetIsLenient(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsLenient:"), value)
}/* debug [instance_properties/setter]: isLenient */


// Determines whether partial string validation is enabled for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/numberformatter/ispartialstringvalidationenabled
func (n_ NumberFormatter) IsPartialStringValidationEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("isPartialStringValidationEnabled"))
	return rv
}/* debug [instance_properties/getter]: isPartialStringValidationEnabled */


// Determines whether partial string validation is enabled for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/numberformatter/ispartialstringvalidationenabled
func (n_ NumberFormatter) SetIsPartialStringValidationEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsPartialStringValidationEnabled:"), value)
}/* debug [instance_properties/setter]: isPartialStringValidationEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSNumberFormatter */



