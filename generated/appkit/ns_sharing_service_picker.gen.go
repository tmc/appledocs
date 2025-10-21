// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SharingServicePicker] class.
var (
	SharingServicePickerClass     _SharingServicePickerClass
	SharingServicePickerClassOnce sync.Once
)

func getSharingServicePickerClass() _SharingServicePickerClass {
	SharingServicePickerClassOnce.Do(func() {
		SharingServicePickerClass = _SharingServicePickerClass{objc.GetClass("NSSharingServicePicker")}
	})
	return SharingServicePickerClass
}

type _SharingServicePickerClass struct {
	class objc.Class
}

// An interface definition for the [SharingServicePicker] class.
type ISharingServicePicker interface {
	objectivec.IObject
	ShowRelativeToRectOfViewPreferredEdge(rect coregraphics.CGRect, view unsafe.Pointer, preferredEdge int)
}

// A list of sharing services that the user can choose from.
//
// An object presents an interface for sharing one or more items using a specific service. In macOS 12 and earlier, this picker displays a menu with a list of services that someone can use to share the item. In macOS 13 and later, the picker displays a popover with a preview of the item and the list of services. When someone chooses a service, the picker automatically shares the proposed item with that service. Create a sharing service picker and configure it with a delegate object to monitor interactions. Your delegate must conform to the protocol. Present the picker from your interface using the method.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker
type SharingServicePicker struct {
	objectivec.Object
}

// SharingServicePickerFrom constructs a [SharingServicePicker] from an unsafe.Pointer.
//
// A list of sharing services that the user can choose from.
func SharingServicePickerFrom(ptr unsafe.Pointer) SharingServicePicker {
	return SharingServicePicker{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SharingServicePickerClass) Alloc() SharingServicePicker {
	rv := objc.Send[SharingServicePicker](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SharingServicePickerClass) New() SharingServicePicker {
	rv := objc.Send[SharingServicePicker](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SharingServicePicker) Init() SharingServicePicker {
	rv := objc.Send[SharingServicePicker](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SharingServicePicker) Autorelease() SharingServicePicker {
	rv := objc.Send[SharingServicePicker](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSharingServicePicker creates a new SharingServicePicker instance.
func NewSharingServicePicker() SharingServicePicker {
	return getSharingServicePickerClass().New()
}


// Shows the picker interface and populates it with the relevant sharing services.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/show(relativeTo:of:preferredEdge:)
func (s_ SharingServicePicker) ShowRelativeToRectOfViewPreferredEdge(rect coregraphics.CGRect, view unsafe.Pointer, preferredEdge int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("showRelativeToRect:ofView:preferredEdge:"), rect, view, preferredEdge)
}

// The object for managing the sharing service picker.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/delegate
func (s_ SharingServicePicker) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The object for managing the sharing service picker.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/delegate
func (s_ SharingServicePicker) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}

// A menu item suitable to display the picker for the specified items.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepicker/standardsharemenuitem
func (s_ SharingServicePicker) StandardShareMenuItem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("standardShareMenuItem"))
	return rv
}


// SetStandardShareMenuItem sets the value of the standardShareMenuItem property.
// A menu item suitable to display the picker for the specified items.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssharingservicepicker/standardsharemenuitem
func (s_ SharingServicePicker) SetStandardShareMenuItem(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStandardShareMenuItem:"), value)
}



