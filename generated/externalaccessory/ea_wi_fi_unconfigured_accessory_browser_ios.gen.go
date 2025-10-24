//go:build darwin && ios

// Code generated from Apple documentation for ExternalAccessory. DO NOT EDIT.

package externalaccessory

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for EAWiFiUnconfiguredAccessoryBrowser


// Begins the configuration process for the specified accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowser/configureAccessory(_:withConfigurationUIOn:)
func (e_ EAWiFiUnconfiguredAccessoryBrowser) ConfigureAccessoryWithConfigurationUIOnViewController(accessory IEAWiFiUnconfiguredAccessory, viewController appkit.ViewController) {
	objc.Send[objc.ID](e_.ID, objc.Sel("configureAccessory:withConfigurationUIOnViewController:"), accessory, viewController)
}

// Starts the search for unconfigured accessories that match the specified predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowser/startSearchingForUnconfiguredAccessories(matching:)
func (e_ EAWiFiUnconfiguredAccessoryBrowser) StartSearchingForUnconfiguredAccessoriesMatchingPredicate(predicate foundation.Predicate) {
	objc.Send[objc.ID](e_.ID, objc.Sel("startSearchingForUnconfiguredAccessoriesMatchingPredicate:"), predicate)
}

// Stops the search for unconfigured accessories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowser/stopSearchingForUnconfiguredAccessories()
func (e_ EAWiFiUnconfiguredAccessoryBrowser) StopSearchingForUnconfiguredAccessories() {
	objc.Send[objc.ID](e_.ID, objc.Sel("stopSearchingForUnconfiguredAccessories"))
}

// iOS-only properties

// The object that acts as the delegate of the browser and receives browser events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowser/delegate
func (e_ EAWiFiUnconfiguredAccessoryBrowser) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("delegate"))
	return rv
}
func (e_ EAWiFiUnconfiguredAccessoryBrowser) SetDelegate(value unsafe.Pointer) {
	e_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// The set of unconfigured accessories that have been discovered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowser/unconfiguredAccessories
func (e_ EAWiFiUnconfiguredAccessoryBrowser) UnconfiguredAccessories() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("unconfiguredAccessories"))
	return rv
}




