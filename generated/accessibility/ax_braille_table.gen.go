// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AXBrailleTable */


/* debug [class_header]: Header for AXBrailleTable */
// The class instance for the [AXBrailleTable] class.
var (
	AXBrailleTableClass     _AXBrailleTableClass
	AXBrailleTableClassOnce sync.Once
)

func getAXBrailleTableClass() _AXBrailleTableClass {
	AXBrailleTableClassOnce.Do(func() {
		AXBrailleTableClass = _AXBrailleTableClass{objc.GetClass("AXBrailleTable")}
	})
	return AXBrailleTableClass
}

type _AXBrailleTableClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXBrailleTable */
// An interface definition for the [AXBrailleTable] class.
type IAXBrailleTable interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AXBrailleTable */
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
	IsEightDot() bool
	Language() objc.IObject /* cross-framework: NSString */
	Locales() unsafe.Pointer
	LocalizedName() objc.IObject /* cross-framework: NSString */
	LocalizedProviderName() objc.IObject /* cross-framework: NSString */
	ProviderIdentifier() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXBrailleTable */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXBrailleTable */
// Alloc allocates a new instance without initialization.
func (ac _AXBrailleTableClass) Alloc() AXBrailleTable {
	rv := objc.Send[AXBrailleTable](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXBrailleTableClass) New() AXBrailleTable {
	rv := objc.Send[AXBrailleTable](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXBrailleTable) Init() AXBrailleTable {
	rv := objc.Send[AXBrailleTable](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXBrailleTable) Autorelease() AXBrailleTable {
	rv := objc.Send[AXBrailleTable](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXBrailleTable creates a new AXBrailleTable instance.
func NewAXBrailleTable() AXBrailleTable {
	return getAXBrailleTableClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXBrailleTable */
// A rule for translating print text to Braille, and back-translating Braille to print text.


// A rule for translating print text to Braille, and back-translating Braille to print text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable
type AXBrailleTable struct {
	objectivec.Object
}

// AXBrailleTableFrom constructs a [AXBrailleTable] from an unsafe.Pointer.
//
// A rule for translating print text to Braille, and back-translating Braille to print text.
func AXBrailleTableFrom(ptr unsafe.Pointer) AXBrailleTable {
	return AXBrailleTable{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXBrailleTable */

// Returns nil if there is no table with the given identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/init(identifier:)
func NewAXBrailleTableWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) AXBrailleTable {
	instance := getAXBrailleTableClass().Alloc()
	rv := objc.Send[AXBrailleTable](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXBrailleTableWithIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXBrailleTable */

// The default table that provides translations for the given locale’s language. Returns nil if there is none.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/defaultTable(for:)
func (ac _AXBrailleTableClass) DefaultTableForLocale(locale foundation.Locale) AXBrailleTable {
	rv := objc.Send[AXBrailleTable](objc.ID(ac.class), objc.Sel("defaultTableForLocale:"), locale)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultTableForLocale) */


// All tables that are not specific to any language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/languageAgnosticTables()
func (ac _AXBrailleTableClass) LanguageAgnosticTables() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("languageAgnosticTables"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LanguageAgnosticTables) */


// All locales supported by existing tables.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/supportedLocales()
func (ac _AXBrailleTableClass) SupportedLocales() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("supportedLocales"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportedLocales) */


// All tables that provide translations for the given locale’s language.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/tables(for:)
func (ac _AXBrailleTableClass) TablesForLocale(locale foundation.Locale) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("tablesForLocale:"), locale)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TablesForLocale) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXBrailleTable */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXBrailleTable */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXBrailleTable */

// A unique string that identifies this table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/identifier
func (a_ AXBrailleTable) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// Returns true if this table makes use of eight dots as opposed to six dots.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/isEightDot
func (a_ AXBrailleTable) IsEightDot() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEightDot"))
	return rv
}/* debug [instance_properties/getter]: isEightDot */


// The 3-character code from ISO 639-2 for the language this Braille table pertains to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/language-3570f
func (a_ AXBrailleTable) Language() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("language"))
	return rv
}/* debug [instance_properties/getter]: language */


// All locales this table supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/locales
func (a_ AXBrailleTable) Locales() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("locales"))
	return rv
}/* debug [instance_properties/getter]: locales */


// The localized name of this table for user display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/localizedName
func (a_ AXBrailleTable) LocalizedName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("localizedName"))
	return rv
}/* debug [instance_properties/getter]: localizedName */


// The localized name of the provider of this table for user display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/localizedProviderName
func (a_ AXBrailleTable) LocalizedProviderName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("localizedProviderName"))
	return rv
}/* debug [instance_properties/getter]: localizedProviderName */


// The identifier of the provider of this table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/providerIdentifier
func (a_ AXBrailleTable) ProviderIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("providerIdentifier"))
	return rv
}/* debug [instance_properties/getter]: providerIdentifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXBrailleTable */


