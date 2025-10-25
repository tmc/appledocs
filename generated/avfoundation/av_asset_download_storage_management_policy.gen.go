// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetDownloadStorageManagementPolicy */


/* debug [class_header]: Header for AVAssetDownloadStorageManagementPolicy */
// The class instance for the [AssetDownloadStorageManagementPolicy] class.
var (
	AssetDownloadStorageManagementPolicyClass     _AssetDownloadStorageManagementPolicyClass
	AssetDownloadStorageManagementPolicyClassOnce sync.Once
)

func getAssetDownloadStorageManagementPolicyClass() _AssetDownloadStorageManagementPolicyClass {
	AssetDownloadStorageManagementPolicyClassOnce.Do(func() {
		AssetDownloadStorageManagementPolicyClass = _AssetDownloadStorageManagementPolicyClass{objc.GetClass("AVAssetDownloadStorageManagementPolicy")}
	})
	return AssetDownloadStorageManagementPolicyClass
}

type _AssetDownloadStorageManagementPolicyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetDownloadStorageManagementPolicy */
// An interface definition for the [AssetDownloadStorageManagementPolicy] class.
type IAssetDownloadStorageManagementPolicy interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetDownloadStorageManagementPolicy */
	// properties:
	ExpirationDate() objc.IObject /* cross-framework: NSDate */
	Priority() AssetDownloadedAssetEvictionPriority /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetDownloadStorageManagementPolicy */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetDownloadStorageManagementPolicy */
// Alloc allocates a new instance without initialization.
func (ac _AssetDownloadStorageManagementPolicyClass) Alloc() AssetDownloadStorageManagementPolicy {
	rv := objc.Send[AssetDownloadStorageManagementPolicy](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetDownloadStorageManagementPolicyClass) New() AssetDownloadStorageManagementPolicy {
	rv := objc.Send[AssetDownloadStorageManagementPolicy](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetDownloadStorageManagementPolicy) Init() AssetDownloadStorageManagementPolicy {
	rv := objc.Send[AssetDownloadStorageManagementPolicy](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetDownloadStorageManagementPolicy) Autorelease() AssetDownloadStorageManagementPolicy {
	rv := objc.Send[AssetDownloadStorageManagementPolicy](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetDownloadStorageManagementPolicy creates a new AssetDownloadStorageManagementPolicy instance.
func NewAssetDownloadStorageManagementPolicy() AssetDownloadStorageManagementPolicy {
	return getAssetDownloadStorageManagementPolicyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetDownloadStorageManagementPolicy */
// An object that defines a policy to automatically manage the storage of downloaded assets.


// An object that defines a policy to automatically manage the storage of downloaded assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManagementPolicy
type AssetDownloadStorageManagementPolicy struct {
	objectivec.Object
}

// AssetDownloadStorageManagementPolicyFrom constructs a [AssetDownloadStorageManagementPolicy] from an unsafe.Pointer.
//
// An object that defines a policy to automatically manage the storage of downloaded assets.
func AssetDownloadStorageManagementPolicyFrom(ptr unsafe.Pointer) AssetDownloadStorageManagementPolicy {
	return AssetDownloadStorageManagementPolicy{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetDownloadStorageManagementPolicy *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetDownloadStorageManagementPolicy */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetDownloadStorageManagementPolicy */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetDownloadStorageManagementPolicy */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetDownloadStorageManagementPolicy */

// The expiration date for an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManagementPolicy/expirationDate
func (a_ AssetDownloadStorageManagementPolicy) ExpirationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](a_.ID, objc.Sel("expirationDate"))
	return rv
}/* debug [instance_properties/getter]: expirationDate */


// The eviction priority for an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadStorageManagementPolicy/priority
func (a_ AssetDownloadStorageManagementPolicy) Priority() AssetDownloadedAssetEvictionPriority /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("priority"))
	return rv
}/* debug [instance_properties/getter]: priority */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetDownloadStorageManagementPolicy */



