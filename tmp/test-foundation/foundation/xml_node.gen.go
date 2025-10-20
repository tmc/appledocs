// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var XMLNodeClass _XMLNodeClass

func init() {
	XMLNodeClass = _XMLNodeClass{objc.GetClass("NSXMLNode")}
}

type _XMLNodeClass struct {
	class objc.Class
}

type XMLNode struct {
	objc.ID
}

func XMLNodeFrom(ptr unsafe.Pointer) XMLNode {
	return XMLNode{
		ID: objc.ID(ptr),
	}
}




