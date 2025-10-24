// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [RPSystemBroadcastPickerView] class.
var (
	RPSystemBroadcastPickerViewClass     _RPSystemBroadcastPickerViewClass
	RPSystemBroadcastPickerViewClassOnce sync.Once
)

func getRPSystemBroadcastPickerViewClass() _RPSystemBroadcastPickerViewClass {
	RPSystemBroadcastPickerViewClassOnce.Do(func() {
		RPSystemBroadcastPickerViewClass = _RPSystemBroadcastPickerViewClass{objc.GetClass("RPSystemBroadcastPickerView")}
	})
	return RPSystemBroadcastPickerViewClass
}

type _RPSystemBroadcastPickerViewClass struct {
	class objc.Class
}

// An interface definition for the [RPSystemBroadcastPickerView] class.
type IRPSystemBroadcastPickerView interface {
	appkit.IView
	// properties:
	// methods:
}

// A view displaying a broadcast button that, when tapped, shows a broadcast picker.
//
// Add this view to your view hierarchy to let users broadcast directly from your app. When a user taps the broadcast button, it displays a broadcast picker, allowing the user to select a broadcast provider. You can limit the picker to a particular broadcast provider by setting to the bundle identifier of a broadcast extension. You can also show or hide the microphone button displayed in the picker by setting the property. Set these properties before presenting , as shown here:


// A view displaying a broadcast button that, when tapped, shows a broadcast picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPSystemBroadcastPickerView
type RPSystemBroadcastPickerView struct {
	appkit.View
}

// RPSystemBroadcastPickerViewFrom constructs a [RPSystemBroadcastPickerView] from an unsafe.Pointer.
//
// A view displaying a broadcast button that, when tapped, shows a broadcast picker.
func RPSystemBroadcastPickerViewFrom(ptr unsafe.Pointer) RPSystemBroadcastPickerView {
	return RPSystemBroadcastPickerView{
		View: appkit.ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RPSystemBroadcastPickerViewClass) Alloc() RPSystemBroadcastPickerView {
	rv := objc.Send[RPSystemBroadcastPickerView](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RPSystemBroadcastPickerViewClass) New() RPSystemBroadcastPickerView {
	rv := objc.Send[RPSystemBroadcastPickerView](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RPSystemBroadcastPickerView) Init() RPSystemBroadcastPickerView {
	rv := objc.Send[RPSystemBroadcastPickerView](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RPSystemBroadcastPickerView) Autorelease() RPSystemBroadcastPickerView {
	rv := objc.Send[RPSystemBroadcastPickerView](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRPSystemBroadcastPickerView creates a new RPSystemBroadcastPickerView instance.
func NewRPSystemBroadcastPickerView() RPSystemBroadcastPickerView {
	return getRPSystemBroadcastPickerViewClass().New()
}



