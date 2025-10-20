// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var decimalNumberClass _DecimalNumberClass

func init() {
	decimalNumberClass = _DecimalNumberClass{objc.GetClass("NSDecimalNumber")}
}

type _DecimalNumberClass struct {
	class objc.Class
}

type DecimalNumber struct {
	objc.ID
}

func DecimalNumberFrom(ptr unsafe.Pointer) DecimalNumber {
	return DecimalNumber{
		ID: objc.ID(ptr),
	}
}




