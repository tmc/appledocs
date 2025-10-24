// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTL4LibraryFunctionDescriptor */


/* debug [class_header]: Header for MTL4LibraryFunctionDescriptor */
// The class instance for the [MTL4LibraryFunctionDescriptor] class.
var (
	MTL4LibraryFunctionDescriptorClass     _MTL4LibraryFunctionDescriptorClass
	MTL4LibraryFunctionDescriptorClassOnce sync.Once
)

func getMTL4LibraryFunctionDescriptorClass() _MTL4LibraryFunctionDescriptorClass {
	MTL4LibraryFunctionDescriptorClassOnce.Do(func() {
		MTL4LibraryFunctionDescriptorClass = _MTL4LibraryFunctionDescriptorClass{objc.GetClass("MTL4LibraryFunctionDescriptor")}
	})
	return MTL4LibraryFunctionDescriptorClass
}

type _MTL4LibraryFunctionDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4LibraryFunctionDescriptor */
// An interface definition for the [MTL4LibraryFunctionDescriptor] class.
type IMTL4LibraryFunctionDescriptor interface {
	IMTL4FunctionDescriptor
	
/* debug [class_interface_properties]: Properties for MTL4LibraryFunctionDescriptor */
	// properties:
	Library() unsafe.Pointer
	SetLibrary(value unsafe.Pointer)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4LibraryFunctionDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4LibraryFunctionDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4LibraryFunctionDescriptorClass) Alloc() MTL4LibraryFunctionDescriptor {
	rv := objc.Send[MTL4LibraryFunctionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4LibraryFunctionDescriptorClass) New() MTL4LibraryFunctionDescriptor {
	rv := objc.Send[MTL4LibraryFunctionDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4LibraryFunctionDescriptor) Init() MTL4LibraryFunctionDescriptor {
	rv := objc.Send[MTL4LibraryFunctionDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4LibraryFunctionDescriptor) Autorelease() MTL4LibraryFunctionDescriptor {
	rv := objc.Send[MTL4LibraryFunctionDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4LibraryFunctionDescriptor creates a new MTL4LibraryFunctionDescriptor instance.
func NewMTL4LibraryFunctionDescriptor() MTL4LibraryFunctionDescriptor {
	return getMTL4LibraryFunctionDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4LibraryFunctionDescriptor */
// Describes a shader function from a Metal library.


// Describes a shader function from a Metal library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryFunctionDescriptor
type MTL4LibraryFunctionDescriptor struct {
	MTL4FunctionDescriptor
}

// MTL4LibraryFunctionDescriptorFrom constructs a [MTL4LibraryFunctionDescriptor] from an unsafe.Pointer.
//
// Describes a shader function from a Metal library.
func MTL4LibraryFunctionDescriptorFrom(ptr unsafe.Pointer) MTL4LibraryFunctionDescriptor {
	return MTL4LibraryFunctionDescriptor{
		MTL4FunctionDescriptor: MTL4FunctionDescriptorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4LibraryFunctionDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4LibraryFunctionDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4LibraryFunctionDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4LibraryFunctionDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4LibraryFunctionDescriptor */

// Returns a reference to the library containing the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryFunctionDescriptor/library
func (m_ MTL4LibraryFunctionDescriptor) Library() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("library"))
	return rv
}/* debug [instance_properties/getter]: library */


// Returns a reference to the library containing the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryFunctionDescriptor/library
func (m_ MTL4LibraryFunctionDescriptor) SetLibrary(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLibrary:"), value)
}/* debug [instance_properties/setter]: library */


// Assigns a name to the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryFunctionDescriptor/name
func (m_ MTL4LibraryFunctionDescriptor) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// Assigns a name to the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryFunctionDescriptor/name
func (m_ MTL4LibraryFunctionDescriptor) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4LibraryFunctionDescriptor */



