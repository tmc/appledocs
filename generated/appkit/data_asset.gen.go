// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DataAsset] class.
var dataAssetClass = _DataAssetClass{objc.GetClass("NSDataAsset")}

type _DataAssetClass struct {
	class objc.Class
}

// An interface definition for the [DataAsset] class.
type IDataAsset interface {
	objectivec.IObject
}

// An object from a data set type stored in an asset catalog. [Full Topic]
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



