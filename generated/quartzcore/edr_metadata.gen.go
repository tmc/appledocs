// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [EDRMetadata] class.
var eDRMetadataClass = _EDRMetadataClass{objc.GetClass("CAEDRMetadata")}

type _EDRMetadataClass struct {
	class objc.Class
}

// Metadata describing how extended dynamic range (EDR) values should be tone mapped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAEDRMetadata

type EDRMetadata struct {
	objectivec.Object
}

// EDRMetadataFrom constructs a [EDRMetadata] from an unsafe.Pointer.
//
// Metadata describing how extended dynamic range (EDR) values should be tone mapped.
func EDRMetadataFrom(ptr unsafe.Pointer) EDRMetadata {
	return EDRMetadata{objectivec.Object{objc.ID(ptr)}}
}



