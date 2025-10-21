// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [DomainState] class.
var (
	DomainStateClass     _DomainStateClass
	DomainStateClassOnce sync.Once
)

func getDomainStateClass() _DomainStateClass {
	DomainStateClassOnce.Do(func() {
		DomainStateClass = _DomainStateClass{objc.GetClass("LADomainState")}
	})
	return DomainStateClass
}

type _DomainStateClass struct {
	class objc.Class
}

// An interface definition for the [DomainState] class.
type IDomainState interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainState
type DomainState struct {
	objectivec.Object
}

// DomainStateFrom constructs a [DomainState] from an unsafe.Pointer.
func DomainStateFrom(ptr unsafe.Pointer) DomainState {
	return DomainState{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DomainStateClass) Alloc() DomainState {
	rv := objc.Send[DomainState](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DomainStateClass) New() DomainState {
	rv := objc.Send[DomainState](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DomainState) Init() DomainState {
	rv := objc.Send[DomainState](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DomainState) Autorelease() DomainState {
	rv := objc.Send[DomainState](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDomainState creates a new DomainState instance.
func NewDomainState() DomainState {
	return getDomainStateClass().New()
}


// Contains biometric domain state.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainState/biometry
func (d_ DomainState) Biometry() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("biometry"))
	return rv
}

// Contains companion domain state.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainState/companion
func (d_ DomainState) Companion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("companion"))
	return rv
}

// Contains combined state hash data for biometry and companion state hashes.
//
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainState/stateHash
func (d_ DomainState) StateHash() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("stateHash"))
	return rv
}



