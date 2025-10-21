// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
func NewAXBrailleTableWithIdentifier(identifier string) AXBrailleTable {
	instance := getAXBrailleTableClass().Alloc()
	rv := objc.Send[AXBrailleTable](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(identifier))
	rv.Autorelease()
	return rv
}


// All tables that provide translations for the given locale’s language.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/tables(for:)
func (ac _AXBrailleTableClass) TablesForLocale(locale unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("tablesForLocale:"), locale)
	return rv
}

// The localized name of the provider of this table for user display.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/localizedProviderName
func (a_ AXBrailleTable) LocalizedProviderName() string {
	rv := objc.Send[string](a_.ID, objc.Sel("localizedProviderName"))
	return rv
}

// The identifier of the provider of this table.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTable/providerIdentifier
func (a_ AXBrailleTable) ProviderIdentifier() string {
	rv := objc.Send[string](a_.ID, objc.Sel("providerIdentifier"))
	return rv
}


