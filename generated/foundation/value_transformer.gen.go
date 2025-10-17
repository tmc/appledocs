// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ValueTransformer] class.
var ValueTransformerClass objc.Class

func init() {
	ValueTransformerClass = objc.GetClass("NSValueTransformer")
}

type ValueTransformer struct {
	objc.ID
}

func ValueTransformerFrom(ptr unsafe.Pointer) ValueTransformer {
	return ValueTransformer{
		ID: objc.ID(ptr),
	}
}



