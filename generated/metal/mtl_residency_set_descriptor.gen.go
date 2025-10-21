// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ResidencySetDescriptor] class.
var (
	ResidencySetDescriptorClass     _ResidencySetDescriptorClass
	ResidencySetDescriptorClassOnce sync.Once
)

func getResidencySetDescriptorClass() _ResidencySetDescriptorClass {
	ResidencySetDescriptorClassOnce.Do(func() {
		ResidencySetDescriptorClass = _ResidencySetDescriptorClass{objc.GetClass("MTLResidencySetDescriptor")}
	})
	return ResidencySetDescriptorClass
}

type _ResidencySetDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [ResidencySetDescriptor] class.
type IResidencySetDescriptor interface {
	objectivec.IObject
}

// A configuration that customizes the behavior for a residency set.
//
// Make an by creating and configuring an instance and pass it to the method of an instance. See for more information.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResidencySetDescriptor
type ResidencySetDescriptor struct {
	objectivec.Object
}

// ResidencySetDescriptorFrom constructs a [ResidencySetDescriptor] from an unsafe.Pointer.
//
// A configuration that customizes the behavior for a residency set.
func ResidencySetDescriptorFrom(ptr unsafe.Pointer) ResidencySetDescriptor {
	return ResidencySetDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _ResidencySetDescriptorClass) Alloc() ResidencySetDescriptor {
	rv := objc.Send[ResidencySetDescriptor](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ResidencySetDescriptorClass) New() ResidencySetDescriptor {
	rv := objc.Send[ResidencySetDescriptor](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ResidencySetDescriptor) Init() ResidencySetDescriptor {
	rv := objc.Send[ResidencySetDescriptor](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ResidencySetDescriptor) Autorelease() ResidencySetDescriptor {
	rv := objc.Send[ResidencySetDescriptor](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewResidencySetDescriptor creates a new ResidencySetDescriptor instance.
func NewResidencySetDescriptor() ResidencySetDescriptor {
	return getResidencySetDescriptorClass().New()
}


// The number of allocations a new residency set can store without reallocating memory.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresidencysetdescriptor/initialcapacity
func (r_ ResidencySetDescriptor) InitialCapacity() int {
	rv := objc.Send[int](r_.ID, objc.Sel("initialCapacity"))
	return rv
}


// SetInitialCapacity sets the value of the initialCapacity property.
// The number of allocations a new residency set can store without reallocating memory.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlresidencysetdescriptor/initialcapacity
func (r_ ResidencySetDescriptor) SetInitialCapacity(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setInitialCapacity:"), value)
}

// An optional name that can help you identify a residency set you create with the descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResidencySetDescriptor/label
func (r_ ResidencySetDescriptor) Label() string {
	rv := objc.Send[string](r_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// An optional name that can help you identify a residency set you create with the descriptor.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResidencySetDescriptor/label
func (r_ ResidencySetDescriptor) SetLabel(value string) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setLabel:"), objc.String(value))
}



