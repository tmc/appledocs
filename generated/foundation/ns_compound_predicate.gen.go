// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCompoundPredicate */


/* debug [class_header]: Header for NSCompoundPredicate */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CompoundPredicate */
// An interface definition for the [CompoundPredicate] class.
type ICompoundPredicate interface {
	IPredicate
	
/* debug [class_interface_properties]: Properties for CompoundPredicate */
	// properties:
	CompoundPredicateType() CompoundPredicateType
	Subpredicates() IArray
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CompoundPredicate */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CompoundPredicate */
// Alloc allocates a new instance without initialization.
func (cc _CompoundPredicateClass) Alloc() CompoundPredicate {
	rv := objc.Send[CompoundPredicate](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CompoundPredicate */
// A specialized predicate that evaluates logical combinations of other predicates.
//
// Use to create an or compound predicate of one or more other predicates, or the of a single predicate. For the logical and operations: An predicate with no subpredicates evaluates to . An predicate with no subpredicates evaluates to . A compound predicate with one or more subpredicates evaluates to the truth of its subpredicates.


// A specialized predicate that evaluates logical combinations of other predicates.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CompoundPredicate */

// Returns a new predicate that you form using an AND operation on the predicates in a specified array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/init(andPredicateWithSubpredicates:)
func NewCompoundPredicateAndPredicateWithSubpredicates(subpredicates []Predicate) CompoundPredicate {
	rv := objc.Send[CompoundPredicate](objc.ID(getCompoundPredicateClass().class), objc.Sel("andPredicateWithSubpredicates:"), subpredicates)
	return rv
}/* debug [class_init_methods/constructor]: NewCompoundPredicateAndPredicateWithSubpredicates */


// Returns a new predicate that you form using a NOT operation on a specified predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/init(notPredicateWithSubpredicate:)
func NewCompoundPredicateNotPredicateWithSubpredicate(predicate IPredicate) CompoundPredicate {
	rv := objc.Send[CompoundPredicate](objc.ID(getCompoundPredicateClass().class), objc.Sel("notPredicateWithSubpredicate:"), predicate)
	return rv
}/* debug [class_init_methods/constructor]: NewCompoundPredicateNotPredicateWithSubpredicate */


// Returns a new predicate that you form using an OR operation on the predicates in a specified array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/init(orPredicateWithSubpredicates:)
func NewCompoundPredicateOrPredicateWithSubpredicates(subpredicates []Predicate) CompoundPredicate {
	rv := objc.Send[CompoundPredicate](objc.ID(getCompoundPredicateClass().class), objc.Sel("orPredicateWithSubpredicates:"), subpredicates)
	return rv
}/* debug [class_init_methods/constructor]: NewCompoundPredicateOrPredicateWithSubpredicates */


// Creates a predicate by decoding from the coder you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/init(coder:)
func NewCompoundPredicateWithCoder(coder ICoder) CompoundPredicate {
	instance := getCompoundPredicateClass().Alloc()
	rv := objc.Send[CompoundPredicate](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCompoundPredicateWithCoder */


// Returns the receiver that a specified type initializes using predicates from a specified array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/init(type:subpredicates:)
func NewCompoundPredicateWithTypeSubpredicates(type_ CompoundPredicateType, subpredicates []Predicate) CompoundPredicate {
	instance := getCompoundPredicateClass().Alloc()
	rv := objc.Send[CompoundPredicate](instance.ID, objc.Sel("initWithType:subpredicates:"), type_, subpredicates)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCompoundPredicateWithTypeSubpredicates */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CompoundPredicate */

// Returns a new predicate that you form using an AND operation on the predicates in a specified array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/init(andPredicateWithSubpredicates:)
func (cc _CompoundPredicateClass) AndPredicateWithSubpredicates(subpredicates []Predicate) ICompoundPredicate {
	rv := objc.Send[CompoundPredicate](objc.ID(cc.class), objc.Sel("andPredicateWithSubpredicates:"), subpredicates)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AndPredicateWithSubpredicates) */


// Returns a new predicate that you form using a NOT operation on a specified predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/init(notPredicateWithSubpredicate:)
func (cc _CompoundPredicateClass) NotPredicateWithSubpredicate(predicate IPredicate) ICompoundPredicate {
	rv := objc.Send[CompoundPredicate](objc.ID(cc.class), objc.Sel("notPredicateWithSubpredicate:"), predicate)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NotPredicateWithSubpredicate) */


// Returns a new predicate that you form using an OR operation on the predicates in a specified array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/init(orPredicateWithSubpredicates:)
func (cc _CompoundPredicateClass) OrPredicateWithSubpredicates(subpredicates []Predicate) ICompoundPredicate {
	rv := objc.Send[CompoundPredicate](objc.ID(cc.class), objc.Sel("orPredicateWithSubpredicates:"), subpredicates)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OrPredicateWithSubpredicates) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CompoundPredicate */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CompoundPredicate */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CompoundPredicate */

// The predicate type for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/compoundPredicateType
func (c_ CompoundPredicate) CompoundPredicateType() CompoundPredicateType {
	rv := objc.Send[CompoundPredicateType](c_.ID, objc.Sel("compoundPredicateType"))
	return rv
}/* debug [instance_properties/getter]: compoundPredicateType */


// The receiver’s subpredicates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/subpredicates
func (c_ CompoundPredicate) Subpredicates() IArray {
	rv := objc.Send[Array](c_.ID, objc.Sel("subpredicates"))
	return rv
}/* debug [instance_properties/getter]: subpredicates */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCompoundPredicate */


