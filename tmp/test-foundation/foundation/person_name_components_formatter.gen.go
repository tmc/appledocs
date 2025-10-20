// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var PersonNameComponentsFormatterClass _PersonNameComponentsFormatterClass

func init() {
	PersonNameComponentsFormatterClass = _PersonNameComponentsFormatterClass{objc.GetClass("NSPersonNameComponentsFormatter")}
}

type _PersonNameComponentsFormatterClass struct {
	class objc.Class
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/personNameComponents(from:)
func (p_ PersonNameComponentsFormatter) PersonNameComponentsFromString(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("personNameComponentsFromString:"), string)
	return rv
}


