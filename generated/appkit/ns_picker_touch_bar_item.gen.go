// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// A bar item that provides a picker control with multiple options.
//
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

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/init(identifier:labels:selectionMode:target:action:)
func NewPickerTouchBarItemWithIdentifierLabelsSelectionModeTargetAction(identifier unsafe.Pointer, labels unsafe.Pointer, selectionMode unsafe.Pointer, target objc.ID, action objc.SEL) PickerTouchBarItem {
	rv := objc.Send[PickerTouchBarItem](objc.ID(getPickerTouchBarItemClass().class), objc.Sel("pickerTouchBarItemWithIdentifier:labels:selectionMode:target:action:"), identifier, labels, selectionMode, target, action)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/init(identifier:labels:selectionMode:target:action:)
func (pc _PickerTouchBarItemClass) PickerTouchBarItemWithIdentifierLabelsSelectionModeTargetAction(identifier unsafe.Pointer, labels unsafe.Pointer, selectionMode unsafe.Pointer, target objc.ID, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("pickerTouchBarItemWithIdentifier:labels:selectionMode:target:action:"), identifier, labels, selectionMode, target, action)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/collapsedRepresentationImage
func (p_ PickerTouchBarItem) CollapsedRepresentationImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("collapsedRepresentationImage"))
	return rv
}

// SetCollapsedRepresentationImage sets the value of the collapsedRepresentationImage property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/collapsedRepresentationImage
func (p_ PickerTouchBarItem) SetCollapsedRepresentationImage(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCollapsedRepresentationImage:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/collapsedRepresentationLabel
func (p_ PickerTouchBarItem) CollapsedRepresentationLabel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("collapsedRepresentationLabel"))
	return rv
}

// SetCollapsedRepresentationLabel sets the value of the collapsedRepresentationLabel property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/collapsedRepresentationLabel
func (p_ PickerTouchBarItem) SetCollapsedRepresentationLabel(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCollapsedRepresentationLabel:"), value)
}
