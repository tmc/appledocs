// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [XMLDTD] class.
var XMLDTDClass = _XMLDTDClass{objc.GetClass("NSXMLDTD")}

type _XMLDTDClass struct {
	class objc.Class
}

type XMLDTD struct {
	objc.ID
}

func XMLDTDFrom(ptr unsafe.Pointer) XMLDTD {
	return XMLDTD{
		ID: objc.ID(ptr),
	}
}




