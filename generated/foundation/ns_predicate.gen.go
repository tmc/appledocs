// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

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
		PredicateClass = _PredicateClass{objc.GetClass("NSPredicate")}
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
	PredicateFormat() string /* primitive/slice/pointer */
	SetPredicateFormat(value string /* primitive/slice/pointer */)
	// methods:
}

// A definition of logical conditions for constraining a search for a fetch or for in-memory filtering.
//
// Predicates represent logical conditions, which you can use to filter collections of objects. Although it’s common to create predicates directly from instances of , , and , you often create predicates from a format string that the class methods parse on . Examples of predicate format strings include: Simple comparisons, such as or Case- and diacritic-insensitive lookups, such as Logical operations, such as Temporal range constraints, such as Relational conditions, such as Aggregate operations, such as For a complete syntax reference, refer to the . You can also create predicates that include variables using the method so that you can predefine the predicate before substituting concrete values at runtime.


// A definition of logical conditions for constraining a search for a fetch or for in-memory filtering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPredicate
type Predicate struct {
	objectivec.Object
}

// PredicateFrom constructs a [Predicate] from an unsafe.Pointer.
//
// A definition of logical conditions for constraining a search for a fetch or for in-memory filtering.
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



// The predicate’s format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspredicate/predicateformat
func (p_ Predicate) PredicateFormat() string /* primitive/slice/pointer */ {
	rv := objc.Send[string](p_.ID, objc.Sel("predicateFormat"))
	return rv
}


// The predicate’s format string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspredicate/predicateformat
func (p_ Predicate) SetPredicateFormat(value string /* primitive/slice/pointer */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPredicateFormat:"), objc.String(value))
}



