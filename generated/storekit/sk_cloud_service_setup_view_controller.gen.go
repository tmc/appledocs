// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [CloudServiceSetupViewController] class.
var (
	CloudServiceSetupViewControllerClass     _CloudServiceSetupViewControllerClass
	CloudServiceSetupViewControllerClassOnce sync.Once
)

func getCloudServiceSetupViewControllerClass() _CloudServiceSetupViewControllerClass {
	CloudServiceSetupViewControllerClassOnce.Do(func() {
		CloudServiceSetupViewControllerClass = _CloudServiceSetupViewControllerClass{objc.GetClass("SKCloudServiceSetupViewController")}
	})
	return CloudServiceSetupViewControllerClass
}

type _CloudServiceSetupViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [CloudServiceSetupViewController] class.
type ICloudServiceSetupViewController interface {
	appkit.IViewController
}

// A view controller that helps people perform setup for a cloud service, like an Apple Music subscription.
//
// Use the view that this view controller presents to allow customers to set up cloud services that are associated with their iTunes Store account, like an Apple Music subscription. To enable the Apple Music subscriber setup flow in particular, you first request the current set of capabilities from . Then, present the setup view controller only when the capability is enabled and the capability is disabled. For information about other capabilities that you can enable by using this view controller, see .
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceSetupViewController
type CloudServiceSetupViewController struct {
	appkit.ViewController
}

// CloudServiceSetupViewControllerFrom constructs a [CloudServiceSetupViewController] from an unsafe.Pointer.
//
// A view controller that helps people perform setup for a cloud service, like an Apple Music subscription.
func CloudServiceSetupViewControllerFrom(ptr unsafe.Pointer) CloudServiceSetupViewController {
	return CloudServiceSetupViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CloudServiceSetupViewControllerClass) Alloc() CloudServiceSetupViewController {
	rv := objc.Send[CloudServiceSetupViewController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CloudServiceSetupViewControllerClass) New() CloudServiceSetupViewController {
	rv := objc.Send[CloudServiceSetupViewController](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CloudServiceSetupViewController) Init() CloudServiceSetupViewController {
	rv := objc.Send[CloudServiceSetupViewController](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CloudServiceSetupViewController) Autorelease() CloudServiceSetupViewController {
	rv := objc.Send[CloudServiceSetupViewController](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCloudServiceSetupViewController creates a new CloudServiceSetupViewController instance.
func NewCloudServiceSetupViewController() CloudServiceSetupViewController {
	return getCloudServiceSetupViewControllerClass().New()
}


// The cloud service view controller’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceSetupViewController/delegate
func (c_ CloudServiceSetupViewController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The cloud service view controller’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKCloudServiceSetupViewController/delegate
func (c_ CloudServiceSetupViewController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}



