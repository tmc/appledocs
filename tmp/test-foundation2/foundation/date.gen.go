// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var dateClass _DateClass

func init() {
	dateClass = _DateClass{objc.GetClass("NSDate")}
}

type _DateClass struct {
	class objc.Class
}

type Date struct {
	objc.ID
}

func DateFrom(ptr unsafe.Pointer) Date {
	return Date{
		ID: objc.ID(ptr),
	}
}




