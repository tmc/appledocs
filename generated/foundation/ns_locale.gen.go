// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSLocale */


/* debug [class_header]: Header for NSLocale */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Locale */
// An interface definition for the [Locale] class.
type ILocale interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Locale */
	// properties:
	AlternateQuotationBeginDelimiter() IString
	AlternateQuotationEndDelimiter() IString
	CalendarIdentifier() IString
	CollationIdentifier() IString
	CollatorIdentifier() IString
	CountryCode() IString
	CurrencyCode() IString
	CurrencySymbol() IString
	DecimalSeparator() IString
	ExemplarCharacterSet() ICharacterSet
	GroupingSeparator() IString
	LanguageCode() IString
	LanguageIdentifier() IString
	LocaleIdentifier() IString
	QuotationBeginDelimiter() IString
	QuotationEndDelimiter() IString
	RegionCode() IString
	ScriptCode() IString
	UsesMetricSystem() bool
	VariantCode() IString
	Locale() ILocale
	SetLocale(value ILocale)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Locale */
	// methods:
	DisplayNameForKeyValue(key LocaleKey, value objc.IObject) IString
	LocalizedStringForCalendarIdentifier(calendarIdentifier IString) IString
	LocalizedStringForCollationIdentifier(collationIdentifier IString) IString
	LocalizedStringForCollatorIdentifier(collatorIdentifier IString) IString
	LocalizedStringForCountryCode(countryCode IString) IString
	LocalizedStringForCurrencyCode(currencyCode IString) IString
	LocalizedStringForLanguageCode(languageCode IString) IString
	LocalizedStringForLocaleIdentifier(localeIdentifier IString) IString
	LocalizedStringForScriptCode(scriptCode IString) IString
	LocalizedStringForVariantCode(variantCode IString) IString
	ObjectForKey(key LocaleKey) objc.ID
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Locale */
// Alloc allocates a new instance without initialization.
func (lc _LocaleClass) Alloc() Locale {
	rv := objc.Send[Locale](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Locale */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Locale */

// Returns a locale initialized from data in the given unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/init(coder:)
func NewLocaleWithCoder(coder ICoder) Locale {
	instance := getLocaleClass().Alloc()
	rv := objc.Send[Locale](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewLocaleWithCoder */


// Initializes a locale using a given locale identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/init(localeIdentifier:)
func NewLocaleWithLocaleIdentifier(string_ IString) Locale {
	instance := getLocaleClass().Alloc()
	rv := objc.Send[Locale](instance.ID, objc.Sel("initWithLocaleIdentifier:"), string_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewLocaleWithLocaleIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Locale */

// Returns a canonical language identifier by mapping an arbitrary locale identification string to the canonical identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/canonicalLanguageIdentifier(from:)
func (lc _LocaleClass) CanonicalLanguageIdentifierFromString(string_ IString) IString {
	rv := objc.Send[String](objc.ID(lc.class), objc.Sel("canonicalLanguageIdentifierFromString:"), string_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CanonicalLanguageIdentifierFromString) */


// Returns the canonical identifier for a given locale identification string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/canonicalLocaleIdentifier(from:)
func (lc _LocaleClass) CanonicalLocaleIdentifierFromString(string_ IString) IString {
	rv := objc.Send[String](objc.ID(lc.class), objc.Sel("canonicalLocaleIdentifierFromString:"), string_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CanonicalLocaleIdentifierFromString) */


// Returns the direction of the sequence of characters in a line for the specified ISO language code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/characterDirection(forLanguage:)
func (lc _LocaleClass) CharacterDirectionForLanguage(isoLangCode IString) LocaleLanguageDirection {
	rv := objc.Send[LocaleLanguageDirection](objc.ID(lc.class), objc.Sel("characterDirectionForLanguage:"), isoLangCode)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CharacterDirectionForLanguage) */


// Returns a dictionary that is the result of parsing a locale ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/components(fromLocaleIdentifier:)
func (lc _LocaleClass) ComponentsFromLocaleIdentifier(string_ IString) IDictionary {
	rv := objc.Send[Dictionary](objc.ID(lc.class), objc.Sel("componentsFromLocaleIdentifier:"), string_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ComponentsFromLocaleIdentifier) */


// Returns the direction of the sequence of lines for the specified ISO language code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/lineDirection(forLanguage:)
func (lc _LocaleClass) LineDirectionForLanguage(isoLangCode IString) LocaleLanguageDirection {
	rv := objc.Send[LocaleLanguageDirection](objc.ID(lc.class), objc.Sel("lineDirectionForLanguage:"), isoLangCode)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LineDirectionForLanguage) */


// Returns a locale identifier from the components specified in a given dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/localeIdentifier(fromComponents:)
func (lc _LocaleClass) LocaleIdentifierFromComponents(dict IDictionary) IString {
	rv := objc.Send[String](objc.ID(lc.class), objc.Sel("localeIdentifierFromComponents:"), dict)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LocaleIdentifierFromComponents) */


// Returns a locale identifier from a Windows locale code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/localeIdentifier(fromWindowsLocaleCode:)
func (lc _LocaleClass) LocaleIdentifierFromWindowsLocaleCode(lcid uint32 /* not a class type */) IString {
	rv := objc.Send[String](objc.ID(lc.class), objc.Sel("localeIdentifierFromWindowsLocaleCode:"), lcid)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LocaleIdentifierFromWindowsLocaleCode) */


// Returns a locale initialized using the given locale identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/localeWithLocaleIdentifier:
func (lc _LocaleClass) LocaleWithLocaleIdentifier(ident IString) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(lc.class), objc.Sel("localeWithLocaleIdentifier:"), ident)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LocaleWithLocaleIdentifier) */


// Returns a Window locale code from the locale identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/windowsLocaleCode(fromLocaleIdentifier:)
func (lc _LocaleClass) WindowsLocaleCodeFromLocaleIdentifier(localeIdentifier IString) uint32 /* not a class type */ {
	rv := objc.Send[uint32](objc.ID(lc.class), objc.Sel("windowsLocaleCodeFromLocaleIdentifier:"), localeIdentifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WindowsLocaleCodeFromLocaleIdentifier) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Locale */

// A locale which tracks the user’s current preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/autoupdatingCurrent
func (lc _LocaleClass) AutoupdatingCurrentLocale() Locale {
	rv := objc.Send[Locale](objc.ID(lc.class), objc.Sel("autoupdatingCurrentLocale"))
	return rv
}/* debug [class_properties_class/property]: autoupdatingCurrentLocale */

// The list of locale identifiers available on the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/availableLocaleIdentifiers
func (lc _LocaleClass) AvailableLocaleIdentifiers() []string {
	rv := objc.Send[[]string](objc.ID(lc.class), objc.Sel("availableLocaleIdentifiers"))
	return rv
}/* debug [class_properties_class/property]: availableLocaleIdentifiers */

// A list of commonly encountered currency codes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/commonISOCurrencyCodes
func (lc _LocaleClass) CommonISOCurrencyCodes() []string {
	rv := objc.Send[[]string](objc.ID(lc.class), objc.Sel("commonISOCurrencyCodes"))
	return rv
}/* debug [class_properties_class/property]: commonISOCurrencyCodes */

// A locale that represents the user’s region settings at the time the property is read.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/current
func (lc _LocaleClass) CurrentLocale() Locale {
	rv := objc.Send[Locale](objc.ID(lc.class), objc.Sel("currentLocale"))
	return rv
}/* debug [class_properties_class/property]: currentLocale */

// The list of known country or region codes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/isoCountryCodes
func (lc _LocaleClass) ISOCountryCodes() []string {
	rv := objc.Send[[]string](objc.ID(lc.class), objc.Sel("ISOCountryCodes"))
	return rv
}/* debug [class_properties_class/property]: ISOCountryCodes */

// The list of known currency codes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/isoCurrencyCodes
func (lc _LocaleClass) ISOCurrencyCodes() []string {
	rv := objc.Send[[]string](objc.ID(lc.class), objc.Sel("ISOCurrencyCodes"))
	return rv
}/* debug [class_properties_class/property]: ISOCurrencyCodes */

// The list of known language codes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/isoLanguageCodes
func (lc _LocaleClass) ISOLanguageCodes() []string {
	rv := objc.Send[[]string](objc.ID(lc.class), objc.Sel("ISOLanguageCodes"))
	return rv
}/* debug [class_properties_class/property]: ISOLanguageCodes */

// An ordered list of the user’s preferred languages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/preferredLanguages
func (lc _LocaleClass) PreferredLanguages() []string {
	rv := objc.Send[[]string](objc.ID(lc.class), objc.Sel("preferredLanguages"))
	return rv
}/* debug [class_properties_class/property]: preferredLanguages */

// A locale representing the generic root values with little localization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/system
func (lc _LocaleClass) SystemLocale() Locale {
	rv := objc.Send[Locale](objc.ID(lc.class), objc.Sel("systemLocale"))
	return rv
}/* debug [class_properties_class/property]: systemLocale */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Locale */

// Returns the display name for the given locale component value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/displayName(forKey:value:)
func (l_ Locale) DisplayNameForKeyValue(key LocaleKey, value objc.IObject) IString {
	rv := objc.Send[String](l_.ID, objc.Sel("displayNameForKey:value:"), key, value)
	return rv
}/* debug [instance_methods/method]: DisplayNameForKeyValue */


// Returns the localized string for the specified calendar identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forCalendarIdentifier:)
func (l_ Locale) LocalizedStringForCalendarIdentifier(calendarIdentifier IString) IString {
	rv := objc.Send[String](l_.ID, objc.Sel("localizedStringForCalendarIdentifier:"), calendarIdentifier)
	return rv
}/* debug [instance_methods/method]: LocalizedStringForCalendarIdentifier */


// Returns the localized string for the specified collation identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forCollationIdentifier:)
func (l_ Locale) LocalizedStringForCollationIdentifier(collationIdentifier IString) IString {
	rv := objc.Send[String](l_.ID, objc.Sel("localizedStringForCollationIdentifier:"), collationIdentifier)
	return rv
}/* debug [instance_methods/method]: LocalizedStringForCollationIdentifier */


// Returns the localized string for the specified collator identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forCollatorIdentifier:)
func (l_ Locale) LocalizedStringForCollatorIdentifier(collatorIdentifier IString) IString {
	rv := objc.Send[String](l_.ID, objc.Sel("localizedStringForCollatorIdentifier:"), collatorIdentifier)
	return rv
}/* debug [instance_methods/method]: LocalizedStringForCollatorIdentifier */


// Returns the localized string for a country or region code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forCountryCode:)
func (l_ Locale) LocalizedStringForCountryCode(countryCode IString) IString {
	rv := objc.Send[String](l_.ID, objc.Sel("localizedStringForCountryCode:"), countryCode)
	return rv
}/* debug [instance_methods/method]: LocalizedStringForCountryCode */


// Returns the localized string for the specified currency code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forCurrencyCode:)
func (l_ Locale) LocalizedStringForCurrencyCode(currencyCode IString) IString {
	rv := objc.Send[String](l_.ID, objc.Sel("localizedStringForCurrencyCode:"), currencyCode)
	return rv
}/* debug [instance_methods/method]: LocalizedStringForCurrencyCode */


// Returns the localized string for the specified language code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forLanguageCode:)
func (l_ Locale) LocalizedStringForLanguageCode(languageCode IString) IString {
	rv := objc.Send[String](l_.ID, objc.Sel("localizedStringForLanguageCode:"), languageCode)
	return rv
}/* debug [instance_methods/method]: LocalizedStringForLanguageCode */


// Returns the localized string for the specified locale identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forLocaleIdentifier:)
func (l_ Locale) LocalizedStringForLocaleIdentifier(localeIdentifier IString) IString {
	rv := objc.Send[String](l_.ID, objc.Sel("localizedStringForLocaleIdentifier:"), localeIdentifier)
	return rv
}/* debug [instance_methods/method]: LocalizedStringForLocaleIdentifier */


// Returns the localized string for the specified script code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forScriptCode:)
func (l_ Locale) LocalizedStringForScriptCode(scriptCode IString) IString {
	rv := objc.Send[String](l_.ID, objc.Sel("localizedStringForScriptCode:"), scriptCode)
	return rv
}/* debug [instance_methods/method]: LocalizedStringForScriptCode */


// Returns the localized string for the specified variant code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/localizedString(forVariantCode:)
func (l_ Locale) LocalizedStringForVariantCode(variantCode IString) IString {
	rv := objc.Send[String](l_.ID, objc.Sel("localizedStringForVariantCode:"), variantCode)
	return rv
}/* debug [instance_methods/method]: LocalizedStringForVariantCode */


// Returns the value of the component corresponding to the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/object(forKey:)
func (l_ Locale) ObjectForKey(key LocaleKey) objc.ID {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("objectForKey:"), key)
	return rv
}/* debug [instance_methods/method]: ObjectForKey */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Locale */

// The alternate begin quotation symbol for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/alternateQuotationBeginDelimiter
func (l_ Locale) AlternateQuotationBeginDelimiter() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("alternateQuotationBeginDelimiter"))
	return rv
}/* debug [instance_properties/getter]: alternateQuotationBeginDelimiter */


// The alternate end quotation symbol for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/alternateQuotationEndDelimiter
func (l_ Locale) AlternateQuotationEndDelimiter() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("alternateQuotationEndDelimiter"))
	return rv
}/* debug [instance_properties/getter]: alternateQuotationEndDelimiter */


// A locale which tracks the user’s current preferences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/autoupdatingCurrent
func (l_ Locale) AutoupdatingCurrentLocale() ILocale {
	rv := objc.Send[Locale](l_.ID, objc.Sel("autoupdatingCurrentLocale"))
	return rv
}/* debug [instance_properties/getter]: autoupdatingCurrentLocale */


// The list of locale identifiers available on the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/availableLocaleIdentifiers
func (l_ Locale) AvailableLocaleIdentifiers() []string {
	rv := objc.Send[[]string](l_.ID, objc.Sel("availableLocaleIdentifiers"))
	return rv
}/* debug [instance_properties/getter]: availableLocaleIdentifiers */


// The calendar identifier for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/calendarIdentifier
func (l_ Locale) CalendarIdentifier() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("calendarIdentifier"))
	return rv
}/* debug [instance_properties/getter]: calendarIdentifier */


// The collation identifier for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/collationIdentifier
func (l_ Locale) CollationIdentifier() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("collationIdentifier"))
	return rv
}/* debug [instance_properties/getter]: collationIdentifier */


// The collator identifier for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/collatorIdentifier
func (l_ Locale) CollatorIdentifier() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("collatorIdentifier"))
	return rv
}/* debug [instance_properties/getter]: collatorIdentifier */


// A list of commonly encountered currency codes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/commonISOCurrencyCodes
func (l_ Locale) CommonISOCurrencyCodes() []string {
	rv := objc.Send[[]string](l_.ID, objc.Sel("commonISOCurrencyCodes"))
	return rv
}/* debug [instance_properties/getter]: commonISOCurrencyCodes */


// The country or region code for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/countryCode
func (l_ Locale) CountryCode() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("countryCode"))
	return rv
}/* debug [instance_properties/getter]: countryCode */


// The currency code for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/currencyCode
func (l_ Locale) CurrencyCode() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("currencyCode"))
	return rv
}/* debug [instance_properties/getter]: currencyCode */


// The currency symbol for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/currencySymbol
func (l_ Locale) CurrencySymbol() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("currencySymbol"))
	return rv
}/* debug [instance_properties/getter]: currencySymbol */


// A locale that represents the user’s region settings at the time the property is read.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/current
func (l_ Locale) CurrentLocale() ILocale {
	rv := objc.Send[Locale](l_.ID, objc.Sel("currentLocale"))
	return rv
}/* debug [instance_properties/getter]: currentLocale */


// The decimal separator for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/decimalSeparator
func (l_ Locale) DecimalSeparator() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("decimalSeparator"))
	return rv
}/* debug [instance_properties/getter]: decimalSeparator */


// The exemplar character set for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/exemplarCharacterSet
func (l_ Locale) ExemplarCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](l_.ID, objc.Sel("exemplarCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: exemplarCharacterSet */


// The grouping separator for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/groupingSeparator
func (l_ Locale) GroupingSeparator() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("groupingSeparator"))
	return rv
}/* debug [instance_properties/getter]: groupingSeparator */


// The list of known country or region codes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/isoCountryCodes
func (l_ Locale) ISOCountryCodes() []string {
	rv := objc.Send[[]string](l_.ID, objc.Sel("ISOCountryCodes"))
	return rv
}/* debug [instance_properties/getter]: ISOCountryCodes */


// The list of known currency codes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/isoCurrencyCodes
func (l_ Locale) ISOCurrencyCodes() []string {
	rv := objc.Send[[]string](l_.ID, objc.Sel("ISOCurrencyCodes"))
	return rv
}/* debug [instance_properties/getter]: ISOCurrencyCodes */


// The list of known language codes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/isoLanguageCodes
func (l_ Locale) ISOLanguageCodes() []string {
	rv := objc.Send[[]string](l_.ID, objc.Sel("ISOLanguageCodes"))
	return rv
}/* debug [instance_properties/getter]: ISOLanguageCodes */


// The language code for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/languageCode
func (l_ Locale) LanguageCode() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("languageCode"))
	return rv
}/* debug [instance_properties/getter]: languageCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/languageIdentifier
func (l_ Locale) LanguageIdentifier() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("languageIdentifier"))
	return rv
}/* debug [instance_properties/getter]: languageIdentifier */


// The identifier for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/localeIdentifier
func (l_ Locale) LocaleIdentifier() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("localeIdentifier"))
	return rv
}/* debug [instance_properties/getter]: localeIdentifier */


// An ordered list of the user’s preferred languages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/preferredLanguages
func (l_ Locale) PreferredLanguages() []string {
	rv := objc.Send[[]string](l_.ID, objc.Sel("preferredLanguages"))
	return rv
}/* debug [instance_properties/getter]: preferredLanguages */


// The begin quotation symbol for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/quotationBeginDelimiter
func (l_ Locale) QuotationBeginDelimiter() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("quotationBeginDelimiter"))
	return rv
}/* debug [instance_properties/getter]: quotationBeginDelimiter */


// The end quotation symbol for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/quotationEndDelimiter
func (l_ Locale) QuotationEndDelimiter() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("quotationEndDelimiter"))
	return rv
}/* debug [instance_properties/getter]: quotationEndDelimiter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/regionCode
func (l_ Locale) RegionCode() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("regionCode"))
	return rv
}/* debug [instance_properties/getter]: regionCode */


// The script code for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/scriptCode
func (l_ Locale) ScriptCode() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("scriptCode"))
	return rv
}/* debug [instance_properties/getter]: scriptCode */


// A locale representing the generic root values with little localization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/system
func (l_ Locale) SystemLocale() ILocale {
	rv := objc.Send[Locale](l_.ID, objc.Sel("systemLocale"))
	return rv
}/* debug [instance_properties/getter]: systemLocale */


// A Boolean value that indicates whether the locale uses the metric system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/usesMetricSystem
func (l_ Locale) UsesMetricSystem() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("usesMetricSystem"))
	return rv
}/* debug [instance_properties/getter]: usesMetricSystem */


// The variant code for the locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/variantCode
func (l_ Locale) VariantCode() IString {
	rv := objc.Send[String](l_.ID, objc.Sel("variantCode"))
	return rv
}/* debug [instance_properties/getter]: variantCode */


// The locale for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/locale
func (l_ Locale) Locale() ILocale {
	rv := objc.Send[Locale](l_.ID, objc.Sel("locale"))
	return rv
}/* debug [instance_properties/getter]: locale */


// The locale for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/dateformatter/locale
func (l_ Locale) SetLocale(value ILocale) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLocale:"), value)
}/* debug [instance_properties/setter]: locale */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSLocale */


