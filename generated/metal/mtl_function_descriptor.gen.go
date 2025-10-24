// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLFunctionDescriptor */


/* debug [class_header]: Header for MTLFunctionDescriptor */
// The class instance for the [FunctionDescriptor] class.
var (
	FunctionDescriptorClass     _FunctionDescriptorClass
	FunctionDescriptorClassOnce sync.Once
)

func getFunctionDescriptorClass() _FunctionDescriptorClass {
	FunctionDescriptorClassOnce.Do(func() {
		FunctionDescriptorClass = _FunctionDescriptorClass{objc.GetClass("MTLFunctionDescriptor")}
	})
	return FunctionDescriptorClass
}

type _FunctionDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FunctionDescriptor */
// An interface definition for the [FunctionDescriptor] class.
type IFunctionDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FunctionDescriptor */
	// properties:
	BinaryArchives() []objc.ID
	SetBinaryArchives(value []objc.ID)
	ConstantValues() IMTLFunctionConstantValues
	SetConstantValues(value IMTLFunctionConstantValues)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Options() FunctionOptions
	SetOptions(value FunctionOptions)
	SpecializedName() objc.IObject /* cross-framework: NSString */
	SetSpecializedName(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FunctionDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FunctionDescriptor */
// Alloc allocates a new instance without initialization.
func (fc _FunctionDescriptorClass) Alloc() FunctionDescriptor {
	rv := objc.Send[FunctionDescriptor](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FunctionDescriptorClass) New() FunctionDescriptor {
	rv := objc.Send[FunctionDescriptor](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FunctionDescriptor) Init() FunctionDescriptor {
	rv := objc.Send[FunctionDescriptor](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FunctionDescriptor) Autorelease() FunctionDescriptor {
	rv := objc.Send[FunctionDescriptor](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFunctionDescriptor creates a new FunctionDescriptor instance.
func NewFunctionDescriptor() FunctionDescriptor {
	return getFunctionDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FunctionDescriptor */
// A description of a function object to create.


// A description of a function object to create.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor
type FunctionDescriptor struct {
	objectivec.Object
}

// FunctionDescriptorFrom constructs a [FunctionDescriptor] from an unsafe.Pointer.
//
// A description of a function object to create.
func FunctionDescriptorFrom(ptr unsafe.Pointer) FunctionDescriptor {
	return FunctionDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FunctionDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FunctionDescriptor */

// Creates a default function descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/functionDescriptor
func (fc _FunctionDescriptorClass) FunctionDescriptor() IFunctionDescriptor {
	rv := objc.Send[FunctionDescriptor](objc.ID(fc.class), objc.Sel("functionDescriptor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FunctionDescriptor) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FunctionDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FunctionDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FunctionDescriptor */

// The binary archives to search for a previously-compiled version of this function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/binaryArchives
func (f_ FunctionDescriptor) BinaryArchives() []objc.ID {
	rv := objc.Send[[]objc.ID](f_.ID, objc.Sel("binaryArchives"))
	return rv
}/* debug [instance_properties/getter]: binaryArchives */


// The binary archives to search for a previously-compiled version of this function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/binaryArchives
func (f_ FunctionDescriptor) SetBinaryArchives(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](f_.ID, objc.Sel("setBinaryArchives:"), nsArray)
}/* debug [instance_properties/setter]: binaryArchives */


// The set of constant values assigned to the function constants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/constantValues
func (f_ FunctionDescriptor) ConstantValues() IMTLFunctionConstantValues {
	rv := objc.Send[FunctionConstantValues](f_.ID, objc.Sel("constantValues"))
	return rv
}/* debug [instance_properties/getter]: constantValues */


// The set of constant values assigned to the function constants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/constantValues
func (f_ FunctionDescriptor) SetConstantValues(value IMTLFunctionConstantValues) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setConstantValues:"), value)
}/* debug [instance_properties/setter]: constantValues */


// The name of the function to fetch from the library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/name
func (f_ FunctionDescriptor) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The name of the function to fetch from the library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/name
func (f_ FunctionDescriptor) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// Flags specifying how Metal should create the new function object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/options
func (f_ FunctionDescriptor) Options() FunctionOptions {
	rv := objc.Send[FunctionOptions](f_.ID, objc.Sel("options"))
	return rv
}/* debug [instance_properties/getter]: options */


// Flags specifying how Metal should create the new function object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/options
func (f_ FunctionDescriptor) SetOptions(value FunctionOptions) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setOptions:"), value)
}/* debug [instance_properties/setter]: options */


// A new name for the created function object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/specializedName
func (f_ FunctionDescriptor) SpecializedName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("specializedName"))
	return rv
}/* debug [instance_properties/getter]: specializedName */


// A new name for the created function object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLFunctionDescriptor/specializedName
func (f_ FunctionDescriptor) SetSpecializedName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setSpecializedName:"), value)
}/* debug [instance_properties/setter]: specializedName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLFunctionDescriptor */



