// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	StandardShareMenuItem() IMenuItem
	// methods:
	Close()
	ShowRelativeToRectOfViewPreferredEdge(rect objc.IObject /* cross-framework: Rect */, view IView, preferredEdge RectEdge /* not a class type */)
}

// A list of sharing services that the user can choose from.
//
// An object presents an interface for sharing one or more items using a specific service. In macOS 12 and earlier, this picker displays a menu with a list of services that someone can use to share the item. In macOS 13 and later, the picker displays a popover with a preview of the item and the list of services. When someone chooses a service, the picker automatically shares the proposed item with that service. Create a sharing service picker and configure it with a delegate object to monitor interactions. Your delegate must conform to the protocol. Present the picker from your interface using the method.


// A list of sharing services that the user can choose from.
//
// [Full Topic]
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



// Creates a new sharing service picker for the selected items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/init(items:)
func NewSharingServicePickerWithItems(items objc.IObject /* cross-framework: NSArray */) SharingServicePicker {
	instance := getSharingServicePickerClass().Alloc()
	rv := objc.Send[SharingServicePicker](instance.ID, objc.Sel("initWithItems:"), items)
	rv.Autorelease()
	return rv
}



// Closes the picker interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/close()
func (s_ SharingServicePicker) Close() {
	objc.Send[objc.ID](s_.ID, objc.Sel("close"))
}


// Shows the picker interface and populates it with the relevant sharing services.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/show(relativeTo:of:preferredEdge:)
func (s_ SharingServicePicker) ShowRelativeToRectOfViewPreferredEdge(rect objc.IObject /* cross-framework: Rect */, view IView, preferredEdge RectEdge /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("showRelativeToRect:ofView:preferredEdge:"), rect, view, preferredEdge)
}


// The object for managing the sharing service picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/delegate
func (s_ SharingServicePicker) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}


// The object for managing the sharing service picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/delegate
func (s_ SharingServicePicker) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}


// A menu item suitable to display the picker for the specified items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingServicePicker/standardShareMenuItem
func (s_ SharingServicePicker) StandardShareMenuItem() IMenuItem {
	rv := objc.Send[MenuItem](s_.ID, objc.Sel("standardShareMenuItem"))
	return rv
}


