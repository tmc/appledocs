// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var DimensionClass _DimensionClass

func init() {
	DimensionClass = _DimensionClass{objc.GetClass("NSDimension")}
}

type _DimensionClass struct {
	class objc.Class
}

type Dimension struct {
	objc.ID
}

func DimensionFrom(ptr unsafe.Pointer) Dimension {
	return Dimension{
		ID: objc.ID(ptr),
	}
}




