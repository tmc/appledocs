// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DisplayManager] class.
var (
	DisplayManagerClass     _DisplayManagerClass
	DisplayManagerClassOnce sync.Once
)

func getDisplayManagerClass() _DisplayManagerClass {
	DisplayManagerClassOnce.Do(func() {
		DisplayManagerClass = _DisplayManagerClass{objc.GetClass("AVDisplayManager")}
	})
	return DisplayManagerClass
}

type _DisplayManagerClass struct {
	class objc.Class
}

// An interface definition for the [DisplayManager] class.
type IDisplayManager interface {
	objectivec.IObject
}

// A tvOS management object that controls whether a TV switches modes to match the video’s native mode.
//
// If you set the display manager’s , when a user enables a Match Content setting, the TV attempts to change modes to match the currently playing video’s native display criteria.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVDisplayManager
type DisplayManager struct {
	objectivec.Object
}

// DisplayManagerFrom constructs a [DisplayManager] from an unsafe.Pointer.
//
// A tvOS management object that controls whether a TV switches modes to match the video’s native mode.
func DisplayManagerFrom(ptr unsafe.Pointer) DisplayManager {
	return DisplayManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DisplayManagerClass) Alloc() DisplayManager {
	rv := objc.Send[DisplayManager](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DisplayManagerClass) New() DisplayManager {
	rv := objc.Send[DisplayManager](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DisplayManager) Init() DisplayManager {
	rv := objc.Send[DisplayManager](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DisplayManager) Autorelease() DisplayManager {
	rv := objc.Send[DisplayManager](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDisplayManager creates a new DisplayManager instance.
func NewDisplayManager() DisplayManager {
	return getDisplayManagerClass().New()
}


// A Boolean value that indicates whether the user has enabled display critera matching.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVDisplayManager/isDisplayCriteriaMatchingEnabled
func (d_ DisplayManager) DisplayCriteriaMatchingEnabled() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("displayCriteriaMatchingEnabled"))
	return rv
}

// A Boolean value that indicates whether a display mode switch is in progress.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVDisplayManager/isDisplayModeSwitchInProgress
func (d_ DisplayManager) DisplayModeSwitchInProgress() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("displayModeSwitchInProgress"))
	return rv
}

// A hint for the TV to set the display mode to best match the currently playing content’s display criteria.
//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVDisplayManager/preferredDisplayCriteria
func (d_ DisplayManager) PreferredDisplayCriteria() avfoundation.DisplayCriteria {
	rv := objc.Send[avfoundation.DisplayCriteria](d_.ID, objc.Sel("preferredDisplayCriteria"))
	return rv
}


// SetPreferredDisplayCriteria sets the value of the preferredDisplayCriteria property.
// A hint for the TV to set the display mode to best match the currently playing content’s display criteria.

//
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVDisplayManager/preferredDisplayCriteria
func (d_ DisplayManager) SetPreferredDisplayCriteria(value avfoundation.IDisplayCriteria) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPreferredDisplayCriteria:"), value)
}

// A Boolean value that indicates whether the user has enabled display critera matching.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avdisplaymanager/isdisplaycriteriamatchingenabled
func (d_ DisplayManager) IsDisplayCriteriaMatchingEnabled() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isDisplayCriteriaMatchingEnabled"))
	return rv
}


// SetIsDisplayCriteriaMatchingEnabled sets the value of the isDisplayCriteriaMatchingEnabled property.
// A Boolean value that indicates whether the user has enabled display critera matching.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avdisplaymanager/isdisplaycriteriamatchingenabled
func (d_ DisplayManager) SetIsDisplayCriteriaMatchingEnabled(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsDisplayCriteriaMatchingEnabled:"), value)
}

// A Boolean value that indicates whether a display mode switch is in progress.
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avdisplaymanager/isdisplaymodeswitchinprogress
func (d_ DisplayManager) IsDisplayModeSwitchInProgress() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isDisplayModeSwitchInProgress"))
	return rv
}


// SetIsDisplayModeSwitchInProgress sets the value of the isDisplayModeSwitchInProgress property.
// A Boolean value that indicates whether a display mode switch is in progress.

//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avdisplaymanager/isdisplaymodeswitchinprogress
func (d_ DisplayManager) SetIsDisplayModeSwitchInProgress(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsDisplayModeSwitchInProgress:"), value)
}



