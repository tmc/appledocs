// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var timeZoneClass _TimeZoneClass

func init() {
	timeZoneClass = _TimeZoneClass{objc.GetClass("NSTimeZone")}
}

type _TimeZoneClass struct {
	class objc.Class
}

type TimeZone struct {
	objc.ID
}

func TimeZoneFrom(ptr unsafe.Pointer) TimeZone {
	return TimeZone{
		ID: objc.ID(ptr),
	}
}




