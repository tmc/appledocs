// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TermOfAddress] class.
var termOfAddressClass = _TermOfAddressClass{objc.GetClass("NSTermOfAddress")}

type _TermOfAddressClass struct {
	class objc.Class
}

// An interface definition for the [TermOfAddress] class.
type ITermOfAddress interface {
	objectivec.IObject
}

// The type for representing grammatical gender in localized text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTermOfAddress

type TermOfAddress struct {
	objectivec.Object
}

// TermOfAddressFrom constructs a [TermOfAddress] from an unsafe.Pointer.
//
// The type for representing grammatical gender in localized text.
func TermOfAddressFrom(ptr unsafe.Pointer) TermOfAddress {
	return TermOfAddress{objectivec.Object{objc.ID(ptr)}}
}



