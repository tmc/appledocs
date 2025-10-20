// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PlayableContentManagerContext] class.
var (
	PlayableContentManagerContextClass     _PlayableContentManagerContextClass
	PlayableContentManagerContextClassOnce sync.Once
)

func getPlayableContentManagerContextClass() _PlayableContentManagerContextClass {
	PlayableContentManagerContextClassOnce.Do(func() {
		PlayableContentManagerContextClass = _PlayableContentManagerContextClass{objc.GetClass("MPPlayableContentManagerContext")}
	})
	return PlayableContentManagerContextClass
}

type _PlayableContentManagerContextClass struct {
	class objc.Class
}

// An interface definition for the [PlayableContentManagerContext] class.
type IPlayableContentManagerContext interface {
	objectivec.IObject
}

// An object representing the current state of the playable endpoint.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPPlayableContentManagerContext
type PlayableContentManagerContext struct {
	objectivec.Object
}

// PlayableContentManagerContextFrom constructs a [PlayableContentManagerContext] from an unsafe.Pointer.
//
// An object representing the current state of the playable endpoint.
func PlayableContentManagerContextFrom(ptr unsafe.Pointer) PlayableContentManagerContext {
	return PlayableContentManagerContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PlayableContentManagerContextClass) Alloc() PlayableContentManagerContext {
	rv := objc.Send[PlayableContentManagerContext](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlayableContentManagerContextClass) New() PlayableContentManagerContext {
	rv := objc.Send[PlayableContentManagerContext](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayableContentManagerContext) Init() PlayableContentManagerContext {
	rv := objc.Send[PlayableContentManagerContext](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayableContentManagerContext) Autorelease() PlayableContentManagerContext {
	rv := objc.Send[PlayableContentManagerContext](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayableContentManagerContext creates a new PlayableContentManagerContext instance.
func NewPlayableContentManagerContext() PlayableContentManagerContext {
	return getPlayableContentManagerContextClass().New()
}




