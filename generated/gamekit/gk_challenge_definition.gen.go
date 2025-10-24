// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ChallengeDefinition] class.
type IChallengeDefinition interface {
	objectivec.IObject
	// properties:
	DurationOptions() []objc.IObject /* cross-framework: DateComponents */
	Details() objc.IObject /* cross-framework: NSString */
	SetDetails(value objc.IObject /* cross-framework: NSString */)
	GroupIdentifier() objc.IObject /* cross-framework: NSString */
	SetGroupIdentifier(value objc.IObject /* cross-framework: NSString */)
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	IsRepeatable() bool
	SetIsRepeatable(value bool)
	Leaderboard() IGKLeaderboard
	SetLeaderboard(value IGKLeaderboard)
	ReleaseState() ReleaseState /* not a class type */
	SetReleaseState(value ReleaseState /* not a class type */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	// methods:
	HasActiveChallengesWithCompletionHandler(completionHandler unsafe.Pointer)
}

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

// Alloc allocates a new instance without initialization.
func (cc _ChallengeDefinitionClass) Alloc() ChallengeDefinition {
	rv := objc.Send[ChallengeDefinition](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Indicates if this definition has active challenges associated with it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeDefinition/hasActiveChallenges(completionHandler:)
func (c_ ChallengeDefinition) HasActiveChallengesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("hasActiveChallengesWithCompletionHandler:"), completionHandler)
}


// The duration options for the challenge, like or .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeDefinition/durationOptions
func (c_ ChallengeDefinition) DurationOptions() []objc.IObject /* cross-framework: DateComponents */ {
	rv := objc.Send[[]foundation.DateComponents](c_.ID, objc.Sel("durationOptions"))
	return rv
}


// A more detailed description of the challenge definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkchallengedefinition/details
func (c_ ChallengeDefinition) Details() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("details"))
	return rv
}


// A more detailed description of the challenge definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkchallengedefinition/details
func (c_ ChallengeDefinition) SetDetails(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDetails:"), value)
}


// The group identifier for the challenge definition, if one exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkchallengedefinition/groupidentifier
func (c_ ChallengeDefinition) GroupIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("groupIdentifier"))
	return rv
}


// The group identifier for the challenge definition, if one exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkchallengedefinition/groupidentifier
func (c_ ChallengeDefinition) SetGroupIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGroupIdentifier:"), value)
}


// The developer defined identifier for a given challenge definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkchallengedefinition/identifier
func (c_ ChallengeDefinition) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("identifier"))
	return rv
}


// The developer defined identifier for a given challenge definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkchallengedefinition/identifier
func (c_ ChallengeDefinition) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIdentifier:"), value)
}


// Indicates if a challenge can be attempted more than once.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkchallengedefinition/isrepeatable
func (c_ ChallengeDefinition) IsRepeatable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isRepeatable"))
	return rv
}


// Indicates if a challenge can be attempted more than once.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkchallengedefinition/isrepeatable
func (c_ ChallengeDefinition) SetIsRepeatable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsRepeatable:"), value)
}


// Scores submitted to this leaderboard will also be submitted as scores in this challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkchallengedefinition/leaderboard
func (c_ ChallengeDefinition) Leaderboard() IGKLeaderboard {
	rv := objc.Send[Leaderboard](c_.ID, objc.Sel("leaderboard"))
	return rv
}


// Scores submitted to this leaderboard will also be submitted as scores in this challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkchallengedefinition/leaderboard
func (c_ ChallengeDefinition) SetLeaderboard(value IGKLeaderboard) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLeaderboard:"), value)
}


// The release state of the challenge definition in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkchallengedefinition/releasestate
func (c_ ChallengeDefinition) ReleaseState() ReleaseState /* not a class type */ {
	rv := objc.Send[ReleaseState](c_.ID, objc.Sel("releaseState"))
	return rv
}


// The release state of the challenge definition in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkchallengedefinition/releasestate
func (c_ ChallengeDefinition) SetReleaseState(value ReleaseState /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setReleaseState:"), value)
}


// A short title for the challenge definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkchallengedefinition/title
func (c_ ChallengeDefinition) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("title"))
	return rv
}


// A short title for the challenge definition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkchallengedefinition/title
func (c_ ChallengeDefinition) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitle:"), value)
}



