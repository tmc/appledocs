// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Dimension] class.
var DimensionClass objc.Class

func init() {
	DimensionClass = objc.GetClass("NSDimension")
}

type Dimension struct {
	objc.ID
}

func DimensionFrom(ptr unsafe.Pointer) Dimension {
	return Dimension{
		ID: objc.ID(ptr),
	}
}




