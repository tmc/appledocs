// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PointerFunctions] class.
var (
	PointerFunctionsClass     _PointerFunctionsClass
	PointerFunctionsClassOnce sync.Once
)

func getPointerFunctionsClass() _PointerFunctionsClass {
	PointerFunctionsClassOnce.Do(func() {
		PointerFunctionsClass = _PointerFunctionsClass{objc.GetClass("NSPointerFunctions")}
	})
	return PointerFunctionsClass
}

type _PointerFunctionsClass struct {
	class objc.Class
}

// An interface definition for the [PointerFunctions] class.
type IPointerFunctions interface {
	objectivec.IObject
	PointerFunctions() NSPointerFunctions
	SetPointerFunctions(value IPointerFunctions)
	DescriptionFunction() string
	SetDescriptionFunction(value string)
	HashFunction() int
	SetHashFunction(value int)
	IsEqualFunction() unsafe.Pointer
	SetIsEqualFunction(value unsafe.Pointer)
	RelinquishFunction() unsafe.Pointer
	SetRelinquishFunction(value unsafe.Pointer)
	UsesStrongWriteBarrier() bool
	SetUsesStrongWriteBarrier(value bool)
	UsesWeakReadAndWriteBarriers() bool
	SetUsesWeakReadAndWriteBarriers(value bool)
}

// An instance of defines callout functions appropriate for managing a pointer reference held somewhere else.
//
// The functions specified by an instance of are separated into two clusters—those that define “personality” such as “object” or “C-string”, and those that describe memory management issues such as a memory deallocation function. There are constants for common personalities and memory manager selections (see ). , , and use an object to define the acquisition and retention behavior for the pointers they manage. Note, however, that not all combinations of personality and memory management behavior are valid for these collections. The pointer collection objects copy the object on input and output, so you cannot usefully subclass .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions
type PointerFunctions struct {
	objectivec.Object
}

// PointerFunctionsFrom constructs a [PointerFunctions] from an unsafe.Pointer.
//
// An instance of defines callout functions appropriate for managing a pointer reference held somewhere else.
func PointerFunctionsFrom(ptr unsafe.Pointer) PointerFunctions {
	return PointerFunctions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PointerFunctionsClass) Alloc() PointerFunctions {
	rv := objc.Send[PointerFunctions](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PointerFunctionsClass) New() PointerFunctions {
	rv := objc.Send[PointerFunctions](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PointerFunctions) Init() PointerFunctions {
	rv := objc.Send[PointerFunctions](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PointerFunctions) Autorelease() PointerFunctions {
	rv := objc.Send[PointerFunctions](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPointerFunctions creates a new PointerFunctions instance.
func NewPointerFunctions() PointerFunctions {
	return getPointerFunctionsClass().New()
}


// The pointer functions for the hash table.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshashtable/pointerfunctions
func (p_ PointerFunctions) PointerFunctions() NSPointerFunctions {
	rv := objc.Send[NSPointerFunctions](p_.ID, objc.Sel("pointerFunctions"))
	return rv
}


// SetPointerFunctions sets the value of the pointerFunctions property.
// The pointer functions for the hash table.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshashtable/pointerfunctions
func (p_ PointerFunctions) SetPointerFunctions(value IPointerFunctions) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPointerFunctions:"), value)
}

// The function used to describe elements.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/descriptionfunction
func (p_ PointerFunctions) DescriptionFunction() string {
	rv := objc.Send[string](p_.ID, objc.Sel("descriptionFunction"))
	return rv
}


// SetDescriptionFunction sets the value of the descriptionFunction property.
// The function used to describe elements.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/descriptionfunction
func (p_ PointerFunctions) SetDescriptionFunction(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDescriptionFunction:"), objc.String(value))
}

// The hash function.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/hashfunction
func (p_ PointerFunctions) HashFunction() int {
	rv := objc.Send[int](p_.ID, objc.Sel("hashFunction"))
	return rv
}


// SetHashFunction sets the value of the hashFunction property.
// The hash function.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/hashfunction
func (p_ PointerFunctions) SetHashFunction(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHashFunction:"), value)
}

// The function used to compare pointers.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/isequalfunction
func (p_ PointerFunctions) IsEqualFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("isEqualFunction"))
	return rv
}


// SetIsEqualFunction sets the value of the isEqualFunction property.
// The function used to compare pointers.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/isequalfunction
func (p_ PointerFunctions) SetIsEqualFunction(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsEqualFunction:"), value)
}

// The function used to relinquish memory.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/relinquishfunction
func (p_ PointerFunctions) RelinquishFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("relinquishFunction"))
	return rv
}


// SetRelinquishFunction sets the value of the relinquishFunction property.
// The function used to relinquish memory.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/relinquishfunction
func (p_ PointerFunctions) SetRelinquishFunction(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRelinquishFunction:"), value)
}

// Specifies whether, in a garbage collected environment, pointers should be assigned using a strong write barrier.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/usesstrongwritebarrier
func (p_ PointerFunctions) UsesStrongWriteBarrier() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("usesStrongWriteBarrier"))
	return rv
}


// SetUsesStrongWriteBarrier sets the value of the usesStrongWriteBarrier property.
// Specifies whether, in a garbage collected environment, pointers should be assigned using a strong write barrier.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/usesstrongwritebarrier
func (p_ PointerFunctions) SetUsesStrongWriteBarrier(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUsesStrongWriteBarrier:"), value)
}

// Specifies whether, in a garbage collected environment, pointers should use weak read and write barriers.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/usesweakreadandwritebarriers
func (p_ PointerFunctions) UsesWeakReadAndWriteBarriers() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("usesWeakReadAndWriteBarriers"))
	return rv
}


// SetUsesWeakReadAndWriteBarriers sets the value of the usesWeakReadAndWriteBarriers property.
// Specifies whether, in a garbage collected environment, pointers should use weak read and write barriers.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerfunctions/usesweakreadandwritebarriers
func (p_ PointerFunctions) SetUsesWeakReadAndWriteBarriers(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUsesWeakReadAndWriteBarriers:"), value)
}



