// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CompoundPredicate] class.
var (
	CompoundPredicateClass     _CompoundPredicateClass
	CompoundPredicateClassOnce sync.Once
)

func getCompoundPredicateClass() _CompoundPredicateClass {
	CompoundPredicateClassOnce.Do(func() {
		CompoundPredicateClass = _CompoundPredicateClass{objc.GetClass("NSCompoundPredicate")}
	})
	return CompoundPredicateClass
}

type _CompoundPredicateClass struct {
	class objc.Class
}

// An interface definition for the [CompoundPredicate] class.
type ICompoundPredicate interface {
	IPredicate
}

// A specialized predicate that evaluates logical combinations of other predicates.
//
// Use to create an or compound predicate of one or more other predicates, or the of a single predicate. For the logical and operations: An predicate with no subpredicates evaluates to . An predicate with no subpredicates evaluates to . A compound predicate with one or more subpredicates evaluates to the truth of its subpredicates.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate
type CompoundPredicate struct {
	Predicate
}

// CompoundPredicateFrom constructs a [CompoundPredicate] from an unsafe.Pointer.
//
// A specialized predicate that evaluates logical combinations of other predicates.
func CompoundPredicateFrom(ptr unsafe.Pointer) CompoundPredicate {
	return CompoundPredicate{
		Predicate: PredicateFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CompoundPredicateClass) Alloc() CompoundPredicate {
	rv := objc.Send[CompoundPredicate](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CompoundPredicateClass) New() CompoundPredicate {
	rv := objc.Send[CompoundPredicate](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CompoundPredicate) Init() CompoundPredicate {
	rv := objc.Send[CompoundPredicate](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CompoundPredicate) Autorelease() CompoundPredicate {
	rv := objc.Send[CompoundPredicate](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCompoundPredicate creates a new CompoundPredicate instance.
func NewCompoundPredicate() CompoundPredicate {
	return getCompoundPredicateClass().New()
}




// Returns a new predicate that you form using an AND operation on the predicates in a specified array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/init(andPredicateWithSubpredicates:)
func NewCompoundPredicateAndPredicateWithSubpredicates(subpredicates unsafe.Pointer) CompoundPredicate {
	rv := objc.Send[CompoundPredicate](objc.ID(getCompoundPredicateClass().class), objc.Sel("andPredicateWithSubpredicates:"), subpredicates)
	return rv
}



// Returns a new predicate that you form using a NOT operation on a specified predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/init(notPredicateWithSubpredicate:)
func NewCompoundPredicateNotPredicateWithSubpredicate(predicate unsafe.Pointer) CompoundPredicate {
	rv := objc.Send[CompoundPredicate](objc.ID(getCompoundPredicateClass().class), objc.Sel("notPredicateWithSubpredicate:"), predicate)
	return rv
}



// Returns the receiver that a specified type initializes using predicates from a specified array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/init(type:subpredicates:)
func NewCompoundPredicateWithTypeSubpredicates(type_ unsafe.Pointer, subpredicates unsafe.Pointer) CompoundPredicate {
	instance := getCompoundPredicateClass().Alloc()
	rv := objc.Send[CompoundPredicate](instance.ID, objc.Sel("initWithType:subpredicates:"), type_, subpredicates)
	rv.Autorelease()
	return rv
}


// Returns a new predicate that you form using an AND operation on the predicates in a specified array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/init(andPredicateWithSubpredicates:)
func (cc _CompoundPredicateClass) AndPredicateWithSubpredicates(subpredicates unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("andPredicateWithSubpredicates:"), subpredicates)
	return rv
}

// Returns a new predicate that you form using a NOT operation on a specified predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/init(notPredicateWithSubpredicate:)
func (cc _CompoundPredicateClass) NotPredicateWithSubpredicate(predicate unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("notPredicateWithSubpredicate:"), predicate)
	return rv
}

// The predicate type for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/compoundPredicateType
func (c_ CompoundPredicate) CompoundPredicateType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("compoundPredicateType"))
	return rv
}


