// Code generated from Apple documentation for StoreKitTest. DO NOT EDIT.

package storekittest

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AdTestPostback] class.
var (
	AdTestPostbackClass     _AdTestPostbackClass
	AdTestPostbackClassOnce sync.Once
)

func getAdTestPostbackClass() _AdTestPostbackClass {
	AdTestPostbackClassOnce.Do(func() {
		AdTestPostbackClass = _AdTestPostbackClass{objc.GetClass("SKAdTestPostback")}
	})
	return AdTestPostbackClass
}

type _AdTestPostbackClass struct {
	class objc.Class
}

// An interface definition for the [AdTestPostback] class.
type IAdTestPostback interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A test postback that contains ad conversion information in the testing environment.
//
// Use this class to create test postbacks to use for unit testing. In the production environment, the system creates a postback after a user installs an advertised app. The advertised app is responsible for registering the installation and may update the conversion value. The system sends the postback after a timer expires. In the testing environment, you can mimic a postback by creating it directly. You control the property values within the postback. Use it to test your app’s ability to register the app installation and update conversion values, and to test your server’s ability to receive postbacks.


// A test postback that contains ad conversion information in the testing environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback
type AdTestPostback struct {
	objectivec.Object
}

// AdTestPostbackFrom constructs a [AdTestPostback] from an unsafe.Pointer.
//
// A test postback that contains ad conversion information in the testing environment.
func AdTestPostbackFrom(ptr unsafe.Pointer) AdTestPostback {
	return AdTestPostback{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AdTestPostbackClass) Alloc() AdTestPostback {
	rv := objc.Send[AdTestPostback](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AdTestPostbackClass) New() AdTestPostback {
	rv := objc.Send[AdTestPostback](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AdTestPostback) Init() AdTestPostback {
	rv := objc.Send[AdTestPostback](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AdTestPostback) Autorelease() AdTestPostback {
	rv := objc.Send[AdTestPostback](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAdTestPostback creates a new AdTestPostback instance.
func NewAdTestPostback() AdTestPostback {
	return getAdTestPostbackClass().New()
}



// Creates a test postback for an in-app ad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/init(version:adNetworkIdentifier:adCampaignIdentifier:appStoreItemIdentifier:sourceAppStoreItemIdentifier:conversionValue:fidelityType:isRedownload:didWin:postbackURL:)
func NewAdTestPostbackWithVersionAdNetworkIdentifierAdCampaignIdentifierAppStoreItemIdentifierSourceAppStoreItemIdentifierConversionValueFidelityTypeIsRedownloadDidWinPostbackURL(version objc.IObject /* cross-framework: AdTestPostbackVersion */, adNetworkIdentifier objc.IObject /* cross-framework: NSString */, adCampaignIdentifier int, appStoreItemIdentifier int, sourceAppStoreItemIdentifier int, conversionValue int, fidelityType int, isRedownload bool, didWin bool, postbackURL objc.IObject /* cross-framework: NSString */) AdTestPostback {
	instance := getAdTestPostbackClass().Alloc()
	rv := objc.Send[AdTestPostback](instance.ID, objc.Sel("initWithVersion:adNetworkIdentifier:adCampaignIdentifier:appStoreItemIdentifier:sourceAppStoreItemIdentifier:conversionValue:fidelityType:isRedownload:didWin:postbackURL:"), version, adNetworkIdentifier, adCampaignIdentifier, appStoreItemIdentifier, sourceAppStoreItemIdentifier, conversionValue, fidelityType, isRedownload, didWin, postbackURL)
	rv.Autorelease()
	return rv
}


// Creates a test postback for a web ad or an in-app ad.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/init(version:adNetworkIdentifier:sourceIdentifier:appStoreItemIdentifier:sourceAppStoreItemIdentifier:sourceDomain:fidelityType:isRedownload:didWin:postbackURL:)
func NewAdTestPostbackWithVersionAdNetworkIdentifierSourceIdentifierAppStoreItemIdentifierSourceAppStoreItemIdentifierSourceDomainFidelityTypeIsRedownloadDidWinPostbackURL(version objc.IObject /* cross-framework: AdTestPostbackVersion */, adNetworkIdentifier objc.IObject /* cross-framework: NSString */, sourceIdentifier objc.IObject /* cross-framework: NSString */, appStoreItemIdentifier int, sourceAppStoreItemIdentifier int, sourceDomain objc.IObject /* cross-framework: NSString */, fidelityType int, isRedownload bool, didWin bool, postbackURL objc.IObject /* cross-framework: NSString */) AdTestPostback {
	instance := getAdTestPostbackClass().Alloc()
	rv := objc.Send[AdTestPostback](instance.ID, objc.Sel("initWithVersion:adNetworkIdentifier:sourceIdentifier:appStoreItemIdentifier:sourceAppStoreItemIdentifier:sourceDomain:fidelityType:isRedownload:didWin:postbackURL:"), version, adNetworkIdentifier, sourceIdentifier, appStoreItemIdentifier, sourceAppStoreItemIdentifier, sourceDomain, fidelityType, isRedownload, didWin, postbackURL)
	rv.Autorelease()
	return rv
}



// Creates a sequence of test postbacks for an in-app or web ad impression.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestPostback/winningPostbacks(withVersion:adNetworkIdentifier:sourceIdentifier:appStoreItemIdentifier:sourceAppStoreItemIdentifier:sourceDomain:fidelityType:isRedownload:postbackURL:)
func (ac _AdTestPostbackClass) WinningPostbacksWithVersionAdNetworkIdentifierSourceIdentifierAppStoreItemIdentifierSourceAppStoreItemIdentifierSourceDomainFidelityTypeIsRedownloadPostbackURL(version objc.IObject /* cross-framework: AdTestPostbackVersion */, adNetworkIdentifier objc.IObject /* cross-framework: NSString */, sourceIdentifier objc.IObject /* cross-framework: NSString */, appStoreItemIdentifier int, sourceAppStoreItemIdentifier int, sourceDomain objc.IObject /* cross-framework: NSString */, fidelityType int, isRedownload bool, postbackURL objc.IObject /* cross-framework: NSString */) []IAdTestPostback {
	rv := objc.Send[[]AdTestPostback](objc.ID(ac.class), objc.Sel("winningPostbacksWithVersion:adNetworkIdentifier:sourceIdentifier:appStoreItemIdentifier:sourceAppStoreItemIdentifier:sourceDomain:fidelityType:isRedownload:postbackURL:"), version, adNetworkIdentifier, sourceIdentifier, appStoreItemIdentifier, sourceAppStoreItemIdentifier, sourceDomain, fidelityType, isRedownload, postbackURL)
	return rv
}


