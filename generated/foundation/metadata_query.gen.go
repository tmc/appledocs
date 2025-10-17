// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MetadataQuery] class.
var metadataQueryClass = _MetadataQueryClass{objc.GetClass("NSMetadataQuery")}

type _MetadataQueryClass struct {
	class objc.Class
}

// A query that you perform against Spotlight metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery

type MetadataQuery struct {
	objectivec.Object
}

// MetadataQueryFrom constructs a [MetadataQuery] from an unsafe.Pointer.
//
// A query that you perform against Spotlight metadata.
func MetadataQueryFrom(ptr unsafe.Pointer) MetadataQuery {
	return MetadataQuery{objectivec.Object{objc.ID(ptr)}}
}



