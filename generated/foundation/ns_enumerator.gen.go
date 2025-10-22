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
	NextObject() unsafe.Pointer
	AllObjects() []objc.ID
}

// An abstract class whose subclasses enumerate collections of objects, such as arrays and dictionaries.
//
// All creation methods are defined in the collection classes—such as , , and —which provide special objects with which to enumerate their contents. For example, has two methods that return an object: and . also has two methods that return an object: and . These methods let you enumerate the contents of a dictionary by key or by value, respectively. You send repeatedly to a newly created object to have it return the next object in the original collection. When the collection is exhausted, is returned. You cannot “reset” an enumerator after it has exhausted its collection. To enumerate a collection again, you need a new enumerator. The enumerator subclasses used by , , and retain the collection during enumeration. When the enumeration is exhausted, the collection is released.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEnumerator
type Enumerator struct {
	objectivec.Object
}

// EnumeratorFrom constructs a [Enumerator] from an unsafe.Pointer.
//
// An abstract class whose subclasses enumerate collections of objects, such as arrays and dictionaries.
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


// Returns the next object from the collection being enumerated.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEnumerator/nextObject()
func (e_ Enumerator) NextObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("nextObject"))
	return rv
}

// The array of unenumerated objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEnumerator/allObjects
func (e_ Enumerator) AllObjects() []objc.ID {
	rv := objc.Send[[]objc.ID](e_.ID, objc.Sel("allObjects"))
	return rv
}



