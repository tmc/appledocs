// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLBlitPassDescriptor */


/* debug [class_header]: Header for MTLBlitPassDescriptor */
// The class instance for the [BlitPassDescriptor] class.
var (
	BlitPassDescriptorClass     _BlitPassDescriptorClass
	BlitPassDescriptorClassOnce sync.Once
)

func getBlitPassDescriptorClass() _BlitPassDescriptorClass {
	BlitPassDescriptorClassOnce.Do(func() {
		BlitPassDescriptorClass = _BlitPassDescriptorClass{objc.GetClass("MTLBlitPassDescriptor")}
	})
	return BlitPassDescriptorClass
}

type _BlitPassDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BlitPassDescriptor */
// An interface definition for the [BlitPassDescriptor] class.
type IBlitPassDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BlitPassDescriptor */
	// properties:
	SampleBufferAttachments() IMTLBlitPassSampleBufferAttachmentDescriptorArray
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BlitPassDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BlitPassDescriptor */
// Alloc allocates a new instance without initialization.
func (bc _BlitPassDescriptorClass) Alloc() BlitPassDescriptor {
	rv := objc.Send[BlitPassDescriptor](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BlitPassDescriptorClass) New() BlitPassDescriptor {
	rv := objc.Send[BlitPassDescriptor](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BlitPassDescriptor) Init() BlitPassDescriptor {
	rv := objc.Send[BlitPassDescriptor](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BlitPassDescriptor) Autorelease() BlitPassDescriptor {
	rv := objc.Send[BlitPassDescriptor](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBlitPassDescriptor creates a new BlitPassDescriptor instance.
func NewBlitPassDescriptor() BlitPassDescriptor {
	return getBlitPassDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BlitPassDescriptor */
// A configuration you create to customize a blit command encoder, which affects the runtime behavior of the blit pass you encode with it.
//
// You can customize an encoder for a blit pass by creating and configuring an instance and passing it to .


// A configuration you create to customize a blit command encoder, which affects the runtime behavior of the blit pass you encode with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitPassDescriptor
type BlitPassDescriptor struct {
	objectivec.Object
}

// BlitPassDescriptorFrom constructs a [BlitPassDescriptor] from an unsafe.Pointer.
//
// A configuration you create to customize a blit command encoder, which affects the runtime behavior of the blit pass you encode with it.
func BlitPassDescriptorFrom(ptr unsafe.Pointer) BlitPassDescriptor {
	return BlitPassDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BlitPassDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BlitPassDescriptor */

// Creates a new blit pass descriptor with a default configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitPassDescriptor/blitPassDescriptor
func (bc _BlitPassDescriptorClass) BlitPassDescriptor() IBlitPassDescriptor {
	rv := objc.Send[BlitPassDescriptor](objc.ID(bc.class), objc.Sel("blitPassDescriptor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=BlitPassDescriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BlitPassDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BlitPassDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BlitPassDescriptor */

// An array of counter sample buffer attachments that you configure for a blit pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLBlitPassDescriptor/sampleBufferAttachments
func (b_ BlitPassDescriptor) SampleBufferAttachments() IMTLBlitPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[BlitPassSampleBufferAttachmentDescriptorArray](b_.ID, objc.Sel("sampleBufferAttachments"))
	return rv
}/* debug [instance_properties/getter]: sampleBufferAttachments */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLBlitPassDescriptor */



