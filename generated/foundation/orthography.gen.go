// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Orthography] class.
var orthographyClass = _OrthographyClass{objc.GetClass("NSOrthography")}

type _OrthographyClass struct {
	class objc.Class
}

// An interface definition for the [Orthography] class.
type IOrthography interface {
	objectivec.IObject
}

// A description of the linguistic content of natural language text, typically used for spelling and grammar checking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOrthography

type Orthography struct {
	objectivec.Object
}

// OrthographyFrom constructs a [Orthography] from an unsafe.Pointer.
//
// A description of the linguistic content of natural language text, typically used for spelling and grammar checking.
func OrthographyFrom(ptr unsafe.Pointer) Orthography {
	return Orthography{objectivec.Object{objc.ID(ptr)}}
}



