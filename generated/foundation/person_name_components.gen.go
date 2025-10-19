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

// An interface definition for the [PersonNameComponents] class.
type IPersonNameComponents interface {
	objectivec.IObject
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
// Alloc allocates a new instance without initialization.
func (pc _PersonNameComponentsClass) Alloc() PersonNameComponents {
	rv := objc.Send[PersonNameComponents](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _PersonNameComponentsClass) New() PersonNameComponents {
	rv := objc.Send[PersonNameComponents](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersonNameComponents) Init() PersonNameComponents {
	rv := objc.Send[PersonNameComponents](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersonNameComponents) Autorelease() PersonNameComponents {
	rv := objc.Send[PersonNameComponents](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersonNameComponents creates a new PersonNameComponents instance.
func NewPersonNameComponents() PersonNameComponents {
	return personNameComponentsClass.New()
}




