// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [streamerVersion] class.
var (
	StreamerVersionClass     _streamerVersionClass
	StreamerVersionClassOnce sync.Once
)

func getstreamerVersionClass() _streamerVersionClass {
	StreamerVersionClassOnce.Do(func() {
		StreamerVersionClass = _streamerVersionClass{objc.GetClass("streamerVersion")}
	})
	return StreamerVersionClass
}

type _streamerVersionClass struct {
	class objc.Class
}

// An interface definition for the [streamerVersion] class.
type IstreamerVersion interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/streamerVersion
type streamerVersion struct {
	objectivec.Object
}

// streamerVersionFrom constructs a [streamerVersion] from an unsafe.Pointer.
func streamerVersionFrom(ptr unsafe.Pointer) streamerVersion {
	return streamerVersion{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _streamerVersionClass) Alloc() streamerVersion {
	rv := objc.Send[streamerVersion](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _streamerVersionClass) New() streamerVersion {
	rv := objc.Send[streamerVersion](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ streamerVersion) Init() streamerVersion {
	rv := objc.Send[streamerVersion](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ streamerVersion) Autorelease() streamerVersion {
	rv := objc.Send[streamerVersion](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewstreamerVersion creates a new streamerVersion instance.
func NewstreamerVersion() streamerVersion {
	return getstreamerVersionClass().New()
}




