// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Hidden() bool
	AchievedDescription() objc.IObject /* cross-framework: NSString */
	SetAchievedDescription(value objc.IObject /* cross-framework: NSString */)
	ActivityIdentifier() objc.IObject /* cross-framework: NSString */
	SetActivityIdentifier(value objc.IObject /* cross-framework: NSString */)
	ActivityProperties() objc.IObject /* cross-framework: NSString */
	SetActivityProperties(value objc.IObject /* cross-framework: NSString */)
	GroupIdentifier() objc.IObject /* cross-framework: NSString */
	SetGroupIdentifier(value objc.IObject /* cross-framework: NSString */)
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	Image() objc.IObject /* cross-framework: Image */
	SetImage(value objc.IObject /* cross-framework: Image */)
	IsHidden() bool
	SetIsHidden(value bool)
	IsReplayable() bool
	SetIsReplayable(value bool)
	MaximumPoints() int
	SetMaximumPoints(value int)
	RarityPercent() float64
	SetRarityPercent(value float64)
	ReleaseState() ReleaseState /* not a class type */
	SetReleaseState(value ReleaseState /* not a class type */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	UnachievedDescription() objc.IObject /* cross-framework: NSString */
	SetUnachievedDescription(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// An object containing the text and artwork used to present an achievement to a player.
//
// To present an achievement to the player in your interface, you can download the localized text and artwork for the achievements that you enter in App Store Connect. To get the localized text, use the class method. GameKit passes an array of objects to the completion handler that contains the text. To get the artwork for an achievement, use the method. To get standard images your game can use to present achievement progress to the player, use the and ) class methods. Alternatively, either add the access point or display the dashboard so that the player can view achievements and navigate to their other Game Center data.


// An object containing the text and artwork used to present an achievement to a player.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/incompleteAchievementImage()
func (ac _AchievementDescriptionClass) IncompleteAchievementImage() objc.IObject /* cross-framework: Image */ {
	rv := objc.Send[appkit.Image](objc.ID(ac.class), objc.Sel("incompleteAchievementImage"))
	return rv
}


// A placeholder image that you can display when the player completes the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/placeholderCompletedAchievementImage()
func (ac _AchievementDescriptionClass) PlaceholderCompletedAchievementImage() objc.IObject /* cross-framework: Image */ {
	rv := objc.Send[appkit.Image](objc.ID(ac.class), objc.Sel("placeholderCompletedAchievementImage"))
	return rv
}


// A Boolean value that states whether the achievement is initially visible to players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/isHidden
func (a_ AchievementDescription) Hidden() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hidden"))
	return rv
}


// A localized description of the achievement that you display when the player completes the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/achieveddescription
func (a_ AchievementDescription) AchievedDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("achievedDescription"))
	return rv
}


// A localized description of the achievement that you display when the player completes the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/achieveddescription
func (a_ AchievementDescription) SetAchievedDescription(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAchievedDescription:"), value)
}


// The identifier of the game activity associated with this achievement, as configured by the developer in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/activityidentifier
func (a_ AchievementDescription) ActivityIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("activityIdentifier"))
	return rv
}


// The identifier of the game activity associated with this achievement, as configured by the developer in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/activityidentifier
func (a_ AchievementDescription) SetActivityIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setActivityIdentifier:"), value)
}


// The properties when associating this achievement with a game activity, as configured by the developer in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/activityproperties
func (a_ AchievementDescription) ActivityProperties() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("activityProperties"))
	return rv
}


// The properties when associating this achievement with a game activity, as configured by the developer in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/activityproperties
func (a_ AchievementDescription) SetActivityProperties(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setActivityProperties:"), value)
}


// The identifier for the group that the achievement description is part of.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/groupidentifier
func (a_ AchievementDescription) GroupIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("groupIdentifier"))
	return rv
}


// The identifier for the group that the achievement description is part of.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/groupidentifier
func (a_ AchievementDescription) SetGroupIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setGroupIdentifier:"), value)
}


// The string you enter in App Store Connect that uniquely identifies the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/identifier
func (a_ AchievementDescription) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("identifier"))
	return rv
}


// The string you enter in App Store Connect that uniquely identifies the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/identifier
func (a_ AchievementDescription) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIdentifier:"), value)
}


// The achievement’s artwork that you display when the player completes the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/image
func (a_ AchievementDescription) Image() objc.IObject /* cross-framework: Image */ {
	rv := objc.Send[appkit.Image](a_.ID, objc.Sel("image"))
	return rv
}


// The achievement’s artwork that you display when the player completes the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/image
func (a_ AchievementDescription) SetImage(value objc.IObject /* cross-framework: Image */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setImage:"), value)
}


// A Boolean value that states whether the achievement is initially visible to players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/ishidden
func (a_ AchievementDescription) IsHidden() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isHidden"))
	return rv
}


// A Boolean value that states whether the achievement is initially visible to players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/ishidden
func (a_ AchievementDescription) SetIsHidden(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsHidden:"), value)
}


// A Boolean value that states whether the player can earn the achievement multiple times.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/isreplayable
func (a_ AchievementDescription) IsReplayable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isReplayable"))
	return rv
}


// A Boolean value that states whether the player can earn the achievement multiple times.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/isreplayable
func (a_ AchievementDescription) SetIsReplayable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsReplayable:"), value)
}


// The number of points that the player earns when completing the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/maximumpoints
func (a_ AchievementDescription) MaximumPoints() int {
	rv := objc.Send[int](a_.ID, objc.Sel("maximumPoints"))
	return rv
}


// The number of points that the player earns when completing the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/maximumpoints
func (a_ AchievementDescription) SetMaximumPoints(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMaximumPoints:"), value)
}


// The percentage of players of this game that earned the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/raritypercent-4bh6k
func (a_ AchievementDescription) RarityPercent() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("rarityPercent"))
	return rv
}


// The percentage of players of this game that earned the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/raritypercent-4bh6k
func (a_ AchievementDescription) SetRarityPercent(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRarityPercent:"), value)
}


// The release state of the achievement in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/releasestate
func (a_ AchievementDescription) ReleaseState() ReleaseState /* not a class type */ {
	rv := objc.Send[ReleaseState](a_.ID, objc.Sel("releaseState"))
	return rv
}


// The release state of the achievement in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/releasestate
func (a_ AchievementDescription) SetReleaseState(value ReleaseState /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setReleaseState:"), value)
}


// A localized title for the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/title
func (a_ AchievementDescription) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("title"))
	return rv
}


// A localized title for the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/title
func (a_ AchievementDescription) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTitle:"), value)
}


// A localized description of the achievement that you display when the player hasn’t completed the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/unachieveddescription
func (a_ AchievementDescription) UnachievedDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("unachievedDescription"))
	return rv
}


// A localized description of the achievement that you display when the player hasn’t completed the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/unachieveddescription
func (a_ AchievementDescription) SetUnachievedDescription(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUnachievedDescription:"), value)
}



