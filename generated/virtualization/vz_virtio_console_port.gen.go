// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [VZVirtioConsolePort] class.
var (
	VZVirtioConsolePortClass     _VZVirtioConsolePortClass
	VZVirtioConsolePortClassOnce sync.Once
)

func getVZVirtioConsolePortClass() _VZVirtioConsolePortClass {
	VZVirtioConsolePortClassOnce.Do(func() {
		VZVirtioConsolePortClass = _VZVirtioConsolePortClass{objc.GetClass("VZVirtioConsolePort")}
	})
	return VZVirtioConsolePortClass
}

type _VZVirtioConsolePortClass struct {
	class objc.Class
}

// An interface definition for the [VZVirtioConsolePort] class.
type IVZVirtioConsolePort interface {
	objectivec.IObject
}

// A class that represents a Virtio console port in a VM.
//
// Don’t instantiate a directly. You retrieve this object from the property.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePort
type VZVirtioConsolePort struct {
	objectivec.Object
}

// VZVirtioConsolePortFrom constructs a [VZVirtioConsolePort] from an unsafe.Pointer.
//
// A class that represents a Virtio console port in a VM.
func VZVirtioConsolePortFrom(ptr unsafe.Pointer) VZVirtioConsolePort {
	return VZVirtioConsolePort{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _VZVirtioConsolePortClass) Alloc() VZVirtioConsolePort {
	rv := objc.Send[VZVirtioConsolePort](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _VZVirtioConsolePortClass) New() VZVirtioConsolePort {
	rv := objc.Send[VZVirtioConsolePort](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VZVirtioConsolePort) Init() VZVirtioConsolePort {
	rv := objc.Send[VZVirtioConsolePort](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VZVirtioConsolePort) Autorelease() VZVirtioConsolePort {
	rv := objc.Send[VZVirtioConsolePort](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsolePort creates a new VZVirtioConsolePort instance.
func NewVZVirtioConsolePort() VZVirtioConsolePort {
	return getVZVirtioConsolePortClass().New()
}


// An array of serial port attachments.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePort/attachment
func (v_ VZVirtioConsolePort) Attachment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("attachment"))
	return rv
}


// SetAttachment sets the value of the attachment property.
// An array of serial port attachments.

//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePort/attachment
func (v_ VZVirtioConsolePort) SetAttachment(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAttachment:"), value)
}
// The name of the port.
//
// [Full Topic]: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePort/name
func (v_ VZVirtioConsolePort) Name() string {
	rv := objc.Send[string](v_.ID, objc.Sel("name"))
	return rv
}



