// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MetadataItem] class.
var MetadataItemClass objc.Class

func init() {
	MetadataItemClass = objc.GetClass("NSMetadataItem")
}

type MetadataItem struct {
	objc.ID
}

func MetadataItemFrom(ptr unsafe.Pointer) MetadataItem {
	return MetadataItem{
		ID: objc.ID(ptr),
	}
}



