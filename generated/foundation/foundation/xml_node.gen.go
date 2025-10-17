// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [XMLNode] class.
var XMLNodeClass objc.Class

func init() {
	XMLNodeClass = objc.GetClass("NSXMLNode")
}

type XMLNode struct {
	objc.ID
}

func XMLNodeFrom(ptr unsafe.Pointer) XMLNode {
	return XMLNode{
		ID: objc.ID(ptr),
	}
}




