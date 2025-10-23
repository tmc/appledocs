// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FSExtentPacker] class.
var (
	FSExtentPackerClass     _FSExtentPackerClass
	FSExtentPackerClassOnce sync.Once
)

func getFSExtentPackerClass() _FSExtentPackerClass {
	FSExtentPackerClassOnce.Do(func() {
		FSExtentPackerClass = _FSExtentPackerClass{objc.GetClass("FSExtentPacker")}
	})
	return FSExtentPackerClass
}

type _FSExtentPackerClass struct {
	class objc.Class
}

// An interface definition for the [FSExtentPacker] class.
type IFSExtentPacker interface {
	objectivec.IObject
	PackExtentWithResourceTypeLogicalOffsetPhysicalOffsetLength(resource IFSBlockDeviceResource, type_ FSExtentType, logicalOffset unsafe.Pointer, physicalOffset unsafe.Pointer, length uintptr) bool
}

// A type that directs the kernel to map space on disk to a specific file managed by this file system.
//
// provide the kernel the logical-to-physical mapping of a given file. An extent describes a physical offset on disk, and a length and a logical offset within the file. Rather than working with extents directly, you use this type’s methods to provide or “pack” extent information, which FSKit then passes to the kernel.


// A type that directs the kernel to map space on disk to a specific file managed by this file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSExtentPacker
type FSExtentPacker struct {
	objectivec.Object
}

// FSExtentPackerFrom constructs a [FSExtentPacker] from an unsafe.Pointer.
//
// A type that directs the kernel to map space on disk to a specific file managed by this file system.
func FSExtentPackerFrom(ptr unsafe.Pointer) FSExtentPacker {
	return FSExtentPacker{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FSExtentPackerClass) Alloc() FSExtentPacker {
	rv := objc.Send[FSExtentPacker](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FSExtentPackerClass) New() FSExtentPacker {
	rv := objc.Send[FSExtentPacker](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSExtentPacker) Init() FSExtentPacker {
	rv := objc.Send[FSExtentPacker](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSExtentPacker) Autorelease() FSExtentPacker {
	rv := objc.Send[FSExtentPacker](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSExtentPacker creates a new FSExtentPacker instance.
func NewFSExtentPacker() FSExtentPacker {
	return getFSExtentPackerClass().New()
}



// Packs a single extent to send to the kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSExtentPacker/packExtent(resource:type:logicalOffset:physicalOffset:length:)
func (f_ FSExtentPacker) PackExtentWithResourceTypeLogicalOffsetPhysicalOffsetLength(resource IFSBlockDeviceResource, type_ FSExtentType, logicalOffset unsafe.Pointer, physicalOffset unsafe.Pointer, length uintptr) bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("packExtentWithResource:type:logicalOffset:physicalOffset:length:"), resource, type_, logicalOffset, physicalOffset, length)
	return rv
}



