// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	Data() foundation.foundation.INSData
	Name() DataAssetName
	TypeIdentifier() foundation.foundation.INSString


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (dc _DataAssetClass) Alloc() DataAsset {
	rv := objc.Send[DataAsset](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An object from a data set type stored in an asset catalog.
//
// The object’s content is stored as a set of one or more files with associated device attributes. These sets can also be tagged for use as on-demand resources.


// An object from a data set type stored in an asset catalog.
//
// [Full Topic]
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






// Initializes and returns an object with a reference to the named data asset in an asset catalog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDataAsset/init(name:)
func NewDataAssetWithName(name DataAssetName) DataAsset {
	instance := getDataAssetClass().Alloc()
	rv := objc.Send[DataAsset](instance.ID, objc.Sel("initWithName:"), name)
	rv.Autorelease()
	return rv
}


// Initializes and returns an object with a reference to the named data asset that’s in an asset catalog in the specified bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDataAsset/init(name:bundle:)
func NewDataAssetWithNameBundle(name DataAssetName, bundle foundation.Bundle) DataAsset {
	instance := getDataAssetClass().Alloc()
	rv := objc.Send[DataAsset](instance.ID, objc.Sel("initWithName:bundle:"), name, bundle)
	rv.Autorelease()
	return rv
}






















// The raw data values in the data asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDataAsset/data
func (d_ DataAsset) Data() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](d_.ID, objc.Sel("data"))
	return rv
}


// The name of the data set in the asset catalog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDataAsset/name-swift.property
func (d_ DataAsset) Name() DataAssetName {
	rv := objc.Send[DataAssetName](d_.ID, objc.Sel("name"))
	return rv
}


// The uniform type identifier for the data asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDataAsset/typeIdentifier
func (d_ DataAsset) TypeIdentifier() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("typeIdentifier"))
	return rv
}







