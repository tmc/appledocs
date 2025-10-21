// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [INInteraction] class.
var (
	INInteractionClass     _INInteractionClass
	INInteractionClassOnce sync.Once
)

func getINInteractionClass() _INInteractionClass {
	INInteractionClassOnce.Do(func() {
		INInteractionClass = _INInteractionClass{objc.GetClass("INInteraction")}
	})
	return INInteractionClass
}

type _INInteractionClass struct {
	class objc.Class
}

// An interface definition for the [INInteraction] class.
type IINInteraction interface {
	objectivec.IObject
	DonateInteractionWithCompletion(completion unsafe.Pointer)
	ParameterValueForParameter(parameter unsafe.Pointer) objc.ID
}

// An interaction between the user and your app involving an intent object.
//
// An object encapsulates information about a SiriKit request and your app’s response. SiriKit creates interaction objects automatically when it needs your app to respond to a specific intent, either by handling the intent or providing an error explaining why your app couldn’t handle the intent. SiriKit places the interaction in an object that the system passes to your app at launch time. You can also create instances of this class in your app and donate relevant interactions to the system. Donating interactions provides contextual information that might be helpful to other apps. Some system apps use donated interactions to improve search results or to anticipate user actions. For example, a ride-booking app could donate an interaction containing the user’s planned ride information. If the user subsequently uses the Maps app to search for restaurants, Maps can show relevant results near the user’s destination. You choose which of your app’s interactions you want to donate to the system. To donate an interaction, create an instance of this class, filling it with your intent object and response, and call the method. You can also use the methods of this class to delete interactions when they are no longer relevant.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INInteraction
type INInteraction struct {
	objectivec.Object
}

// INInteractionFrom constructs a [INInteraction] from an unsafe.Pointer.
//
// An interaction between the user and your app involving an intent object.
func INInteractionFrom(ptr unsafe.Pointer) INInteraction {
	return INInteraction{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INInteractionClass) Alloc() INInteraction {
	rv := objc.Send[INInteraction](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INInteractionClass) New() INInteraction {
	rv := objc.Send[INInteraction](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INInteraction) Init() INInteraction {
	rv := objc.Send[INInteraction](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INInteraction) Autorelease() INInteraction {
	rv := objc.Send[INInteraction](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINInteraction creates a new INInteraction instance.
func NewINInteraction() INInteraction {
	return getINInteractionClass().New()
}


// Deletes the interactions with the specified group identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INInteraction/delete(with:completion:)-tcq9
func (ic _INInteractionClass) DeleteInteractionsWithGroupIdentifierCompletion(groupIdentifier string, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ic.class), objc.Sel("deleteInteractionsWithGroupIdentifier:completion:"), objc.String(groupIdentifier), completion)
}

// Donates this interaction object to the system.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INInteraction/donate(completion:)
func (i_ INInteraction) DonateInteractionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("donateInteractionWithCompletion:"), completion)
}

// Returns the value of the specified parameter of this interaction object.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INInteraction/parameterValue(for:)
func (i_ INInteraction) ParameterValueForParameter(parameter unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("parameterValueForParameter:"), parameter)
	return rv
}

// The unique identifier of the interaction’s group.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INInteraction/groupIdentifier
func (i_ INInteraction) GroupIdentifier() string {
	rv := objc.Send[string](i_.ID, objc.Sel("groupIdentifier"))
	return rv
}


// SetGroupIdentifier sets the value of the groupIdentifier property.
// The unique identifier of the interaction’s group.

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INInteraction/groupIdentifier
func (i_ INInteraction) SetGroupIdentifier(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGroupIdentifier:"), objc.String(value))
}
// The unique identifier of the interaction.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INInteraction/identifier
func (i_ INInteraction) Identifier() string {
	rv := objc.Send[string](i_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// The unique identifier of the interaction.

//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INInteraction/identifier
func (i_ INInteraction) SetIdentifier(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}
// The current state of the interaction.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INInteraction/intentHandlingStatus
func (i_ INInteraction) IntentHandlingStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("intentHandlingStatus"))
	return rv
}



