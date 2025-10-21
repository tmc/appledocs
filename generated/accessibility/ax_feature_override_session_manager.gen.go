// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AXFeatureOverrideSessionManager] class.
var (
	AXFeatureOverrideSessionManagerClass     _AXFeatureOverrideSessionManagerClass
	AXFeatureOverrideSessionManagerClassOnce sync.Once
)

func getAXFeatureOverrideSessionManagerClass() _AXFeatureOverrideSessionManagerClass {
	AXFeatureOverrideSessionManagerClassOnce.Do(func() {
		AXFeatureOverrideSessionManagerClass = _AXFeatureOverrideSessionManagerClass{objc.GetClass("AXFeatureOverrideSessionManager")}
	})
	return AXFeatureOverrideSessionManagerClass
}

type _AXFeatureOverrideSessionManagerClass struct {
	class objc.Class
}

// An interface definition for the [AXFeatureOverrideSessionManager] class.
type IAXFeatureOverrideSessionManager interface {
	objectivec.IObject
	BeginOverrideSessionEnablingOptionsDisablingOptionsError(enableOptions unsafe.Pointer, disableOptions unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	EndOverrideSessionError(session unsafe.Pointer, error_ unsafe.Pointer) bool
}

// A manager class to begin and end accessibility feature override sessions. Multiple override sessions are reconciled by combining the requests, preferring feature enablement. Ending all sessions restores the prior state of Accessibility feature enablement. Your app must be entitled with com.apple.developer.accessibility.merchant-api-control.
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSessionManager
type AXFeatureOverrideSessionManager struct {
	objectivec.Object
}

// AXFeatureOverrideSessionManagerFrom constructs a [AXFeatureOverrideSessionManager] from an unsafe.Pointer.
//
// A manager class to begin and end accessibility feature override sessions. Multiple override sessions are reconciled by combining the requests, preferring feature enablement. Ending all sessions restores the prior state of Accessibility feature enablement. Your app must be entitled with com.apple.developer.accessibility.merchant-api-control.
func AXFeatureOverrideSessionManagerFrom(ptr unsafe.Pointer) AXFeatureOverrideSessionManager {
	return AXFeatureOverrideSessionManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AXFeatureOverrideSessionManagerClass) Alloc() AXFeatureOverrideSessionManager {
	rv := objc.Send[AXFeatureOverrideSessionManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXFeatureOverrideSessionManagerClass) New() AXFeatureOverrideSessionManager {
	rv := objc.Send[AXFeatureOverrideSessionManager](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXFeatureOverrideSessionManager) Init() AXFeatureOverrideSessionManager {
	rv := objc.Send[AXFeatureOverrideSessionManager](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXFeatureOverrideSessionManager) Autorelease() AXFeatureOverrideSessionManager {
	rv := objc.Send[AXFeatureOverrideSessionManager](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXFeatureOverrideSessionManager creates a new AXFeatureOverrideSessionManager instance.
func NewAXFeatureOverrideSessionManager() AXFeatureOverrideSessionManager {
	return getAXFeatureOverrideSessionManagerClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSessionManager/sharedInstance
func (ac _AXFeatureOverrideSessionManagerClass) SharedInstance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("sharedInstance"))
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSessionManager/beginOverrideSession(enabling:disabling:)
func (a_ AXFeatureOverrideSessionManager) BeginOverrideSessionEnablingOptionsDisablingOptionsError(enableOptions unsafe.Pointer, disableOptions unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("beginOverrideSessionEnablingOptions:disablingOptions:error:"), enableOptions, disableOptions, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSessionManager/end(_:)
func (a_ AXFeatureOverrideSessionManager) EndOverrideSessionError(session unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("endOverrideSession:error:"), session, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSessionManager/sharedInstance
func (a_ AXFeatureOverrideSessionManager) SharedInstance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("sharedInstance"))
	return rv
}



