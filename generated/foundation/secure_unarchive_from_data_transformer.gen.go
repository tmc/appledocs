// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SecureUnarchiveFromDataTransformer] class.
var secureUnarchiveFromDataTransformerClass = _SecureUnarchiveFromDataTransformerClass{objc.GetClass("NSSecureUnarchiveFromDataTransformer")}

type _SecureUnarchiveFromDataTransformerClass struct {
	class objc.Class
}

// An interface definition for the [SecureUnarchiveFromDataTransformer] class.
type ISecureUnarchiveFromDataTransformer interface {
	IValueTransformer
}

// A value transformer that converts data to and from classes that support secure coding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSecureUnarchiveFromDataTransformer

type SecureUnarchiveFromDataTransformer struct {
	ValueTransformer
}

// SecureUnarchiveFromDataTransformerFrom constructs a [SecureUnarchiveFromDataTransformer] from an unsafe.Pointer.
//
// A value transformer that converts data to and from classes that support secure coding.
func SecureUnarchiveFromDataTransformerFrom(ptr unsafe.Pointer) SecureUnarchiveFromDataTransformer {
	return SecureUnarchiveFromDataTransformer{
		ValueTransformer: ValueTransformerFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _SecureUnarchiveFromDataTransformerClass) Alloc() SecureUnarchiveFromDataTransformer {
	rv := objc.Send[SecureUnarchiveFromDataTransformer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SecureUnarchiveFromDataTransformerClass) New() SecureUnarchiveFromDataTransformer {
	rv := objc.Send[SecureUnarchiveFromDataTransformer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SecureUnarchiveFromDataTransformer) Init() SecureUnarchiveFromDataTransformer {
	rv := objc.Send[SecureUnarchiveFromDataTransformer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SecureUnarchiveFromDataTransformer) Autorelease() SecureUnarchiveFromDataTransformer {
	rv := objc.Send[SecureUnarchiveFromDataTransformer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSecureUnarchiveFromDataTransformer creates a new SecureUnarchiveFromDataTransformer instance.
func NewSecureUnarchiveFromDataTransformer() SecureUnarchiveFromDataTransformer {
	return secureUnarchiveFromDataTransformerClass.New()
}




