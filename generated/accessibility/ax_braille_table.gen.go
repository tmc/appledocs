// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	LocalizedProviderName() objc.IObject /* cross-framework: NSString */
	ProviderIdentifier() objc.IObject /* cross-framework: NSString */
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	IsEightDot() bool
	SetIsEightDot(value bool)
	Language() unsafe.Pointer
	SetLanguage(value unsafe.Pointer)
	Locales() objc.IObject /* cross-framework: Locale */
	SetLocales(value objc.IObject /* cross-framework: Locale */)
	LocalizedName() objc.IObject /* cross-framework: NSString */
	SetLocalizedName(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

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



// The default table that provides translations for the given locale’s language. Returns nil if there is none.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/defaultTable(for:)
func (ac _AXBrailleTableClass) DefaultTableForLocale(locale objc.IObject /* cross-framework: Locale */) AXBrailleTable {
	rv := objc.Send[AXBrailleTable](objc.ID(ac.class), objc.Sel("defaultTableForLocale:"), locale)
	return rv
}


// The localized name of the provider of this table for user display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/localizedProviderName
func (a_ AXBrailleTable) LocalizedProviderName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("localizedProviderName"))
	return rv
}


// The identifier of the provider of this table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/providerIdentifier
func (a_ AXBrailleTable) ProviderIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("providerIdentifier"))
	return rv
}


// A unique string that identifies this table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axbrailletable/identifier
func (a_ AXBrailleTable) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("identifier"))
	return rv
}


// A unique string that identifies this table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axbrailletable/identifier
func (a_ AXBrailleTable) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIdentifier:"), value)
}


// Returns true if this table makes use of eight dots as opposed to six dots.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axbrailletable/iseightdot
func (a_ AXBrailleTable) IsEightDot() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isEightDot"))
	return rv
}


// Returns true if this table makes use of eight dots as opposed to six dots.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axbrailletable/iseightdot
func (a_ AXBrailleTable) SetIsEightDot(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsEightDot:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axbrailletable/language-3stsd
func (a_ AXBrailleTable) Language() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("language"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axbrailletable/language-3stsd
func (a_ AXBrailleTable) SetLanguage(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLanguage:"), value)
}


// All locales this table supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axbrailletable/locales
func (a_ AXBrailleTable) Locales() objc.IObject /* cross-framework: Locale */ {
	rv := objc.Send[foundation.Locale](a_.ID, objc.Sel("locales"))
	return rv
}


// All locales this table supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axbrailletable/locales
func (a_ AXBrailleTable) SetLocales(value objc.IObject /* cross-framework: Locale */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLocales:"), value)
}


// The localized name of this table for user display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axbrailletable/localizedname
func (a_ AXBrailleTable) LocalizedName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("localizedName"))
	return rv
}


// The localized name of this table for user display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axbrailletable/localizedname
func (a_ AXBrailleTable) SetLocalizedName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLocalizedName:"), value)
}



