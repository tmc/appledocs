// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [DataAsset] class.
var (
	DataAssetClass     _DataAssetClass
	DataAssetClassOnce sync.Once
)

func getDataAssetClass() _DataAssetClass {
	DataAssetClassOnce.Do(func() {
		DataAssetClass = _DataAssetClass{objc.GetClass("NSDataAsset")}
	})
	return DataAssetClass
}

type _DataAssetClass struct {
	class objc.Class
}

// An interface definition for the [DataAsset] class.
type IDataAsset interface {
	objectivec.IObject
}

// An object from a data set type stored in an asset catalog.
//
// The object’s content is stored as a set of one or more files with associated device attributes. These sets can also be tagged for use as on-demand resources.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDataAsset
type DataAsset struct {
	objectivec.Object
}

// DataAssetFrom constructs a [DataAsset] from an unsafe.Pointer.
//
// An object from a data set type stored in an asset catalog.
func DataAssetFrom(ptr unsafe.Pointer) DataAsset {
	return DataAsset{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DataAssetClass) Alloc() DataAsset {
	rv := objc.Send[DataAsset](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DataAssetClass) New() DataAsset {
	rv := objc.Send[DataAsset](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DataAsset) Init() DataAsset {
	rv := objc.Send[DataAsset](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DataAsset) Autorelease() DataAsset {
	rv := objc.Send[DataAsset](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDataAsset creates a new DataAsset instance.
func NewDataAsset() DataAsset {
	return getDataAssetClass().New()
}




