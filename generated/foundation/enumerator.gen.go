// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Enumerator] class.
var enumeratorClass = _EnumeratorClass{objc.GetClass("NSEnumerator")}

type _EnumeratorClass struct {
	class objc.Class
}

// An interface definition for the [Enumerator] class.
type IEnumerator interface {
	objectivec.IObject
}

// An abstract class whose subclasses enumerate collections of objects, such as arrays and dictionaries. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return enumeratorClass.New()
}




