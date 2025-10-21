// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CloudPlayer] class.
var (
	CloudPlayerClass     _CloudPlayerClass
	CloudPlayerClassOnce sync.Once
)

func getCloudPlayerClass() _CloudPlayerClass {
	CloudPlayerClassOnce.Do(func() {
		CloudPlayerClass = _CloudPlayerClass{objc.GetClass("GKCloudPlayer")}
	})
	return CloudPlayerClass
}

type _CloudPlayerClass struct {
	class objc.Class
}

// An interface definition for the [CloudPlayer] class.
type ICloudPlayer interface {
	IBasePlayer
}

// The object representing the currently signed-in iCloud user.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKCloudPlayer
type CloudPlayer struct {
	BasePlayer
}

// CloudPlayerFrom constructs a [CloudPlayer] from an unsafe.Pointer.
//
// The object representing the currently signed-in iCloud user.
func CloudPlayerFrom(ptr unsafe.Pointer) CloudPlayer {
	return CloudPlayer{
		BasePlayer: BasePlayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CloudPlayerClass) Alloc() CloudPlayer {
	rv := objc.Send[CloudPlayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CloudPlayerClass) New() CloudPlayer {
	rv := objc.Send[CloudPlayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CloudPlayer) Init() CloudPlayer {
	rv := objc.Send[CloudPlayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CloudPlayer) Autorelease() CloudPlayer {
	rv := objc.Send[CloudPlayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCloudPlayer creates a new CloudPlayer instance.
func NewCloudPlayer() CloudPlayer {
	return getCloudPlayerClass().New()
}




