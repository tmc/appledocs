// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersonNameComponents] class.
var personNameComponentsClass = _PersonNameComponentsClass{objc.GetClass("NSPersonNameComponents")}

type _PersonNameComponentsClass struct {
	class objc.Class
}

// An object that manages the separate parts of a person’s name to allow locale-aware formatting. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPersonNameComponents

type PersonNameComponents struct {
	objectivec.Object
}

// PersonNameComponentsFrom constructs a [PersonNameComponents] from an unsafe.Pointer.
//
// An object that manages the separate parts of a person’s name to allow locale-aware formatting.
func PersonNameComponentsFrom(ptr unsafe.Pointer) PersonNameComponents {
	return PersonNameComponents{objectivec.Object{objc.ID(ptr)}}
}



