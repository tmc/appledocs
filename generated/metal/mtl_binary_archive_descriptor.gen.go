// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BinaryArchiveDescriptor] class.
var (
	BinaryArchiveDescriptorClass     _BinaryArchiveDescriptorClass
	BinaryArchiveDescriptorClassOnce sync.Once
)

func getBinaryArchiveDescriptorClass() _BinaryArchiveDescriptorClass {
	BinaryArchiveDescriptorClassOnce.Do(func() {
		BinaryArchiveDescriptorClass = _BinaryArchiveDescriptorClass{objc.GetClass("MTLBinaryArchiveDescriptor")}
	})
	return BinaryArchiveDescriptorClass
}

type _BinaryArchiveDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [BinaryArchiveDescriptor] class.
type IBinaryArchiveDescriptor interface {
	objectivec.IObject
}

// A description of a binary shader archive that you want to create.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBinaryArchiveDescriptor
type BinaryArchiveDescriptor struct {
	objectivec.Object
}

// BinaryArchiveDescriptorFrom constructs a [BinaryArchiveDescriptor] from an unsafe.Pointer.
//
// A description of a binary shader archive that you want to create.
func BinaryArchiveDescriptorFrom(ptr unsafe.Pointer) BinaryArchiveDescriptor {
	return BinaryArchiveDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BinaryArchiveDescriptorClass) Alloc() BinaryArchiveDescriptor {
	rv := objc.Send[BinaryArchiveDescriptor](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BinaryArchiveDescriptorClass) New() BinaryArchiveDescriptor {
	rv := objc.Send[BinaryArchiveDescriptor](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BinaryArchiveDescriptor) Init() BinaryArchiveDescriptor {
	rv := objc.Send[BinaryArchiveDescriptor](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BinaryArchiveDescriptor) Autorelease() BinaryArchiveDescriptor {
	rv := objc.Send[BinaryArchiveDescriptor](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBinaryArchiveDescriptor creates a new BinaryArchiveDescriptor instance.
func NewBinaryArchiveDescriptor() BinaryArchiveDescriptor {
	return getBinaryArchiveDescriptorClass().New()
}




