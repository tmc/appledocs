// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var metadataItemClass _MetadataItemClass

func init() {
	metadataItemClass = _MetadataItemClass{objc.GetClass("NSMetadataItem")}
}

type _MetadataItemClass struct {
	class objc.Class
}

type MetadataItem struct {
	objc.ID
}

func MetadataItemFrom(ptr unsafe.Pointer) MetadataItem {
	return MetadataItem{
		ID: objc.ID(ptr),
	}
}




