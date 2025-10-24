// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [odPredicate] class.
var (
	OdPredicateClass     _odPredicateClass
	OdPredicateClassOnce sync.Once
)

func getodPredicateClass() _odPredicateClass {
	OdPredicateClassOnce.Do(func() {
		OdPredicateClass = _odPredicateClass{objc.GetClass("odPredicate")}
	})
	return OdPredicateClass
}

type _odPredicateClass struct {
	class objc.Class
}

// An interface definition for the [odPredicate] class.
type IodPredicate interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/odPredicate-c.ivar
type odPredicate struct {
	objectivec.Object
}

// odPredicateFrom constructs a [odPredicate] from an unsafe.Pointer.
func odPredicateFrom(ptr unsafe.Pointer) odPredicate {
	return odPredicate{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _odPredicateClass) Alloc() odPredicate {
	rv := objc.Send[odPredicate](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _odPredicateClass) New() odPredicate {
	rv := objc.Send[odPredicate](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ odPredicate) Init() odPredicate {
	rv := objc.Send[odPredicate](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ odPredicate) Autorelease() odPredicate {
	rv := objc.Send[odPredicate](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewodPredicate creates a new odPredicate instance.
func NewodPredicate() odPredicate {
	return getodPredicateClass().New()
}




