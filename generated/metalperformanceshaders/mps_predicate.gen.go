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
	PredicateBuffer() unsafe.Pointer
	SetPredicateBuffer(value unsafe.Pointer)
	PredicateOffset() int
	SetPredicateOffset(value int)
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSPredicate/init(device:)
func NewPredicateWithDevice(device objectivec.IObject) Predicate {
	instance := getPredicateClass().Alloc()
	rv := objc.Send[Predicate](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate/predicatebuffer
func (p_ Predicate) PredicateBuffer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("predicateBuffer"))
	return rv
}


// SetPredicateBuffer sets the value of the predicateBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate/predicatebuffer
func (p_ Predicate) SetPredicateBuffer(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPredicateBuffer:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate/predicateoffset
func (p_ Predicate) PredicateOffset() int {
	rv := objc.Send[int](p_.ID, objc.Sel("predicateOffset"))
	return rv
}


// SetPredicateOffset sets the value of the predicateOffset property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspredicate/predicateoffset
func (p_ Predicate) SetPredicateOffset(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPredicateOffset:"), value)
}


