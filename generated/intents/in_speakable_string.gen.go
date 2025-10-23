// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INSpeakableString] class.
var (
	INSpeakableStringClass     _INSpeakableStringClass
	INSpeakableStringClassOnce sync.Once
)

func getINSpeakableStringClass() _INSpeakableStringClass {
	INSpeakableStringClassOnce.Do(func() {
		INSpeakableStringClass = _INSpeakableStringClass{objc.GetClass("INSpeakableString")}
	})
	return INSpeakableStringClass
}

type _INSpeakableStringClass struct {
	class objc.Class
}

// An interface definition for the [INSpeakableString] class.
type IINSpeakableString interface {
	objectivec.IObject
	// properties:
	Identifier() string
	SetIdentifier(value string)
	// methods:
}

// A custom phrase to be resolved by an Intents extension.
//
// When creating your Intents extension, you can define custom vocabulary for some types of intents. For example, a workout app may provide a custom file with the names of standard workouts that are shared by all users of the app. In places where those terms might be used, the intent object contains an object. Use the contents of an object to resolve the specified term during the handling of an intent. If the user spoke a term that is defined in your file, Siri includes the identifier of that term in the string’s property. For unrecognized terms, the identifier is .


// A custom phrase to be resolved by an Intents extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSpeakableString
type INSpeakableString struct {
	objectivec.Object
}

// INSpeakableStringFrom constructs a [INSpeakableString] from an unsafe.Pointer.
//
// A custom phrase to be resolved by an Intents extension.
func INSpeakableStringFrom(ptr unsafe.Pointer) INSpeakableString {
	return INSpeakableString{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INSpeakableStringClass) Alloc() INSpeakableString {
	rv := objc.Send[INSpeakableString](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSpeakableStringClass) New() INSpeakableString {
	rv := objc.Send[INSpeakableString](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSpeakableString) Init() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSpeakableString) Autorelease() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSpeakableString creates a new INSpeakableString instance.
func NewINSpeakableString() INSpeakableString {
	return getINSpeakableStringClass().New()
}



// The identifier associated with the string in your app’s custom vocabulary file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inspeakable/identifier
func (i_ INSpeakableString) Identifier() string {
	rv := objc.Send[string](i_.ID, objc.Sel("identifier"))
	return rv
}


// The identifier associated with the string in your app’s custom vocabulary file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/inspeakable/identifier
func (i_ INSpeakableString) SetIdentifier(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}



