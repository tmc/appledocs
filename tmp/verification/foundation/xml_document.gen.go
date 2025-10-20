// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var xMLDocumentClass _XMLDocumentClass

func init() {
	xMLDocumentClass = _XMLDocumentClass{objc.GetClass("NSXMLDocument")}
}

type _XMLDocumentClass struct {
	class objc.Class
}

type XMLDocument struct {
	objc.ID
}

func XMLDocumentFrom(ptr unsafe.Pointer) XMLDocument {
	return XMLDocument{
		ID: objc.ID(ptr),
	}
}




