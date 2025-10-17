// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MetadataItem] class.
var metadataItemClass = _MetadataItemClass{objc.GetClass("NSMetadataItem")}

type _MetadataItemClass struct {
	class objc.Class
}

// An interface definition for the [MetadataItem] class.
type IMetadataItem interface {
	objectivec.IObject
}

// The metadata associated with a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataItem

type MetadataItem struct {
	objectivec.Object
}

// MetadataItemFrom constructs a [MetadataItem] from an unsafe.Pointer.
//
// The metadata associated with a file.
func MetadataItemFrom(ptr unsafe.Pointer) MetadataItem {
	return MetadataItem{objectivec.Object{objc.ID(ptr)}}
}



