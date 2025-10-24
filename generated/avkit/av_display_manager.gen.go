// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVDisplayManager */


/* debug [class_header]: Header for AVDisplayManager */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DisplayManager */
// An interface definition for the [DisplayManager] class.
type IDisplayManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DisplayManager */
	// properties:
	IsDisplayCriteriaMatchingEnabled() bool
	SetIsDisplayCriteriaMatchingEnabled(value bool)
	IsDisplayModeSwitchInProgress() bool
	SetIsDisplayModeSwitchInProgress(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DisplayManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DisplayManager */
// Alloc allocates a new instance without initialization.
func (dc _DisplayManagerClass) Alloc() DisplayManager {
	rv := objc.Send[DisplayManager](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DisplayManager */
// A tvOS management object that controls whether a TV switches modes to match the video’s native mode.
//
// If you set the display manager’s , when a user enables a Match Content setting, the TV attempts to change modes to match the currently playing video’s native display criteria.


// A tvOS management object that controls whether a TV switches modes to match the video’s native mode.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DisplayManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DisplayManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DisplayManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DisplayManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DisplayManager */

// A Boolean value that indicates whether the user has enabled display critera matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avdisplaymanager/isdisplaycriteriamatchingenabled
func (d_ DisplayManager) IsDisplayCriteriaMatchingEnabled() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isDisplayCriteriaMatchingEnabled"))
	return rv
}/* debug [instance_properties/getter]: isDisplayCriteriaMatchingEnabled */


// A Boolean value that indicates whether the user has enabled display critera matching.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avdisplaymanager/isdisplaycriteriamatchingenabled
func (d_ DisplayManager) SetIsDisplayCriteriaMatchingEnabled(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsDisplayCriteriaMatchingEnabled:"), value)
}/* debug [instance_properties/setter]: isDisplayCriteriaMatchingEnabled */


// A Boolean value that indicates whether a display mode switch is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avdisplaymanager/isdisplaymodeswitchinprogress
func (d_ DisplayManager) IsDisplayModeSwitchInProgress() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isDisplayModeSwitchInProgress"))
	return rv
}/* debug [instance_properties/getter]: isDisplayModeSwitchInProgress */


// A Boolean value that indicates whether a display mode switch is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avkit/avdisplaymanager/isdisplaymodeswitchinprogress
func (d_ DisplayManager) SetIsDisplayModeSwitchInProgress(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsDisplayModeSwitchInProgress:"), value)
}/* debug [instance_properties/setter]: isDisplayModeSwitchInProgress */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVDisplayManager */


