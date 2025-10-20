// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var XMLDocumentClass _XMLDocumentClass

func init() {
	XMLDocumentClass = _XMLDocumentClass{objc.GetClass("NSXMLDocument")}
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




