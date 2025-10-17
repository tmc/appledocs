// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PersonNameComponentsFormatter] class.
var personNameComponentsFormatterClass = _PersonNameComponentsFormatterClass{objc.GetClass("NSPersonNameComponentsFormatter")}

type _PersonNameComponentsFormatterClass struct {
	class objc.Class
}

// An interface definition for the [PersonNameComponentsFormatter] class.
type IPersonNameComponentsFormatter interface {
	IFormatter
	PersonNameComponentsFromString(string string) unsafe.Pointer
}

// A formatter that provides localized representations of the components of a person’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter

type PersonNameComponentsFormatter struct {
	Formatter
}

// PersonNameComponentsFormatterFrom constructs a [PersonNameComponentsFormatter] from an unsafe.Pointer.
//
// A formatter that provides localized representations of the components of a person’s name.
func PersonNameComponentsFormatterFrom(ptr unsafe.Pointer) PersonNameComponentsFormatter {
	return PersonNameComponentsFormatter{
		Formatter: FormatterFrom(ptr),
	}
}

// Returns a person name components object from a given string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/personNameComponents(from:)
func (p_ PersonNameComponentsFormatter) PersonNameComponentsFromString(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("personNameComponentsFromString:"), string)
	return rv
}


