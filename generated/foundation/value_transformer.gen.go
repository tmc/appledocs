// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ValueTransformer] class.
var valueTransformerClass = _ValueTransformerClass{objc.GetClass("NSValueTransformer")}

type _ValueTransformerClass struct {
	class objc.Class
}

// An interface definition for the [ValueTransformer] class.
type IValueTransformer interface {
	objectivec.IObject
}

// An abstract class used to transform values from one representation to another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ValueTransformer

type ValueTransformer struct {
	objectivec.Object
}

// ValueTransformerFrom constructs a [ValueTransformer] from an unsafe.Pointer.
//
// An abstract class used to transform values from one representation to another.
func ValueTransformerFrom(ptr unsafe.Pointer) ValueTransformer {
	return ValueTransformer{objectivec.Object{objc.ID(ptr)}}
}



