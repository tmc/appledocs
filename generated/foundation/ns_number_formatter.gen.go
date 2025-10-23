// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [NumberFormatter] class.
type INumberFormatter interface {
	IFormatter
	// properties:
	AllowsFloats() bool /* primitive/slice/pointer */
	SetAllowsFloats(value bool /* primitive/slice/pointer */)
	AlwaysShowsDecimalSeparator() bool /* primitive/slice/pointer */
	SetAlwaysShowsDecimalSeparator(value bool /* primitive/slice/pointer */)
	AttributedStringForNil() IAttributedString
	SetAttributedStringForNil(value IAttributedString)
	AttributedStringForNotANumber() IAttributedString
	SetAttributedStringForNotANumber(value IAttributedString)
	AttributedStringForZero() IAttributedString
	SetAttributedStringForZero(value IAttributedString)
	CurrencyCode() string /* primitive/slice/pointer */
	SetCurrencyCode(value string /* primitive/slice/pointer */)
	CurrencyDecimalSeparator() string /* primitive/slice/pointer */
	SetCurrencyDecimalSeparator(value string /* primitive/slice/pointer */)
	CurrencyGroupingSeparator() string /* primitive/slice/pointer */
	SetCurrencyGroupingSeparator(value string /* primitive/slice/pointer */)
	CurrencySymbol() string /* primitive/slice/pointer */
	SetCurrencySymbol(value string /* primitive/slice/pointer */)
	DecimalSeparator() string /* primitive/slice/pointer */
	SetDecimalSeparator(value string /* primitive/slice/pointer */)
	ExponentSymbol() string /* primitive/slice/pointer */
	SetExponentSymbol(value string /* primitive/slice/pointer */)
	Format() string /* primitive/slice/pointer */
	SetFormat(value string /* primitive/slice/pointer */)
	FormatterBehavior() NumberFormatterBehavior
	SetFormatterBehavior(value NumberFormatterBehavior)
	FormattingContext() int /* primitive/slice/pointer */
	SetFormattingContext(value int /* primitive/slice/pointer */)
	FormatWidth() uint /* primitive/slice/pointer */
	SetFormatWidth(value uint /* primitive/slice/pointer */)
	GeneratesDecimalNumbers() bool /* primitive/slice/pointer */
	SetGeneratesDecimalNumbers(value bool /* primitive/slice/pointer */)
	GroupingSeparator() string /* primitive/slice/pointer */
	SetGroupingSeparator(value string /* primitive/slice/pointer */)
	GroupingSize() uint /* primitive/slice/pointer */
	SetGroupingSize(value uint /* primitive/slice/pointer */)
	HasThousandSeparators() bool /* primitive/slice/pointer */
	SetHasThousandSeparators(value bool /* primitive/slice/pointer */)
	InternationalCurrencySymbol() string /* primitive/slice/pointer */
	SetInternationalCurrencySymbol(value string /* primitive/slice/pointer */)
	Lenient() bool /* primitive/slice/pointer */
	SetLenient(value bool /* primitive/slice/pointer */)
	PartialStringValidationEnabled() bool /* primitive/slice/pointer */
	SetPartialStringValidationEnabled(value bool /* primitive/slice/pointer */)
	Locale() ILocale
	SetLocale(value ILocale)
	LocalizesFormat() bool /* primitive/slice/pointer */
	SetLocalizesFormat(value bool /* primitive/slice/pointer */)
	Maximum() Number /* not a class type */
	SetMaximum(value Number /* not a class type */)
	MaximumFractionDigits() uint /* primitive/slice/pointer */
	SetMaximumFractionDigits(value uint /* primitive/slice/pointer */)
	MaximumIntegerDigits() uint /* primitive/slice/pointer */
	SetMaximumIntegerDigits(value uint /* primitive/slice/pointer */)
	MaximumSignificantDigits() uint /* primitive/slice/pointer */
	SetMaximumSignificantDigits(value uint /* primitive/slice/pointer */)
	Minimum() Number /* not a class type */
	SetMinimum(value Number /* not a class type */)
	MinimumFractionDigits() uint /* primitive/slice/pointer */
	SetMinimumFractionDigits(value uint /* primitive/slice/pointer */)
	MinimumGroupingDigits() int /* primitive/slice/pointer */
	SetMinimumGroupingDigits(value int /* primitive/slice/pointer */)
	MinimumIntegerDigits() uint /* primitive/slice/pointer */
	SetMinimumIntegerDigits(value uint /* primitive/slice/pointer */)
	MinimumSignificantDigits() uint /* primitive/slice/pointer */
	SetMinimumSignificantDigits(value uint /* primitive/slice/pointer */)
	MinusSign() string /* primitive/slice/pointer */
	SetMinusSign(value string /* primitive/slice/pointer */)
	Multiplier() Number /* not a class type */
	SetMultiplier(value Number /* not a class type */)
	NegativeFormat() string /* primitive/slice/pointer */
	SetNegativeFormat(value string /* primitive/slice/pointer */)
	NegativeInfinitySymbol() string /* primitive/slice/pointer */
	SetNegativeInfinitySymbol(value string /* primitive/slice/pointer */)
	NegativePrefix() string /* primitive/slice/pointer */
	SetNegativePrefix(value string /* primitive/slice/pointer */)
	NegativeSuffix() string /* primitive/slice/pointer */
	SetNegativeSuffix(value string /* primitive/slice/pointer */)
	NilSymbol() string /* primitive/slice/pointer */
	SetNilSymbol(value string /* primitive/slice/pointer */)
	NotANumberSymbol() string /* primitive/slice/pointer */
	SetNotANumberSymbol(value string /* primitive/slice/pointer */)
	NumberStyle() NumberFormatterStyle
	SetNumberStyle(value NumberFormatterStyle)
	PaddingCharacter() string /* primitive/slice/pointer */
	SetPaddingCharacter(value string /* primitive/slice/pointer */)
	PaddingPosition() NumberFormatterPadPosition
	SetPaddingPosition(value NumberFormatterPadPosition)
	PercentSymbol() string /* primitive/slice/pointer */
	SetPercentSymbol(value string /* primitive/slice/pointer */)
	PerMillSymbol() string /* primitive/slice/pointer */
	SetPerMillSymbol(value string /* primitive/slice/pointer */)
	PlusSign() string /* primitive/slice/pointer */
	SetPlusSign(value string /* primitive/slice/pointer */)
	PositiveFormat() string /* primitive/slice/pointer */
	SetPositiveFormat(value string /* primitive/slice/pointer */)
	PositiveInfinitySymbol() string /* primitive/slice/pointer */
	SetPositiveInfinitySymbol(value string /* primitive/slice/pointer */)
	PositivePrefix() string /* primitive/slice/pointer */
	SetPositivePrefix(value string /* primitive/slice/pointer */)
	PositiveSuffix() string /* primitive/slice/pointer */
	SetPositiveSuffix(value string /* primitive/slice/pointer */)
	RoundingBehavior() IDecimalNumberHandler
	SetRoundingBehavior(value IDecimalNumberHandler)
	RoundingIncrement() Number /* not a class type */
	SetRoundingIncrement(value Number /* not a class type */)
	RoundingMode() NumberFormatterRoundingMode
	SetRoundingMode(value NumberFormatterRoundingMode)
	SecondaryGroupingSize() uint /* primitive/slice/pointer */
	SetSecondaryGroupingSize(value uint /* primitive/slice/pointer */)
	TextAttributesForNegativeInfinity() IDictionary /* already interface */
	SetTextAttributesForNegativeInfinity(value IDictionary /* already interface */)
	TextAttributesForNegativeValues() IDictionary /* already interface */
	SetTextAttributesForNegativeValues(value IDictionary /* already interface */)
	TextAttributesForNil() IDictionary /* already interface */
	SetTextAttributesForNil(value IDictionary /* already interface */)
	TextAttributesForNotANumber() IDictionary /* already interface */
	SetTextAttributesForNotANumber(value IDictionary /* already interface */)
	TextAttributesForPositiveInfinity() IDictionary /* already interface */
	SetTextAttributesForPositiveInfinity(value IDictionary /* already interface */)
	TextAttributesForPositiveValues() IDictionary /* already interface */
	SetTextAttributesForPositiveValues(value IDictionary /* already interface */)
	TextAttributesForZero() IDictionary /* already interface */
	SetTextAttributesForZero(value IDictionary /* already interface */)
	ThousandSeparator() string /* primitive/slice/pointer */
	SetThousandSeparator(value string /* primitive/slice/pointer */)
	UsesGroupingSeparator() bool /* primitive/slice/pointer */
	SetUsesGroupingSeparator(value bool /* primitive/slice/pointer */)
	UsesSignificantDigits() bool /* primitive/slice/pointer */
	SetUsesSignificantDigits(value bool /* primitive/slice/pointer */)
	ZeroSymbol() string /* primitive/slice/pointer */
	SetZeroSymbol(value string /* primitive/slice/pointer */)
	IsLenient() bool /* primitive/slice/pointer */
	SetIsLenient(value bool /* primitive/slice/pointer */)
	IsPartialStringValidationEnabled() bool /* primitive/slice/pointer */
	SetIsPartialStringValidationEnabled(value bool /* primitive/slice/pointer */)
	// methods:
	GetObjectValueForStringRangeError(obj objectivec.IObject, string_ string /* primitive/slice/pointer */, rangep Range /* not a class type */, error_ unsafe.Pointer) bool /* primitive/slice/pointer */
	NumberFromString(string_ string /* primitive/slice/pointer */) Number /* not a class type */
	StringFromNumber(number Number /* not a class type */) String /* not a class type */
}

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

// Alloc allocates a new instance without initialization.
func (nc _NumberFormatterClass) Alloc() NumberFormatter {
	rv := objc.Send[NumberFormatter](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Returns an constant that indicates default formatter behavior for new instances of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/defaultFormatterBehavior()
func (nc _NumberFormatterClass) DefaultFormatterBehavior() NumberFormatterBehavior {
	rv := objc.Send[NumberFormatterBehavior](objc.ID(nc.class), objc.Sel("defaultFormatterBehavior"))
	return rv
}


// Returns a localized number string with the specified style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/localizedString(from:number:)
func (nc _NumberFormatterClass) LocalizedStringFromNumberNumberStyle(num Number /* not a class type */, nstyle NumberFormatterStyle) String /* not a class type */ {
	rv := objc.Send[String](objc.ID(nc.class), objc.Sel("localizedStringFromNumber:numberStyle:"), num, nstyle)
	return rv
}


// Sets the default formatter behavior for new instances of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/setDefaultFormatterBehavior(_:)
func (nc _NumberFormatterClass) SetDefaultFormatterBehavior(behavior NumberFormatterBehavior) {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("setDefaultFormatterBehavior:"), behavior)
}


// Returns by reference a cell-content object after creating it from a range of characters in a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/getObjectValue(_:for:range:)
func (n_ NumberFormatter) GetObjectValueForStringRangeError(obj objectivec.IObject, string_ string /* primitive/slice/pointer */, rangep Range /* not a class type */, error_ unsafe.Pointer) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](n_.ID, objc.Sel("getObjectValue:forString:range:error:"), obj, objc.String(string_), rangep, error_)
	return rv
}


// Returns an object created by parsing a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/number(from:)
func (n_ NumberFormatter) NumberFromString(string_ string /* primitive/slice/pointer */) Number /* not a class type */ {
	rv := objc.Send[Number](n_.ID, objc.Sel("numberFromString:"), objc.String(string_))
	return rv
}


// Returns a string containing the formatted value of the provided number object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/string(from:)
func (n_ NumberFormatter) StringFromNumber(number Number /* not a class type */) String /* not a class type */ {
	rv := objc.Send[String](n_.ID, objc.Sel("stringFromNumber:"), number)
	return rv
}


// Determines whether the receiver allows as input floating-point values (that is, values that include the period character [ ]).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/allowsFloats
func (n_ NumberFormatter) AllowsFloats() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](n_.ID, objc.Sel("allowsFloats"))
	return rv
}


// Determines whether the receiver allows as input floating-point values (that is, values that include the period character [ ]).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/allowsFloats
func (n_ NumberFormatter) SetAllowsFloats(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAllowsFloats:"), value)
}


// Determines whether the receiver always shows the decimal separator, even for integer numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/alwaysShowsDecimalSeparator
func (n_ NumberFormatter) AlwaysShowsDecimalSeparator() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](n_.ID, objc.Sel("alwaysShowsDecimalSeparator"))
	return rv
}


// Determines whether the receiver always shows the decimal separator, even for integer numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/alwaysShowsDecimalSeparator
func (n_ NumberFormatter) SetAlwaysShowsDecimalSeparator(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAlwaysShowsDecimalSeparator:"), value)
}


// The attributed string the receiver uses to display values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForNil
func (n_ NumberFormatter) AttributedStringForNil() IAttributedString {
	rv := objc.Send[AttributedString](n_.ID, objc.Sel("attributedStringForNil"))
	return rv
}


// The attributed string the receiver uses to display values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForNil
func (n_ NumberFormatter) SetAttributedStringForNil(value IAttributedString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAttributedStringForNil:"), value)
}


// The attributed string the receiver uses to display “not a number” values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForNotANumber
func (n_ NumberFormatter) AttributedStringForNotANumber() IAttributedString {
	rv := objc.Send[AttributedString](n_.ID, objc.Sel("attributedStringForNotANumber"))
	return rv
}


// The attributed string the receiver uses to display “not a number” values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForNotANumber
func (n_ NumberFormatter) SetAttributedStringForNotANumber(value IAttributedString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAttributedStringForNotANumber:"), value)
}


// The attributed string that the receiver uses to display zero values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForZero
func (n_ NumberFormatter) AttributedStringForZero() IAttributedString {
	rv := objc.Send[AttributedString](n_.ID, objc.Sel("attributedStringForZero"))
	return rv
}


// The attributed string that the receiver uses to display zero values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForZero
func (n_ NumberFormatter) SetAttributedStringForZero(value IAttributedString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAttributedStringForZero:"), value)
}


// The receiver’s currency code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyCode
func (n_ NumberFormatter) CurrencyCode() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("currencyCode"))
	return rv
}


// The receiver’s currency code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyCode
func (n_ NumberFormatter) SetCurrencyCode(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCurrencyCode:"), objc.String(value))
}


// The string used by the receiver as a currency decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyDecimalSeparator
func (n_ NumberFormatter) CurrencyDecimalSeparator() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("currencyDecimalSeparator"))
	return rv
}


// The string used by the receiver as a currency decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyDecimalSeparator
func (n_ NumberFormatter) SetCurrencyDecimalSeparator(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCurrencyDecimalSeparator:"), objc.String(value))
}


// The currency grouping separator for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyGroupingSeparator
func (n_ NumberFormatter) CurrencyGroupingSeparator() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("currencyGroupingSeparator"))
	return rv
}


// The currency grouping separator for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyGroupingSeparator
func (n_ NumberFormatter) SetCurrencyGroupingSeparator(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCurrencyGroupingSeparator:"), objc.String(value))
}


// The string used by the receiver as a local currency symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencySymbol
func (n_ NumberFormatter) CurrencySymbol() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("currencySymbol"))
	return rv
}


// The string used by the receiver as a local currency symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencySymbol
func (n_ NumberFormatter) SetCurrencySymbol(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCurrencySymbol:"), objc.String(value))
}


// The character the receiver uses as a decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/decimalSeparator
func (n_ NumberFormatter) DecimalSeparator() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("decimalSeparator"))
	return rv
}


// The character the receiver uses as a decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/decimalSeparator
func (n_ NumberFormatter) SetDecimalSeparator(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDecimalSeparator:"), objc.String(value))
}


// The string used to represent an exponent symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/exponentSymbol
func (n_ NumberFormatter) ExponentSymbol() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("exponentSymbol"))
	return rv
}


// The string used to represent an exponent symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/exponentSymbol
func (n_ NumberFormatter) SetExponentSymbol(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExponentSymbol:"), objc.String(value))
}


// The receiver’s format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/format
func (n_ NumberFormatter) Format() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("format"))
	return rv
}


// The receiver’s format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/format
func (n_ NumberFormatter) SetFormat(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFormat:"), objc.String(value))
}


// The formatter behavior of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/formatterBehavior
func (n_ NumberFormatter) FormatterBehavior() NumberFormatterBehavior {
	rv := objc.Send[NumberFormatterBehavior](n_.ID, objc.Sel("formatterBehavior"))
	return rv
}


// The formatter behavior of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/formatterBehavior
func (n_ NumberFormatter) SetFormatterBehavior(value NumberFormatterBehavior) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFormatterBehavior:"), value)
}


// The capitalization formatting context used when formatting a number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/formattingContext
func (n_ NumberFormatter) FormattingContext() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](n_.ID, objc.Sel("formattingContext"))
	return rv
}


// The capitalization formatting context used when formatting a number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/formattingContext
func (n_ NumberFormatter) SetFormattingContext(value int /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFormattingContext:"), value)
}


// The format width used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/formatWidth
func (n_ NumberFormatter) FormatWidth() uint /* primitive/slice/pointer */ {
	rv := objc.Send[uint](n_.ID, objc.Sel("formatWidth"))
	return rv
}


// The format width used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/formatWidth
func (n_ NumberFormatter) SetFormatWidth(value uint /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFormatWidth:"), value)
}


// Determines whether the receiver creates instances of when it converts strings to number objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/generatesDecimalNumbers
func (n_ NumberFormatter) GeneratesDecimalNumbers() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](n_.ID, objc.Sel("generatesDecimalNumbers"))
	return rv
}


// Determines whether the receiver creates instances of when it converts strings to number objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/generatesDecimalNumbers
func (n_ NumberFormatter) SetGeneratesDecimalNumbers(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setGeneratesDecimalNumbers:"), value)
}


// The string used by the receiver for a grouping separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/groupingSeparator
func (n_ NumberFormatter) GroupingSeparator() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("groupingSeparator"))
	return rv
}


// The string used by the receiver for a grouping separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/groupingSeparator
func (n_ NumberFormatter) SetGroupingSeparator(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setGroupingSeparator:"), objc.String(value))
}


// The grouping size of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/groupingSize
func (n_ NumberFormatter) GroupingSize() uint /* primitive/slice/pointer */ {
	rv := objc.Send[uint](n_.ID, objc.Sel("groupingSize"))
	return rv
}


// The grouping size of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/groupingSize
func (n_ NumberFormatter) SetGroupingSize(value uint /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setGroupingSize:"), value)
}


// Determines whether the receiver uses thousand separators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/hasThousandSeparators
func (n_ NumberFormatter) HasThousandSeparators() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](n_.ID, objc.Sel("hasThousandSeparators"))
	return rv
}


// Determines whether the receiver uses thousand separators.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/hasThousandSeparators
func (n_ NumberFormatter) SetHasThousandSeparators(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHasThousandSeparators:"), value)
}


// The international currency symbol used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/internationalCurrencySymbol
func (n_ NumberFormatter) InternationalCurrencySymbol() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("internationalCurrencySymbol"))
	return rv
}


// The international currency symbol used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/internationalCurrencySymbol
func (n_ NumberFormatter) SetInternationalCurrencySymbol(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setInternationalCurrencySymbol:"), objc.String(value))
}


// Determines whether the receiver will use heuristics to guess at the number which is intended by a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/isLenient
func (n_ NumberFormatter) Lenient() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](n_.ID, objc.Sel("lenient"))
	return rv
}


// Determines whether the receiver will use heuristics to guess at the number which is intended by a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/isLenient
func (n_ NumberFormatter) SetLenient(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLenient:"), value)
}


// Determines whether partial string validation is enabled for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/isPartialStringValidationEnabled
func (n_ NumberFormatter) PartialStringValidationEnabled() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](n_.ID, objc.Sel("partialStringValidationEnabled"))
	return rv
}


// Determines whether partial string validation is enabled for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/isPartialStringValidationEnabled
func (n_ NumberFormatter) SetPartialStringValidationEnabled(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPartialStringValidationEnabled:"), value)
}


// The locale of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/locale
func (n_ NumberFormatter) Locale() ILocale {
	rv := objc.Send[Locale](n_.ID, objc.Sel("locale"))
	return rv
}


// The locale of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/locale
func (n_ NumberFormatter) SetLocale(value ILocale) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocale:"), value)
}


// Determines whether the dollar sign character ( ), decimal separator character ( ), and thousand separator character ( ) are converted to appropriately localized characters as specified by the user’s localization preference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/localizesFormat
func (n_ NumberFormatter) LocalizesFormat() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](n_.ID, objc.Sel("localizesFormat"))
	return rv
}


// Determines whether the dollar sign character ( ), decimal separator character ( ), and thousand separator character ( ) are converted to appropriately localized characters as specified by the user’s localization preference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/localizesFormat
func (n_ NumberFormatter) SetLocalizesFormat(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocalizesFormat:"), value)
}


// The highest number allowed as input by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximum
func (n_ NumberFormatter) Maximum() Number /* not a class type */ {
	rv := objc.Send[Number](n_.ID, objc.Sel("maximum"))
	return rv
}


// The highest number allowed as input by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximum
func (n_ NumberFormatter) SetMaximum(value Number /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMaximum:"), value)
}


// The maximum number of digits after the decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumFractionDigits
func (n_ NumberFormatter) MaximumFractionDigits() uint /* primitive/slice/pointer */ {
	rv := objc.Send[uint](n_.ID, objc.Sel("maximumFractionDigits"))
	return rv
}


// The maximum number of digits after the decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumFractionDigits
func (n_ NumberFormatter) SetMaximumFractionDigits(value uint /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMaximumFractionDigits:"), value)
}


// The maximum number of digits before the decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumIntegerDigits
func (n_ NumberFormatter) MaximumIntegerDigits() uint /* primitive/slice/pointer */ {
	rv := objc.Send[uint](n_.ID, objc.Sel("maximumIntegerDigits"))
	return rv
}


// The maximum number of digits before the decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumIntegerDigits
func (n_ NumberFormatter) SetMaximumIntegerDigits(value uint /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMaximumIntegerDigits:"), value)
}


// The maximum number of significant digits for the number formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumSignificantDigits
func (n_ NumberFormatter) MaximumSignificantDigits() uint /* primitive/slice/pointer */ {
	rv := objc.Send[uint](n_.ID, objc.Sel("maximumSignificantDigits"))
	return rv
}


// The maximum number of significant digits for the number formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumSignificantDigits
func (n_ NumberFormatter) SetMaximumSignificantDigits(value uint /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMaximumSignificantDigits:"), value)
}


// The lowest number allowed as input by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimum
func (n_ NumberFormatter) Minimum() Number /* not a class type */ {
	rv := objc.Send[Number](n_.ID, objc.Sel("minimum"))
	return rv
}


// The lowest number allowed as input by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimum
func (n_ NumberFormatter) SetMinimum(value Number /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimum:"), value)
}


// The minimum number of digits after the decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumFractionDigits
func (n_ NumberFormatter) MinimumFractionDigits() uint /* primitive/slice/pointer */ {
	rv := objc.Send[uint](n_.ID, objc.Sel("minimumFractionDigits"))
	return rv
}


// The minimum number of digits after the decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumFractionDigits
func (n_ NumberFormatter) SetMinimumFractionDigits(value uint /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimumFractionDigits:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumGroupingDigits
func (n_ NumberFormatter) MinimumGroupingDigits() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](n_.ID, objc.Sel("minimumGroupingDigits"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumGroupingDigits
func (n_ NumberFormatter) SetMinimumGroupingDigits(value int /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimumGroupingDigits:"), value)
}


// The minimum number of digits before the decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumIntegerDigits
func (n_ NumberFormatter) MinimumIntegerDigits() uint /* primitive/slice/pointer */ {
	rv := objc.Send[uint](n_.ID, objc.Sel("minimumIntegerDigits"))
	return rv
}


// The minimum number of digits before the decimal separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumIntegerDigits
func (n_ NumberFormatter) SetMinimumIntegerDigits(value uint /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimumIntegerDigits:"), value)
}


// The minimum number of significant digits for the number formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumSignificantDigits
func (n_ NumberFormatter) MinimumSignificantDigits() uint /* primitive/slice/pointer */ {
	rv := objc.Send[uint](n_.ID, objc.Sel("minimumSignificantDigits"))
	return rv
}


// The minimum number of significant digits for the number formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumSignificantDigits
func (n_ NumberFormatter) SetMinimumSignificantDigits(value uint /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimumSignificantDigits:"), value)
}


// The string used to represent a minus sign.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minusSign
func (n_ NumberFormatter) MinusSign() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("minusSign"))
	return rv
}


// The string used to represent a minus sign.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minusSign
func (n_ NumberFormatter) SetMinusSign(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinusSign:"), objc.String(value))
}


// The multiplier of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/multiplier
func (n_ NumberFormatter) Multiplier() Number /* not a class type */ {
	rv := objc.Send[Number](n_.ID, objc.Sel("multiplier"))
	return rv
}


// The multiplier of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/multiplier
func (n_ NumberFormatter) SetMultiplier(value Number /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMultiplier:"), value)
}


// The format the receiver uses to display negative values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeFormat
func (n_ NumberFormatter) NegativeFormat() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("negativeFormat"))
	return rv
}


// The format the receiver uses to display negative values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeFormat
func (n_ NumberFormatter) SetNegativeFormat(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNegativeFormat:"), objc.String(value))
}


// The string used to represent a negative infinity symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeInfinitySymbol
func (n_ NumberFormatter) NegativeInfinitySymbol() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("negativeInfinitySymbol"))
	return rv
}


// The string used to represent a negative infinity symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeInfinitySymbol
func (n_ NumberFormatter) SetNegativeInfinitySymbol(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNegativeInfinitySymbol:"), objc.String(value))
}


// The string the receiver uses as a prefix for negative values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativePrefix
func (n_ NumberFormatter) NegativePrefix() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("negativePrefix"))
	return rv
}


// The string the receiver uses as a prefix for negative values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativePrefix
func (n_ NumberFormatter) SetNegativePrefix(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNegativePrefix:"), objc.String(value))
}


// The string the receiver uses as a suffix for negative values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeSuffix
func (n_ NumberFormatter) NegativeSuffix() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("negativeSuffix"))
	return rv
}


// The string the receiver uses as a suffix for negative values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeSuffix
func (n_ NumberFormatter) SetNegativeSuffix(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNegativeSuffix:"), objc.String(value))
}


// The string used to represent a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/nilSymbol
func (n_ NumberFormatter) NilSymbol() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("nilSymbol"))
	return rv
}


// The string used to represent a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/nilSymbol
func (n_ NumberFormatter) SetNilSymbol(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNilSymbol:"), objc.String(value))
}


// The string used to represent a NaN (“not a number”) value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/notANumberSymbol
func (n_ NumberFormatter) NotANumberSymbol() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("notANumberSymbol"))
	return rv
}


// The string used to represent a NaN (“not a number”) value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/notANumberSymbol
func (n_ NumberFormatter) SetNotANumberSymbol(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNotANumberSymbol:"), objc.String(value))
}


// The number style used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/numberStyle
func (n_ NumberFormatter) NumberStyle() NumberFormatterStyle {
	rv := objc.Send[NumberFormatterStyle](n_.ID, objc.Sel("numberStyle"))
	return rv
}


// The number style used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/numberStyle
func (n_ NumberFormatter) SetNumberStyle(value NumberFormatterStyle) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNumberStyle:"), value)
}


// The string that the receiver uses to pad numbers in the formatted string representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/paddingCharacter
func (n_ NumberFormatter) PaddingCharacter() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("paddingCharacter"))
	return rv
}


// The string that the receiver uses to pad numbers in the formatted string representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/paddingCharacter
func (n_ NumberFormatter) SetPaddingCharacter(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPaddingCharacter:"), objc.String(value))
}


// The padding position used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/paddingPosition
func (n_ NumberFormatter) PaddingPosition() NumberFormatterPadPosition {
	rv := objc.Send[NumberFormatterPadPosition](n_.ID, objc.Sel("paddingPosition"))
	return rv
}


// The padding position used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/paddingPosition
func (n_ NumberFormatter) SetPaddingPosition(value NumberFormatterPadPosition) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPaddingPosition:"), value)
}


// The string used to represent a percent symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/percentSymbol
func (n_ NumberFormatter) PercentSymbol() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("percentSymbol"))
	return rv
}


// The string used to represent a percent symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/percentSymbol
func (n_ NumberFormatter) SetPercentSymbol(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPercentSymbol:"), objc.String(value))
}


// The string used to represent a per-mill (per-thousand) symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/perMillSymbol
func (n_ NumberFormatter) PerMillSymbol() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("perMillSymbol"))
	return rv
}


// The string used to represent a per-mill (per-thousand) symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/perMillSymbol
func (n_ NumberFormatter) SetPerMillSymbol(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPerMillSymbol:"), objc.String(value))
}


// The string used to represent a plus sign.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/plusSign
func (n_ NumberFormatter) PlusSign() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("plusSign"))
	return rv
}


// The string used to represent a plus sign.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/plusSign
func (n_ NumberFormatter) SetPlusSign(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPlusSign:"), objc.String(value))
}


// The format the receiver uses to display positive values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveFormat
func (n_ NumberFormatter) PositiveFormat() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("positiveFormat"))
	return rv
}


// The format the receiver uses to display positive values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveFormat
func (n_ NumberFormatter) SetPositiveFormat(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPositiveFormat:"), objc.String(value))
}


// The string used to represent a positive infinity symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveInfinitySymbol
func (n_ NumberFormatter) PositiveInfinitySymbol() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("positiveInfinitySymbol"))
	return rv
}


// The string used to represent a positive infinity symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveInfinitySymbol
func (n_ NumberFormatter) SetPositiveInfinitySymbol(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPositiveInfinitySymbol:"), objc.String(value))
}


// The string the receiver uses as the prefix for positive values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positivePrefix
func (n_ NumberFormatter) PositivePrefix() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("positivePrefix"))
	return rv
}


// The string the receiver uses as the prefix for positive values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positivePrefix
func (n_ NumberFormatter) SetPositivePrefix(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPositivePrefix:"), objc.String(value))
}


// The string the receiver uses as the suffix for positive values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveSuffix
func (n_ NumberFormatter) PositiveSuffix() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("positiveSuffix"))
	return rv
}


// The string the receiver uses as the suffix for positive values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveSuffix
func (n_ NumberFormatter) SetPositiveSuffix(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPositiveSuffix:"), objc.String(value))
}


// The rounding behavior used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingBehavior
func (n_ NumberFormatter) RoundingBehavior() IDecimalNumberHandler {
	rv := objc.Send[DecimalNumberHandler](n_.ID, objc.Sel("roundingBehavior"))
	return rv
}


// The rounding behavior used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingBehavior
func (n_ NumberFormatter) SetRoundingBehavior(value IDecimalNumberHandler) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRoundingBehavior:"), value)
}


// The rounding increment used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingIncrement
func (n_ NumberFormatter) RoundingIncrement() Number /* not a class type */ {
	rv := objc.Send[Number](n_.ID, objc.Sel("roundingIncrement"))
	return rv
}


// The rounding increment used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingIncrement
func (n_ NumberFormatter) SetRoundingIncrement(value Number /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRoundingIncrement:"), value)
}


// The rounding mode used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingMode-swift.property
func (n_ NumberFormatter) RoundingMode() NumberFormatterRoundingMode {
	rv := objc.Send[NumberFormatterRoundingMode](n_.ID, objc.Sel("roundingMode"))
	return rv
}


// The rounding mode used by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingMode-swift.property
func (n_ NumberFormatter) SetRoundingMode(value NumberFormatterRoundingMode) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRoundingMode:"), value)
}


// The secondary grouping size of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/secondaryGroupingSize
func (n_ NumberFormatter) SecondaryGroupingSize() uint /* primitive/slice/pointer */ {
	rv := objc.Send[uint](n_.ID, objc.Sel("secondaryGroupingSize"))
	return rv
}


// The secondary grouping size of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/secondaryGroupingSize
func (n_ NumberFormatter) SetSecondaryGroupingSize(value uint /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSecondaryGroupingSize:"), value)
}


// The text attributes used to display the negative infinity symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNegativeInfinity
func (n_ NumberFormatter) TextAttributesForNegativeInfinity() IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](n_.ID, objc.Sel("textAttributesForNegativeInfinity"))
	return rv
}


// The text attributes used to display the negative infinity symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNegativeInfinity
func (n_ NumberFormatter) SetTextAttributesForNegativeInfinity(value IDictionary /* already interface */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTextAttributesForNegativeInfinity:"), value)
}


// The text attributes to be used in displaying negative values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNegativeValues
func (n_ NumberFormatter) TextAttributesForNegativeValues() IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](n_.ID, objc.Sel("textAttributesForNegativeValues"))
	return rv
}


// The text attributes to be used in displaying negative values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNegativeValues
func (n_ NumberFormatter) SetTextAttributesForNegativeValues(value IDictionary /* already interface */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTextAttributesForNegativeValues:"), value)
}


// The text attributes used to display the symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNil
func (n_ NumberFormatter) TextAttributesForNil() IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](n_.ID, objc.Sel("textAttributesForNil"))
	return rv
}


// The text attributes used to display the symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNil
func (n_ NumberFormatter) SetTextAttributesForNil(value IDictionary /* already interface */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTextAttributesForNil:"), value)
}


// The text attributes used to display the NaN (“not a number”) string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNotANumber
func (n_ NumberFormatter) TextAttributesForNotANumber() IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](n_.ID, objc.Sel("textAttributesForNotANumber"))
	return rv
}


// The text attributes used to display the NaN (“not a number”) string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNotANumber
func (n_ NumberFormatter) SetTextAttributesForNotANumber(value IDictionary /* already interface */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTextAttributesForNotANumber:"), value)
}


// The text attributes used to display the positive infinity symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForPositiveInfinity
func (n_ NumberFormatter) TextAttributesForPositiveInfinity() IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](n_.ID, objc.Sel("textAttributesForPositiveInfinity"))
	return rv
}


// The text attributes used to display the positive infinity symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForPositiveInfinity
func (n_ NumberFormatter) SetTextAttributesForPositiveInfinity(value IDictionary /* already interface */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTextAttributesForPositiveInfinity:"), value)
}


// The text attributes to be used in displaying positive values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForPositiveValues
func (n_ NumberFormatter) TextAttributesForPositiveValues() IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](n_.ID, objc.Sel("textAttributesForPositiveValues"))
	return rv
}


// The text attributes to be used in displaying positive values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForPositiveValues
func (n_ NumberFormatter) SetTextAttributesForPositiveValues(value IDictionary /* already interface */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTextAttributesForPositiveValues:"), value)
}


// The text attributes used to display a zero value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForZero
func (n_ NumberFormatter) TextAttributesForZero() IDictionary /* already interface */ {
	rv := objc.Send[IDictionary](n_.ID, objc.Sel("textAttributesForZero"))
	return rv
}


// The text attributes used to display a zero value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForZero
func (n_ NumberFormatter) SetTextAttributesForZero(value IDictionary /* already interface */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTextAttributesForZero:"), value)
}


// The character the receiver uses as a thousand separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/thousandSeparator
func (n_ NumberFormatter) ThousandSeparator() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("thousandSeparator"))
	return rv
}


// The character the receiver uses as a thousand separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/thousandSeparator
func (n_ NumberFormatter) SetThousandSeparator(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setThousandSeparator:"), objc.String(value))
}


// Determines whether the receiver displays the group separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/usesGroupingSeparator
func (n_ NumberFormatter) UsesGroupingSeparator() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](n_.ID, objc.Sel("usesGroupingSeparator"))
	return rv
}


// Determines whether the receiver displays the group separator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/usesGroupingSeparator
func (n_ NumberFormatter) SetUsesGroupingSeparator(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUsesGroupingSeparator:"), value)
}


// A Boolean value indicating whether the formatter uses minimum and maximum significant digits when formatting numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/usesSignificantDigits
func (n_ NumberFormatter) UsesSignificantDigits() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](n_.ID, objc.Sel("usesSignificantDigits"))
	return rv
}


// A Boolean value indicating whether the formatter uses minimum and maximum significant digits when formatting numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/usesSignificantDigits
func (n_ NumberFormatter) SetUsesSignificantDigits(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUsesSignificantDigits:"), value)
}


// The string used to represent a zero value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/zeroSymbol
func (n_ NumberFormatter) ZeroSymbol() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](n_.ID, objc.Sel("zeroSymbol"))
	return rv
}


// The string used to represent a zero value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/zeroSymbol
func (n_ NumberFormatter) SetZeroSymbol(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setZeroSymbol:"), objc.String(value))
}


// Determines whether the receiver will use heuristics to guess at the number which is intended by a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/numberformatter/islenient
func (n_ NumberFormatter) IsLenient() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](n_.ID, objc.Sel("isLenient"))
	return rv
}


// Determines whether the receiver will use heuristics to guess at the number which is intended by a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/numberformatter/islenient
func (n_ NumberFormatter) SetIsLenient(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsLenient:"), value)
}


// Determines whether partial string validation is enabled for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/numberformatter/ispartialstringvalidationenabled
func (n_ NumberFormatter) IsPartialStringValidationEnabled() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](n_.ID, objc.Sel("isPartialStringValidationEnabled"))
	return rv
}


// Determines whether partial string validation is enabled for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/numberformatter/ispartialstringvalidationenabled
func (n_ NumberFormatter) SetIsPartialStringValidationEnabled(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIsPartialStringValidationEnabled:"), value)
}



