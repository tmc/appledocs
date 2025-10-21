// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SFContentBlockerState] class.
var (
	SFContentBlockerStateClass     _SFContentBlockerStateClass
	SFContentBlockerStateClassOnce sync.Once
)

func getSFContentBlockerStateClass() _SFContentBlockerStateClass {
	SFContentBlockerStateClassOnce.Do(func() {
		SFContentBlockerStateClass = _SFContentBlockerStateClass{objc.GetClass("SFContentBlockerState")}
	})
	return SFContentBlockerStateClass
}

type _SFContentBlockerStateClass struct {
	class objc.Class
}

// An interface definition for the [SFContentBlockerState] class.
type ISFContentBlockerState interface {
	objectivec.IObject
}

// The state of a content blocker extension.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFContentBlockerState
type SFContentBlockerState struct {
	objectivec.Object
}

// SFContentBlockerStateFrom constructs a [SFContentBlockerState] from an unsafe.Pointer.
//
// The state of a content blocker extension.
func SFContentBlockerStateFrom(ptr unsafe.Pointer) SFContentBlockerState {
	return SFContentBlockerState{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFContentBlockerStateClass) Alloc() SFContentBlockerState {
	rv := objc.Send[SFContentBlockerState](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFContentBlockerStateClass) New() SFContentBlockerState {
	rv := objc.Send[SFContentBlockerState](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFContentBlockerState) Init() SFContentBlockerState {
	rv := objc.Send[SFContentBlockerState](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFContentBlockerState) Autorelease() SFContentBlockerState {
	rv := objc.Send[SFContentBlockerState](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFContentBlockerState creates a new SFContentBlockerState instance.
func NewSFContentBlockerState() SFContentBlockerState {
	return getSFContentBlockerStateClass().New()
}


// A Boolean value that indicates whether the content blocker is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFContentBlockerState/isEnabled
func (s_ SFContentBlockerState) Enabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("enabled"))
	return rv
}



