// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var hostClass _HostClass

func init() {
	hostClass = _HostClass{objc.GetClass("NSHost")}
}

type _HostClass struct {
	class objc.Class
}

type Host struct {
	objc.ID
}

func HostFrom(ptr unsafe.Pointer) Host {
	return Host{
		ID: objc.ID(ptr),
	}
}




