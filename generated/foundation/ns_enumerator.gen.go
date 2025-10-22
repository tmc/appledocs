// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Enumerator] class.
var (
	EnumeratorClass     _EnumeratorClass
	EnumeratorClassOnce sync.Once
)

func getEnumeratorClass() _EnumeratorClass {
	EnumeratorClassOnce.Do(func() {
		EnumeratorClass = _EnumeratorClass{objc.GetClass("NSEnumerator")}
	})
	return EnumeratorClass
}

type _EnumeratorClass struct {
	class objc.Class
}

// An interface definition for the [Enumerator] class.
type IEnumerator interface {
	objectivec.IObject
}

// A parent class referenced by other Foundation classes.


// A parent class referenced by other Foundation classes. [Full Topic]
type Enumerator struct {
	objectivec.Object
}

// EnumeratorFrom constructs a [Enumerator] from an unsafe.Pointer.
//
// A parent class referenced by other Foundation classes.
func EnumeratorFrom(ptr unsafe.Pointer) Enumerator {
	return Enumerator{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _EnumeratorClass) Alloc() Enumerator {
	rv := objc.Send[Enumerator](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _EnumeratorClass) New() Enumerator {
	rv := objc.Send[Enumerator](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ Enumerator) Init() Enumerator {
	rv := objc.Send[Enumerator](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ Enumerator) Autorelease() Enumerator {
	rv := objc.Send[Enumerator](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEnumerator creates a new Enumerator instance.
func NewEnumerator() Enumerator {
	return getEnumeratorClass().New()
}




