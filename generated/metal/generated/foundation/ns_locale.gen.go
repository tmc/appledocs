// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Locale] class.
var (
	LocaleClass     _LocaleClass
	LocaleClassOnce sync.Once
)

func getLocaleClass() _LocaleClass {
	LocaleClassOnce.Do(func() {
		LocaleClass = _LocaleClass{objc.GetClass("NSLocale")}
	})
	return LocaleClass
}

type _LocaleClass struct {
	class objc.Class
}

// An interface definition for the [Locale] class.
type ILocale interface {
	objectivec.IObject
	// properties:
	LanguageCode() IString
	Locale() ILocale
	SetLocale(value ILocale)
	AlternateQuotationBeginDelimiter() IString
	SetAlternateQuotationBeginDelimiter(value IString)
	AlternateQuotationEndDelimiter() IString
	SetAlternateQuotationEndDelimiter(value IString)
	CalendarIdentifier() IString
	SetCalendarIdentifier(value IString)
	CollationIdentifier() IString
	SetCollationIdentifier(value IString)
	CollatorIdentifier() IString
	SetCollatorIdentifier(value IString)
	CountryCode() IString
	SetCountryCode(value IString)
	CurrencyCode() IString
	SetCurrencyCode(value IString)
	CurrencySymbol() IString
	SetCurrencySymbol(value IString)
	DecimalSeparator() IString
	SetDecimalSeparator(value IString)
	ExemplarCharacterSet() ICharacterSet
	SetExemplarCharacterSet(value ICharacterSet)
	GroupingSeparator() IString
	SetGroupingSeparator(value IString)
	LanguageIdentifier() IString
	SetLanguageIdentifier(value IString)
	LocaleIdentifier() IString
	SetLocaleIdentifier(value IString)
	QuotationBeginDelimiter() IString
	SetQuotationBeginDelimiter(value IString)
	QuotationEndDelimiter() IString
	SetQuotationEndDelimiter(value IString)
	RegionCode() IString
	SetRegionCode(value IString)
	ScriptCode() IString
	SetScriptCode(value IString)
	UsesMetricSystem() bool
	SetUsesMetricSystem(value bool)
	VariantCode() IString
	SetVariantCode(value IString)
	// methods:
}

// Information about linguistic, cultural, and technological conventions for use in formatting data for presentation.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. You typically use a locale to format and interpret information about and according to the user’s customs and preferences. You can initialize any number of locale instances with using one of the locale identifiers found in the array. However, you usually use a locale configured to match the preferences of the current user. Use the property to get the locale matching the current user’s preferences. If you need to be alerted when the user does make changes to region settings, register for the notification. Alternatively, you can use the property to get a locale that automatically updates with the user’s configuration settings: You can inspect a locale by reading its properties, as listed in Getting Information About a Locale. For properties containing a code or identifier, you can then obtain a string suitable for presentation to the user with the methods listed in Getting Display Information About a Locale. For example, you can report the user’s language as a string localized in that language using the autoupdating locale obtained in the previous example: You frequently use a locale in conjunction with a formatter. For example, the class has a property that ensures dates are converted to strings that match the user’s expectations about date formatting. By default, this property indicates the user’s current locale, which is usually the behavior you want, but you can instead set it to another locale instance to obtain a different output. See for more information about working with formatters. is with its Core Foundation counterpart, . See for more information on toll-free bridging.


// Information about linguistic, cultural, and technological conventions for use in formatting data for presentation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale
type Locale struct {
	objectivec.Object
}

// LocaleFrom constructs a [Locale] from an unsafe.Pointer.
//
// Information about linguistic, cultural, and technological conventions for use in formatting data for presentation.
func LocaleFrom(ptr unsafe.Pointer) Locale {
	return Locale{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LocaleClass) Alloc() Locale {
	rv := objc.Send[Locale](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LocaleClass) New() Locale {
	rv := objc.Send[Locale](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ Locale) Init() Locale {
	rv := objc.Send[Locale](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ Locale) Autorelease() Locale {
	rv := objc.Send[Locale](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLocale creates a new Locale instance.
func NewLocale() Locale {
	return getLocaleClass().New()
}



// The language code for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/languageCode
func (l_ Locale) LanguageCode() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("languageCode"))
	return rv
}


// The locale for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/locale
func (l_ Locale) Locale() ILocale {
	rv := objc.Send[Locale](l_.ID, objc.Sel("locale"))
	return rv
}


// The locale for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/locale
func (l_ Locale) SetLocale(value ILocale) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLocale:"), value)
}


// The alternate begin quotation symbol for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/alternatequotationbegindelimiter
func (l_ Locale) AlternateQuotationBeginDelimiter() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("alternateQuotationBeginDelimiter"))
	return rv
}


// The alternate begin quotation symbol for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/alternatequotationbegindelimiter
func (l_ Locale) SetAlternateQuotationBeginDelimiter(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAlternateQuotationBeginDelimiter:"), value)
}


// The alternate end quotation symbol for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/alternatequotationenddelimiter
func (l_ Locale) AlternateQuotationEndDelimiter() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("alternateQuotationEndDelimiter"))
	return rv
}


// The alternate end quotation symbol for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/alternatequotationenddelimiter
func (l_ Locale) SetAlternateQuotationEndDelimiter(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAlternateQuotationEndDelimiter:"), value)
}


// The calendar identifier for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/calendaridentifier
func (l_ Locale) CalendarIdentifier() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("calendarIdentifier"))
	return rv
}


// The calendar identifier for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/calendaridentifier
func (l_ Locale) SetCalendarIdentifier(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCalendarIdentifier:"), value)
}


// The collation identifier for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/collationidentifier
func (l_ Locale) CollationIdentifier() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("collationIdentifier"))
	return rv
}


// The collation identifier for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/collationidentifier
func (l_ Locale) SetCollationIdentifier(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCollationIdentifier:"), value)
}


// The collator identifier for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/collatoridentifier
func (l_ Locale) CollatorIdentifier() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("collatorIdentifier"))
	return rv
}


// The collator identifier for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/collatoridentifier
func (l_ Locale) SetCollatorIdentifier(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCollatorIdentifier:"), value)
}


// The country or region code for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/countrycode
func (l_ Locale) CountryCode() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("countryCode"))
	return rv
}


// The country or region code for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/countrycode
func (l_ Locale) SetCountryCode(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCountryCode:"), value)
}


// The currency code for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/currencycode
func (l_ Locale) CurrencyCode() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("currencyCode"))
	return rv
}


// The currency code for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/currencycode
func (l_ Locale) SetCurrencyCode(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCurrencyCode:"), value)
}


// The currency symbol for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/currencysymbol
func (l_ Locale) CurrencySymbol() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("currencySymbol"))
	return rv
}


// The currency symbol for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/currencysymbol
func (l_ Locale) SetCurrencySymbol(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCurrencySymbol:"), value)
}


// The decimal separator for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/decimalseparator
func (l_ Locale) DecimalSeparator() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("decimalSeparator"))
	return rv
}


// The decimal separator for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/decimalseparator
func (l_ Locale) SetDecimalSeparator(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDecimalSeparator:"), value)
}


// The exemplar character set for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/exemplarcharacterset
func (l_ Locale) ExemplarCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](l_.ID, objc.Sel("exemplarCharacterSet"))
	return rv
}


// The exemplar character set for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/exemplarcharacterset
func (l_ Locale) SetExemplarCharacterSet(value ICharacterSet) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setExemplarCharacterSet:"), value)
}


// The grouping separator for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/groupingseparator
func (l_ Locale) GroupingSeparator() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("groupingSeparator"))
	return rv
}


// The grouping separator for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/groupingseparator
func (l_ Locale) SetGroupingSeparator(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setGroupingSeparator:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/languageidentifier
func (l_ Locale) LanguageIdentifier() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("languageIdentifier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/languageidentifier
func (l_ Locale) SetLanguageIdentifier(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLanguageIdentifier:"), value)
}


// The identifier for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/localeidentifier
func (l_ Locale) LocaleIdentifier() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("localeIdentifier"))
	return rv
}


// The identifier for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/localeidentifier
func (l_ Locale) SetLocaleIdentifier(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLocaleIdentifier:"), value)
}


// The begin quotation symbol for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/quotationbegindelimiter
func (l_ Locale) QuotationBeginDelimiter() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("quotationBeginDelimiter"))
	return rv
}


// The begin quotation symbol for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/quotationbegindelimiter
func (l_ Locale) SetQuotationBeginDelimiter(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setQuotationBeginDelimiter:"), value)
}


// The end quotation symbol for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/quotationenddelimiter
func (l_ Locale) QuotationEndDelimiter() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("quotationEndDelimiter"))
	return rv
}


// The end quotation symbol for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/quotationenddelimiter
func (l_ Locale) SetQuotationEndDelimiter(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setQuotationEndDelimiter:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/regioncode
func (l_ Locale) RegionCode() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("regionCode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/regioncode
func (l_ Locale) SetRegionCode(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setRegionCode:"), value)
}


// The script code for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/scriptcode
func (l_ Locale) ScriptCode() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("scriptCode"))
	return rv
}


// The script code for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/scriptcode
func (l_ Locale) SetScriptCode(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setScriptCode:"), value)
}


// A Boolean value that indicates whether the locale uses the metric system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/usesmetricsystem
func (l_ Locale) UsesMetricSystem() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("usesMetricSystem"))
	return rv
}


// A Boolean value that indicates whether the locale uses the metric system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/usesmetricsystem
func (l_ Locale) SetUsesMetricSystem(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setUsesMetricSystem:"), value)
}


// The variant code for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/variantcode
func (l_ Locale) VariantCode() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("variantCode"))
	return rv
}


// The variant code for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nslocale/variantcode
func (l_ Locale) SetVariantCode(value IString) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setVariantCode:"), value)
}



