// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AdNetwork] class.
var (
	AdNetworkClass     _AdNetworkClass
	AdNetworkClassOnce sync.Once
)

func getAdNetworkClass() _AdNetworkClass {
	AdNetworkClassOnce.Do(func() {
		AdNetworkClass = _AdNetworkClass{objc.GetClass("SKAdNetwork")}
	})
	return AdNetworkClass
}

type _AdNetworkClass struct {
	class objc.Class
}

// An interface definition for the [AdNetwork] class.
type IAdNetwork interface {
	objectivec.IObject
}

// A class that validates advertisement-driven app installations.
//
// The ad network API helps advertisers measure the success of ad campaigns while maintaining user privacy. The API involves three participants: that sign ads and receive install-validation postbacks after ads result in conversions that display ads from the ad networks, or websites that display the ads in Safari that update conversion values as people engage with the app Ad networks register with Apple to get an ad network ID and to use the API. Developers configure their apps to accept attributable ads from ad networks, and to receive copies of winning postbacks. For information about setup, see , , and . For information about displaying ads in Safari, see . The following diagram shows the path of an ad impression that wins ad attribution. The ad network serves an ad that an app or Safari web page displays. A user taps the ad and downloads the advertised app. Apple determines a postback data tier for the app download, and the device uses the tier later to determine the level of detail the postback can contain to ensure crowd anonymity. For more information about the postback contents and the data tiers, see . If the user launches the app within an attribution time-window, the ad impression is eligible for install-attribution postbacks. As the user engages with the app, the app updates the conversion value. Starting in iOS 16.1, apps can update conversion values during three conversion windows, which results in up to three postbacks for an ad signed using version 4. The system sends the postbacks to the ad network, and to the app’s developer if they opt in to receive postbacks. Devices send install-validation postbacks to multiple ad networks that sign their ads using version 3 or later. One ad network receives a postback with a  parameter value of  for the ad impression that wins the ad attribution. Up to five other ad networks receive a postback with a   parameter value of   if their ad impressions qualify for the attribution, but don’t win. The following diagram shows the path of ad impressions that qualify for, but don’t win, the ad attribution. Up to five ad networks receive a single nonwinning postback. For more information about receiving ad attributions, including time-window details and other constraints, see . The information in the postback that Apple cryptographically signs doesn’t include user- or device-specific data. It may include values from the ad network and the advertised app if providing those values meets Apple’s privacy threshold. For more information about postback values and postback data tiers, see . For more information about the contents of postbacks for each SKAdNetwork version, see .
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdNetwork
type AdNetwork struct {
	objectivec.Object
}

// AdNetworkFrom constructs a [AdNetwork] from an unsafe.Pointer.
//
// A class that validates advertisement-driven app installations.
func AdNetworkFrom(ptr unsafe.Pointer) AdNetwork {
	return AdNetwork{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AdNetworkClass) Alloc() AdNetwork {
	rv := objc.Send[AdNetwork](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AdNetworkClass) New() AdNetwork {
	rv := objc.Send[AdNetwork](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AdNetwork) Init() AdNetwork {
	rv := objc.Send[AdNetwork](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AdNetwork) Autorelease() AdNetwork {
	rv := objc.Send[AdNetwork](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAdNetwork creates a new AdNetwork instance.
func NewAdNetwork() AdNetwork {
	return getAdNetworkClass().New()
}


// Indicates that your app is no longer presenting a view-through ad to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdNetwork/endImpression(_:completionHandler:)
func (ac _AdNetworkClass) EndImpressionCompletionHandler(impression ISKAdImpression, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("endImpression:completionHandler:"), impression, completion)
}

// Verifies the first launch of an app installed as a result of an ad.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdNetwork/registerAppForAdNetworkAttribution()
func (ac _AdNetworkClass) RegisterAppForAdNetworkAttribution() {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("registerAppForAdNetworkAttribution"))
}

// Indicates that your app is presenting a view-through ad to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdNetwork/startImpression(_:completionHandler:)
func (ac _AdNetworkClass) StartImpressionCompletionHandler(impression ISKAdImpression, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("startImpression:completionHandler:"), impression, completion)
}

// Updates the conversion value and verifies the first launch of an app installed as a result of an ad.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdNetwork/updateConversionValue(_:)
func (ac _AdNetworkClass) UpdateConversionValue(conversionValue int) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("updateConversionValue:"), conversionValue)
}

// Updates the fine and coarse conversion values, and calls a completion handler if the update fails.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdNetwork/updatePostbackConversionValue(_:coarseValue:completionHandler:)
func (ac _AdNetworkClass) UpdatePostbackConversionValueCoarseValueCompletionHandler(fineValue int, coarseValue IAdNetworkCoarseConversionValue, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("updatePostbackConversionValue:coarseValue:completionHandler:"), fineValue, coarseValue, completion)
}

// Updates the fine and coarse conversion values and indicates whether to send the postback before the conversion window ends, and calls a completion handler if the update fails.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdNetwork/updatePostbackConversionValue(_:coarseValue:lockWindow:completionHandler:)
func (ac _AdNetworkClass) UpdatePostbackConversionValueCoarseValueLockWindowCompletionHandler(fineValue int, coarseValue IAdNetworkCoarseConversionValue, lockWindow bool, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("updatePostbackConversionValue:coarseValue:lockWindow:completionHandler:"), fineValue, coarseValue, lockWindow, completion)
}

// Verifies the first launch of an advertised app and, on subsequent calls, updates the conversion value or calls a completion handler if the update fails.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKAdNetwork/updatePostbackConversionValue(_:completionHandler:)
func (ac _AdNetworkClass) UpdatePostbackConversionValueCompletionHandler(conversionValue int, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("updatePostbackConversionValue:completionHandler:"), conversionValue, completion)
}



