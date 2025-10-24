// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPointerArray */


/* debug [class_header]: Header for NSPointerArray */
// The class instance for the [PointerArray] class.
var (
	PointerArrayClass     _PointerArrayClass
	PointerArrayClassOnce sync.Once
)

func getPointerArrayClass() _PointerArrayClass {
	PointerArrayClassOnce.Do(func() {
		PointerArrayClass = _PointerArrayClass{objc.GetClass("NSPointerArray")}
	})
	return PointerArrayClass
}

type _PointerArrayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PointerArray */
// An interface definition for the [PointerArray] class.
type IPointerArray interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PointerArray */
	// properties:
	AllObjects() IArray
	Count() uint
	SetCount(value uint)
	PointerFunctions() IPointerFunctions
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PointerArray */
	// methods:
	AddPointer(pointer objectivec.IObject)
	Compact()
	InsertPointerAtIndex(item objectivec.IObject, index uint)
	PointerAtIndex(index uint)
	RemovePointerAtIndex(index uint)
	ReplacePointerAtIndexWithPointer(index uint, item objectivec.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PointerArray */
// Alloc allocates a new instance without initialization.
func (pc _PointerArrayClass) Alloc() PointerArray {
	rv := objc.Send[PointerArray](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PointerArrayClass) New() PointerArray {
	rv := objc.Send[PointerArray](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PointerArray) Init() PointerArray {
	rv := objc.Send[PointerArray](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PointerArray) Autorelease() PointerArray {
	rv := objc.Send[PointerArray](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPointerArray creates a new PointerArray instance.
func NewPointerArray() PointerArray {
	return getPointerArrayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PointerArray */
// A collection similar to an array, but with a broader range of available memory semantics.
//
// The pointer array class is modeled after , but can also hold values. You can insert or remove values which contribute to the array’s . A pointer array can be initialized to maintain strong or weak references to objects, or according to any of the memory or personality options defined by . The and protocols are applicable only when a pointer array is initialized to maintain strong or weak references to objects. When enumerating a pointer array with using , the loop will yield any values present in the array. See in for more information.


// A collection similar to an array, but with a broader range of available memory semantics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray
type PointerArray struct {
	objectivec.Object
}

// PointerArrayFrom constructs a [PointerArray] from an unsafe.Pointer.
//
// A collection similar to an array, but with a broader range of available memory semantics.
func PointerArrayFrom(ptr unsafe.Pointer) PointerArray {
	return PointerArray{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PointerArray */

// Initializes the receiver to use the given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/init(options:)
func NewPointerArrayWithOptions(options PointerFunctionsOptions) PointerArray {
	instance := getPointerArrayClass().Alloc()
	rv := objc.Send[PointerArray](instance.ID, objc.Sel("initWithOptions:"), options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPointerArrayWithOptions */


// Initializes the receiver to use the given functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/init(pointerFunctions:)
func NewPointerArrayWithPointerFunctions(functions IPointerFunctions) PointerArray {
	instance := getPointerArrayClass().Alloc()
	rv := objc.Send[PointerArray](instance.ID, objc.Sel("initWithPointerFunctions:"), functions)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPointerArrayWithPointerFunctions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PointerArray */

// Returns a new pointer array initialized to use the given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/pointerArrayWithOptions:
func (pc _PointerArrayClass) PointerArrayWithOptions(options PointerFunctionsOptions) IPointerArray {
	rv := objc.Send[PointerArray](objc.ID(pc.class), objc.Sel("pointerArrayWithOptions:"), options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PointerArrayWithOptions) */


// A new pointer array initialized to use the given functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/pointerArrayWithPointerFunctions:
func (pc _PointerArrayClass) PointerArrayWithPointerFunctions(functions IPointerFunctions) IPointerArray {
	rv := objc.Send[PointerArray](objc.ID(pc.class), objc.Sel("pointerArrayWithPointerFunctions:"), functions)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PointerArrayWithPointerFunctions) */


// Returns a new pointer array that maintains strong references to its elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/pointerArrayWithStrongObjects
func (pc _PointerArrayClass) PointerArrayWithStrongObjects() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("pointerArrayWithStrongObjects"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PointerArrayWithStrongObjects) */


// Returns a new pointer array that maintains weak references to its elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/pointerArrayWithWeakObjects
func (pc _PointerArrayClass) PointerArrayWithWeakObjects() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("pointerArrayWithWeakObjects"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PointerArrayWithWeakObjects) */


// Returns a new pointer array that maintains strong references to its elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/strongObjects()
func (pc _PointerArrayClass) StrongObjectsPointerArray() IPointerArray {
	rv := objc.Send[PointerArray](objc.ID(pc.class), objc.Sel("strongObjectsPointerArray"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StrongObjectsPointerArray) */


// Returns a new pointer array that maintains weak references to its elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/weakObjects()
func (pc _PointerArrayClass) WeakObjectsPointerArray() IPointerArray {
	rv := objc.Send[PointerArray](objc.ID(pc.class), objc.Sel("weakObjectsPointerArray"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WeakObjectsPointerArray) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PointerArray */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PointerArray */

// Adds a given pointer to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/addPointer(_:)
func (p_ PointerArray) AddPointer(pointer objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addPointer:"), pointer)
}/* debug [instance_methods/method]: AddPointer */


// Removes values from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/compact()
func (p_ PointerArray) Compact() {
	objc.Send[objc.ID](p_.ID, objc.Sel("compact"))
}/* debug [instance_methods/method]: Compact */


// Inserts a pointer at a given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/insertPointer(_:at:)
func (p_ PointerArray) InsertPointerAtIndex(item objectivec.IObject, index uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("insertPointer:atIndex:"), item, index)
}/* debug [instance_methods/method]: InsertPointerAtIndex */


// Returns the pointer at a given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/pointer(at:)
func (p_ PointerArray) PointerAtIndex(index uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("pointerAtIndex:"), index)
}/* debug [instance_methods/method]: PointerAtIndex */


// Removes the pointer at a given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/removePointer(at:)
func (p_ PointerArray) RemovePointerAtIndex(index uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removePointerAtIndex:"), index)
}/* debug [instance_methods/method]: RemovePointerAtIndex */


// Replaces the pointer at a given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/replacePointer(at:withPointer:)
func (p_ PointerArray) ReplacePointerAtIndexWithPointer(index uint, item objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("replacePointerAtIndex:withPointer:"), index, item)
}/* debug [instance_methods/method]: ReplacePointerAtIndexWithPointer */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PointerArray */

// All the objects in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/allObjects
func (p_ PointerArray) AllObjects() IArray {
	rv := objc.Send[Array](p_.ID, objc.Sel("allObjects"))
	return rv
}/* debug [instance_properties/getter]: allObjects */


// The number of elements in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/count
func (p_ PointerArray) Count() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("count"))
	return rv
}/* debug [instance_properties/getter]: count */


// The number of elements in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/count
func (p_ PointerArray) SetCount(value uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCount:"), value)
}/* debug [instance_properties/setter]: count */


// The functions in use by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/pointerFunctions
func (p_ PointerArray) PointerFunctions() IPointerFunctions {
	rv := objc.Send[PointerFunctions](p_.ID, objc.Sel("pointerFunctions"))
	return rv
}/* debug [instance_properties/getter]: pointerFunctions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPointerArray */


