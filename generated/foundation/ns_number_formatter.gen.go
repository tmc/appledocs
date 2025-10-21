// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	GetObjectValueForStringRangeError(obj objc.ID, string_ string, rangep Range, error_ unsafe.Pointer) bool
	NumberFromString(string_ string) unsafe.Pointer
	StringFromNumber(number unsafe.Pointer) string
}

// A formatter that converts between numeric values and their textual representations.
//
// Instances of format the textual representation of cells that contain objects and convert textual representations of numeric values into objects. The representation encompasses integers, floats, and doubles; floats and doubles can be formatted to a specified decimal position. objects can also impose ranges on the numeric values cells can accept.
//
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/defaultFormatterBehavior()
func (nc _NumberFormatterClass) DefaultFormatterBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("defaultFormatterBehavior"))
	return rv
}

// Returns a localized number string with the specified style.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/localizedString(from:number:)
func (nc _NumberFormatterClass) LocalizedStringFromNumberNumberStyle(num unsafe.Pointer, nstyle unsafe.Pointer) string {
	rv := objc.Send[string](objc.ID(nc.class), objc.Sel("localizedStringFromNumber:numberStyle:"), num, nstyle)
	return rv
}

// Sets the default formatter behavior for new instances of .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/setDefaultFormatterBehavior(_:)
func (nc _NumberFormatterClass) SetDefaultFormatterBehavior(behavior unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(nc.class), objc.Sel("setDefaultFormatterBehavior:"), behavior)
}

// Returns by reference a cell-content object after creating it from a range of characters in a given string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/getObjectValue(_:for:range:)
func (n_ NumberFormatter) GetObjectValueForStringRangeError(obj objc.ID, string_ string, rangep Range, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("getObjectValue:forString:range:error:"), obj, objc.String(string_), rangep, error_)
	return rv
}

// Returns an object created by parsing a given string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/number(from:)
func (n_ NumberFormatter) NumberFromString(string_ string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("numberFromString:"), objc.String(string_))
	return rv
}

// Returns a string containing the formatted value of the provided number object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/string(from:)
func (n_ NumberFormatter) StringFromNumber(number unsafe.Pointer) string {
	rv := objc.Send[string](n_.ID, objc.Sel("stringFromNumber:"), number)
	return rv
}

// Determines whether the receiver allows as input floating-point values (that is, values that include the period character [ ]).
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/allowsFloats
func (n_ NumberFormatter) AllowsFloats() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("allowsFloats"))
	return rv
}


// SetAllowsFloats sets the value of the allowsFloats property.
// Determines whether the receiver allows as input floating-point values (that is, values that include the period character [ ]).

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/allowsFloats
func (n_ NumberFormatter) SetAllowsFloats(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAllowsFloats:"), value)
}
// Determines whether the receiver always shows the decimal separator, even for integer numbers.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/alwaysShowsDecimalSeparator
func (n_ NumberFormatter) AlwaysShowsDecimalSeparator() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("alwaysShowsDecimalSeparator"))
	return rv
}


// SetAlwaysShowsDecimalSeparator sets the value of the alwaysShowsDecimalSeparator property.
// Determines whether the receiver always shows the decimal separator, even for integer numbers.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/alwaysShowsDecimalSeparator
func (n_ NumberFormatter) SetAlwaysShowsDecimalSeparator(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAlwaysShowsDecimalSeparator:"), value)
}
// The attributed string the receiver uses to display values.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForNil
func (n_ NumberFormatter) AttributedStringForNil() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("attributedStringForNil"))
	return rv
}


// SetAttributedStringForNil sets the value of the attributedStringForNil property.
// The attributed string the receiver uses to display values.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForNil
func (n_ NumberFormatter) SetAttributedStringForNil(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAttributedStringForNil:"), value)
}
// The attributed string the receiver uses to display “not a number” values.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForNotANumber
func (n_ NumberFormatter) AttributedStringForNotANumber() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("attributedStringForNotANumber"))
	return rv
}


// SetAttributedStringForNotANumber sets the value of the attributedStringForNotANumber property.
// The attributed string the receiver uses to display “not a number” values.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForNotANumber
func (n_ NumberFormatter) SetAttributedStringForNotANumber(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAttributedStringForNotANumber:"), value)
}
// The attributed string that the receiver uses to display zero values.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForZero
func (n_ NumberFormatter) AttributedStringForZero() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("attributedStringForZero"))
	return rv
}


// SetAttributedStringForZero sets the value of the attributedStringForZero property.
// The attributed string that the receiver uses to display zero values.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/attributedStringForZero
func (n_ NumberFormatter) SetAttributedStringForZero(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setAttributedStringForZero:"), value)
}
// The receiver’s currency code.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyCode
func (n_ NumberFormatter) CurrencyCode() string {
	rv := objc.Send[string](n_.ID, objc.Sel("currencyCode"))
	return rv
}


// SetCurrencyCode sets the value of the currencyCode property.
// The receiver’s currency code.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyCode
func (n_ NumberFormatter) SetCurrencyCode(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCurrencyCode:"), objc.String(value))
}
// The string used by the receiver as a currency decimal separator.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyDecimalSeparator
func (n_ NumberFormatter) CurrencyDecimalSeparator() string {
	rv := objc.Send[string](n_.ID, objc.Sel("currencyDecimalSeparator"))
	return rv
}


// SetCurrencyDecimalSeparator sets the value of the currencyDecimalSeparator property.
// The string used by the receiver as a currency decimal separator.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyDecimalSeparator
func (n_ NumberFormatter) SetCurrencyDecimalSeparator(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCurrencyDecimalSeparator:"), objc.String(value))
}
// The currency grouping separator for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyGroupingSeparator
func (n_ NumberFormatter) CurrencyGroupingSeparator() string {
	rv := objc.Send[string](n_.ID, objc.Sel("currencyGroupingSeparator"))
	return rv
}


// SetCurrencyGroupingSeparator sets the value of the currencyGroupingSeparator property.
// The currency grouping separator for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencyGroupingSeparator
func (n_ NumberFormatter) SetCurrencyGroupingSeparator(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCurrencyGroupingSeparator:"), objc.String(value))
}
// The string used by the receiver as a local currency symbol.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencySymbol
func (n_ NumberFormatter) CurrencySymbol() string {
	rv := objc.Send[string](n_.ID, objc.Sel("currencySymbol"))
	return rv
}


// SetCurrencySymbol sets the value of the currencySymbol property.
// The string used by the receiver as a local currency symbol.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/currencySymbol
func (n_ NumberFormatter) SetCurrencySymbol(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setCurrencySymbol:"), objc.String(value))
}
// The character the receiver uses as a decimal separator.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/decimalSeparator
func (n_ NumberFormatter) DecimalSeparator() string {
	rv := objc.Send[string](n_.ID, objc.Sel("decimalSeparator"))
	return rv
}


// SetDecimalSeparator sets the value of the decimalSeparator property.
// The character the receiver uses as a decimal separator.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/decimalSeparator
func (n_ NumberFormatter) SetDecimalSeparator(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDecimalSeparator:"), objc.String(value))
}
// The string used to represent an exponent symbol.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/exponentSymbol
func (n_ NumberFormatter) ExponentSymbol() string {
	rv := objc.Send[string](n_.ID, objc.Sel("exponentSymbol"))
	return rv
}


// SetExponentSymbol sets the value of the exponentSymbol property.
// The string used to represent an exponent symbol.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/exponentSymbol
func (n_ NumberFormatter) SetExponentSymbol(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setExponentSymbol:"), objc.String(value))
}
// The receiver’s format.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/format
func (n_ NumberFormatter) Format() string {
	rv := objc.Send[string](n_.ID, objc.Sel("format"))
	return rv
}


// SetFormat sets the value of the format property.
// The receiver’s format.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/format
func (n_ NumberFormatter) SetFormat(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFormat:"), objc.String(value))
}
// The format width used by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/formatWidth
func (n_ NumberFormatter) FormatWidth() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("formatWidth"))
	return rv
}


// SetFormatWidth sets the value of the formatWidth property.
// The format width used by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/formatWidth
func (n_ NumberFormatter) SetFormatWidth(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFormatWidth:"), value)
}
// The formatter behavior of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/formatterBehavior
func (n_ NumberFormatter) FormatterBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("formatterBehavior"))
	return rv
}


// SetFormatterBehavior sets the value of the formatterBehavior property.
// The formatter behavior of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/formatterBehavior
func (n_ NumberFormatter) SetFormatterBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFormatterBehavior:"), value)
}
// The capitalization formatting context used when formatting a number.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/formattingContext
func (n_ NumberFormatter) FormattingContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("formattingContext"))
	return rv
}


// SetFormattingContext sets the value of the formattingContext property.
// The capitalization formatting context used when formatting a number.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/formattingContext
func (n_ NumberFormatter) SetFormattingContext(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setFormattingContext:"), value)
}
// The string used by the receiver for a grouping separator.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/groupingSeparator
func (n_ NumberFormatter) GroupingSeparator() string {
	rv := objc.Send[string](n_.ID, objc.Sel("groupingSeparator"))
	return rv
}


// SetGroupingSeparator sets the value of the groupingSeparator property.
// The string used by the receiver for a grouping separator.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/groupingSeparator
func (n_ NumberFormatter) SetGroupingSeparator(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setGroupingSeparator:"), objc.String(value))
}
// The grouping size of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/groupingSize
func (n_ NumberFormatter) GroupingSize() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("groupingSize"))
	return rv
}


// SetGroupingSize sets the value of the groupingSize property.
// The grouping size of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/groupingSize
func (n_ NumberFormatter) SetGroupingSize(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setGroupingSize:"), value)
}
// Determines whether the receiver uses thousand separators.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/hasThousandSeparators
func (n_ NumberFormatter) HasThousandSeparators() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("hasThousandSeparators"))
	return rv
}


// SetHasThousandSeparators sets the value of the hasThousandSeparators property.
// Determines whether the receiver uses thousand separators.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/hasThousandSeparators
func (n_ NumberFormatter) SetHasThousandSeparators(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setHasThousandSeparators:"), value)
}
// The international currency symbol used by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/internationalCurrencySymbol
func (n_ NumberFormatter) InternationalCurrencySymbol() string {
	rv := objc.Send[string](n_.ID, objc.Sel("internationalCurrencySymbol"))
	return rv
}


// SetInternationalCurrencySymbol sets the value of the internationalCurrencySymbol property.
// The international currency symbol used by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/internationalCurrencySymbol
func (n_ NumberFormatter) SetInternationalCurrencySymbol(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setInternationalCurrencySymbol:"), objc.String(value))
}
// Determines whether the receiver will use heuristics to guess at the number which is intended by a string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/isLenient
func (n_ NumberFormatter) Lenient() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("lenient"))
	return rv
}


// SetLenient sets the value of the lenient property.
// Determines whether the receiver will use heuristics to guess at the number which is intended by a string.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/isLenient
func (n_ NumberFormatter) SetLenient(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLenient:"), value)
}
// Determines whether partial string validation is enabled for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/isPartialStringValidationEnabled
func (n_ NumberFormatter) PartialStringValidationEnabled() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("partialStringValidationEnabled"))
	return rv
}


// SetPartialStringValidationEnabled sets the value of the partialStringValidationEnabled property.
// Determines whether partial string validation is enabled for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/isPartialStringValidationEnabled
func (n_ NumberFormatter) SetPartialStringValidationEnabled(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPartialStringValidationEnabled:"), value)
}
// The locale of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/locale
func (n_ NumberFormatter) Locale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("locale"))
	return rv
}


// SetLocale sets the value of the locale property.
// The locale of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/locale
func (n_ NumberFormatter) SetLocale(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocale:"), value)
}
// Determines whether the dollar sign character ( ), decimal separator character ( ), and thousand separator character ( ) are converted to appropriately localized characters as specified by the user’s localization preference.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/localizesFormat
func (n_ NumberFormatter) LocalizesFormat() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("localizesFormat"))
	return rv
}


// SetLocalizesFormat sets the value of the localizesFormat property.
// Determines whether the dollar sign character ( ), decimal separator character ( ), and thousand separator character ( ) are converted to appropriately localized characters as specified by the user’s localization preference.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/localizesFormat
func (n_ NumberFormatter) SetLocalizesFormat(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setLocalizesFormat:"), value)
}
// The highest number allowed as input by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximum
func (n_ NumberFormatter) Maximum() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("maximum"))
	return rv
}


// SetMaximum sets the value of the maximum property.
// The highest number allowed as input by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximum
func (n_ NumberFormatter) SetMaximum(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMaximum:"), value)
}
// The maximum number of digits after the decimal separator.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumFractionDigits
func (n_ NumberFormatter) MaximumFractionDigits() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("maximumFractionDigits"))
	return rv
}


// SetMaximumFractionDigits sets the value of the maximumFractionDigits property.
// The maximum number of digits after the decimal separator.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumFractionDigits
func (n_ NumberFormatter) SetMaximumFractionDigits(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMaximumFractionDigits:"), value)
}
// The maximum number of digits before the decimal separator.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumIntegerDigits
func (n_ NumberFormatter) MaximumIntegerDigits() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("maximumIntegerDigits"))
	return rv
}


// SetMaximumIntegerDigits sets the value of the maximumIntegerDigits property.
// The maximum number of digits before the decimal separator.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumIntegerDigits
func (n_ NumberFormatter) SetMaximumIntegerDigits(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMaximumIntegerDigits:"), value)
}
// The maximum number of significant digits for the number formatter.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumSignificantDigits
func (n_ NumberFormatter) MaximumSignificantDigits() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("maximumSignificantDigits"))
	return rv
}


// SetMaximumSignificantDigits sets the value of the maximumSignificantDigits property.
// The maximum number of significant digits for the number formatter.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/maximumSignificantDigits
func (n_ NumberFormatter) SetMaximumSignificantDigits(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMaximumSignificantDigits:"), value)
}
// The lowest number allowed as input by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimum
func (n_ NumberFormatter) Minimum() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("minimum"))
	return rv
}


// SetMinimum sets the value of the minimum property.
// The lowest number allowed as input by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimum
func (n_ NumberFormatter) SetMinimum(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimum:"), value)
}
// The minimum number of digits after the decimal separator.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumFractionDigits
func (n_ NumberFormatter) MinimumFractionDigits() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("minimumFractionDigits"))
	return rv
}


// SetMinimumFractionDigits sets the value of the minimumFractionDigits property.
// The minimum number of digits after the decimal separator.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumFractionDigits
func (n_ NumberFormatter) SetMinimumFractionDigits(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimumFractionDigits:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumGroupingDigits
func (n_ NumberFormatter) MinimumGroupingDigits() int {
	rv := objc.Send[int](n_.ID, objc.Sel("minimumGroupingDigits"))
	return rv
}


// SetMinimumGroupingDigits sets the value of the minimumGroupingDigits property.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumGroupingDigits
func (n_ NumberFormatter) SetMinimumGroupingDigits(value int) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimumGroupingDigits:"), value)
}
// The minimum number of digits before the decimal separator.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumIntegerDigits
func (n_ NumberFormatter) MinimumIntegerDigits() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("minimumIntegerDigits"))
	return rv
}


// SetMinimumIntegerDigits sets the value of the minimumIntegerDigits property.
// The minimum number of digits before the decimal separator.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumIntegerDigits
func (n_ NumberFormatter) SetMinimumIntegerDigits(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimumIntegerDigits:"), value)
}
// The minimum number of significant digits for the number formatter.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumSignificantDigits
func (n_ NumberFormatter) MinimumSignificantDigits() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("minimumSignificantDigits"))
	return rv
}


// SetMinimumSignificantDigits sets the value of the minimumSignificantDigits property.
// The minimum number of significant digits for the number formatter.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minimumSignificantDigits
func (n_ NumberFormatter) SetMinimumSignificantDigits(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinimumSignificantDigits:"), value)
}
// The string used to represent a minus sign.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minusSign
func (n_ NumberFormatter) MinusSign() string {
	rv := objc.Send[string](n_.ID, objc.Sel("minusSign"))
	return rv
}


// SetMinusSign sets the value of the minusSign property.
// The string used to represent a minus sign.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/minusSign
func (n_ NumberFormatter) SetMinusSign(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMinusSign:"), objc.String(value))
}
// The multiplier of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/multiplier
func (n_ NumberFormatter) Multiplier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("multiplier"))
	return rv
}


// SetMultiplier sets the value of the multiplier property.
// The multiplier of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/multiplier
func (n_ NumberFormatter) SetMultiplier(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMultiplier:"), value)
}
// The format the receiver uses to display negative values.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeFormat
func (n_ NumberFormatter) NegativeFormat() string {
	rv := objc.Send[string](n_.ID, objc.Sel("negativeFormat"))
	return rv
}


// SetNegativeFormat sets the value of the negativeFormat property.
// The format the receiver uses to display negative values.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeFormat
func (n_ NumberFormatter) SetNegativeFormat(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNegativeFormat:"), objc.String(value))
}
// The string used to represent a negative infinity symbol.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeInfinitySymbol
func (n_ NumberFormatter) NegativeInfinitySymbol() string {
	rv := objc.Send[string](n_.ID, objc.Sel("negativeInfinitySymbol"))
	return rv
}


// SetNegativeInfinitySymbol sets the value of the negativeInfinitySymbol property.
// The string used to represent a negative infinity symbol.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeInfinitySymbol
func (n_ NumberFormatter) SetNegativeInfinitySymbol(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNegativeInfinitySymbol:"), objc.String(value))
}
// The string the receiver uses as a prefix for negative values.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativePrefix
func (n_ NumberFormatter) NegativePrefix() string {
	rv := objc.Send[string](n_.ID, objc.Sel("negativePrefix"))
	return rv
}


// SetNegativePrefix sets the value of the negativePrefix property.
// The string the receiver uses as a prefix for negative values.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativePrefix
func (n_ NumberFormatter) SetNegativePrefix(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNegativePrefix:"), objc.String(value))
}
// The string the receiver uses as a suffix for negative values.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeSuffix
func (n_ NumberFormatter) NegativeSuffix() string {
	rv := objc.Send[string](n_.ID, objc.Sel("negativeSuffix"))
	return rv
}


// SetNegativeSuffix sets the value of the negativeSuffix property.
// The string the receiver uses as a suffix for negative values.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/negativeSuffix
func (n_ NumberFormatter) SetNegativeSuffix(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNegativeSuffix:"), objc.String(value))
}
// The string used to represent a value.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/nilSymbol
func (n_ NumberFormatter) NilSymbol() string {
	rv := objc.Send[string](n_.ID, objc.Sel("nilSymbol"))
	return rv
}


// SetNilSymbol sets the value of the nilSymbol property.
// The string used to represent a value.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/nilSymbol
func (n_ NumberFormatter) SetNilSymbol(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNilSymbol:"), objc.String(value))
}
// The string used to represent a NaN (“not a number”) value.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/notANumberSymbol
func (n_ NumberFormatter) NotANumberSymbol() string {
	rv := objc.Send[string](n_.ID, objc.Sel("notANumberSymbol"))
	return rv
}


// SetNotANumberSymbol sets the value of the notANumberSymbol property.
// The string used to represent a NaN (“not a number”) value.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/notANumberSymbol
func (n_ NumberFormatter) SetNotANumberSymbol(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNotANumberSymbol:"), objc.String(value))
}
// The number style used by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/numberStyle
func (n_ NumberFormatter) NumberStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("numberStyle"))
	return rv
}


// SetNumberStyle sets the value of the numberStyle property.
// The number style used by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/numberStyle
func (n_ NumberFormatter) SetNumberStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNumberStyle:"), value)
}
// The string that the receiver uses to pad numbers in the formatted string representation.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/paddingCharacter
func (n_ NumberFormatter) PaddingCharacter() string {
	rv := objc.Send[string](n_.ID, objc.Sel("paddingCharacter"))
	return rv
}


// SetPaddingCharacter sets the value of the paddingCharacter property.
// The string that the receiver uses to pad numbers in the formatted string representation.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/paddingCharacter
func (n_ NumberFormatter) SetPaddingCharacter(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPaddingCharacter:"), objc.String(value))
}
// The padding position used by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/paddingPosition
func (n_ NumberFormatter) PaddingPosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("paddingPosition"))
	return rv
}


// SetPaddingPosition sets the value of the paddingPosition property.
// The padding position used by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/paddingPosition
func (n_ NumberFormatter) SetPaddingPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPaddingPosition:"), value)
}
// The string used to represent a per-mill (per-thousand) symbol.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/perMillSymbol
func (n_ NumberFormatter) PerMillSymbol() string {
	rv := objc.Send[string](n_.ID, objc.Sel("perMillSymbol"))
	return rv
}


// SetPerMillSymbol sets the value of the perMillSymbol property.
// The string used to represent a per-mill (per-thousand) symbol.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/perMillSymbol
func (n_ NumberFormatter) SetPerMillSymbol(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPerMillSymbol:"), objc.String(value))
}
// The string used to represent a percent symbol.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/percentSymbol
func (n_ NumberFormatter) PercentSymbol() string {
	rv := objc.Send[string](n_.ID, objc.Sel("percentSymbol"))
	return rv
}


// SetPercentSymbol sets the value of the percentSymbol property.
// The string used to represent a percent symbol.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/percentSymbol
func (n_ NumberFormatter) SetPercentSymbol(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPercentSymbol:"), objc.String(value))
}
// The string used to represent a plus sign.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/plusSign
func (n_ NumberFormatter) PlusSign() string {
	rv := objc.Send[string](n_.ID, objc.Sel("plusSign"))
	return rv
}


// SetPlusSign sets the value of the plusSign property.
// The string used to represent a plus sign.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/plusSign
func (n_ NumberFormatter) SetPlusSign(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPlusSign:"), objc.String(value))
}
// The format the receiver uses to display positive values.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveFormat
func (n_ NumberFormatter) PositiveFormat() string {
	rv := objc.Send[string](n_.ID, objc.Sel("positiveFormat"))
	return rv
}


// SetPositiveFormat sets the value of the positiveFormat property.
// The format the receiver uses to display positive values.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveFormat
func (n_ NumberFormatter) SetPositiveFormat(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPositiveFormat:"), objc.String(value))
}
// The string used to represent a positive infinity symbol.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveInfinitySymbol
func (n_ NumberFormatter) PositiveInfinitySymbol() string {
	rv := objc.Send[string](n_.ID, objc.Sel("positiveInfinitySymbol"))
	return rv
}


// SetPositiveInfinitySymbol sets the value of the positiveInfinitySymbol property.
// The string used to represent a positive infinity symbol.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveInfinitySymbol
func (n_ NumberFormatter) SetPositiveInfinitySymbol(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPositiveInfinitySymbol:"), objc.String(value))
}
// The string the receiver uses as the suffix for positive values.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveSuffix
func (n_ NumberFormatter) PositiveSuffix() string {
	rv := objc.Send[string](n_.ID, objc.Sel("positiveSuffix"))
	return rv
}


// SetPositiveSuffix sets the value of the positiveSuffix property.
// The string the receiver uses as the suffix for positive values.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/positiveSuffix
func (n_ NumberFormatter) SetPositiveSuffix(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setPositiveSuffix:"), objc.String(value))
}
// The rounding behavior used by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingBehavior
func (n_ NumberFormatter) RoundingBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("roundingBehavior"))
	return rv
}


// SetRoundingBehavior sets the value of the roundingBehavior property.
// The rounding behavior used by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingBehavior
func (n_ NumberFormatter) SetRoundingBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRoundingBehavior:"), value)
}
// The rounding increment used by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingIncrement
func (n_ NumberFormatter) RoundingIncrement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("roundingIncrement"))
	return rv
}


// SetRoundingIncrement sets the value of the roundingIncrement property.
// The rounding increment used by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingIncrement
func (n_ NumberFormatter) SetRoundingIncrement(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRoundingIncrement:"), value)
}
// The rounding mode used by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingMode-swift.property
func (n_ NumberFormatter) RoundingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("roundingMode"))
	return rv
}


// SetRoundingMode sets the value of the roundingMode property.
// The rounding mode used by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/roundingMode-swift.property
func (n_ NumberFormatter) SetRoundingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setRoundingMode:"), value)
}
// The secondary grouping size of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/secondaryGroupingSize
func (n_ NumberFormatter) SecondaryGroupingSize() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("secondaryGroupingSize"))
	return rv
}


// SetSecondaryGroupingSize sets the value of the secondaryGroupingSize property.
// The secondary grouping size of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/secondaryGroupingSize
func (n_ NumberFormatter) SetSecondaryGroupingSize(value uint) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSecondaryGroupingSize:"), value)
}
// The text attributes used to display the negative infinity symbol.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNegativeInfinity
func (n_ NumberFormatter) TextAttributesForNegativeInfinity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("textAttributesForNegativeInfinity"))
	return rv
}


// SetTextAttributesForNegativeInfinity sets the value of the textAttributesForNegativeInfinity property.
// The text attributes used to display the negative infinity symbol.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNegativeInfinity
func (n_ NumberFormatter) SetTextAttributesForNegativeInfinity(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTextAttributesForNegativeInfinity:"), value)
}
// The text attributes to be used in displaying negative values.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNegativeValues
func (n_ NumberFormatter) TextAttributesForNegativeValues() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("textAttributesForNegativeValues"))
	return rv
}


// SetTextAttributesForNegativeValues sets the value of the textAttributesForNegativeValues property.
// The text attributes to be used in displaying negative values.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNegativeValues
func (n_ NumberFormatter) SetTextAttributesForNegativeValues(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTextAttributesForNegativeValues:"), value)
}
// The text attributes used to display the symbol.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNil
func (n_ NumberFormatter) TextAttributesForNil() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("textAttributesForNil"))
	return rv
}


// SetTextAttributesForNil sets the value of the textAttributesForNil property.
// The text attributes used to display the symbol.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForNil
func (n_ NumberFormatter) SetTextAttributesForNil(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTextAttributesForNil:"), value)
}
// The text attributes used to display the positive infinity symbol.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForPositiveInfinity
func (n_ NumberFormatter) TextAttributesForPositiveInfinity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("textAttributesForPositiveInfinity"))
	return rv
}


// SetTextAttributesForPositiveInfinity sets the value of the textAttributesForPositiveInfinity property.
// The text attributes used to display the positive infinity symbol.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForPositiveInfinity
func (n_ NumberFormatter) SetTextAttributesForPositiveInfinity(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTextAttributesForPositiveInfinity:"), value)
}
// The text attributes to be used in displaying positive values.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForPositiveValues
func (n_ NumberFormatter) TextAttributesForPositiveValues() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("textAttributesForPositiveValues"))
	return rv
}


// SetTextAttributesForPositiveValues sets the value of the textAttributesForPositiveValues property.
// The text attributes to be used in displaying positive values.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForPositiveValues
func (n_ NumberFormatter) SetTextAttributesForPositiveValues(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTextAttributesForPositiveValues:"), value)
}
// The text attributes used to display a zero value.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForZero
func (n_ NumberFormatter) TextAttributesForZero() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("textAttributesForZero"))
	return rv
}


// SetTextAttributesForZero sets the value of the textAttributesForZero property.
// The text attributes used to display a zero value.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/textAttributesForZero
func (n_ NumberFormatter) SetTextAttributesForZero(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setTextAttributesForZero:"), value)
}
// The character the receiver uses as a thousand separator.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/thousandSeparator
func (n_ NumberFormatter) ThousandSeparator() string {
	rv := objc.Send[string](n_.ID, objc.Sel("thousandSeparator"))
	return rv
}


// SetThousandSeparator sets the value of the thousandSeparator property.
// The character the receiver uses as a thousand separator.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/thousandSeparator
func (n_ NumberFormatter) SetThousandSeparator(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setThousandSeparator:"), objc.String(value))
}
// Determines whether the receiver displays the group separator.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/usesGroupingSeparator
func (n_ NumberFormatter) UsesGroupingSeparator() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("usesGroupingSeparator"))
	return rv
}


// SetUsesGroupingSeparator sets the value of the usesGroupingSeparator property.
// Determines whether the receiver displays the group separator.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/usesGroupingSeparator
func (n_ NumberFormatter) SetUsesGroupingSeparator(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUsesGroupingSeparator:"), value)
}
// A Boolean value indicating whether the formatter uses minimum and maximum significant digits when formatting numbers.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/usesSignificantDigits
func (n_ NumberFormatter) UsesSignificantDigits() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("usesSignificantDigits"))
	return rv
}


// SetUsesSignificantDigits sets the value of the usesSignificantDigits property.
// A Boolean value indicating whether the formatter uses minimum and maximum significant digits when formatting numbers.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/usesSignificantDigits
func (n_ NumberFormatter) SetUsesSignificantDigits(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUsesSignificantDigits:"), value)
}
// The string used to represent a zero value.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/zeroSymbol
func (n_ NumberFormatter) ZeroSymbol() string {
	rv := objc.Send[string](n_.ID, objc.Sel("zeroSymbol"))
	return rv
}


// SetZeroSymbol sets the value of the zeroSymbol property.
// The string used to represent a zero value.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NumberFormatter/zeroSymbol
func (n_ NumberFormatter) SetZeroSymbol(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setZeroSymbol:"), objc.String(value))
}


