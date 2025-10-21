// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RPBroadcastActivityController] class.
var (
	RPBroadcastActivityControllerClass     _RPBroadcastActivityControllerClass
	RPBroadcastActivityControllerClassOnce sync.Once
)

func getRPBroadcastActivityControllerClass() _RPBroadcastActivityControllerClass {
	RPBroadcastActivityControllerClassOnce.Do(func() {
		RPBroadcastActivityControllerClass = _RPBroadcastActivityControllerClass{objc.GetClass("RPBroadcastActivityController")}
	})
	return RPBroadcastActivityControllerClass
}

type _RPBroadcastActivityControllerClass struct {
	class objc.Class
}

// An interface definition for the [RPBroadcastActivityController] class.
type IRPBroadcastActivityController interface {
	objectivec.IObject
}

// A controller object that presents the macOS broadcast picker.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastActivityController
type RPBroadcastActivityController struct {
	objectivec.Object
}

// RPBroadcastActivityControllerFrom constructs a [RPBroadcastActivityController] from an unsafe.Pointer.
//
// A controller object that presents the macOS broadcast picker.
func RPBroadcastActivityControllerFrom(ptr unsafe.Pointer) RPBroadcastActivityController {
	return RPBroadcastActivityController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RPBroadcastActivityControllerClass) Alloc() RPBroadcastActivityController {
	rv := objc.Send[RPBroadcastActivityController](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RPBroadcastActivityControllerClass) New() RPBroadcastActivityController {
	rv := objc.Send[RPBroadcastActivityController](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RPBroadcastActivityController) Init() RPBroadcastActivityController {
	rv := objc.Send[RPBroadcastActivityController](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RPBroadcastActivityController) Autorelease() RPBroadcastActivityController {
	rv := objc.Send[RPBroadcastActivityController](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRPBroadcastActivityController creates a new RPBroadcastActivityController instance.
func NewRPBroadcastActivityController() RPBroadcastActivityController {
	return getRPBroadcastActivityControllerClass().New()
}


// Presents a list of available broadcast services for the user to select.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastActivityController/showBroadcastPicker(at:from:preferredExtensionIdentifier:completionHandler:)
func (rc _RPBroadcastActivityControllerClass) ShowBroadcastPickerAtPointFromWindowPreferredExtensionIdentifierCompletionHandler(point coregraphics.CGPoint, window unsafe.Pointer, preferredExtension string, handler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(rc.class), objc.Sel("showBroadcastPickerAtPoint:fromWindow:preferredExtensionIdentifier:completionHandler:"), point, window, objc.String(preferredExtension), handler)
}

// The broadcast activity controller’s delegate object.
//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastActivityController/delegate
func (r_ RPBroadcastActivityController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The broadcast activity controller’s delegate object.

//
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastActivityController/delegate
func (r_ RPBroadcastActivityController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDelegate:"), value)
}



