// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Predicate] class.
var (
	PredicateClass     _PredicateClass
	PredicateClassOnce sync.Once
)

func getPredicateClass() _PredicateClass {
	PredicateClassOnce.Do(func() {
		PredicateClass = _PredicateClass{objc.GetClass("MPSPredicate")}
	})
	return PredicateClass
}

type _PredicateClass struct {
	class objc.Class
}

// An interface definition for the [Predicate] class.
type IPredicate interface {
	objectivec.IObject
	// properties:
	PredicateBuffer() Buffer /* not a class type */
	SetPredicateBuffer(value Buffer /* not a class type */)
	PredicateOffset() int
	SetPredicateOffset(value int)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPredicate
type Predicate struct {
	objectivec.Object
}

// PredicateFrom constructs a [Predicate] from an unsafe.Pointer.
func PredicateFrom(ptr unsafe.Pointer) Predicate {
	return Predicate{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PredicateClass) Alloc() Predicate {
	rv := objc.Send[Predicate](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PredicateClass) New() Predicate {
	rv := objc.Send[Predicate](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Predicate) Init() Predicate {
	rv := objc.Send[Predicate](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Predicate) Autorelease() Predicate {
	rv := objc.Send[Predicate](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPredicate creates a new Predicate instance.
func NewPredicate() Predicate {
	return getPredicateClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPredicate/init(device:)
func NewPredicateWithDevice(device objectivec.IObject) Predicate {
	instance := getPredicateClass().Alloc()
	rv := objc.Send[Predicate](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate/predicatebuffer
func (p_ Predicate) PredicateBuffer() Buffer /* not a class type */ {
	rv := objc.Send[Buffer](p_.ID, objc.Sel("predicateBuffer"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate/predicatebuffer
func (p_ Predicate) SetPredicateBuffer(value Buffer /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPredicateBuffer:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate/predicateoffset
func (p_ Predicate) PredicateOffset() int {
	rv := objc.Send[int](p_.ID, objc.Sel("predicateOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate/predicateoffset
func (p_ Predicate) SetPredicateOffset(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPredicateOffset:"), value)
}


