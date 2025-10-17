// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Typesetter] class.
var typesetterClass = _TypesetterClass{objc.GetClass("NSTypesetter")}

type _TypesetterClass struct {
	class objc.Class
}

// An interface definition for the [Typesetter] class.
type ITypesetter interface {
	objectivec.IObject
}

// An abstract class that performs various type layout tasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetter

type Typesetter struct {
	objectivec.Object
}

// TypesetterFrom constructs a [Typesetter] from an unsafe.Pointer.
//
// An abstract class that performs various type layout tasks.
func TypesetterFrom(ptr unsafe.Pointer) Typesetter {
	return Typesetter{objectivec.Object{objc.ID(ptr)}}
}



