// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AXBrailleTable] class.
type IAXBrailleTable interface {
	objectivec.IObject
}

// A rule for translating print text to Braille, and back-translating Braille to print text.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AXBrailleTableClass) Alloc() AXBrailleTable {
	rv := objc.Send[AXBrailleTable](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Returns nil if there is no table with the given identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/init(identifier:)
func NewAXBrailleTableWithIdentifier(identifier appkit.string) AXBrailleTable {
	instance := getAXBrailleTableClass().Alloc()
	rv := objc.Send[AXBrailleTable](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}


// All tables that provide translations for the given locale’s language.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/tables(for:)
func (ac _AXBrailleTableClass) TablesForLocale(locale foundation.ILocale) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("tablesForLocale:"), locale)
	return rv
}

// The localized name of the provider of this table for user display.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/localizedProviderName
func (a_ AXBrailleTable) LocalizedProviderName() appkit.string {
	rv := objc.Send[appkit.string](a_.ID, objc.Sel("localizedProviderName"))
	return rv
}

// The identifier of the provider of this table.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/providerIdentifier
func (a_ AXBrailleTable) ProviderIdentifier() appkit.string {
	rv := objc.Send[appkit.string](a_.ID, objc.Sel("providerIdentifier"))
	return rv
}

// A unique string that identifies this table.
//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axbrailletable/identifier
func (a_ AXBrailleTable) Identifier() appkit.string {
	rv := objc.Send[appkit.string](a_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// A unique string that identifies this table.

//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axbrailletable/identifier
func (a_ AXBrailleTable) SetIdentifier(value appkit.string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIdentifier:"), value)
}

// Returns true if this table makes use of eight dots as opposed to six dots.
//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axbrailletable/iseightdot
func (a_ AXBrailleTable) IsEightDot() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEightDot"))
	return rv
}


// SetIsEightDot sets the value of the isEightDot property.
// Returns true if this table makes use of eight dots as opposed to six dots.

//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axbrailletable/iseightdot
func (a_ AXBrailleTable) SetIsEightDot(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsEightDot:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axbrailletable/language-3stsd
func (a_ AXBrailleTable) Language() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("language"))
	return rv
}


// SetLanguage sets the value of the language property.
//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axbrailletable/language-3stsd
func (a_ AXBrailleTable) SetLanguage(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLanguage:"), value)
}

// All locales this table supports.
//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axbrailletable/locales
func (a_ AXBrailleTable) Locales() foundation.Locale {
	rv := objc.Send[foundation.Locale](a_.ID, objc.Sel("locales"))
	return rv
}


// SetLocales sets the value of the locales property.
// All locales this table supports.

//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axbrailletable/locales
func (a_ AXBrailleTable) SetLocales(value foundation.ILocale) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLocales:"), value)
}

// The localized name of this table for user display.
//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axbrailletable/localizedname
func (a_ AXBrailleTable) LocalizedName() appkit.string {
	rv := objc.Send[appkit.string](a_.ID, objc.Sel("localizedName"))
	return rv
}


// SetLocalizedName sets the value of the localizedName property.
// The localized name of this table for user display.

//
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axbrailletable/localizedname
func (a_ AXBrailleTable) SetLocalizedName(value appkit.string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLocalizedName:"), value)
}


