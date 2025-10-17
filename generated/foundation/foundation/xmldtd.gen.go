// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [XMLDTD] class.
var XMLDTDClass objc.Class

func init() {
	XMLDTDClass = objc.GetClass("NSXMLDTD")
}

type XMLDTD struct {
	objc.ID
}

func XMLDTDFrom(ptr unsafe.Pointer) XMLDTD {
	return XMLDTD{
		ID: objc.ID(ptr),
	}
}




