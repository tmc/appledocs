// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [LinkedFunctions] class.
type ILinkedFunctions interface {
	objectivec.IObject
	BinaryFunctions() []objc.ID
	SetBinaryFunctions(value []objc.ID)
	Functions() []objc.ID
	SetFunctions(value []objc.ID)
	Groups() unsafe.Pointer
	SetGroups(value unsafe.Pointer)
	PrivateFunctions() []objc.ID
	SetPrivateFunctions(value []objc.ID)
	BinaryArchives() unsafe.Pointer
	SetBinaryArchives(value unsafe.Pointer)
	ConstantValues() unsafe.Pointer
	SetConstantValues(value unsafe.Pointer)
	Name() string
	SetName(value string)
	Options() FunctionOptions
	SetOptions(value FunctionOptions)
	SpecializedName() string
	SetSpecializedName(value string)
}

// A set of related functions that Metal links to when necessary to create the function instance.
//
// When you create a Metal function instance using an , you specify additional functions that Metal needs to link to when it compiles and links the underlying shader code. Most often, you need to do this if your shader takes a visible function table as one or more of its arguments. For Metal to create the instance, it needs a complete list of functions that your shader can call so that it can resolve any dependencies and generate the correct code to run on the GPU.
//
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

// Alloc allocates a new instance without initialization.
func (lc _LinkedFunctionsClass) Alloc() LinkedFunctions {
	rv := objc.Send[LinkedFunctions](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates an empty linked functions object.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/linkedFunctions
func (lc _LinkedFunctionsClass) LinkedFunctions() LinkedFunctions {
	rv := objc.Send[LinkedFunctions](objc.ID(lc.class), objc.Sel("linkedFunctions"))
	return rv
}

// An array of function objects already compiled to a binary representation to link.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/binaryFunctions
func (l_ LinkedFunctions) BinaryFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](l_.ID, objc.Sel("binaryFunctions"))
	return rv
}


// SetBinaryFunctions sets the value of the binaryFunctions property.
// An array of function objects already compiled to a binary representation to link.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/binaryFunctions
func (l_ LinkedFunctions) SetBinaryFunctions(value []objc.ID) {
	// Convert Go slice to NSArray
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
}

// An array of function objects to link to the new function.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/functions
func (l_ LinkedFunctions) Functions() []objc.ID {
	rv := objc.Send[[]objc.ID](l_.ID, objc.Sel("functions"))
	return rv
}


// SetFunctions sets the value of the functions property.
// An array of function objects to link to the new function.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/functions
func (l_ LinkedFunctions) SetFunctions(value []objc.ID) {
	// Convert Go slice to NSArray
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
}

// An optional list of groups specifying which functions your shader can call at each call site.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/groups
func (l_ LinkedFunctions) Groups() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("groups"))
	return rv
}


// SetGroups sets the value of the groups property.
// An optional list of groups specifying which functions your shader can call at each call site.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/groups
func (l_ LinkedFunctions) SetGroups(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setGroups:"), value)
}

// An array of function objects to link to the new function, without exporting the functions publicly.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/privateFunctions
func (l_ LinkedFunctions) PrivateFunctions() []objc.ID {
	rv := objc.Send[[]objc.ID](l_.ID, objc.Sel("privateFunctions"))
	return rv
}


// SetPrivateFunctions sets the value of the privateFunctions property.
// An array of function objects to link to the new function, without exporting the functions publicly.

//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLLinkedFunctions/privateFunctions
func (l_ LinkedFunctions) SetPrivateFunctions(value []objc.ID) {
	// Convert Go slice to NSArray
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
}

// The binary archives to search for a previously-compiled version of this function.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/binaryarchives
func (l_ LinkedFunctions) BinaryArchives() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("binaryArchives"))
	return rv
}


// SetBinaryArchives sets the value of the binaryArchives property.
// The binary archives to search for a previously-compiled version of this function.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/binaryarchives
func (l_ LinkedFunctions) SetBinaryArchives(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBinaryArchives:"), value)
}

// The set of constant values assigned to the function constants.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/constantvalues
func (l_ LinkedFunctions) ConstantValues() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("constantValues"))
	return rv
}


// SetConstantValues sets the value of the constantValues property.
// The set of constant values assigned to the function constants.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/constantvalues
func (l_ LinkedFunctions) SetConstantValues(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setConstantValues:"), value)
}

// The name of the function to fetch from the library.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/name
func (l_ LinkedFunctions) Name() string {
	rv := objc.Send[string](l_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name of the function to fetch from the library.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/name
func (l_ LinkedFunctions) SetName(value string) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setName:"), objc.String(value))
}

// Flags specifying how Metal should create the new function object.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/options
func (l_ LinkedFunctions) Options() FunctionOptions {
	rv := objc.Send[FunctionOptions](l_.ID, objc.Sel("options"))
	return rv
}


// SetOptions sets the value of the options property.
// Flags specifying how Metal should create the new function object.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/options
func (l_ LinkedFunctions) SetOptions(value FunctionOptions) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOptions:"), value)
}

// A new name for the created function object.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/specializedname
func (l_ LinkedFunctions) SpecializedName() string {
	rv := objc.Send[string](l_.ID, objc.Sel("specializedName"))
	return rv
}


// SetSpecializedName sets the value of the specializedName property.
// A new name for the created function object.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctiondescriptor/specializedname
func (l_ LinkedFunctions) SetSpecializedName(value string) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setSpecializedName:"), objc.String(value))
}



