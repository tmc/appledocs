// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureDeskViewApplicationLaunchConfiguration */


/* debug [class_header]: Header for AVCaptureDeskViewApplicationLaunchConfiguration */
// The class instance for the [CaptureDeskViewApplicationLaunchConfiguration] class.
var (
	CaptureDeskViewApplicationLaunchConfigurationClass     _CaptureDeskViewApplicationLaunchConfigurationClass
	CaptureDeskViewApplicationLaunchConfigurationClassOnce sync.Once
)

func getCaptureDeskViewApplicationLaunchConfigurationClass() _CaptureDeskViewApplicationLaunchConfigurationClass {
	CaptureDeskViewApplicationLaunchConfigurationClassOnce.Do(func() {
		CaptureDeskViewApplicationLaunchConfigurationClass = _CaptureDeskViewApplicationLaunchConfigurationClass{objc.GetClass("AVCaptureDeskViewApplicationLaunchConfiguration")}
	})
	return CaptureDeskViewApplicationLaunchConfigurationClass
}

type _CaptureDeskViewApplicationLaunchConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureDeskViewApplicationLaunchConfiguration */
// An interface definition for the [CaptureDeskViewApplicationLaunchConfiguration] class.
type ICaptureDeskViewApplicationLaunchConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureDeskViewApplicationLaunchConfiguration */
	// properties:
	MainWindowFrame() corefoundation.CGRect
	SetMainWindowFrame(value corefoundation.CGRect)
	RequiresSetUpModeCompletion() bool
	SetRequiresSetUpModeCompletion(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureDeskViewApplicationLaunchConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureDeskViewApplicationLaunchConfiguration */
// Alloc allocates a new instance without initialization.
func (cc _CaptureDeskViewApplicationLaunchConfigurationClass) Alloc() CaptureDeskViewApplicationLaunchConfiguration {
	rv := objc.Send[CaptureDeskViewApplicationLaunchConfiguration](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureDeskViewApplicationLaunchConfigurationClass) New() CaptureDeskViewApplicationLaunchConfiguration {
	rv := objc.Send[CaptureDeskViewApplicationLaunchConfiguration](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureDeskViewApplicationLaunchConfiguration) Init() CaptureDeskViewApplicationLaunchConfiguration {
	rv := objc.Send[CaptureDeskViewApplicationLaunchConfiguration](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureDeskViewApplicationLaunchConfiguration) Autorelease() CaptureDeskViewApplicationLaunchConfiguration {
	rv := objc.Send[CaptureDeskViewApplicationLaunchConfiguration](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureDeskViewApplicationLaunchConfiguration creates a new CaptureDeskViewApplicationLaunchConfiguration instance.
func NewCaptureDeskViewApplicationLaunchConfiguration() CaptureDeskViewApplicationLaunchConfiguration {
	return getCaptureDeskViewApplicationLaunchConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureDeskViewApplicationLaunchConfiguration */
// An object that configures how to present Desk View.
//
// Use this object to specify the frame for Desk View when it launches, and when to execute the completion handler. You can specify whether to perform the completion handler as soon as Desk View is visible to the user, or only after they start Desk View.


// An object that configures how to present Desk View.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeskViewApplication/LaunchConfiguration
type CaptureDeskViewApplicationLaunchConfiguration struct {
	objectivec.Object
}

// CaptureDeskViewApplicationLaunchConfigurationFrom constructs a [CaptureDeskViewApplicationLaunchConfiguration] from an unsafe.Pointer.
//
// An object that configures how to present Desk View.
func CaptureDeskViewApplicationLaunchConfigurationFrom(ptr unsafe.Pointer) CaptureDeskViewApplicationLaunchConfiguration {
	return CaptureDeskViewApplicationLaunchConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureDeskViewApplicationLaunchConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureDeskViewApplicationLaunchConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureDeskViewApplicationLaunchConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureDeskViewApplicationLaunchConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureDeskViewApplicationLaunchConfiguration */

// The frame for Desk View after it launches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeskViewApplication/LaunchConfiguration/mainWindowFrame
func (c_ CaptureDeskViewApplicationLaunchConfiguration) MainWindowFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c_.ID, objc.Sel("mainWindowFrame"))
	return rv
}/* debug [instance_properties/getter]: mainWindowFrame */


// The frame for Desk View after it launches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeskViewApplication/LaunchConfiguration/mainWindowFrame
func (c_ CaptureDeskViewApplicationLaunchConfiguration) SetMainWindowFrame(value corefoundation.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMainWindowFrame:"), value)
}/* debug [instance_properties/setter]: mainWindowFrame */


// A Boolean value that specifies whether the system requires the user to complete setup mode before it executes the completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeskViewApplication/LaunchConfiguration/requiresSetUpModeCompletion
func (c_ CaptureDeskViewApplicationLaunchConfiguration) RequiresSetUpModeCompletion() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("requiresSetUpModeCompletion"))
	return rv
}/* debug [instance_properties/getter]: requiresSetUpModeCompletion */


// A Boolean value that specifies whether the system requires the user to complete setup mode before it executes the completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeskViewApplication/LaunchConfiguration/requiresSetUpModeCompletion
func (c_ CaptureDeskViewApplicationLaunchConfiguration) SetRequiresSetUpModeCompletion(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRequiresSetUpModeCompletion:"), value)
}/* debug [instance_properties/setter]: requiresSetUpModeCompletion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureDeskViewApplicationLaunchConfiguration */



