// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Host] class.
var hostClass = _HostClass{objc.GetClass("NSHost")}

type _HostClass struct {
	class objc.Class
}

// An interface definition for the [Host] class.
type IHost interface {
	objectivec.IObject
}

// A representation of an individual host on the network. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Host

type Host struct {
	objectivec.Object
}

// HostFrom constructs a [Host] from an unsafe.Pointer.
//
// A representation of an individual host on the network.
func HostFrom(ptr unsafe.Pointer) Host {
	return Host{objectivec.Object{objc.ID(ptr)}}
}



