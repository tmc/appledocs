// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MusicPlayerPlayParameters] class.
var (
	MusicPlayerPlayParametersClass     _MusicPlayerPlayParametersClass
	MusicPlayerPlayParametersClassOnce sync.Once
)

func getMusicPlayerPlayParametersClass() _MusicPlayerPlayParametersClass {
	MusicPlayerPlayParametersClassOnce.Do(func() {
		MusicPlayerPlayParametersClass = _MusicPlayerPlayParametersClass{objc.GetClass("MPMusicPlayerPlayParameters")}
	})
	return MusicPlayerPlayParametersClass
}

type _MusicPlayerPlayParametersClass struct {
	class objc.Class
}

// An interface definition for the [MusicPlayerPlayParameters] class.
type IMusicPlayerPlayParameters interface {
	objectivec.IObject
}

// The MusicKit parameters that describe items to play.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMusicPlayerPlayParameters
type MusicPlayerPlayParameters struct {
	objectivec.Object
}

// MusicPlayerPlayParametersFrom constructs a [MusicPlayerPlayParameters] from an unsafe.Pointer.
//
// The MusicKit parameters that describe items to play.
func MusicPlayerPlayParametersFrom(ptr unsafe.Pointer) MusicPlayerPlayParameters {
	return MusicPlayerPlayParameters{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MusicPlayerPlayParametersClass) Alloc() MusicPlayerPlayParameters {
	rv := objc.Send[MusicPlayerPlayParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MusicPlayerPlayParametersClass) New() MusicPlayerPlayParameters {
	rv := objc.Send[MusicPlayerPlayParameters](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MusicPlayerPlayParameters) Init() MusicPlayerPlayParameters {
	rv := objc.Send[MusicPlayerPlayParameters](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MusicPlayerPlayParameters) Autorelease() MusicPlayerPlayParameters {
	rv := objc.Send[MusicPlayerPlayParameters](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMusicPlayerPlayParameters creates a new MusicPlayerPlayParameters instance.
func NewMusicPlayerPlayParameters() MusicPlayerPlayParameters {
	return getMusicPlayerPlayParametersClass().New()
}




