// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var valueTransformerClass _ValueTransformerClass

func init() {
	valueTransformerClass = _ValueTransformerClass{objc.GetClass("NSValueTransformer")}
}

type _ValueTransformerClass struct {
	class objc.Class
}

type ValueTransformer struct {
	objc.ID
}

func ValueTransformerFrom(ptr unsafe.Pointer) ValueTransformer {
	return ValueTransformer{
		ID: objc.ID(ptr),
	}
}




