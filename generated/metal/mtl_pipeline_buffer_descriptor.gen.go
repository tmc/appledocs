// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLPipelineBufferDescriptor */


/* debug [class_header]: Header for MTLPipelineBufferDescriptor */
// The class instance for the [PipelineBufferDescriptor] class.
var (
	PipelineBufferDescriptorClass     _PipelineBufferDescriptorClass
	PipelineBufferDescriptorClassOnce sync.Once
)

func getPipelineBufferDescriptorClass() _PipelineBufferDescriptorClass {
	PipelineBufferDescriptorClassOnce.Do(func() {
		PipelineBufferDescriptorClass = _PipelineBufferDescriptorClass{objc.GetClass("MTLPipelineBufferDescriptor")}
	})
	return PipelineBufferDescriptorClass
}

type _PipelineBufferDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PipelineBufferDescriptor */
// An interface definition for the [PipelineBufferDescriptor] class.
type IPipelineBufferDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PipelineBufferDescriptor */
	// properties:
	Mutability() Mutability
	SetMutability(value Mutability)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PipelineBufferDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PipelineBufferDescriptor */
// Alloc allocates a new instance without initialization.
func (pc _PipelineBufferDescriptorClass) Alloc() PipelineBufferDescriptor {
	rv := objc.Send[PipelineBufferDescriptor](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PipelineBufferDescriptorClass) New() PipelineBufferDescriptor {
	rv := objc.Send[PipelineBufferDescriptor](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PipelineBufferDescriptor) Init() PipelineBufferDescriptor {
	rv := objc.Send[PipelineBufferDescriptor](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PipelineBufferDescriptor) Autorelease() PipelineBufferDescriptor {
	rv := objc.Send[PipelineBufferDescriptor](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPipelineBufferDescriptor creates a new PipelineBufferDescriptor instance.
func NewPipelineBufferDescriptor() PipelineBufferDescriptor {
	return getPipelineBufferDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PipelineBufferDescriptor */
// The mutability options for a buffer that a render or compute pipeline uses.
//
// Metal can perform additional optimizations if you guarantee that neither the CPU nor the GPU modify a buffer’s contents before starting a pass. Use immutable buffers as much as possible to take advantage of Metal optimizations. To declare that a buffer is immutable, set the property of their associated object to .


// The mutability options for a buffer that a render or compute pipeline uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineBufferDescriptor
type PipelineBufferDescriptor struct {
	objectivec.Object
}

// PipelineBufferDescriptorFrom constructs a [PipelineBufferDescriptor] from an unsafe.Pointer.
//
// The mutability options for a buffer that a render or compute pipeline uses.
func PipelineBufferDescriptorFrom(ptr unsafe.Pointer) PipelineBufferDescriptor {
	return PipelineBufferDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PipelineBufferDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PipelineBufferDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PipelineBufferDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PipelineBufferDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PipelineBufferDescriptor */

// A mutability option that determines whether you can update a buffer’s contents before related commands use the buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineBufferDescriptor/mutability
func (p_ PipelineBufferDescriptor) Mutability() Mutability {
	rv := objc.Send[Mutability](p_.ID, objc.Sel("mutability"))
	return rv
}/* debug [instance_properties/getter]: mutability */


// A mutability option that determines whether you can update a buffer’s contents before related commands use the buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPipelineBufferDescriptor/mutability
func (p_ PipelineBufferDescriptor) SetMutability(value Mutability) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMutability:"), value)
}/* debug [instance_properties/setter]: mutability */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLPipelineBufferDescriptor */



