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

/* debug [class.gen.go]: Generating class GKAchievementDescription */


/* debug [class_header]: Header for GKAchievementDescription */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AchievementDescription */
// An interface definition for the [AchievementDescription] class.
type IAchievementDescription interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AchievementDescription */
	// properties:
	AchievedDescription() objc.IObject /* cross-framework: NSString */
	ActivityIdentifier() objc.IObject /* cross-framework: NSString */
	ActivityProperties() foundation.IDictionary
	GroupIdentifier() objc.IObject /* cross-framework: NSString */
	Identifier() objc.IObject /* cross-framework: NSString */
	Image() appkit.Image
	Hidden() bool
	Replayable() bool
	MaximumPoints() int
	RarityPercent() objc.IObject /* cross-framework: NSNumber */
	ReleaseState() ReleaseState
	Title() objc.IObject /* cross-framework: NSString */
	UnachievedDescription() objc.IObject /* cross-framework: NSString */
	IsHidden() bool
	SetIsHidden(value bool)
	IsReplayable() bool
	SetIsReplayable(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AchievementDescription */
	// methods:
	LoadImageWithCompletionHandler(completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AchievementDescription */
// Alloc allocates a new instance without initialization.
func (ac _AchievementDescriptionClass) Alloc() AchievementDescription {
	rv := objc.Send[AchievementDescription](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AchievementDescription */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AchievementDescription *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AchievementDescription */

// A common image that you can display when the player hasn’t completed the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/incompleteAchievementImage()
func (ac _AchievementDescriptionClass) IncompleteAchievementImage() appkit.Image {
	rv := objc.Send[appkit.Image](objc.ID(ac.class), objc.Sel("incompleteAchievementImage"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IncompleteAchievementImage) */


// Downloads the localized descriptions of achievements from Game Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/loadAchievementDescriptions(completionHandler:)
func (ac _AchievementDescriptionClass) LoadAchievementDescriptionsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("loadAchievementDescriptionsWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadAchievementDescriptionsWithCompletionHandler) */


// A placeholder image that you can display when the player completes the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/placeholderCompletedAchievementImage()
func (ac _AchievementDescriptionClass) PlaceholderCompletedAchievementImage() appkit.Image {
	rv := objc.Send[appkit.Image](objc.ID(ac.class), objc.Sel("placeholderCompletedAchievementImage"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PlaceholderCompletedAchievementImage) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AchievementDescription */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AchievementDescription */

// Loads the image to display when the player completes the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/loadImage(completionHandler:)
func (a_ AchievementDescription) LoadImageWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("loadImageWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: LoadImageWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AchievementDescription */

// A localized description of the achievement that you display when the player completes the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/achievedDescription
func (a_ AchievementDescription) AchievedDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("achievedDescription"))
	return rv
}/* debug [instance_properties/getter]: achievedDescription */


// The identifier of the game activity associated with this achievement, as configured by the developer in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/activityIdentifier
func (a_ AchievementDescription) ActivityIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("activityIdentifier"))
	return rv
}/* debug [instance_properties/getter]: activityIdentifier */


// The properties when associating this achievement with a game activity, as configured by the developer in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/activityProperties
func (a_ AchievementDescription) ActivityProperties() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("activityProperties"))
	return rv
}/* debug [instance_properties/getter]: activityProperties */


// The identifier for the group that the achievement description is part of.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/groupIdentifier
func (a_ AchievementDescription) GroupIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("groupIdentifier"))
	return rv
}/* debug [instance_properties/getter]: groupIdentifier */


// The string you enter in App Store Connect that uniquely identifies the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/identifier
func (a_ AchievementDescription) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The achievement’s artwork that you display when the player completes the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/image
func (a_ AchievementDescription) Image() appkit.Image {
	rv := objc.Send[appkit.Image](a_.ID, objc.Sel("image"))
	return rv
}/* debug [instance_properties/getter]: image */


// A Boolean value that states whether the achievement is initially visible to players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/isHidden
func (a_ AchievementDescription) Hidden() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("hidden"))
	return rv
}/* debug [instance_properties/getter]: hidden */


// A Boolean value that states whether the player can earn the achievement multiple times.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/isReplayable
func (a_ AchievementDescription) Replayable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("replayable"))
	return rv
}/* debug [instance_properties/getter]: replayable */


// The number of points that the player earns when completing the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/maximumPoints
func (a_ AchievementDescription) MaximumPoints() int {
	rv := objc.Send[int](a_.ID, objc.Sel("maximumPoints"))
	return rv
}/* debug [instance_properties/getter]: maximumPoints */


// The percentage of players of this game that earned the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/rarityPercent-3zqw6
func (a_ AchievementDescription) RarityPercent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](a_.ID, objc.Sel("rarityPercent"))
	return rv
}/* debug [instance_properties/getter]: rarityPercent */


// The release state of the achievement in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/releaseState
func (a_ AchievementDescription) ReleaseState() ReleaseState {
	rv := objc.Send[ReleaseState](a_.ID, objc.Sel("releaseState"))
	return rv
}/* debug [instance_properties/getter]: releaseState */


// A localized title for the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/title
func (a_ AchievementDescription) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// A localized description of the achievement that you display when the player hasn’t completed the achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementDescription/unachievedDescription
func (a_ AchievementDescription) UnachievedDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("unachievedDescription"))
	return rv
}/* debug [instance_properties/getter]: unachievedDescription */


// A Boolean value that states whether the achievement is initially visible to players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/ishidden
func (a_ AchievementDescription) IsHidden() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isHidden"))
	return rv
}/* debug [instance_properties/getter]: isHidden */


// A Boolean value that states whether the achievement is initially visible to players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/ishidden
func (a_ AchievementDescription) SetIsHidden(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsHidden:"), value)
}/* debug [instance_properties/setter]: isHidden */


// A Boolean value that states whether the player can earn the achievement multiple times.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/isreplayable
func (a_ AchievementDescription) IsReplayable() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isReplayable"))
	return rv
}/* debug [instance_properties/getter]: isReplayable */


// A Boolean value that states whether the player can earn the achievement multiple times.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkachievementdescription/isreplayable
func (a_ AchievementDescription) SetIsReplayable(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsReplayable:"), value)
}/* debug [instance_properties/setter]: isReplayable */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKAchievementDescription */



