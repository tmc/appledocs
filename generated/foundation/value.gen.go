// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Value] class.
var valueClass = _ValueClass{objc.GetClass("NSValue")}

type _ValueClass struct {
	class objc.Class
}

// An interface definition for the [Value] class.
type IValue interface {
	objectivec.IObject
}

// A simple container for a single C or Objective-C data item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSValue

type Value struct {
	objectivec.Object
}

// ValueFrom constructs a [Value] from an unsafe.Pointer.
//
// A simple container for a single C or Objective-C data item.
func ValueFrom(ptr unsafe.Pointer) Value {
	return Value{objectivec.Object{objc.ID(ptr)}}
}



