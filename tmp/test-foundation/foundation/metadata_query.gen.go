// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var MetadataQueryClass _MetadataQueryClass

func init() {
	MetadataQueryClass = _MetadataQueryClass{objc.GetClass("NSMetadataQuery")}
}

type _MetadataQueryClass struct {
	class objc.Class
}

type MetadataQuery struct {
	objc.ID
}

func MetadataQueryFrom(ptr unsafe.Pointer) MetadataQuery {
	return MetadataQuery{
		ID: objc.ID(ptr),
	}
}




