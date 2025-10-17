
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DataAsset] class.
var DataAssetClass _DataAssetClass

func init() {
	DataAssetClass = _DataAssetClass{objc.GetClass("NSDataAsset")}
}

type _DataAssetClass struct {
	objc.Class
}

// An interface definition for the [DataAsset] class.
type IDataAsset interface {
	ID() objc.ID
}

type DataAsset struct {
	id objc.ID
}

func DataAssetFrom(ptr unsafe.Pointer) DataAsset {
	return DataAsset{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ DataAsset) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _DataAssetClass) Alloc() DataAsset {
	rv := objc.Send[DataAsset](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _DataAssetClass) New() DataAsset {
	rv := objc.Send[DataAsset](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewDataAsset creates and returns a new initialized instance.
func NewDataAsset() DataAsset {
	return DataAssetClass.New()
}

// Init initializes the instance.
func (d_ DataAsset) Init() DataAsset {
	rv := objc.Send[DataAsset](d_.ID(), selInit)
	return rv
}
