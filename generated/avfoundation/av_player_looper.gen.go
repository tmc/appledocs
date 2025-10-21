// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PlayerLooper] class.
var (
	PlayerLooperClass     _PlayerLooperClass
	PlayerLooperClassOnce sync.Once
)

func getPlayerLooperClass() _PlayerLooperClass {
	PlayerLooperClassOnce.Do(func() {
		PlayerLooperClass = _PlayerLooperClass{objc.GetClass("AVPlayerLooper")}
	})
	return PlayerLooperClass
}

type _PlayerLooperClass struct {
	class objc.Class
}

// An interface definition for the [PlayerLooper] class.
type IPlayerLooper interface {
	objectivec.IObject
}

// An object that loops media content using a queue player.
//
// You can manually implement looping playback in your app using , but provides a much simpler interface to loop a single . You create a player looper by passing it a reference to your and a template and the looper automatically manages the looping playback of this content (see example).
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper
type PlayerLooper struct {
	objectivec.Object
}

// PlayerLooperFrom constructs a [PlayerLooper] from an unsafe.Pointer.
//
// An object that loops media content using a queue player.
func PlayerLooperFrom(ptr unsafe.Pointer) PlayerLooper {
	return PlayerLooper{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PlayerLooperClass) Alloc() PlayerLooper {
	rv := objc.Send[PlayerLooper](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlayerLooperClass) New() PlayerLooper {
	rv := objc.Send[PlayerLooper](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlayerLooper) Init() PlayerLooper {
	rv := objc.Send[PlayerLooper](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlayerLooper) Autorelease() PlayerLooper {
	rv := objc.Send[PlayerLooper](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayerLooper creates a new PlayerLooper instance.
func NewPlayerLooper() PlayerLooper {
	return getPlayerLooperClass().New()
}




