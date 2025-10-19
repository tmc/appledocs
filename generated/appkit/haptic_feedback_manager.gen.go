// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HapticFeedbackManager] class.
var hapticFeedbackManagerClass = _HapticFeedbackManagerClass{objc.GetClass("NSHapticFeedbackManager")}

type _HapticFeedbackManagerClass struct {
	class objc.Class
}

// An interface definition for the [HapticFeedbackManager] class.
type IHapticFeedbackManager interface {
	objectivec.IObject
}

// An object that provides access to the haptic feedback management attributes on a system with a Force Touch trackpad. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager

type HapticFeedbackManager struct {
	objectivec.Object
}

// HapticFeedbackManagerFrom constructs a [HapticFeedbackManager] from an unsafe.Pointer.
//
// An object that provides access to the haptic feedback management attributes on a system with a Force Touch trackpad.
func HapticFeedbackManagerFrom(ptr unsafe.Pointer) HapticFeedbackManager {
	return HapticFeedbackManager{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (hc _HapticFeedbackManagerClass) Alloc() HapticFeedbackManager {
	rv := objc.Send[HapticFeedbackManager](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (hc _HapticFeedbackManagerClass) New() HapticFeedbackManager {
	rv := objc.Send[HapticFeedbackManager](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HapticFeedbackManager) Init() HapticFeedbackManager {
	rv := objc.Send[HapticFeedbackManager](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HapticFeedbackManager) Autorelease() HapticFeedbackManager {
	rv := objc.Send[HapticFeedbackManager](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHapticFeedbackManager creates a new HapticFeedbackManager instance.
func NewHapticFeedbackManager() HapticFeedbackManager {
	return hapticFeedbackManagerClass.New()
}




