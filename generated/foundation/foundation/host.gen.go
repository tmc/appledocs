// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Host] class.
var HostClass objc.Class

func init() {
	HostClass = objc.GetClass("NSHost")
}

type Host struct {
	objc.ID
}

func HostFrom(ptr unsafe.Pointer) Host {
	return Host{
		ID: objc.ID(ptr),
	}
}




