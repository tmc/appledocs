// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MutableAssetDownloadStorageManagementPolicy] class.
var (
	MutableAssetDownloadStorageManagementPolicyClass     _MutableAssetDownloadStorageManagementPolicyClass
	MutableAssetDownloadStorageManagementPolicyClassOnce sync.Once
)

func getMutableAssetDownloadStorageManagementPolicyClass() _MutableAssetDownloadStorageManagementPolicyClass {
	MutableAssetDownloadStorageManagementPolicyClassOnce.Do(func() {
		MutableAssetDownloadStorageManagementPolicyClass = _MutableAssetDownloadStorageManagementPolicyClass{objc.GetClass("AVMutableAssetDownloadStorageManagementPolicy")}
	})
	return MutableAssetDownloadStorageManagementPolicyClass
}

type _MutableAssetDownloadStorageManagementPolicyClass struct {
	class objc.Class
}





// An interface definition for the [MutableAssetDownloadStorageManagementPolicy] class.
type IMutableAssetDownloadStorageManagementPolicy interface {
	IAssetDownloadStorageManagementPolicy
	

	// properties:
	ExpirationDate() objc.IObject /* cross-framework: NSDate */
	SetExpirationDate(value objc.IObject /* cross-framework: NSDate */)
	Priority() AssetDownloadedAssetEvictionPriority /* typedef */
	SetPriority(value AssetDownloadedAssetEvictionPriority /* typedef */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MutableAssetDownloadStorageManagementPolicyClass) Alloc() MutableAssetDownloadStorageManagementPolicy {
	rv := objc.Send[MutableAssetDownloadStorageManagementPolicy](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MutableAssetDownloadStorageManagementPolicyClass) New() MutableAssetDownloadStorageManagementPolicy {
	rv := objc.Send[MutableAssetDownloadStorageManagementPolicy](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableAssetDownloadStorageManagementPolicy) Init() MutableAssetDownloadStorageManagementPolicy {
	rv := objc.Send[MutableAssetDownloadStorageManagementPolicy](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableAssetDownloadStorageManagementPolicy) Autorelease() MutableAssetDownloadStorageManagementPolicy {
	rv := objc.Send[MutableAssetDownloadStorageManagementPolicy](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableAssetDownloadStorageManagementPolicy creates a new MutableAssetDownloadStorageManagementPolicy instance.
func NewMutableAssetDownloadStorageManagementPolicy() MutableAssetDownloadStorageManagementPolicy {
	return getMutableAssetDownloadStorageManagementPolicyClass().New()
}





// A mutable object that you use to create a new storage management policy.


// A mutable object that you use to create a new storage management policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAssetDownloadStorageManagementPolicy
type MutableAssetDownloadStorageManagementPolicy struct {
	AssetDownloadStorageManagementPolicy
}

// MutableAssetDownloadStorageManagementPolicyFrom constructs a [MutableAssetDownloadStorageManagementPolicy] from an unsafe.Pointer.
//
// A mutable object that you use to create a new storage management policy.
func MutableAssetDownloadStorageManagementPolicyFrom(ptr unsafe.Pointer) MutableAssetDownloadStorageManagementPolicy {
	return MutableAssetDownloadStorageManagementPolicy{
		AssetDownloadStorageManagementPolicy: AssetDownloadStorageManagementPolicyFrom(ptr),
	}
}

























// The expiration date for an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAssetDownloadStorageManagementPolicy/expirationDate
func (m_ MutableAssetDownloadStorageManagementPolicy) ExpirationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](m_.ID, objc.Sel("expirationDate"))
	return rv
}


// The expiration date for an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAssetDownloadStorageManagementPolicy/expirationDate
func (m_ MutableAssetDownloadStorageManagementPolicy) SetExpirationDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExpirationDate:"), value)
}


// The eviction priority for a downloaded asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAssetDownloadStorageManagementPolicy/priority
func (m_ MutableAssetDownloadStorageManagementPolicy) Priority() AssetDownloadedAssetEvictionPriority /* typedef */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("priority"))
	return rv
}


// The eviction priority for a downloaded asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableAssetDownloadStorageManagementPolicy/priority
func (m_ MutableAssetDownloadStorageManagementPolicy) SetPriority(value AssetDownloadedAssetEvictionPriority /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPriority:"), value)
}








