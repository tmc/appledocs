// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PersonNameComponentsFormatter] class.
var (
	personNameComponentsFormatterClass     _PersonNameComponentsFormatterClass
	personNameComponentsFormatterClassOnce sync.Once
)

func getPersonNameComponentsFormatterClass() _PersonNameComponentsFormatterClass {
	personNameComponentsFormatterClassOnce.Do(func() {
		personNameComponentsFormatterClass = _PersonNameComponentsFormatterClass{objc.GetClass("NSPersonNameComponentsFormatter")}
	})
	return personNameComponentsFormatterClass
}

type _PersonNameComponentsFormatterClass struct {
	class objc.Class
}

// An interface definition for the [PersonNameComponentsFormatter] class.
type IPersonNameComponentsFormatter interface {
	IFormatter
	PersonNameComponentsFromString(string string) unsafe.Pointer
}

// A formatter that provides localized representations of the components of a person’s name.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getPersonNameComponentsFormatterClass().New()
}


// Returns a person name components object from a given string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/personNameComponents(from:)
func (p_ PersonNameComponentsFormatter) PersonNameComponentsFromString(string string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("personNameComponentsFromString:"), objc.String(string))
	return rv
}


