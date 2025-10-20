// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var xMLNodeClass _XMLNodeClass

func init() {
	xMLNodeClass = _XMLNodeClass{objc.GetClass("NSXMLNode")}
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




