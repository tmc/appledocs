// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTL4ArgumentTableDescriptor */


/* debug [class_header]: Header for MTL4ArgumentTableDescriptor */
// The class instance for the [MTL4ArgumentTableDescriptor] class.
var (
	MTL4ArgumentTableDescriptorClass     _MTL4ArgumentTableDescriptorClass
	MTL4ArgumentTableDescriptorClassOnce sync.Once
)

func getMTL4ArgumentTableDescriptorClass() _MTL4ArgumentTableDescriptorClass {
	MTL4ArgumentTableDescriptorClassOnce.Do(func() {
		MTL4ArgumentTableDescriptorClass = _MTL4ArgumentTableDescriptorClass{objc.GetClass("MTL4ArgumentTableDescriptor")}
	})
	return MTL4ArgumentTableDescriptorClass
}

type _MTL4ArgumentTableDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4ArgumentTableDescriptor */
// An interface definition for the [MTL4ArgumentTableDescriptor] class.
type IMTL4ArgumentTableDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTL4ArgumentTableDescriptor */
	// properties:
	InitializeBindings() bool
	SetInitializeBindings(value bool)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	MaxBufferBindCount() uint
	SetMaxBufferBindCount(value uint)
	MaxSamplerStateBindCount() uint
	SetMaxSamplerStateBindCount(value uint)
	MaxTextureBindCount() uint
	SetMaxTextureBindCount(value uint)
	SupportAttributeStrides() bool
	SetSupportAttributeStrides(value bool)
	MTL4CommandQueueErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4ArgumentTableDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4ArgumentTableDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4ArgumentTableDescriptorClass) Alloc() MTL4ArgumentTableDescriptor {
	rv := objc.Send[MTL4ArgumentTableDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4ArgumentTableDescriptorClass) New() MTL4ArgumentTableDescriptor {
	rv := objc.Send[MTL4ArgumentTableDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4ArgumentTableDescriptor) Init() MTL4ArgumentTableDescriptor {
	rv := objc.Send[MTL4ArgumentTableDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4ArgumentTableDescriptor) Autorelease() MTL4ArgumentTableDescriptor {
	rv := objc.Send[MTL4ArgumentTableDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4ArgumentTableDescriptor creates a new MTL4ArgumentTableDescriptor instance.
func NewMTL4ArgumentTableDescriptor() MTL4ArgumentTableDescriptor {
	return getMTL4ArgumentTableDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4ArgumentTableDescriptor */
// Groups parameters for the creation of a Metal argument table.
//
// Argument tables provide resource bindings to your Metal pipeline states.


// Groups parameters for the creation of a Metal argument table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor
type MTL4ArgumentTableDescriptor struct {
	objectivec.Object
}

// MTL4ArgumentTableDescriptorFrom constructs a [MTL4ArgumentTableDescriptor] from an unsafe.Pointer.
//
// Groups parameters for the creation of a Metal argument table.
func MTL4ArgumentTableDescriptorFrom(ptr unsafe.Pointer) MTL4ArgumentTableDescriptor {
	return MTL4ArgumentTableDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4ArgumentTableDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4ArgumentTableDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4ArgumentTableDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4ArgumentTableDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4ArgumentTableDescriptor */

// Configures whether Metal initializes the bindings to nil values upon creation of argument table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor/initializeBindings
func (m_ MTL4ArgumentTableDescriptor) InitializeBindings() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("initializeBindings"))
	return rv
}/* debug [instance_properties/getter]: initializeBindings */


// Configures whether Metal initializes the bindings to nil values upon creation of argument table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor/initializeBindings
func (m_ MTL4ArgumentTableDescriptor) SetInitializeBindings(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInitializeBindings:"), value)
}/* debug [instance_properties/setter]: initializeBindings */


// Assigns an optional label with the argument table for debug purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor/label
func (m_ MTL4ArgumentTableDescriptor) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// Assigns an optional label with the argument table for debug purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor/label
func (m_ MTL4ArgumentTableDescriptor) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// Determines the number of buffer-binding slots for the argument table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor/maxBufferBindCount
func (m_ MTL4ArgumentTableDescriptor) MaxBufferBindCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxBufferBindCount"))
	return rv
}/* debug [instance_properties/getter]: maxBufferBindCount */


// Determines the number of buffer-binding slots for the argument table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor/maxBufferBindCount
func (m_ MTL4ArgumentTableDescriptor) SetMaxBufferBindCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxBufferBindCount:"), value)
}/* debug [instance_properties/setter]: maxBufferBindCount */


// Determines the number of sampler state-binding slots for the argument table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor/maxSamplerStateBindCount
func (m_ MTL4ArgumentTableDescriptor) MaxSamplerStateBindCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxSamplerStateBindCount"))
	return rv
}/* debug [instance_properties/getter]: maxSamplerStateBindCount */


// Determines the number of sampler state-binding slots for the argument table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor/maxSamplerStateBindCount
func (m_ MTL4ArgumentTableDescriptor) SetMaxSamplerStateBindCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxSamplerStateBindCount:"), value)
}/* debug [instance_properties/setter]: maxSamplerStateBindCount */


// Determines the number of texture-binding slots for the argument table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor/maxTextureBindCount
func (m_ MTL4ArgumentTableDescriptor) MaxTextureBindCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maxTextureBindCount"))
	return rv
}/* debug [instance_properties/getter]: maxTextureBindCount */


// Determines the number of texture-binding slots for the argument table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor/maxTextureBindCount
func (m_ MTL4ArgumentTableDescriptor) SetMaxTextureBindCount(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTextureBindCount:"), value)
}/* debug [instance_properties/setter]: maxTextureBindCount */


// Controls whether Metal should reserve memory for attribute strides in the argument table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor/supportAttributeStrides
func (m_ MTL4ArgumentTableDescriptor) SupportAttributeStrides() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportAttributeStrides"))
	return rv
}/* debug [instance_properties/getter]: supportAttributeStrides */


// Controls whether Metal should reserve memory for attribute strides in the argument table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor/supportAttributeStrides
func (m_ MTL4ArgumentTableDescriptor) SetSupportAttributeStrides(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportAttributeStrides:"), value)
}/* debug [instance_properties/setter]: supportAttributeStrides */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4commandqueueerrordomain
func (m_ MTL4ArgumentTableDescriptor) MTL4CommandQueueErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MTL4CommandQueueErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: MTL4CommandQueueErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4ArgumentTableDescriptor */



