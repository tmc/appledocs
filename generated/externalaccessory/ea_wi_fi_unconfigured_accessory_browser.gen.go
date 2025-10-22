// Code generated from Apple documentation for ExternalAccessory. DO NOT EDIT.

package externalaccessory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EAWiFiUnconfiguredAccessoryBrowser] class.
var (
	EAWiFiUnconfiguredAccessoryBrowserClass     _EAWiFiUnconfiguredAccessoryBrowserClass
	EAWiFiUnconfiguredAccessoryBrowserClassOnce sync.Once
)

func getEAWiFiUnconfiguredAccessoryBrowserClass() _EAWiFiUnconfiguredAccessoryBrowserClass {
	EAWiFiUnconfiguredAccessoryBrowserClassOnce.Do(func() {
		EAWiFiUnconfiguredAccessoryBrowserClass = _EAWiFiUnconfiguredAccessoryBrowserClass{objc.GetClass("EAWiFiUnconfiguredAccessoryBrowser")}
	})
	return EAWiFiUnconfiguredAccessoryBrowserClass
}

type _EAWiFiUnconfiguredAccessoryBrowserClass struct {
	class objc.Class
}

// An interface definition for the [EAWiFiUnconfiguredAccessoryBrowser] class.
type IEAWiFiUnconfiguredAccessoryBrowser interface {
	objectivec.IObject
	ConfigureAccessoryWithConfigurationUIOnViewController(accessory IEAWiFiUnconfiguredAccessory, viewController appkit.IViewController)
	StartSearchingForUnconfiguredAccessoriesMatchingPredicate(predicate foundation.IPredicate)
	StopSearchingForUnconfiguredAccessories()
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	UnconfiguredAccessories() unsafe.Pointer
}

// An object you use to scan for wireless accessories and configure them for use with the user’s app.
//
// The class gives your app access to the MFi Wireless Accessory Configuration process. You use a browser object to scan for unconfigured accessories, connect them to the user’s Wi-Fi infrastructure, and configure attributes of the accessories. An accessory is represented by an instance of .
//
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowser
type EAWiFiUnconfiguredAccessoryBrowser struct {
	objectivec.Object
}

// EAWiFiUnconfiguredAccessoryBrowserFrom constructs a [EAWiFiUnconfiguredAccessoryBrowser] from an unsafe.Pointer.
//
// An object you use to scan for wireless accessories and configure them for use with the user’s app.
func EAWiFiUnconfiguredAccessoryBrowserFrom(ptr unsafe.Pointer) EAWiFiUnconfiguredAccessoryBrowser {
	return EAWiFiUnconfiguredAccessoryBrowser{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EAWiFiUnconfiguredAccessoryBrowserClass) Alloc() EAWiFiUnconfiguredAccessoryBrowser {
	rv := objc.Send[EAWiFiUnconfiguredAccessoryBrowser](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EAWiFiUnconfiguredAccessoryBrowserClass) New() EAWiFiUnconfiguredAccessoryBrowser {
	rv := objc.Send[EAWiFiUnconfiguredAccessoryBrowser](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EAWiFiUnconfiguredAccessoryBrowser) Init() EAWiFiUnconfiguredAccessoryBrowser {
	rv := objc.Send[EAWiFiUnconfiguredAccessoryBrowser](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EAWiFiUnconfiguredAccessoryBrowser) Autorelease() EAWiFiUnconfiguredAccessoryBrowser {
	rv := objc.Send[EAWiFiUnconfiguredAccessoryBrowser](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEAWiFiUnconfiguredAccessoryBrowser creates a new EAWiFiUnconfiguredAccessoryBrowser instance.
func NewEAWiFiUnconfiguredAccessoryBrowser() EAWiFiUnconfiguredAccessoryBrowser {
	return getEAWiFiUnconfiguredAccessoryBrowserClass().New()
}




// Creates a browser object that scans for unconfigured accessories.
//
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowser/init(delegate:queue:)
func NewEAWiFiUnconfiguredAccessoryBrowserWithDelegateQueue(delegate objectivec.IObject, queue unsafe.Pointer) EAWiFiUnconfiguredAccessoryBrowser {
	instance := getEAWiFiUnconfiguredAccessoryBrowserClass().Alloc()
	rv := objc.Send[EAWiFiUnconfiguredAccessoryBrowser](instance.ID, objc.Sel("initWithDelegate:queue:"), delegate, queue)
	rv.Autorelease()
	return rv
}


// Begins the configuration process for the specified accessory.
//
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowser/configureAccessory(_:withConfigurationUIOn:)
func (e_ EAWiFiUnconfiguredAccessoryBrowser) ConfigureAccessoryWithConfigurationUIOnViewController(accessory IEAWiFiUnconfiguredAccessory, viewController appkit.IViewController) {
	objc.Send[objc.ID](e_.ID, objc.Sel("configureAccessory:withConfigurationUIOnViewController:"), accessory, viewController)
}

// Starts the search for unconfigured accessories that match the specified predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowser/startSearchingForUnconfiguredAccessories(matching:)
func (e_ EAWiFiUnconfiguredAccessoryBrowser) StartSearchingForUnconfiguredAccessoriesMatchingPredicate(predicate foundation.IPredicate) {
	objc.Send[objc.ID](e_.ID, objc.Sel("startSearchingForUnconfiguredAccessoriesMatchingPredicate:"), predicate)
}

// Stops the search for unconfigured accessories.
//
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowser/stopSearchingForUnconfiguredAccessories()
func (e_ EAWiFiUnconfiguredAccessoryBrowser) StopSearchingForUnconfiguredAccessories() {
	objc.Send[objc.ID](e_.ID, objc.Sel("stopSearchingForUnconfiguredAccessories"))
}

// The object that acts as the delegate of the browser and receives browser events.
//
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowser/delegate
func (e_ EAWiFiUnconfiguredAccessoryBrowser) Delegate() objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The object that acts as the delegate of the browser and receives browser events.

//
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowser/delegate
func (e_ EAWiFiUnconfiguredAccessoryBrowser) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setDelegate:"), value)
}

// The set of unconfigured accessories that have been discovered.
//
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowser/unconfiguredAccessories
func (e_ EAWiFiUnconfiguredAccessoryBrowser) UnconfiguredAccessories() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("unconfiguredAccessories"))
	return rv
}


