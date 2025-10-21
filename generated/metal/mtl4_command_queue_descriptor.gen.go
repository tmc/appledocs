// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTL4CommandQueueDescriptor] class.
var (
	MTL4CommandQueueDescriptorClass     _MTL4CommandQueueDescriptorClass
	MTL4CommandQueueDescriptorClassOnce sync.Once
)

func getMTL4CommandQueueDescriptorClass() _MTL4CommandQueueDescriptorClass {
	MTL4CommandQueueDescriptorClassOnce.Do(func() {
		MTL4CommandQueueDescriptorClass = _MTL4CommandQueueDescriptorClass{objc.GetClass("MTL4CommandQueueDescriptor")}
	})
	return MTL4CommandQueueDescriptorClass
}

type _MTL4CommandQueueDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4CommandQueueDescriptor] class.
type IMTL4CommandQueueDescriptor interface {
	objectivec.IObject
}

// Groups together parameters for the creation of a new command queue.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandQueueDescriptor
type MTL4CommandQueueDescriptor struct {
	objectivec.Object
}

// MTL4CommandQueueDescriptorFrom constructs a [MTL4CommandQueueDescriptor] from an unsafe.Pointer.
//
// Groups together parameters for the creation of a new command queue.
func MTL4CommandQueueDescriptorFrom(ptr unsafe.Pointer) MTL4CommandQueueDescriptor {
	return MTL4CommandQueueDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4CommandQueueDescriptorClass) Alloc() MTL4CommandQueueDescriptor {
	rv := objc.Send[MTL4CommandQueueDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4CommandQueueDescriptorClass) New() MTL4CommandQueueDescriptor {
	rv := objc.Send[MTL4CommandQueueDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4CommandQueueDescriptor) Init() MTL4CommandQueueDescriptor {
	rv := objc.Send[MTL4CommandQueueDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4CommandQueueDescriptor) Autorelease() MTL4CommandQueueDescriptor {
	rv := objc.Send[MTL4CommandQueueDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4CommandQueueDescriptor creates a new MTL4CommandQueueDescriptor instance.
func NewMTL4CommandQueueDescriptor() MTL4CommandQueueDescriptor {
	return getMTL4CommandQueueDescriptorClass().New()
}


// Assigns a dispatch queue to which Metal submits feedback notification blocks.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4commandqueuedescriptor/feedbackqueue
func (m_ MTL4CommandQueueDescriptor) FeedbackQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("feedbackQueue"))
	return rv
}


// SetFeedbackQueue sets the value of the feedbackQueue property.
// Assigns a dispatch queue to which Metal submits feedback notification blocks.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4commandqueuedescriptor/feedbackqueue
func (m_ MTL4CommandQueueDescriptor) SetFeedbackQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFeedbackQueue:"), value)
}

// Assigns an optional label to the command queue instance for debugging purposes.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4commandqueuedescriptor/label
func (m_ MTL4CommandQueueDescriptor) Label() string {
	rv := objc.Send[string](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// Assigns an optional label to the command queue instance for debugging purposes.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4commandqueuedescriptor/label
func (m_ MTL4CommandQueueDescriptor) SetLabel(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4commandqueueerrordomain
func (m_ MTL4CommandQueueDescriptor) MTL4CommandQueueErrorDomain() string {
	rv := objc.Send[string](m_.ID, objc.Sel("MTL4CommandQueueErrorDomain"))
	return rv
}



