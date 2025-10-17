// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [XMLElement] class.
var XMLElementClass objc.Class

func init() {
	XMLElementClass = objc.GetClass("NSXMLElement")
}

type XMLElement struct {
	objc.ID
}

func XMLElementFrom(ptr unsafe.Pointer) XMLElement {
	return XMLElement{
		ID: objc.ID(ptr),
	}
}



