// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
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
func (ac _AchievementDescriptionClass) IncompleteAchievementImage() appkit.Image {
	rv := objc.Send[appkit.Image](objc.ID(ac.class), objc.Sel("incompleteAchievementImage"))
	return rv
}

// A placeholder image that you can display when the player completes the achievement.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/placeholderCompletedAchievementImage()
func (ac _AchievementDescriptionClass) PlaceholderCompletedAchievementImage() appkit.Image {
	rv := objc.Send[appkit.Image](objc.ID(ac.class), objc.Sel("placeholderCompletedAchievementImage"))
	return rv
}

// A localized description of the achievement that you display when the player completes the achievement.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/achieveddescription
func (a_ AchievementDescription) AchievedDescription() appkit.string {
	rv := objc.Send[appkit.string](a_.ID, objc.Sel("achievedDescription"))
	return rv
}


// SetAchievedDescription sets the value of the achievedDescription property.
// A localized description of the achievement that you display when the player completes the achievement.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/achieveddescription
func (a_ AchievementDescription) SetAchievedDescription(value appkit.string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAchievedDescription:"), value)
}

// The identifier of the game activity associated with this achievement, as configured by the developer in App Store Connect.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/activityidentifier
func (a_ AchievementDescription) ActivityIdentifier() appkit.string {
	rv := objc.Send[appkit.string](a_.ID, objc.Sel("activityIdentifier"))
	return rv
}


// SetActivityIdentifier sets the value of the activityIdentifier property.
// The identifier of the game activity associated with this achievement, as configured by the developer in App Store Connect.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/activityidentifier
func (a_ AchievementDescription) SetActivityIdentifier(value appkit.string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setActivityIdentifier:"), value)
}

// The properties when associating this achievement with a game activity, as configured by the developer in App Store Connect.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/activityproperties
func (a_ AchievementDescription) ActivityProperties() appkit.string {
	rv := objc.Send[appkit.string](a_.ID, objc.Sel("activityProperties"))
	return rv
}


// SetActivityProperties sets the value of the activityProperties property.
// The properties when associating this achievement with a game activity, as configured by the developer in App Store Connect.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/activityproperties
func (a_ AchievementDescription) SetActivityProperties(value appkit.string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setActivityProperties:"), value)
}

// The identifier for the group that the achievement description is part of.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/groupidentifier
func (a_ AchievementDescription) GroupIdentifier() appkit.string {
	rv := objc.Send[appkit.string](a_.ID, objc.Sel("groupIdentifier"))
	return rv
}


// SetGroupIdentifier sets the value of the groupIdentifier property.
// The identifier for the group that the achievement description is part of.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/groupidentifier
func (a_ AchievementDescription) SetGroupIdentifier(value appkit.string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setGroupIdentifier:"), value)
}

// The string you enter in App Store Connect that uniquely identifies the achievement.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/identifier
func (a_ AchievementDescription) Identifier() appkit.string {
	rv := objc.Send[appkit.string](a_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// The string you enter in App Store Connect that uniquely identifies the achievement.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/identifier
func (a_ AchievementDescription) SetIdentifier(value appkit.string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIdentifier:"), value)
}

// The achievement’s artwork that you display when the player completes the achievement.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/image
func (a_ AchievementDescription) Image() appkit.Image {
	rv := objc.Send[appkit.Image](a_.ID, objc.Sel("image"))
	return rv
}


// SetImage sets the value of the image property.
// The achievement’s artwork that you display when the player completes the achievement.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/image
func (a_ AchievementDescription) SetImage(value appkit.IImage) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setImage:"), value)
}

// A Boolean value that states whether the achievement is initially visible to players.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/ishidden
func (a_ AchievementDescription) IsHidden() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isHidden"))
	return rv
}


// SetIsHidden sets the value of the isHidden property.
// A Boolean value that states whether the achievement is initially visible to players.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/ishidden
func (a_ AchievementDescription) SetIsHidden(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsHidden:"), value)
}

// A Boolean value that states whether the player can earn the achievement multiple times.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/isreplayable
func (a_ AchievementDescription) IsReplayable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isReplayable"))
	return rv
}


// SetIsReplayable sets the value of the isReplayable property.
// A Boolean value that states whether the player can earn the achievement multiple times.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/isreplayable
func (a_ AchievementDescription) SetIsReplayable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsReplayable:"), value)
}

// The number of points that the player earns when completing the achievement.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/maximumpoints
func (a_ AchievementDescription) MaximumPoints() int {
	rv := objc.Send[int](a_.ID, objc.Sel("maximumPoints"))
	return rv
}


// SetMaximumPoints sets the value of the maximumPoints property.
// The number of points that the player earns when completing the achievement.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/maximumpoints
func (a_ AchievementDescription) SetMaximumPoints(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaximumPoints:"), value)
}

// The percentage of players of this game that earned the achievement.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/raritypercent-4bh6k
func (a_ AchievementDescription) RarityPercent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("rarityPercent"))
	return rv
}


// SetRarityPercent sets the value of the rarityPercent property.
// The percentage of players of this game that earned the achievement.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/raritypercent-4bh6k
func (a_ AchievementDescription) SetRarityPercent(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRarityPercent:"), value)
}

// The release state of the achievement in App Store Connect.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/releasestate
func (a_ AchievementDescription) ReleaseState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("releaseState"))
	return rv
}


// SetReleaseState sets the value of the releaseState property.
// The release state of the achievement in App Store Connect.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/releasestate
func (a_ AchievementDescription) SetReleaseState(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setReleaseState:"), value)
}

// A localized title for the achievement.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/title
func (a_ AchievementDescription) Title() appkit.string {
	rv := objc.Send[appkit.string](a_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// A localized title for the achievement.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/title
func (a_ AchievementDescription) SetTitle(value appkit.string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTitle:"), value)
}

// A localized description of the achievement that you display when the player hasn’t completed the achievement.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/unachieveddescription
func (a_ AchievementDescription) UnachievedDescription() appkit.string {
	rv := objc.Send[appkit.string](a_.ID, objc.Sel("unachievedDescription"))
	return rv
}


// SetUnachievedDescription sets the value of the unachievedDescription property.
// A localized description of the achievement that you display when the player hasn’t completed the achievement.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/unachieveddescription
func (a_ AchievementDescription) SetUnachievedDescription(value appkit.string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUnachievedDescription:"), value)
}



