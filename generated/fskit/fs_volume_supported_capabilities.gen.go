// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [FSVolumeSupportedCapabilities] class.
var (
	FSVolumeSupportedCapabilitiesClass     _FSVolumeSupportedCapabilitiesClass
	FSVolumeSupportedCapabilitiesClassOnce sync.Once
)

func getFSVolumeSupportedCapabilitiesClass() _FSVolumeSupportedCapabilitiesClass {
	FSVolumeSupportedCapabilitiesClassOnce.Do(func() {
		FSVolumeSupportedCapabilitiesClass = _FSVolumeSupportedCapabilitiesClass{objc.GetClass("FSVolumeSupportedCapabilities")}
	})
	return FSVolumeSupportedCapabilitiesClass
}

type _FSVolumeSupportedCapabilitiesClass struct {
	class objc.Class
}

// An interface definition for the [FSVolumeSupportedCapabilities] class.
type IFSVolumeSupportedCapabilities interface {
	objectivec.IObject
}

// A type that represents capabillities supported by a volume, such as hard and symbolic links, journaling, and large file sizes.
//
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSVolume/SupportedCapabilities
type FSVolumeSupportedCapabilities struct {
	objectivec.Object
}

// FSVolumeSupportedCapabilitiesFrom constructs a [FSVolumeSupportedCapabilities] from an unsafe.Pointer.
//
// A type that represents capabillities supported by a volume, such as hard and symbolic links, journaling, and large file sizes.
func FSVolumeSupportedCapabilitiesFrom(ptr unsafe.Pointer) FSVolumeSupportedCapabilities {
	return FSVolumeSupportedCapabilities{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FSVolumeSupportedCapabilitiesClass) Alloc() FSVolumeSupportedCapabilities {
	rv := objc.Send[FSVolumeSupportedCapabilities](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FSVolumeSupportedCapabilitiesClass) New() FSVolumeSupportedCapabilities {
	rv := objc.Send[FSVolumeSupportedCapabilities](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSVolumeSupportedCapabilities) Init() FSVolumeSupportedCapabilities {
	rv := objc.Send[FSVolumeSupportedCapabilities](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSVolumeSupportedCapabilities) Autorelease() FSVolumeSupportedCapabilities {
	rv := objc.Send[FSVolumeSupportedCapabilities](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSVolumeSupportedCapabilities creates a new FSVolumeSupportedCapabilities instance.
func NewFSVolumeSupportedCapabilities() FSVolumeSupportedCapabilities {
	return getFSVolumeSupportedCapabilitiesClass().New()
}




