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
	DisplayNameForKeyValue(key unsafe.Pointer, value objc.ID) unsafe.Pointer
	LocalizedStringForCollatorIdentifier(collatorIdentifier string) unsafe.Pointer
	LocalizedStringForCountryCode(countryCode string) unsafe.Pointer
	LocalizedStringForLanguageCode(languageCode string) unsafe.Pointer
	ObjectForKey(key unsafe.Pointer) objc.ID
}

// Information about linguistic, cultural, and technological conventions for use in formatting data for presentation.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. You typically use a locale to format and interpret information about and according to the user’s customs and preferences. You can initialize any number of locale instances with using one of the locale identifiers found in the array. However, you usually use a locale configured to match the preferences of the current user. Use the property to get the locale matching the current user’s preferences. If you need to be alerted when the user does make changes to region settings, register for the notification. Alternatively, you can use the property to get a locale that automatically updates with the user’s configuration settings: You can inspect a locale by reading its properties, as listed in Getting Information About a Locale. For properties containing a code or identifier, you can then obtain a string suitable for presentation to the user with the methods listed in Getting Display Information About a Locale. For example, you can report the user’s language as a string localized in that language using the autoupdating locale obtained in the previous example: You frequently use a locale in conjunction with a formatter. For example, the class has a property that ensures dates are converted to strings that match the user’s expectations about date formatting. By default, this property indicates the user’s current locale, which is usually the behavior you want, but you can instead set it to another locale instance to obtain a different output. See for more information about working with formatters. is with its Core Foundation counterpart, . See for more information on toll-free bridging.
//
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


// Returns a locale initialized from data in the given unarchiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/init(coder:)
func NewLocaleWithCoder(coder unsafe.Pointer) Locale {
	instance := getLocaleClass().Alloc()
	rv := objc.Send[Locale](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}

// Initializes a locale using a given locale identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/init(localeIdentifier:)
func NewLocaleWithLocaleIdentifier(string string) Locale {
	instance := getLocaleClass().Alloc()
	rv := objc.Send[Locale](instance.ID, objc.Sel("initWithLocaleIdentifier:"), objc.String(string))
	rv.Autorelease()
	return rv
}


// Returns the canonical identifier for a given locale identification string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/canonicalLocaleIdentifier(from:)
func (lc _LocaleClass) CanonicalLocaleIdentifierFromString(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(lc.class), objc.Sel("canonicalLocaleIdentifierFromString:"), objc.String(string))
	return rv
}

// Returns the direction of the sequence of characters in a line for the specified ISO language code.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/characterDirection(forLanguage:)
func (lc _LocaleClass) CharacterDirectionForLanguage(isoLangCode string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(lc.class), objc.Sel("characterDirectionForLanguage:"), objc.String(isoLangCode))
	return rv
}

// Returns a locale identifier from a Windows locale code.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/localeIdentifier(fromWindowsLocaleCode:)
func (lc _LocaleClass) LocaleIdentifierFromWindowsLocaleCode(lcid unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(lc.class), objc.Sel("localeIdentifierFromWindowsLocaleCode:"), lcid)
	return rv
}

// Returns a Window locale code from the locale identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/windowsLocaleCode(fromLocaleIdentifier:)
func (lc _LocaleClass) WindowsLocaleCodeFromLocaleIdentifier(localeIdentifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(lc.class), objc.Sel("windowsLocaleCodeFromLocaleIdentifier:"), objc.String(localeIdentifier))
	return rv
}

// Returns the display name for the given locale component value.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/displayName(forKey:value:)
func (l_ Locale) DisplayNameForKeyValue(key unsafe.Pointer, value objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("displayNameForKey:value:"), key, value)
	return rv
}

// Returns the localized string for the specified collator identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forCollatorIdentifier:)
func (l_ Locale) LocalizedStringForCollatorIdentifier(collatorIdentifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("localizedStringForCollatorIdentifier:"), objc.String(collatorIdentifier))
	return rv
}

// Returns the localized string for a country or region code.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forCountryCode:)
func (l_ Locale) LocalizedStringForCountryCode(countryCode string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("localizedStringForCountryCode:"), objc.String(countryCode))
	return rv
}

// Returns the localized string for the specified language code.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forLanguageCode:)
func (l_ Locale) LocalizedStringForLanguageCode(languageCode string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("localizedStringForLanguageCode:"), objc.String(languageCode))
	return rv
}

// Returns the value of the component corresponding to the specified key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/object(forKey:)
func (l_ Locale) ObjectForKey(key unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("objectForKey:"), key)
	return rv
}

// The alternate end quotation symbol for the locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/alternateQuotationEndDelimiter
func (l_ Locale) AlternateQuotationEndDelimiter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("alternateQuotationEndDelimiter"))
	return rv
}

// The calendar identifier for the locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/calendarIdentifier
func (l_ Locale) CalendarIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("calendarIdentifier"))
	return rv
}

// The collation identifier for the locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/collationIdentifier
func (l_ Locale) CollationIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("collationIdentifier"))
	return rv
}

// The collator identifier for the locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/collatorIdentifier
func (l_ Locale) CollatorIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("collatorIdentifier"))
	return rv
}

// The currency code for the locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/currencyCode
func (l_ Locale) CurrencyCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("currencyCode"))
	return rv
}

// The decimal separator for the locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/decimalSeparator
func (l_ Locale) DecimalSeparator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("decimalSeparator"))
	return rv
}

// The grouping separator for the locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/groupingSeparator
func (l_ Locale) GroupingSeparator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("groupingSeparator"))
	return rv
}

// The language code for the locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/languageCode
func (l_ Locale) LanguageCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("languageCode"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/languageIdentifier
func (l_ Locale) LanguageIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("languageIdentifier"))
	return rv
}

// The begin quotation symbol for the locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/quotationBeginDelimiter
func (l_ Locale) QuotationBeginDelimiter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("quotationBeginDelimiter"))
	return rv
}

// The end quotation symbol for the locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/quotationEndDelimiter
func (l_ Locale) QuotationEndDelimiter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("quotationEndDelimiter"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/regionCode
func (l_ Locale) RegionCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("regionCode"))
	return rv
}

// The script code for the locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/scriptCode
func (l_ Locale) ScriptCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("scriptCode"))
	return rv
}

// A Boolean value that indicates whether the locale uses the metric system.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/usesMetricSystem
func (l_ Locale) UsesMetricSystem() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("usesMetricSystem"))
	return rv
}


