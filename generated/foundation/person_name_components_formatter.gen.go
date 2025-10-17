// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PersonNameComponentsFormatter] class.
var PersonNameComponentsFormatterClass objc.Class

func init() {
	PersonNameComponentsFormatterClass = objc.GetClass("NSPersonNameComponentsFormatter")
}

type PersonNameComponentsFormatter struct {
	objc.ID
}

func PersonNameComponentsFormatterFrom(ptr unsafe.Pointer) PersonNameComponentsFormatter {
	return PersonNameComponentsFormatter{
		ID: objc.ID(ptr),
	}
}


// Returns a person name components object from a given string. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/PersonNameComponentsFormatter/personNameComponents(from:)
func (p_ PersonNameComponentsFormatter) PersonNameComponentsFromString(string string) unsafe.Pointer {
	sel := objc.RegisterName("personNameComponentsFromString:")
	ret := p_.ID.Send(sel, string)
	return unsafe.Pointer(ret)
}

