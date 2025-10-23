// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INAnswerCallIntent] class.
var (
	INAnswerCallIntentClass     _INAnswerCallIntentClass
	INAnswerCallIntentClassOnce sync.Once
)

func getINAnswerCallIntentClass() _INAnswerCallIntentClass {
	INAnswerCallIntentClassOnce.Do(func() {
		INAnswerCallIntentClass = _INAnswerCallIntentClass{objc.GetClass("INAnswerCallIntent")}
	})
	return INAnswerCallIntentClass
}

type _INAnswerCallIntentClass struct {
	class objc.Class
}

// An interface definition for the [INAnswerCallIntent] class.
type IINAnswerCallIntent interface {
	IINIntent
	AudioRoute() INCallAudioRoute
	SetAudioRoute(value INCallAudioRoute)
	CallIdentifier() string
	SetCallIdentifier(value string)
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INAnswerCallIntent
type INAnswerCallIntent struct {
	INIntent
}

// INAnswerCallIntentFrom constructs a [INAnswerCallIntent] from an unsafe.Pointer.
func INAnswerCallIntentFrom(ptr unsafe.Pointer) INAnswerCallIntent {
	return INAnswerCallIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INAnswerCallIntentClass) Alloc() INAnswerCallIntent {
	rv := objc.Send[INAnswerCallIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INAnswerCallIntentClass) New() INAnswerCallIntent {
	rv := objc.Send[INAnswerCallIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INAnswerCallIntent) Init() INAnswerCallIntent {
	rv := objc.Send[INAnswerCallIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INAnswerCallIntent) Autorelease() INAnswerCallIntent {
	rv := objc.Send[INAnswerCallIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINAnswerCallIntent creates a new INAnswerCallIntent instance.
func NewINAnswerCallIntent() INAnswerCallIntent {
	return getINAnswerCallIntentClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inanswercallintent/audioroute
func (i_ INAnswerCallIntent) AudioRoute() INCallAudioRoute {
	rv := objc.Send[INCallAudioRoute](i_.ID, objc.Sel("audioRoute"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inanswercallintent/audioroute
func (i_ INAnswerCallIntent) SetAudioRoute(value INCallAudioRoute) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAudioRoute:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inanswercallintent/callidentifier
func (i_ INAnswerCallIntent) CallIdentifier() string {
	rv := objc.Send[string](i_.ID, objc.Sel("callIdentifier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inanswercallintent/callidentifier
func (i_ INAnswerCallIntent) SetCallIdentifier(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCallIdentifier:"), objc.String(value))
}



