// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SFSafariViewControllerPrewarmingToken] class.
var (
	SFSafariViewControllerPrewarmingTokenClass     _SFSafariViewControllerPrewarmingTokenClass
	SFSafariViewControllerPrewarmingTokenClassOnce sync.Once
)

func getSFSafariViewControllerPrewarmingTokenClass() _SFSafariViewControllerPrewarmingTokenClass {
	SFSafariViewControllerPrewarmingTokenClassOnce.Do(func() {
		SFSafariViewControllerPrewarmingTokenClass = _SFSafariViewControllerPrewarmingTokenClass{objc.GetClass("SFSafariViewControllerPrewarmingToken")}
	})
	return SFSafariViewControllerPrewarmingTokenClass
}

type _SFSafariViewControllerPrewarmingTokenClass struct {
	class objc.Class
}

// An interface definition for the [SFSafariViewControllerPrewarmingToken] class.
type ISFSafariViewControllerPrewarmingToken interface {
	objectivec.IObject
	Invalidate()
}

//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/PrewarmingToken
type SFSafariViewControllerPrewarmingToken struct {
	objectivec.Object
}

// SFSafariViewControllerPrewarmingTokenFrom constructs a [SFSafariViewControllerPrewarmingToken] from an unsafe.Pointer.
func SFSafariViewControllerPrewarmingTokenFrom(ptr unsafe.Pointer) SFSafariViewControllerPrewarmingToken {
	return SFSafariViewControllerPrewarmingToken{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SFSafariViewControllerPrewarmingTokenClass) Alloc() SFSafariViewControllerPrewarmingToken {
	rv := objc.Send[SFSafariViewControllerPrewarmingToken](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SFSafariViewControllerPrewarmingTokenClass) New() SFSafariViewControllerPrewarmingToken {
	rv := objc.Send[SFSafariViewControllerPrewarmingToken](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariViewControllerPrewarmingToken) Init() SFSafariViewControllerPrewarmingToken {
	rv := objc.Send[SFSafariViewControllerPrewarmingToken](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariViewControllerPrewarmingToken) Autorelease() SFSafariViewControllerPrewarmingToken {
	rv := objc.Send[SFSafariViewControllerPrewarmingToken](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariViewControllerPrewarmingToken creates a new SFSafariViewControllerPrewarmingToken instance.
func NewSFSafariViewControllerPrewarmingToken() SFSafariViewControllerPrewarmingToken {
	return getSFSafariViewControllerPrewarmingTokenClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/PrewarmingToken/invalidate()
func (s_ SFSafariViewControllerPrewarmingToken) Invalidate() {
	objc.Send[objc.ID](s_.ID, objc.Sel("invalidate"))
}



