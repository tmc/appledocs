// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [DomainStateCompanion] class.
var (
	DomainStateCompanionClass     _DomainStateCompanionClass
	DomainStateCompanionClassOnce sync.Once
)

func getDomainStateCompanionClass() _DomainStateCompanionClass {
	DomainStateCompanionClassOnce.Do(func() {
		DomainStateCompanionClass = _DomainStateCompanionClass{objc.GetClass("LADomainStateCompanion")}
	})
	return DomainStateCompanionClass
}

type _DomainStateCompanionClass struct {
	class objc.Class
}

// An interface definition for the [DomainStateCompanion] class.
type IDomainStateCompanion interface {
	objectivec.IObject
	StateHashForCompanionType(companionType unsafe.Pointer) unsafe.Pointer
}

//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateCompanion
type DomainStateCompanion struct {
	objectivec.Object
}

// DomainStateCompanionFrom constructs a [DomainStateCompanion] from an unsafe.Pointer.
func DomainStateCompanionFrom(ptr unsafe.Pointer) DomainStateCompanion {
	return DomainStateCompanion{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DomainStateCompanionClass) Alloc() DomainStateCompanion {
	rv := objc.Send[DomainStateCompanion](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DomainStateCompanionClass) New() DomainStateCompanion {
	rv := objc.Send[DomainStateCompanion](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DomainStateCompanion) Init() DomainStateCompanion {
	rv := objc.Send[DomainStateCompanion](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DomainStateCompanion) Autorelease() DomainStateCompanion {
	rv := objc.Send[DomainStateCompanion](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDomainStateCompanion creates a new DomainStateCompanion instance.
func NewDomainStateCompanion() DomainStateCompanion {
	return getDomainStateCompanionClass().New()
}


// Returns state hash data for the given companion type.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateCompanion/stateHash(for:)
func (d_ DomainStateCompanion) StateHashForCompanionType(companionType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("stateHashForCompanionType:"), companionType)
	return rv
}

// Indicates types of companions paired with the device. The elements are NSNumber-wrapped instances of @c .
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateCompanion/availableCompanionTypes-1ggnh
func (d_ DomainStateCompanion) AvailableCompanionTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("availableCompanionTypes"))
	return rv
}

// Contains combined state hash data for all available companion types. . Returns if no companion devices are paired.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateCompanion/stateHash
func (d_ DomainStateCompanion) StateHash() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("stateHash"))
	return rv
}



