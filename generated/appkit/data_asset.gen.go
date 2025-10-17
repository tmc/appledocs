// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DataAsset] class.
var DataAssetClass objc.Class

func init() {
	DataAssetClass = objc.GetClass("NSDataAsset")
}

type DataAsset struct {
	objc.ID
}

func DataAssetFrom(ptr unsafe.Pointer) DataAsset {
	return DataAsset{
		ID: objc.ID(ptr),
	}
}



