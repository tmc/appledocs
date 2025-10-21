// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Achievement] class.
var (
	AchievementClass     _AchievementClass
	AchievementClassOnce sync.Once
)

func getAchievementClass() _AchievementClass {
	AchievementClassOnce.Do(func() {
		AchievementClass = _AchievementClass{objc.GetClass("GKAchievement")}
	})
	return AchievementClass
}

type _AchievementClass struct {
	class objc.Class
}

// An interface definition for the [Achievement] class.
type IAchievement interface {
	objectivec.IObject
	ChallengeComposeControllerWithPlayersMessageCompletionHandler(playerIDs unsafe.Pointer, message string, completionHandler unsafe.Pointer) unsafe.Pointer
}

// An achievement you can award a player as they make progress toward and reach a goal in your game.
//
// Before using this class, configure your game achievements in App Store Connect. Then the dashboard shows the achievements initially locked and you can access them in your code. Use the method to load all the achievements that the local player is progressing toward. If an achievement doesn’t load, then it’s the first time you’re reporting the player’s progress toward it, and you must create a object to represent it. Next, set the percentage complete of the achievement using the property. You can report the player’s progress for one or more achievements to Game Center using the method. The dashboard changes the appearance of the achievements to show the current percentages. If you set the percentage of an achievement to 100, the dashboard shows it as completed. To reset the player’s progress on all achievements, use the class method.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement
type Achievement struct {
	objectivec.Object
}

// AchievementFrom constructs a [Achievement] from an unsafe.Pointer.
//
// An achievement you can award a player as they make progress toward and reach a goal in your game.
func AchievementFrom(ptr unsafe.Pointer) Achievement {
	return Achievement{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AchievementClass) Alloc() Achievement {
	rv := objc.Send[Achievement](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AchievementClass) New() Achievement {
	rv := objc.Send[Achievement](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Achievement) Init() Achievement {
	rv := objc.Send[Achievement](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Achievement) Autorelease() Achievement {
	rv := objc.Send[Achievement](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAchievement creates a new Achievement instance.
func NewAchievement() Achievement {
	return getAchievementClass().New()
}


// Provides a challenge compose view controller with preselected player identifiers and a message.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievement/challengeComposeController(withPlayers:message:completionHandler:)
func (a_ Achievement) ChallengeComposeControllerWithPlayersMessageCompletionHandler(playerIDs unsafe.Pointer, message string, completionHandler unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("challengeComposeControllerWithPlayers:message:completionHandler:"), playerIDs, objc.String(message), completionHandler)
	return rv
}



