// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TimeZone] class.
var TimeZoneClass objc.Class

func init() {
	TimeZoneClass = objc.GetClass("NSTimeZone")
}

type TimeZone struct {
	objc.ID
}

func TimeZoneFrom(ptr unsafe.Pointer) TimeZone {
	return TimeZone{
		ID: objc.ID(ptr),
	}
}




