// Code generated from Apple documentation for AdSupport. DO NOT EDIT.

package adsupport

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [IdentifierManager] class.
var (
	IdentifierManagerClass     _IdentifierManagerClass
	IdentifierManagerClassOnce sync.Once
)

func getIdentifierManagerClass() _IdentifierManagerClass {
	IdentifierManagerClassOnce.Do(func() {
		IdentifierManagerClass = _IdentifierManagerClass{objc.GetClass("ASIdentifierManager")}
	})
	return IdentifierManagerClass
}

type _IdentifierManagerClass struct {
	class objc.Class
}

// An interface definition for the [IdentifierManager] class.
type IIdentifierManager interface {
	objectivec.IObject
	AdvertisingIdentifier() foundation.UUID
	AdvertisingTrackingEnabled() bool
	IsAdvertisingTrackingEnabled() bool
	SetIsAdvertisingTrackingEnabled(value bool)
}

// The object that contains the advertising identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AdSupport/ASIdentifierManager
type IdentifierManager struct {
	objectivec.Object
}

// IdentifierManagerFrom constructs a [IdentifierManager] from an unsafe.Pointer.
//
// The object that contains the advertising identifier.
func IdentifierManagerFrom(ptr unsafe.Pointer) IdentifierManager {
	return IdentifierManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _IdentifierManagerClass) Alloc() IdentifierManager {
	rv := objc.Send[IdentifierManager](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IdentifierManagerClass) New() IdentifierManager {
	rv := objc.Send[IdentifierManager](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IdentifierManager) Init() IdentifierManager {
	rv := objc.Send[IdentifierManager](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IdentifierManager) Autorelease() IdentifierManager {
	rv := objc.Send[IdentifierManager](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIdentifierManager creates a new IdentifierManager instance.
func NewIdentifierManager() IdentifierManager {
	return getIdentifierManagerClass().New()
}


// The shared instance of the identifier manager class.
//
// [Full Topic]: https://developer.apple.com/documentation/AdSupport/ASIdentifierManager/shared()
func (ic _IdentifierManagerClass) SharedManager() IdentifierManager {
	rv := objc.Send[IdentifierManager](objc.ID(ic.class), objc.Sel("sharedManager"))
	return rv
}

// The UUID that is specific to a device.
//
// [Full Topic]: https://developer.apple.com/documentation/AdSupport/ASIdentifierManager/advertisingIdentifier
func (i_ IdentifierManager) AdvertisingIdentifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](i_.ID, objc.Sel("advertisingIdentifier"))
	return rv
}

// A Boolean value that indicates whether the user has limited ad tracking enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/AdSupport/ASIdentifierManager/isAdvertisingTrackingEnabled
func (i_ IdentifierManager) AdvertisingTrackingEnabled() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("advertisingTrackingEnabled"))
	return rv
}

// A Boolean value that indicates whether the user has limited ad tracking
//
// [Full Topic]: https://developer.apple.com/documentation/adsupport/asidentifiermanager/isadvertisingtrackingenabled
func (i_ IdentifierManager) IsAdvertisingTrackingEnabled() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isAdvertisingTrackingEnabled"))
	return rv
}


// SetIsAdvertisingTrackingEnabled sets the value of the isAdvertisingTrackingEnabled property.
// A Boolean value that indicates whether the user has limited ad tracking

//
// [Full Topic]: https://developer.apple.com/documentation/adsupport/asidentifiermanager/isadvertisingtrackingenabled
func (i_ IdentifierManager) SetIsAdvertisingTrackingEnabled(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsAdvertisingTrackingEnabled:"), value)
}




