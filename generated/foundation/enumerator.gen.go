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



