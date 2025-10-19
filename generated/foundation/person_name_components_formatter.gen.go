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
// Alloc allocates a new instance without initialization.
func (pc _PersonNameComponentsFormatterClass) Alloc() PersonNameComponentsFormatter {
	rv := objc.Send[PersonNameComponentsFormatter](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _PersonNameComponentsFormatterClass) New() PersonNameComponentsFormatter {
	rv := objc.Send[PersonNameComponentsFormatter](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersonNameComponentsFormatter) Init() PersonNameComponentsFormatter {
	rv := objc.Send[PersonNameComponentsFormatter](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersonNameComponentsFormatter) Autorelease() PersonNameComponentsFormatter {
	rv := objc.Send[PersonNameComponentsFormatter](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersonNameComponentsFormatter creates a new PersonNameComponentsFormatter instance.
func NewPersonNameComponentsFormatter() PersonNameComponentsFormatter {
	return personNameComponentsFormatterClass.New()
}


// Returns a person name components object from a given string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/personNameComponents(from:)
func (p_ PersonNameComponentsFormatter) PersonNameComponentsFromString(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("personNameComponentsFromString:"), string)
	return rv
}


