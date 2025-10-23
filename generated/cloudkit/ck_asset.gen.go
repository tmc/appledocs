// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CKAsset] class.
var (
	CKAssetClass     _CKAssetClass
	CKAssetClassOnce sync.Once
)

func getCKAssetClass() _CKAssetClass {
	CKAssetClassOnce.Do(func() {
		CKAssetClass = _CKAssetClass{objc.GetClass("CKAsset")}
	})
	return CKAssetClass
}

type _CKAssetClass struct {
	class objc.Class
}

// An interface definition for the [CKAsset] class.
type ICKAsset interface {
	objectivec.IObject
	// properties:
	FileURL() foundation.objc.IObject /* cross-framework: URL */
	SetFileURL(value foundation.objc.IObject /* cross-framework: URL */)
	// methods:
}

// An external file that belongs to a record.
//
// Use assets to incorporate external files into your app’s records, such as photos, videos, and binary files. Alternatively, use assets when a field’s value is more than a few kilobytes in size. To associate an instance of with a record, assign it to one of its fields. CloudKit stores an asset’s data separately from a record that references it, but maintains an association with that record. When you save a record that has an asset, CloudKit saves both the record and the asset to the server. Similarly, when you fetch the record, the server returns the record and the asset. When you fetch a record that contains an asset, CloudKit stores the asset’s data in a staging area accessible to your app. Use the asset’s property to access its staged location. The system regularly deletes files in the staging area to reclaim disk space. To avoid this behavior, move the data into your app’s container as soon as you fetch it. If you don’t require the asset when retrieving records, use the operation’s property to exclude the field. For more information, see , , and . If you no longer require an asset that’s on the server, you don’t delete it. Instead, orphan the asset by setting any fields that contain the asset to and then saving the record. CloudKit periodically deletes orphaned assets from the server.


// An external file that belongs to a record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CloudKit/CKAsset
type CKAsset struct {
	objectivec.Object
}

// CKAssetFrom constructs a [CKAsset] from an unsafe.Pointer.
//
// An external file that belongs to a record.
func CKAssetFrom(ptr unsafe.Pointer) CKAsset {
	return CKAsset{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CKAssetClass) Alloc() CKAsset {
	rv := objc.Send[CKAsset](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CKAssetClass) New() CKAsset {
	rv := objc.Send[CKAsset](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CKAsset) Init() CKAsset {
	rv := objc.Send[CKAsset](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CKAsset) Autorelease() CKAsset {
	rv := objc.Send[CKAsset](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKAsset creates a new CKAsset instance.
func NewCKAsset() CKAsset {
	return getCKAssetClass().New()
}



// The URL for accessing the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckasset/fileurl
func (c_ CKAsset) FileURL() foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("fileURL"))
	return rv
}


// The URL for accessing the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/cloudkit/ckasset/fileurl
func (c_ CKAsset) SetFileURL(value foundation.objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFileURL:"), value)
}



