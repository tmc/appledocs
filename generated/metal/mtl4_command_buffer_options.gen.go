// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTL4CommandBufferOptions */


/* debug [class_header]: Header for MTL4CommandBufferOptions */
// The class instance for the [MTL4CommandBufferOptions] class.
var (
	MTL4CommandBufferOptionsClass     _MTL4CommandBufferOptionsClass
	MTL4CommandBufferOptionsClassOnce sync.Once
)

func getMTL4CommandBufferOptionsClass() _MTL4CommandBufferOptionsClass {
	MTL4CommandBufferOptionsClassOnce.Do(func() {
		MTL4CommandBufferOptionsClass = _MTL4CommandBufferOptionsClass{objc.GetClass("MTL4CommandBufferOptions")}
	})
	return MTL4CommandBufferOptionsClass
}

type _MTL4CommandBufferOptionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4CommandBufferOptions */
// An interface definition for the [MTL4CommandBufferOptions] class.
type IMTL4CommandBufferOptions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTL4CommandBufferOptions */
	// properties:
	LogState() unsafe.Pointer
	SetLogState(value unsafe.Pointer)
	MTL4CommandQueueErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4CommandBufferOptions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4CommandBufferOptions */
// Alloc allocates a new instance without initialization.
func (mc _MTL4CommandBufferOptionsClass) Alloc() MTL4CommandBufferOptions {
	rv := objc.Send[MTL4CommandBufferOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4CommandBufferOptionsClass) New() MTL4CommandBufferOptions {
	rv := objc.Send[MTL4CommandBufferOptions](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4CommandBufferOptions) Init() MTL4CommandBufferOptions {
	rv := objc.Send[MTL4CommandBufferOptions](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4CommandBufferOptions) Autorelease() MTL4CommandBufferOptions {
	rv := objc.Send[MTL4CommandBufferOptions](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4CommandBufferOptions creates a new MTL4CommandBufferOptions instance.
func NewMTL4CommandBufferOptions() MTL4CommandBufferOptions {
	return getMTL4CommandBufferOptionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4CommandBufferOptions */
// Options to configure a command buffer before encoding work into it.


// Options to configure a command buffer before encoding work into it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandBufferOptions
type MTL4CommandBufferOptions struct {
	objectivec.Object
}

// MTL4CommandBufferOptionsFrom constructs a [MTL4CommandBufferOptions] from an unsafe.Pointer.
//
// Options to configure a command buffer before encoding work into it.
func MTL4CommandBufferOptionsFrom(ptr unsafe.Pointer) MTL4CommandBufferOptions {
	return MTL4CommandBufferOptions{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4CommandBufferOptions *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4CommandBufferOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4CommandBufferOptions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4CommandBufferOptions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4CommandBufferOptions */

// Contains information related to shader logging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandBufferOptions/logState
func (m_ MTL4CommandBufferOptions) LogState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("logState"))
	return rv
}/* debug [instance_properties/getter]: logState */


// Contains information related to shader logging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CommandBufferOptions/logState
func (m_ MTL4CommandBufferOptions) SetLogState(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLogState:"), value)
}/* debug [instance_properties/setter]: logState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4commandqueueerrordomain
func (m_ MTL4CommandBufferOptions) MTL4CommandQueueErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MTL4CommandQueueErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: MTL4CommandQueueErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4CommandBufferOptions */



