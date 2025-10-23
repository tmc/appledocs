// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	DateInterval() foundation.objc.IObject /* cross-framework: DateInterval */
	SetDateInterval(value foundation.objc.IObject /* cross-framework: DateInterval */)
	Direction() unsafe.Pointer
	SetDirection(value unsafe.Pointer)
	GroupIdentifier() string /* primitive/slice/pointer. */
	SetGroupIdentifier(value string /* primitive/slice/pointer. */)
	Identifier() string /* primitive/slice/pointer. */
	SetIdentifier(value string /* primitive/slice/pointer. */)
	Intent() INIntent /* already interface */
	SetIntent(value INIntent /* already interface */)
	IntentHandlingStatus() unsafe.Pointer
	SetIntentHandlingStatus(value unsafe.Pointer)
	IntentResponse() INIntentResponse /* already interface */
	SetIntentResponse(value INIntentResponse /* already interface */)
	// methods:
	DonateInteractionWithCompletion(completion unsafe.Pointer)
	ParameterValueForParameter(parameter INParameter /* already interface */) objc.ID
}

// An interaction between the user and your app involving an intent object.
//
// An object encapsulates information about a SiriKit request and your app’s response. SiriKit creates interaction objects automatically when it needs your app to respond to a specific intent, either by handling the intent or providing an error explaining why your app couldn’t handle the intent. SiriKit places the interaction in an object that the system passes to your app at launch time. You can also create instances of this class in your app and donate relevant interactions to the system. Donating interactions provides contextual information that might be helpful to other apps. Some system apps use donated interactions to improve search results or to anticipate user actions. For example, a ride-booking app could donate an interaction containing the user’s planned ride information. If the user subsequently uses the Maps app to search for restaurants, Maps can show relevant results near the user’s destination. You choose which of your app’s interactions you want to donate to the system. To donate an interaction, create an instance of this class, filling it with your intent object and response, and call the method. You can also use the methods of this class to delete interactions when they are no longer relevant.


// An interaction between the user and your app involving an intent object.
//
// [Full Topic]
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



// Deletes the specified interactions that were donated by the calling app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INInteraction/delete(with:completion:)-2d1gs
func (ic _INInteractionClass) DeleteInteractionsWithIdentifiersCompletion(identifiers []string /* primitive/slice/pointer. */, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ic.class), objc.Sel("deleteInteractionsWithIdentifiers:completion:"), identifiers, completion)
}


// Donates this interaction object to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INInteraction/donate(completion:)
func (i_ INInteraction) DonateInteractionWithCompletion(completion unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("donateInteractionWithCompletion:"), completion)
}


// Returns the value of the specified parameter of this interaction object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INInteraction/parameterValue(for:)
func (i_ INInteraction) ParameterValueForParameter(parameter INParameter /* already interface */) objc.ID {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("parameterValueForParameter:"), parameter)
	return rv
}


// The time at which the interaction started and its duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/ininteraction/dateinterval
func (i_ INInteraction) DateInterval() foundation.objc.IObject /* cross-framework: DateInterval */ {
	rv := objc.Send[foundation.DateInterval](i_.ID, objc.Sel("dateInterval"))
	return rv
}


// The time at which the interaction started and its duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/ininteraction/dateinterval
func (i_ INInteraction) SetDateInterval(value foundation.objc.IObject /* cross-framework: DateInterval */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDateInterval:"), value)
}


// The direction in which information flowed to or from the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/ininteraction/direction
func (i_ INInteraction) Direction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("direction"))
	return rv
}


// The direction in which information flowed to or from the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/ininteraction/direction
func (i_ INInteraction) SetDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDirection:"), value)
}


// The unique identifier of the interaction’s group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/ininteraction/groupidentifier
func (i_ INInteraction) GroupIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](i_.ID, objc.Sel("groupIdentifier"))
	return rv
}


// The unique identifier of the interaction’s group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/ininteraction/groupidentifier
func (i_ INInteraction) SetGroupIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGroupIdentifier:"), objc.String(value))
}


// The unique identifier of the interaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/ininteraction/identifier
func (i_ INInteraction) Identifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](i_.ID, objc.Sel("identifier"))
	return rv
}


// The unique identifier of the interaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/ininteraction/identifier
func (i_ INInteraction) SetIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// The intent object that describes the user’s request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/ininteraction/intent
func (i_ INInteraction) Intent() INIntent /* already interface */ {
	rv := objc.Send[INIntent](i_.ID, objc.Sel("intent"))
	return rv
}


// The intent object that describes the user’s request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/ininteraction/intent
func (i_ INInteraction) SetIntent(value INIntent /* already interface */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIntent:"), value)
}


// The current state of the interaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/ininteraction/intenthandlingstatus
func (i_ INInteraction) IntentHandlingStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("intentHandlingStatus"))
	return rv
}


// The current state of the interaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/ininteraction/intenthandlingstatus
func (i_ INInteraction) SetIntentHandlingStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIntentHandlingStatus:"), value)
}


// The response object that your app created in response to the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/ininteraction/intentresponse
func (i_ INInteraction) IntentResponse() INIntentResponse /* already interface */ {
	rv := objc.Send[INIntentResponse](i_.ID, objc.Sel("intentResponse"))
	return rv
}


// The response object that your app created in response to the request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/ininteraction/intentresponse
func (i_ INInteraction) SetIntentResponse(value INIntentResponse /* already interface */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIntentResponse:"), value)
}



