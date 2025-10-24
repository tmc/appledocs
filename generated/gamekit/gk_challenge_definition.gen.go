// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKChallengeDefinition */


/* debug [class_header]: Header for GKChallengeDefinition */
// The class instance for the [ChallengeDefinition] class.
var (
	ChallengeDefinitionClass     _ChallengeDefinitionClass
	ChallengeDefinitionClassOnce sync.Once
)

func getChallengeDefinitionClass() _ChallengeDefinitionClass {
	ChallengeDefinitionClassOnce.Do(func() {
		ChallengeDefinitionClass = _ChallengeDefinitionClass{objc.GetClass("GKChallengeDefinition")}
	})
	return ChallengeDefinitionClass
}

type _ChallengeDefinitionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ChallengeDefinition */
// An interface definition for the [ChallengeDefinition] class.
type IChallengeDefinition interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ChallengeDefinition */
	// properties:
	Details() objc.IObject /* cross-framework: NSString */
	DurationOptions() []foundation.DateComponents
	GroupIdentifier() objc.IObject /* cross-framework: NSString */
	Identifier() objc.IObject /* cross-framework: NSString */
	IsRepeatable() bool
	Leaderboard() IGKLeaderboard
	ReleaseState() ReleaseState
	Title() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ChallengeDefinition */
	// methods:
	HasActiveChallengesWithCompletionHandler(completionHandler unsafe.Pointer)
	LoadImageWithCompletionHandler(completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ChallengeDefinition */
// Alloc allocates a new instance without initialization.
func (cc _ChallengeDefinitionClass) Alloc() ChallengeDefinition {
	rv := objc.Send[ChallengeDefinition](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ChallengeDefinitionClass) New() ChallengeDefinition {
	rv := objc.Send[ChallengeDefinition](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChallengeDefinition) Init() ChallengeDefinition {
	rv := objc.Send[ChallengeDefinition](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChallengeDefinition) Autorelease() ChallengeDefinition {
	rv := objc.Send[ChallengeDefinition](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChallengeDefinition creates a new ChallengeDefinition instance.
func NewChallengeDefinition() ChallengeDefinition {
	return getChallengeDefinitionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ChallengeDefinition */
// An object that represents the static metadata you define for the challenge.


// An object that represents the static metadata you define for the challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeDefinition
type ChallengeDefinition struct {
	objectivec.Object
}

// ChallengeDefinitionFrom constructs a [ChallengeDefinition] from an unsafe.Pointer.
//
// An object that represents the static metadata you define for the challenge.
func ChallengeDefinitionFrom(ptr unsafe.Pointer) ChallengeDefinition {
	return ChallengeDefinition{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ChallengeDefinition *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ChallengeDefinition */

// Loads all the challenge definitions for the current game, returns an empty array if none exist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeDefinition/loadChallengeDefinitions(completionHandler:)
func (cc _ChallengeDefinitionClass) LoadChallengeDefinitionsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("loadChallengeDefinitionsWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadChallengeDefinitionsWithCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ChallengeDefinition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ChallengeDefinition */

// Indicates if this definition has active challenges associated with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeDefinition/hasActiveChallenges(completionHandler:)
func (c_ ChallengeDefinition) HasActiveChallengesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("hasActiveChallengesWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: HasActiveChallengesWithCompletionHandler */


// Loads the image set on the challenge definition, which may be if none was set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeDefinition/loadImage(completionHandler:)
func (c_ ChallengeDefinition) LoadImageWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("loadImageWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: LoadImageWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ChallengeDefinition */

// A more detailed description of the challenge definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeDefinition/details
func (c_ ChallengeDefinition) Details() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("details"))
	return rv
}/* debug [instance_properties/getter]: details */


// The duration options for the challenge, like or .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeDefinition/durationOptions
func (c_ ChallengeDefinition) DurationOptions() []foundation.DateComponents {
	rv := objc.Send[[]foundation.DateComponents](c_.ID, objc.Sel("durationOptions"))
	return rv
}/* debug [instance_properties/getter]: durationOptions */


// The group identifier for the challenge definition, if one exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeDefinition/groupIdentifier
func (c_ ChallengeDefinition) GroupIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("groupIdentifier"))
	return rv
}/* debug [instance_properties/getter]: groupIdentifier */


// The developer defined identifier for a given challenge definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeDefinition/identifier
func (c_ ChallengeDefinition) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// Indicates if a challenge can be attempted more than once.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeDefinition/isRepeatable
func (c_ ChallengeDefinition) IsRepeatable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isRepeatable"))
	return rv
}/* debug [instance_properties/getter]: isRepeatable */


// Scores submitted to this leaderboard will also be submitted as scores in this challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeDefinition/leaderboard
func (c_ ChallengeDefinition) Leaderboard() IGKLeaderboard {
	rv := objc.Send[Leaderboard](c_.ID, objc.Sel("leaderboard"))
	return rv
}/* debug [instance_properties/getter]: leaderboard */


// The release state of the challenge definition in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeDefinition/releaseState
func (c_ ChallengeDefinition) ReleaseState() ReleaseState {
	rv := objc.Send[ReleaseState](c_.ID, objc.Sel("releaseState"))
	return rv
}/* debug [instance_properties/getter]: releaseState */


// A short title for the challenge definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeDefinition/title
func (c_ ChallengeDefinition) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKChallengeDefinition */



