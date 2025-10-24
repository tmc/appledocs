// Code generated from Apple documentation for AdSupport. DO NOT EDIT.

package adsupport

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASIdentifierManager */


/* debug [class_header]: Header for ASIdentifierManager */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for IdentifierManager */
// An interface definition for the [IdentifierManager] class.
type IIdentifierManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for IdentifierManager */
	// properties:
	AdvertisingIdentifier() foundation.UUID
	AdvertisingTrackingEnabled() bool
	IsAdvertisingTrackingEnabled() bool
	SetIsAdvertisingTrackingEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for IdentifierManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for IdentifierManager */
// Alloc allocates a new instance without initialization.
func (ic _IdentifierManagerClass) Alloc() IdentifierManager {
	rv := objc.Send[IdentifierManager](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for IdentifierManager */
// The object that contains the advertising identifier.


// The object that contains the advertising identifier.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for IdentifierManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for IdentifierManager */

// The shared instance of the identifier manager class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AdSupport/ASIdentifierManager/shared()
func (ic _IdentifierManagerClass) SharedManager() IIdentifierManager {
	rv := objc.Send[IdentifierManager](objc.ID(ic.class), objc.Sel("sharedManager"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedManager) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for IdentifierManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for IdentifierManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for IdentifierManager */

// The UUID that is specific to a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AdSupport/ASIdentifierManager/advertisingIdentifier
func (i_ IdentifierManager) AdvertisingIdentifier() foundation.UUID {
	rv := objc.Send[foundation.UUID](i_.ID, objc.Sel("advertisingIdentifier"))
	return rv
}/* debug [instance_properties/getter]: advertisingIdentifier */


// A Boolean value that indicates whether the user has limited ad tracking enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AdSupport/ASIdentifierManager/isAdvertisingTrackingEnabled
func (i_ IdentifierManager) AdvertisingTrackingEnabled() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("advertisingTrackingEnabled"))
	return rv
}/* debug [instance_properties/getter]: advertisingTrackingEnabled */


// A Boolean value that indicates whether the user has limited ad tracking
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/adsupport/asidentifiermanager/isadvertisingtrackingenabled
func (i_ IdentifierManager) IsAdvertisingTrackingEnabled() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("isAdvertisingTrackingEnabled"))
	return rv
}/* debug [instance_properties/getter]: isAdvertisingTrackingEnabled */


// A Boolean value that indicates whether the user has limited ad tracking
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/adsupport/asidentifiermanager/isadvertisingtrackingenabled
func (i_ IdentifierManager) SetIsAdvertisingTrackingEnabled(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIsAdvertisingTrackingEnabled:"), value)
}/* debug [instance_properties/setter]: isAdvertisingTrackingEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASIdentifierManager */






