// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AchievementDescription] class.
var (
	AchievementDescriptionClass     _AchievementDescriptionClass
	AchievementDescriptionClassOnce sync.Once
)

func getAchievementDescriptionClass() _AchievementDescriptionClass {
	AchievementDescriptionClassOnce.Do(func() {
		AchievementDescriptionClass = _AchievementDescriptionClass{objc.GetClass("GKAchievementDescription")}
	})
	return AchievementDescriptionClass
}

type _AchievementDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [AchievementDescription] class.
type IAchievementDescription interface {
	objectivec.IObject
}

// An object containing the text and artwork used to present an achievement to a player.
//
// To present an achievement to the player in your interface, you can download the localized text and artwork for the achievements that you enter in App Store Connect. To get the localized text, use the class method. GameKit passes an array of objects to the completion handler that contains the text. To get the artwork for an achievement, use the method. To get standard images your game can use to present achievement progress to the player, use the and ) class methods. Alternatively, either add the access point or display the dashboard so that the player can view achievements and navigate to their other Game Center data.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription
type AchievementDescription struct {
	objectivec.Object
}

// AchievementDescriptionFrom constructs a [AchievementDescription] from an unsafe.Pointer.
//
// An object containing the text and artwork used to present an achievement to a player.
func AchievementDescriptionFrom(ptr unsafe.Pointer) AchievementDescription {
	return AchievementDescription{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AchievementDescriptionClass) Alloc() AchievementDescription {
	rv := objc.Send[AchievementDescription](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AchievementDescriptionClass) New() AchievementDescription {
	rv := objc.Send[AchievementDescription](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AchievementDescription) Init() AchievementDescription {
	rv := objc.Send[AchievementDescription](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AchievementDescription) Autorelease() AchievementDescription {
	rv := objc.Send[AchievementDescription](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAchievementDescription creates a new AchievementDescription instance.
func NewAchievementDescription() AchievementDescription {
	return getAchievementDescriptionClass().New()
}


// A common image that you can display when the player hasn’t completed the achievement.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/incompleteAchievementImage()
func (ac _AchievementDescriptionClass) IncompleteAchievementImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("incompleteAchievementImage"))
	return rv
}

// A placeholder image that you can display when the player completes the achievement.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/placeholderCompletedAchievementImage()
func (ac _AchievementDescriptionClass) PlaceholderCompletedAchievementImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("placeholderCompletedAchievementImage"))
	return rv
}



