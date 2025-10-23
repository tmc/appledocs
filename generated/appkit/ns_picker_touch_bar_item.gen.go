// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PickerTouchBarItem] class.
var (
	PickerTouchBarItemClass     _PickerTouchBarItemClass
	PickerTouchBarItemClassOnce sync.Once
)

func getPickerTouchBarItemClass() _PickerTouchBarItemClass {
	PickerTouchBarItemClassOnce.Do(func() {
		PickerTouchBarItemClass = _PickerTouchBarItemClass{objc.GetClass("NSPickerTouchBarItem")}
	})
	return PickerTouchBarItemClass
}

type _PickerTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [PickerTouchBarItem] class.
type IPickerTouchBarItem interface {
	ITouchBarItem
	CollapsedRepresentationImage() Image
	SetCollapsedRepresentationImage(value IImage)
	CollapsedRepresentationLabel() string
	SetCollapsedRepresentationLabel(value string)
	Action() unsafe.Pointer
	SetAction(value unsafe.Pointer)
	ControlRepresentation() unsafe.Pointer
	SetControlRepresentation(value unsafe.Pointer)
	CustomizationLabel() string
	SetCustomizationLabel(value string)
	IsEnabled() bool
	SetIsEnabled(value bool)
	NumberOfOptions() int
	SetNumberOfOptions(value int)
	SelectedIndex() int
	SetSelectedIndex(value int)
	SelectionColor() Color
	SetSelectionColor(value IColor)
	SelectionMode() unsafe.Pointer
	SetSelectionMode(value unsafe.Pointer)
	Target() unsafe.Pointer
	SetTarget(value unsafe.Pointer)
}

// A bar item that provides a picker control with multiple options.


// A bar item that provides a picker control with multiple options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem
type PickerTouchBarItem struct {
	TouchBarItem
}

// PickerTouchBarItemFrom constructs a [PickerTouchBarItem] from an unsafe.Pointer.
//
// A bar item that provides a picker control with multiple options.
func PickerTouchBarItemFrom(ptr unsafe.Pointer) PickerTouchBarItem {
	return PickerTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PickerTouchBarItemClass) Alloc() PickerTouchBarItem {
	rv := objc.Send[PickerTouchBarItem](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PickerTouchBarItemClass) New() PickerTouchBarItem {
	rv := objc.Send[PickerTouchBarItem](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PickerTouchBarItem) Init() PickerTouchBarItem {
	rv := objc.Send[PickerTouchBarItem](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PickerTouchBarItem) Autorelease() PickerTouchBarItem {
	rv := objc.Send[PickerTouchBarItem](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPickerTouchBarItem creates a new PickerTouchBarItem instance.
func NewPickerTouchBarItem() PickerTouchBarItem {
	return getPickerTouchBarItemClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/init(identifier:labels:selectionMode:target:action:)
func NewPickerTouchBarItemWithIdentifierLabelsSelectionModeTargetAction(identifier ITouchBarItemIdentifier, labels []string, selectionMode PickerTouchBarItemSelectionMode, target objectivec.IObject, action objc.SEL) PickerTouchBarItem {
	rv := objc.Send[PickerTouchBarItem](objc.ID(getPickerTouchBarItemClass().class), objc.Sel("pickerTouchBarItemWithIdentifier:labels:selectionMode:target:action:"), identifier, labels, selectionMode, target, action)
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/init(identifier:labels:selectionMode:target:action:)
func (pc _PickerTouchBarItemClass) PickerTouchBarItemWithIdentifierLabelsSelectionModeTargetAction(identifier ITouchBarItemIdentifier, labels []string, selectionMode PickerTouchBarItemSelectionMode, target objectivec.IObject, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("pickerTouchBarItemWithIdentifier:labels:selectionMode:target:action:"), identifier, labels, selectionMode, target, action)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/collapsedRepresentationImage
func (p_ PickerTouchBarItem) CollapsedRepresentationImage() Image {
	rv := objc.Send[Image](p_.ID, objc.Sel("collapsedRepresentationImage"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/collapsedRepresentationImage
func (p_ PickerTouchBarItem) SetCollapsedRepresentationImage(value IImage) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCollapsedRepresentationImage:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/collapsedRepresentationLabel
func (p_ PickerTouchBarItem) CollapsedRepresentationLabel() string {
	rv := objc.Send[string](p_.ID, objc.Sel("collapsedRepresentationLabel"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/collapsedRepresentationLabel
func (p_ PickerTouchBarItem) SetCollapsedRepresentationLabel(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCollapsedRepresentationLabel:"), objc.String(value))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/action
func (p_ PickerTouchBarItem) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("action"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/action
func (p_ PickerTouchBarItem) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAction:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/controlrepresentation-swift.property
func (p_ PickerTouchBarItem) ControlRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("controlRepresentation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/controlrepresentation-swift.property
func (p_ PickerTouchBarItem) SetControlRepresentation(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setControlRepresentation:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/customizationlabel
func (p_ PickerTouchBarItem) CustomizationLabel() string {
	rv := objc.Send[string](p_.ID, objc.Sel("customizationLabel"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/customizationlabel
func (p_ PickerTouchBarItem) SetCustomizationLabel(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCustomizationLabel:"), objc.String(value))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/isenabled
func (p_ PickerTouchBarItem) IsEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/isenabled
func (p_ PickerTouchBarItem) SetIsEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/numberofoptions
func (p_ PickerTouchBarItem) NumberOfOptions() int {
	rv := objc.Send[int](p_.ID, objc.Sel("numberOfOptions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/numberofoptions
func (p_ PickerTouchBarItem) SetNumberOfOptions(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setNumberOfOptions:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/selectedindex
func (p_ PickerTouchBarItem) SelectedIndex() int {
	rv := objc.Send[int](p_.ID, objc.Sel("selectedIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/selectedindex
func (p_ PickerTouchBarItem) SetSelectedIndex(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSelectedIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/selectioncolor
func (p_ PickerTouchBarItem) SelectionColor() Color {
	rv := objc.Send[Color](p_.ID, objc.Sel("selectionColor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/selectioncolor
func (p_ PickerTouchBarItem) SetSelectionColor(value IColor) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSelectionColor:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/selectionmode-swift.property
func (p_ PickerTouchBarItem) SelectionMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("selectionMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/selectionmode-swift.property
func (p_ PickerTouchBarItem) SetSelectionMode(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSelectionMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/target
func (p_ PickerTouchBarItem) Target() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("target"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspickertouchbaritem/target
func (p_ PickerTouchBarItem) SetTarget(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTarget:"), value)
}


