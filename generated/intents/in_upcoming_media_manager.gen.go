// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INUpcomingMediaManager] class.
var (
	INUpcomingMediaManagerClass     _INUpcomingMediaManagerClass
	INUpcomingMediaManagerClassOnce sync.Once
)

func getINUpcomingMediaManagerClass() _INUpcomingMediaManagerClass {
	INUpcomingMediaManagerClassOnce.Do(func() {
		INUpcomingMediaManagerClass = _INUpcomingMediaManagerClass{objc.GetClass("INUpcomingMediaManager")}
	})
	return INUpcomingMediaManagerClass
}

type _INUpcomingMediaManagerClass struct {
	class objc.Class
}

// An interface definition for the [INUpcomingMediaManager] class.
type IINUpcomingMediaManager interface {
	objectivec.IObject
	SetPredictionModeForType(mode INUpcomingMediaPredictionMode, type_ unsafe.Pointer)
	SetSuggestedMediaIntents(intents unsafe.Pointer)
}

// The manager object you use to suggest media to the user.
//
// Use this class to provide Siri a list of media intents for content that the user hasn’t listened to or watched, but might be interested in. For example, a podcast app may provide the latest episodes of the podcast, or a video app may provide the most recent episodes of TV shows, or suggest new movies.

// The manager object you use to suggest media to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INUpcomingMediaManager
type INUpcomingMediaManager struct {
	objectivec.Object
}

// INUpcomingMediaManagerFrom constructs a [INUpcomingMediaManager] from an unsafe.Pointer.
//
// The manager object you use to suggest media to the user.
func INUpcomingMediaManagerFrom(ptr unsafe.Pointer) INUpcomingMediaManager {
	return INUpcomingMediaManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INUpcomingMediaManagerClass) Alloc() INUpcomingMediaManager {
	rv := objc.Send[INUpcomingMediaManager](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INUpcomingMediaManagerClass) New() INUpcomingMediaManager {
	rv := objc.Send[INUpcomingMediaManager](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INUpcomingMediaManager) Init() INUpcomingMediaManager {
	rv := objc.Send[INUpcomingMediaManager](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INUpcomingMediaManager) Autorelease() INUpcomingMediaManager {
	rv := objc.Send[INUpcomingMediaManager](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINUpcomingMediaManager creates a new INUpcomingMediaManager instance.
func NewINUpcomingMediaManager() INUpcomingMediaManager {
	return getINUpcomingMediaManagerClass().New()
}

// The shared upcoming media manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INUpcomingMediaManager/shared
func (ic _INUpcomingMediaManagerClass) SharedManager() INUpcomingMediaManager {
	rv := objc.Send[INUpcomingMediaManager](objc.ID(ic.class), objc.Sel("sharedManager"))
	return rv
}

// Suggests how Siri should predict media intent shortcuts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INUpcomingMediaManager/setPredictionMode(_:for:)
func (i_ INUpcomingMediaManager) SetPredictionModeForType(mode INUpcomingMediaPredictionMode, type_ unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPredictionMode:forType:"), mode, type_)
}

// Provides Siri with a list of media intents to suggest to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INUpcomingMediaManager/setSuggestedMediaIntents(_:)
func (i_ INUpcomingMediaManager) SetSuggestedMediaIntents(intents unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSuggestedMediaIntents:"), intents)
}

// The shared upcoming media manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INUpcomingMediaManager/shared
func (i_ INUpcomingMediaManager) SharedManager() INUpcomingMediaManager {
	rv := objc.Send[INUpcomingMediaManager](i_.ID, objc.Sel("sharedManager"))
	return rv
}
