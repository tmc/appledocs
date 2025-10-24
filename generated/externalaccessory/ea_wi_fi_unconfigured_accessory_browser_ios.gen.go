//go:build darwin && ios

// Code generated from Apple documentation for ExternalAccessory. DO NOT EDIT.

package externalaccessory

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for EAWiFiUnconfiguredAccessoryBrowser


// Starts the search for unconfigured accessories that match the specified predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowser/startSearchingForUnconfiguredAccessories(matching:)
func (e_ EAWiFiUnconfiguredAccessoryBrowser) StartSearchingForUnconfiguredAccessoriesMatchingPredicate(predicate objc.IObject /* cross-framework: Predicate */) {
	objc.Send[objc.ID](e_.ID, objc.Sel("startSearchingForUnconfiguredAccessoriesMatchingPredicate:"), predicate)
}

// iOS-only properties

// The object that acts as the delegate of the browser and receives browser events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessoryBrowser/delegate
func (e_ EAWiFiUnconfiguredAccessoryBrowser) Delegate() objc.ID {
	rv := objc.Send[objc.ID](e_.ID, objc.Sel("delegate"))
	return rv
}
func (e_ EAWiFiUnconfiguredAccessoryBrowser) SetDelegate(value objc.ID) {
	e_.ID.Send(objc.RegisterName("setDelegate:"), value)
}






