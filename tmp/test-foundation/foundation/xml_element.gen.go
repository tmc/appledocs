// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var XMLElementClass _XMLElementClass

func init() {
	XMLElementClass = _XMLElementClass{objc.GetClass("NSXMLElement")}
}

type _XMLElementClass struct {
	class objc.Class
}

type XMLElement struct {
	objc.ID
}

func XMLElementFrom(ptr unsafe.Pointer) XMLElement {
	return XMLElement{
		ID: objc.ID(ptr),
	}
}




