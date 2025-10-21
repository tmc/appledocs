// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SFSafariExtensionState] class.
var (
	SFSafariExtensionStateClass     _SFSafariExtensionStateClass
	SFSafariExtensionStateClassOnce sync.Once
)

func getSFSafariExtensionStateClass() _SFSafariExtensionStateClass {
	SFSafariExtensionStateClassOnce.Do(func() {
		SFSafariExtensionStateClass = _SFSafariExtensionStateClass{objc.GetClass("SFSafariExtensionState")}
	})
	return SFSafariExtensionStateClass
}

type _SFSafariExtensionStateClass struct {
	class objc.Class
}

// An interface definition for the [SFSafariExtensionState] class.
type ISFSafariExtensionState interface {
	objectivec.IObject
}

// The state of a Safari app extension.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariExtensionState
type SFSafariExtensionState struct {
	objectivec.Object
}

// SFSafariExtensionStateFrom constructs a [SFSafariExtensionState] from an unsafe.Pointer.
//
// The state of a Safari app extension.
func SFSafariExtensionStateFrom(ptr unsafe.Pointer) SFSafariExtensionState {
	return SFSafariExtensionState{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFSafariExtensionStateClass) Alloc() SFSafariExtensionState {
	rv := objc.Send[SFSafariExtensionState](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFSafariExtensionStateClass) New() SFSafariExtensionState {
	rv := objc.Send[SFSafariExtensionState](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariExtensionState) Init() SFSafariExtensionState {
	rv := objc.Send[SFSafariExtensionState](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariExtensionState) Autorelease() SFSafariExtensionState {
	rv := objc.Send[SFSafariExtensionState](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariExtensionState creates a new SFSafariExtensionState instance.
func NewSFSafariExtensionState() SFSafariExtensionState {
	return getSFSafariExtensionStateClass().New()
}


// A Boolean value that indicates whether the user has enabled the app extension.
//
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfsafariextensionstate/isenabled
func (s_ SFSafariExtensionState) IsEnabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isEnabled"))
	return rv
}


// SetIsEnabled sets the value of the isEnabled property.
// A Boolean value that indicates whether the user has enabled the app extension.

//
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfsafariextensionstate/isenabled
func (s_ SFSafariExtensionState) SetIsEnabled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsEnabled:"), value)
}

// A string the system uses as a key in a user info dictionary to identify a profile identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfextensionprofilekey
func (s_ SFSafariExtensionState) SFExtensionProfileKey() string {
	rv := objc.Send[string](s_.ID, objc.Sel("SFExtensionProfileKey"))
	return rv
}

// A Boolean value that indicates whether the user has enabled the app extension.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariExtensionState/isEnabled
func (s_ SFSafariExtensionState) Enabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("enabled"))
	return rv
}



