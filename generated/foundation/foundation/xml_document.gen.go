// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [XMLDocument] class.
var XMLDocumentClass objc.Class

func init() {
	XMLDocumentClass = objc.GetClass("NSXMLDocument")
}

type XMLDocument struct {
	objc.ID
}

func XMLDocumentFrom(ptr unsafe.Pointer) XMLDocument {
	return XMLDocument{
		ID: objc.ID(ptr),
	}
}




