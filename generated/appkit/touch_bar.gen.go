
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TouchBar] class.
var TouchBarClass _TouchBarClass

func init() {
	TouchBarClass = _TouchBarClass{objc.GetClass("NSTouchBar")}
}

type _TouchBarClass struct {
	objc.Class
}

// An interface definition for the [TouchBar] class.
type ITouchBar interface {
	ID() objc.ID
	InitWithCoder(coder unsafe.Pointer) unsafe.Pointer
	ItemForIdentifier(identifier unsafe.Pointer) unsafe.Pointer
}

type TouchBar struct {
	id objc.ID
}

func TouchBarFrom(ptr unsafe.Pointer) TouchBar {
	return TouchBar{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TouchBar) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TouchBarClass) Alloc() TouchBar {
	rv := objc.Send[TouchBar](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TouchBarClass) New() TouchBar {
	rv := objc.Send[TouchBar](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTouchBar creates and returns a new initialized instance.
func NewTouchBar() TouchBar {
	return TouchBarClass.New()
}

// Init initializes the instance.
func (t_ TouchBar) Init() TouchBar {
	rv := objc.Send[TouchBar](t_.ID(), selInit)
	return rv
}
// Creates a Touch Bar object from a coder object provided by a storyboard or NIB file. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/init(coder:)
func (t_ TouchBar) InitWithCoder(coder unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("initWithCoder:"), coder)
	return rv
}
// Returns the Touch Bar item that corresponds to a given identifier. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/item(forIdentifier:)
func (t_ TouchBar) ItemForIdentifier(identifier unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("itemForIdentifier:"), identifier)
	return rv
}
// A list of identifiers for items to show in the Touch Bar’s customization UI. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/customizationAllowedItemIdentifiers
func (t_ TouchBar) CustomizationAllowedItemIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("customizationAllowedItemIdentifiers"))
	return rv
}
// SetCustomizationAllowedItemIdentifiers sets the value of the customizationAllowedItemIdentifiers property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/customizationAllowedItemIdentifiers
func (t_ TouchBar) SetCustomizationAllowedItemIdentifiers(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setCustomizationAllowedItemIdentifiers:"), value)
}
// A globally unique string that makes the Touch Bar eligible for user customization. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/customizationIdentifier-swift.property
func (t_ TouchBar) CustomizationIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("customizationIdentifier"))
	return rv
}
// SetCustomizationIdentifier sets the value of the customizationIdentifier property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/customizationIdentifier-swift.property
func (t_ TouchBar) SetCustomizationIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setCustomizationIdentifier:"), value)
}
// An optional list of identifiers for items you want to always appear in the Touch Bar and which the user can’t remove during customization. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/customizationRequiredItemIdentifiers
func (t_ TouchBar) CustomizationRequiredItemIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("customizationRequiredItemIdentifiers"))
	return rv
}
// SetCustomizationRequiredItemIdentifiers sets the value of the customizationRequiredItemIdentifiers property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/customizationRequiredItemIdentifiers
func (t_ TouchBar) SetCustomizationRequiredItemIdentifiers(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setCustomizationRequiredItemIdentifiers:"), value)
}
// A required list of identifiers for items that you want to appear in the Touch Bar after instantiating it. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/defaultItemIdentifiers
func (t_ TouchBar) DefaultItemIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("defaultItemIdentifiers"))
	return rv
}
// SetDefaultItemIdentifiers sets the value of the defaultItemIdentifiers property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/defaultItemIdentifiers
func (t_ TouchBar) SetDefaultItemIdentifiers(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setDefaultItemIdentifiers:"), value)
}
// The delegate that provides items to the Touch Bar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/delegate
func (t_ TouchBar) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("delegate"))
	return rv
}
// SetDelegate sets the value of the delegate property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/delegate
func (t_ TouchBar) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setDelegate:"), value)
}
// The identifier of an item that replaces the system-provided button in the Touch Bar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/escapeKeyReplacementItemIdentifier
func (t_ TouchBar) EscapeKeyReplacementItemIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("escapeKeyReplacementItemIdentifier"))
	return rv
}
// SetEscapeKeyReplacementItemIdentifier sets the value of the escapeKeyReplacementItemIdentifier property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/escapeKeyReplacementItemIdentifier
func (t_ TouchBar) SetEscapeKeyReplacementItemIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setEscapeKeyReplacementItemIdentifier:"), value)
}
// A Boolean value that Indicates whether the Touch Bar is eligible for display. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/isVisible
func (t_ TouchBar) Visible() bool {
	rv := objc.Send[bool](t_.ID(), objc.RegisterName("visible"))
	return rv
}
// The list of identifiers for the current items in the Touch Bar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/itemIdentifiers
func (t_ TouchBar) ItemIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("itemIdentifiers"))
	return rv
}
// The identifier of an item you want the system to center in the Touch Bar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/principalItemIdentifier
func (t_ TouchBar) PrincipalItemIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("principalItemIdentifier"))
	return rv
}
// SetPrincipalItemIdentifier sets the value of the principalItemIdentifier property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/principalItemIdentifier
func (t_ TouchBar) SetPrincipalItemIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setPrincipalItemIdentifier:"), value)
}
// The primary source of items that the Touch Bar uses to fill its private items array, unless you provide items using a delegate. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/templateItems
func (t_ TouchBar) TemplateItems() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID(), objc.RegisterName("templateItems"))
	return rv
}
// SetTemplateItems sets the value of the templateItems property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSTouchBar/templateItems
func (t_ TouchBar) SetTemplateItems(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID(), objc.RegisterName("setTemplateItems:"), value)
}
