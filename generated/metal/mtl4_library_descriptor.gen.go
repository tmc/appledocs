// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTL4LibraryDescriptor */


/* debug [class_header]: Header for MTL4LibraryDescriptor */
// The class instance for the [MTL4LibraryDescriptor] class.
var (
	MTL4LibraryDescriptorClass     _MTL4LibraryDescriptorClass
	MTL4LibraryDescriptorClassOnce sync.Once
)

func getMTL4LibraryDescriptorClass() _MTL4LibraryDescriptorClass {
	MTL4LibraryDescriptorClassOnce.Do(func() {
		MTL4LibraryDescriptorClass = _MTL4LibraryDescriptorClass{objc.GetClass("MTL4LibraryDescriptor")}
	})
	return MTL4LibraryDescriptorClass
}

type _MTL4LibraryDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTL4LibraryDescriptor */
// An interface definition for the [MTL4LibraryDescriptor] class.
type IMTL4LibraryDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTL4LibraryDescriptor */
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Options() IMTLCompileOptions
	SetOptions(value IMTLCompileOptions)
	Source() objc.IObject /* cross-framework: NSString */
	SetSource(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTL4LibraryDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTL4LibraryDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTL4LibraryDescriptorClass) Alloc() MTL4LibraryDescriptor {
	rv := objc.Send[MTL4LibraryDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTL4LibraryDescriptorClass) New() MTL4LibraryDescriptor {
	rv := objc.Send[MTL4LibraryDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4LibraryDescriptor) Init() MTL4LibraryDescriptor {
	rv := objc.Send[MTL4LibraryDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4LibraryDescriptor) Autorelease() MTL4LibraryDescriptor {
	rv := objc.Send[MTL4LibraryDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4LibraryDescriptor creates a new MTL4LibraryDescriptor instance.
func NewMTL4LibraryDescriptor() MTL4LibraryDescriptor {
	return getMTL4LibraryDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTL4LibraryDescriptor */
// Serves as the base descriptor for creating a Metal library.


// Serves as the base descriptor for creating a Metal library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryDescriptor
type MTL4LibraryDescriptor struct {
	objectivec.Object
}

// MTL4LibraryDescriptorFrom constructs a [MTL4LibraryDescriptor] from an unsafe.Pointer.
//
// Serves as the base descriptor for creating a Metal library.
func MTL4LibraryDescriptorFrom(ptr unsafe.Pointer) MTL4LibraryDescriptor {
	return MTL4LibraryDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTL4LibraryDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTL4LibraryDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTL4LibraryDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTL4LibraryDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTL4LibraryDescriptor */

// Assigns an optional name to the Metal library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryDescriptor/name
func (m_ MTL4LibraryDescriptor) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// Assigns an optional name to the Metal library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryDescriptor/name
func (m_ MTL4LibraryDescriptor) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// Provides compile-time options for the Metal library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryDescriptor/options
func (m_ MTL4LibraryDescriptor) Options() IMTLCompileOptions {
	rv := objc.Send[CompileOptions](m_.ID, objc.Sel("options"))
	return rv
}/* debug [instance_properties/getter]: options */


// Provides compile-time options for the Metal library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryDescriptor/options
func (m_ MTL4LibraryDescriptor) SetOptions(value IMTLCompileOptions) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptions:"), value)
}/* debug [instance_properties/setter]: options */


// Assigns an optional string containing the source code of the shader language program to compile into a Metal library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryDescriptor/source
func (m_ MTL4LibraryDescriptor) Source() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("source"))
	return rv
}/* debug [instance_properties/getter]: source */


// Assigns an optional string containing the source code of the shader language program to compile into a Metal library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4LibraryDescriptor/source
func (m_ MTL4LibraryDescriptor) SetSource(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSource:"), value)
}/* debug [instance_properties/setter]: source */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTL4LibraryDescriptor */



