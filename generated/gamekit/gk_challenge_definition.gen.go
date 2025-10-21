// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
	HasActiveChallengesWithCompletionHandler(completionHandler unsafe.Pointer)
}

// An object that represents the static metadata you define for the challenge.
//
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
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeDefinition/hasActiveChallenges(completionHandler:)
func (c_ ChallengeDefinition) HasActiveChallengesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("hasActiveChallengesWithCompletionHandler:"), completionHandler)
}

// The release state of the challenge definition in App Store Connect.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeDefinition/releaseState
func (c_ ChallengeDefinition) ReleaseState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("releaseState"))
	return rv
}



