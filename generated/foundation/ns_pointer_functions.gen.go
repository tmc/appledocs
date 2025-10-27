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
	

	// properties:
	UsesStrongWriteBarrier() bool
	SetUsesStrongWriteBarrier(value bool)
	UsesWeakReadAndWriteBarriers() bool
	SetUsesWeakReadAndWriteBarriers(value bool)
	PointerFunctions() IPointerFunctions
	SetPointerFunctions(value IPointerFunctions)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (pc _PointerFunctionsClass) Alloc() PointerFunctions {
	rv := objc.Send[PointerFunctions](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// An instance of defines callout functions appropriate for managing a pointer reference held somewhere else.
//
// The functions specified by an instance of are separated into two clusters—those that define “personality” such as “object” or “C-string”, and those that describe memory management issues such as a memory deallocation function. There are constants for common personalities and memory manager selections (see ). , , and use an object to define the acquisition and retention behavior for the pointers they manage. Note, however, that not all combinations of personality and memory management behavior are valid for these collections. The pointer collection objects copy the object on input and output, so you cannot usefully subclass .


// An instance of defines callout functions appropriate for managing a pointer reference held somewhere else.
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






// Returns an object initialized with the given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/init(options:)
func NewPointerFunctionsWithOptions(options PointerFunctionsOptions) PointerFunctions {
	instance := getPointerFunctionsClass().Alloc()
	rv := objc.Send[PointerFunctions](instance.ID, objc.Sel("initWithOptions:"), options)
	rv.Autorelease()
	return rv
}







// Returns a new object initialized with the given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/pointerFunctionsWithOptions:
func (pc _PointerFunctionsClass) PointerFunctionsWithOptions(options PointerFunctionsOptions) IPointerFunctions {
	rv := objc.Send[PointerFunctions](objc.ID(pc.class), objc.Sel("pointerFunctionsWithOptions:"), options)
	return rv
}

















// Specifies whether, in a garbage collected environment, pointers should be assigned using a strong write barrier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/usesStrongWriteBarrier
func (p_ PointerFunctions) UsesStrongWriteBarrier() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("usesStrongWriteBarrier"))
	return rv
}


// Specifies whether, in a garbage collected environment, pointers should be assigned using a strong write barrier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/usesStrongWriteBarrier
func (p_ PointerFunctions) SetUsesStrongWriteBarrier(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUsesStrongWriteBarrier:"), value)
}


// Specifies whether, in a garbage collected environment, pointers should use weak read and write barriers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/usesWeakReadAndWriteBarriers
func (p_ PointerFunctions) UsesWeakReadAndWriteBarriers() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("usesWeakReadAndWriteBarriers"))
	return rv
}


// Specifies whether, in a garbage collected environment, pointers should use weak read and write barriers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/usesWeakReadAndWriteBarriers
func (p_ PointerFunctions) SetUsesWeakReadAndWriteBarriers(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUsesWeakReadAndWriteBarriers:"), value)
}


// The pointer functions for the hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshashtable/pointerfunctions
func (p_ PointerFunctions) PointerFunctions() IPointerFunctions {
	rv := objc.Send[PointerFunctions](p_.ID, objc.Sel("pointerFunctions"))
	return rv
}


// The pointer functions for the hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nshashtable/pointerfunctions
func (p_ PointerFunctions) SetPointerFunctions(value IPointerFunctions) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPointerFunctions:"), value)
}







