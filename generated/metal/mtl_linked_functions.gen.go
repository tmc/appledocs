// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLLinkedFunctions */


/* debug [class_header]: Header for MTLLinkedFunctions */
// The class instance for the [LinkedFunctions] class.
var (
	LinkedFunctionsClass     _LinkedFunctionsClass
	LinkedFunctionsClassOnce sync.Once
)

func getLinkedFunctionsClass() _LinkedFunctionsClass {
	LinkedFunctionsClassOnce.Do(func() {
		LinkedFunctionsClass = _LinkedFunctionsClass{objc.GetClass("MTLLinkedFunctions")}
	})
	return LinkedFunctionsClass
}

type _LinkedFunctionsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LinkedFunctions */
// An interface definition for the [LinkedFunctions] class.
type ILinkedFunctions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for LinkedFunctions */
	// properties:
	BinaryFunctions() []objc.ID
	SetBinaryFunctions(value []objc.ID)
	Functions() []objc.ID
	SetFunctions(value []objc.ID)
	Groups() foundation.IDictionary
	SetGroups(value foundation.IDictionary)
	PrivateFunctions() []objc.ID
	SetPrivateFunctions(value []objc.ID)
	BinaryArchives() BinaryArchive /* not a class type */
	SetBinaryArchives(value BinaryArchive /* not a class type */)
	ConstantValues() IMTLFunctionConstantValues
	SetConstantValues(value IMTLFunctionConstantValues)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Options() FunctionOptions
	SetOptions(value FunctionOptions)
	SpecializedName() objc.IObject /* cross-framework: NSString */
	SetSpecializedName(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LinkedFunctions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LinkedFunctions */
// Alloc allocates a new instance without initialization.
func (lc _LinkedFunctionsClass) Alloc() LinkedFunctions {
	rv := objc.Send[LinkedFunctions](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LinkedFunctionsClass) New() LinkedFunctions {
	rv := objc.Send[LinkedFunctions](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LinkedFunctions) Init() LinkedFunctions {
	rv := objc.Send[LinkedFunctions](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LinkedFunctions) Autorelease() LinkedFunctions {
	rv := objc.Send[LinkedFunctions](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLinkedFunctions creates a new LinkedFunctions instance.
func NewLinkedFunctions() LinkedFunctions {
	return getLinkedFunctionsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LinkedFunctions */
// A set of related functions that Metal links to when necessary to create the function instance.
//
// When you create a Metal function instance using an , you specify additional functions that Metal needs to link to when it compiles and links the underlying shader code. Most often, you need to do this if your shader takes a visible function table as one or more of its arguments. For Metal to create the instance, it needs a complete list of functions that your shader can call so that it can resolve any dependencies and generate the correct code to run on the GPU.


// A set of related functions that Metal links to when necessary to create the function instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions
type LinkedFunctions struct {
	objectivec.Object
}

// LinkedFunctionsFrom constructs a [LinkedFunctions] from an unsafe.Pointer.
//
// A set of related functions that Metal links to when necessary to create the function instance.
func LinkedFunctionsFrom(ptr unsafe.Pointer) LinkedFunctions {
	return LinkedFunctions{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LinkedFunctions *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LinkedFunctions */

// Creates an empty linked functions object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/linkedFunctions
func (lc _LinkedFunctionsClass) LinkedFunctions() ILinkedFunctions {
	rv := objc.Send[LinkedFunctions](objc.ID(lc.class), objc.Sel("linkedFunctions"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LinkedFunctions) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LinkedFunctions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LinkedFunctions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LinkedFunctions */

// An array of function objects already compiled to a binary representation to link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/binaryFunctions
func (l_ LinkedFunctions) BinaryFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](l_.ID, objc.Sel("binaryFunctions"))
	return rv
}/* debug [instance_properties/getter]: binaryFunctions */


// An array of function objects already compiled to a binary representation to link.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/binaryFunctions
func (l_ LinkedFunctions) SetBinaryFunctions(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](l_.ID, objc.Sel("setBinaryFunctions:"), nsArray)
}/* debug [instance_properties/setter]: binaryFunctions */


// An array of function objects to link to the new function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/functions
func (l_ LinkedFunctions) Functions() []objc.ID {
	rv := objc.Send[[]objc.ID](l_.ID, objc.Sel("functions"))
	return rv
}/* debug [instance_properties/getter]: functions */


// An array of function objects to link to the new function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/functions
func (l_ LinkedFunctions) SetFunctions(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](l_.ID, objc.Sel("setFunctions:"), nsArray)
}/* debug [instance_properties/setter]: functions */


// An optional list of groups specifying which functions your shader can call at each call site.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/groups
func (l_ LinkedFunctions) Groups() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](l_.ID, objc.Sel("groups"))
	return rv
}/* debug [instance_properties/getter]: groups */


// An optional list of groups specifying which functions your shader can call at each call site.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/groups
func (l_ LinkedFunctions) SetGroups(value foundation.IDictionary) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setGroups:"), value)
}/* debug [instance_properties/setter]: groups */


// An array of function objects to link to the new function, without exporting the functions publicly.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/privateFunctions
func (l_ LinkedFunctions) PrivateFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](l_.ID, objc.Sel("privateFunctions"))
	return rv
}/* debug [instance_properties/getter]: privateFunctions */


// An array of function objects to link to the new function, without exporting the functions publicly.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/privateFunctions
func (l_ LinkedFunctions) SetPrivateFunctions(value []objc.ID) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](l_.ID, objc.Sel("setPrivateFunctions:"), nsArray)
}/* debug [instance_properties/setter]: privateFunctions */


// The binary archives to search for a previously-compiled version of this function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/binaryarchives
func (l_ LinkedFunctions) BinaryArchives() BinaryArchive /* not a class type */ {
	rv := objc.Send[BinaryArchive](l_.ID, objc.Sel("binaryArchives"))
	return rv
}/* debug [instance_properties/getter]: binaryArchives */


// The binary archives to search for a previously-compiled version of this function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/binaryarchives
func (l_ LinkedFunctions) SetBinaryArchives(value BinaryArchive /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBinaryArchives:"), value)
}/* debug [instance_properties/setter]: binaryArchives */


// The set of constant values assigned to the function constants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/constantvalues
func (l_ LinkedFunctions) ConstantValues() IMTLFunctionConstantValues {
	rv := objc.Send[FunctionConstantValues](l_.ID, objc.Sel("constantValues"))
	return rv
}/* debug [instance_properties/getter]: constantValues */


// The set of constant values assigned to the function constants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/constantvalues
func (l_ LinkedFunctions) SetConstantValues(value IMTLFunctionConstantValues) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setConstantValues:"), value)
}/* debug [instance_properties/setter]: constantValues */


// The name of the function to fetch from the library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/name
func (l_ LinkedFunctions) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The name of the function to fetch from the library.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/name
func (l_ LinkedFunctions) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// Flags specifying how Metal should create the new function object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/options
func (l_ LinkedFunctions) Options() FunctionOptions {
	rv := objc.Send[FunctionOptions](l_.ID, objc.Sel("options"))
	return rv
}/* debug [instance_properties/getter]: options */


// Flags specifying how Metal should create the new function object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/options
func (l_ LinkedFunctions) SetOptions(value FunctionOptions) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOptions:"), value)
}/* debug [instance_properties/setter]: options */


// A new name for the created function object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/specializedname
func (l_ LinkedFunctions) SpecializedName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("specializedName"))
	return rv
}/* debug [instance_properties/getter]: specializedName */


// A new name for the created function object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/specializedname
func (l_ LinkedFunctions) SetSpecializedName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setSpecializedName:"), value)
}/* debug [instance_properties/setter]: specializedName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLLinkedFunctions */



