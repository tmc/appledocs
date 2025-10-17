// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MetadataQuery] class.
var MetadataQueryClass objc.Class

func init() {
	MetadataQueryClass = objc.GetClass("NSMetadataQuery")
}

type MetadataQuery struct {
	objc.ID
}

func MetadataQueryFrom(ptr unsafe.Pointer) MetadataQuery {
	return MetadataQuery{
		ID: objc.ID(ptr),
	}
}




