// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptureDeskViewApplication] class.
var (
	CaptureDeskViewApplicationClass     _CaptureDeskViewApplicationClass
	CaptureDeskViewApplicationClassOnce sync.Once
)

func getCaptureDeskViewApplicationClass() _CaptureDeskViewApplicationClass {
	CaptureDeskViewApplicationClassOnce.Do(func() {
		CaptureDeskViewApplicationClass = _CaptureDeskViewApplicationClass{objc.GetClass("AVCaptureDeskViewApplication")}
	})
	return CaptureDeskViewApplicationClass
}

type _CaptureDeskViewApplicationClass struct {
	class objc.Class
}





// An interface definition for the [CaptureDeskViewApplication] class.
type ICaptureDeskViewApplication interface {
	objectivec.IObject
	

	// properties:


	

	// methods:
	PresentWithCompletionHandler(completionHandler unsafe.Pointer)
	PresentWithLaunchConfigurationCompletionHandler(launchConfiguration IAVCaptureDeskViewApplicationLaunchConfiguration, completionHandler unsafe.Pointer)


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureDeskViewApplicationClass) Alloc() CaptureDeskViewApplication {
	rv := objc.Send[CaptureDeskViewApplication](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureDeskViewApplicationClass) New() CaptureDeskViewApplication {
	rv := objc.Send[CaptureDeskViewApplication](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureDeskViewApplication) Init() CaptureDeskViewApplication {
	rv := objc.Send[CaptureDeskViewApplication](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureDeskViewApplication) Autorelease() CaptureDeskViewApplication {
	rv := objc.Send[CaptureDeskViewApplication](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureDeskViewApplication creates a new CaptureDeskViewApplication instance.
func NewCaptureDeskViewApplication() CaptureDeskViewApplication {
	return getCaptureDeskViewApplicationClass().New()
}





// An object that programmatically presents Desk View.
//
// Use this class to programmatically launch Desk View from your app. You can optionally customize the presentation and specifiy an action to take afterward. The following example shows how to configure and present Desk View with a completion handler:


// An object that programmatically presents Desk View.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeskViewApplication
type CaptureDeskViewApplication struct {
	objectivec.Object
}

// CaptureDeskViewApplicationFrom constructs a [CaptureDeskViewApplication] from an unsafe.Pointer.
//
// An object that programmatically presents Desk View.
func CaptureDeskViewApplicationFrom(ptr unsafe.Pointer) CaptureDeskViewApplication {
	return CaptureDeskViewApplication{objectivec.Object{objc.ID(ptr)}}
}




















// Launches Desk View with no additional configuration and then performs a completion handler if you specify it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeskViewApplication/present(completionHandler:)
func (c_ CaptureDeskViewApplication) PresentWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("presentWithCompletionHandler:"), completionHandler)
}


// Launches Desk View with the configuration and completion handler that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeskViewApplication/present(launchConfiguration:completionHandler:)
func (c_ CaptureDeskViewApplication) PresentWithLaunchConfigurationCompletionHandler(launchConfiguration IAVCaptureDeskViewApplicationLaunchConfiguration, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("presentWithLaunchConfiguration:completionHandler:"), launchConfiguration, completionHandler)
}













