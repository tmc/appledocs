// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INObject] class.
var (
	INObjectClass     _INObjectClass
	INObjectClassOnce sync.Once
)

func getINObjectClass() _INObjectClass {
	INObjectClassOnce.Do(func() {
		INObjectClass = _INObjectClass{objc.GetClass("INObject")}
	})
	return INObjectClass
}

type _INObjectClass struct {
	class objc.Class
}

// An interface definition for the [INObject] class.
type IINObject interface {
	objectivec.IObject
}

// A representation of a custom intent parameter or response property.
//
// Use to create custom parameters and response properties for intent data that doesn’t fit into one of the System Types, such as Boolean, Duration, or Location. Define custom types and associate them with your custom intents and responses in the Intent Definition file. Xcode uses the type defined in the Intent Definition file to generate a subclass of . Create an instance of this subclass to structure data in intents and intent responses.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INObject
type INObject struct {
	objectivec.Object
}

// INObjectFrom constructs a [INObject] from an unsafe.Pointer.
//
// A representation of a custom intent parameter or response property.
func INObjectFrom(ptr unsafe.Pointer) INObject {
	return INObject{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INObjectClass) Alloc() INObject {
	rv := objc.Send[INObject](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INObjectClass) New() INObject {
	rv := objc.Send[INObject](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INObject) Init() INObject {
	rv := objc.Send[INObject](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INObject) Autorelease() INObject {
	rv := objc.Send[INObject](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINObject creates a new INObject instance.
func NewINObject() INObject {
	return getINObjectClass().New()
}


// Additional details about the custom intent object.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inobject/subtitlestring
func (i_ INObject) SubtitleString() string {
	rv := objc.Send[string](i_.ID, objc.Sel("subtitleString"))
	return rv
}


// SetSubtitleString sets the value of the subtitleString property.
// Additional details about the custom intent object.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inobject/subtitlestring
func (i_ INObject) SetSubtitleString(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSubtitleString:"), objc.String(value))
}

// A string that identifies the custom intent object within your app.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inobject/identifier
func (i_ INObject) Identifier() string {
	rv := objc.Send[string](i_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// A string that identifies the custom intent object within your app.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inobject/identifier
func (i_ INObject) SetIdentifier(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}

// An array of alternative speakable strings that identify the object.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inobject/alternativespeakablematches
func (i_ INObject) AlternativeSpeakableMatches() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("alternativeSpeakableMatches"))
	return rv
}


// SetAlternativeSpeakableMatches sets the value of the alternativeSpeakableMatches property.
// An array of alternative speakable strings that identify the object.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inobject/alternativespeakablematches
func (i_ INObject) SetAlternativeSpeakableMatches(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAlternativeSpeakableMatches:"), value)
}

// A name or description for the custom intent object.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inobject/displaystring
func (i_ INObject) DisplayString() string {
	rv := objc.Send[string](i_.ID, objc.Sel("displayString"))
	return rv
}


// SetDisplayString sets the value of the displayString property.
// A name or description for the custom intent object.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inobject/displaystring
func (i_ INObject) SetDisplayString(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDisplayString:"), objc.String(value))
}

// An image to display alongside the custom intent object’s text.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inobject/displayimage
func (i_ INObject) DisplayImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("displayImage"))
	return rv
}


// SetDisplayImage sets the value of the displayImage property.
// An image to display alongside the custom intent object’s text.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inobject/displayimage
func (i_ INObject) SetDisplayImage(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDisplayImage:"), value)
}

// A hint that describes how to pronounce the custom intent object.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inobject/pronunciationhint
func (i_ INObject) PronunciationHint() string {
	rv := objc.Send[string](i_.ID, objc.Sel("pronunciationHint"))
	return rv
}


// SetPronunciationHint sets the value of the pronunciationHint property.
// A hint that describes how to pronounce the custom intent object.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inobject/pronunciationhint
func (i_ INObject) SetPronunciationHint(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPronunciationHint:"), objc.String(value))
}



