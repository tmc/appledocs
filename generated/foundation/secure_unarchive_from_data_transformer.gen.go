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



