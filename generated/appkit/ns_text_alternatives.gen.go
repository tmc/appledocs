// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTextAlternatives */


/* debug [class_header]: Header for NSTextAlternatives */
// The class instance for the [TextAlternatives] class.
var (
	TextAlternativesClass     _TextAlternativesClass
	TextAlternativesClassOnce sync.Once
)

func getTextAlternativesClass() _TextAlternativesClass {
	TextAlternativesClassOnce.Do(func() {
		TextAlternativesClass = _TextAlternativesClass{objc.GetClass("NSTextAlternatives")}
	})
	return TextAlternativesClass
}

type _TextAlternativesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextAlternatives */
// An interface definition for the [TextAlternatives] class.
type ITextAlternatives interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TextAlternatives */
	// properties:
	AlternativeStrings() objc.IObject /* cross-framework: NSString */
	SetAlternativeStrings(value objc.IObject /* cross-framework: NSString */)
	PrimaryString() objc.IObject /* cross-framework: NSString */
	SetPrimaryString(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextAlternatives */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextAlternatives */
// Alloc allocates a new instance without initialization.
func (tc _TextAlternativesClass) Alloc() TextAlternatives {
	rv := objc.Send[TextAlternatives](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TextAlternativesClass) New() TextAlternatives {
	rv := objc.Send[TextAlternatives](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextAlternatives) Init() TextAlternatives {
	rv := objc.Send[TextAlternatives](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextAlternatives) Autorelease() TextAlternatives {
	rv := objc.Send[TextAlternatives](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextAlternatives creates a new TextAlternatives instance.
func NewTextAlternatives() TextAlternatives {
	return getTextAlternativesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextAlternatives */
// A list of alternative strings for a piece of text.
//
// is an immutable value class that stores a list of alternatives for a piece of text and communicates the user’s selection of an alternative via a notification to your app. To support dictation, for example, you might use to present a list of alternative interpretations for a word or phrase the user speaks. If the user chooses to replace the initial interpretation with an alternative, notifies you of the choice so that you can update the text appropriately. instances are attached to attributed strings as the value of a text attribute, .


// A list of alternative strings for a piece of text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAlternatives
type TextAlternatives struct {
	objectivec.Object
}

// TextAlternativesFrom constructs a [TextAlternatives] from an unsafe.Pointer.
//
// A list of alternative strings for a piece of text.
func TextAlternativesFrom(ptr unsafe.Pointer) TextAlternatives {
	return TextAlternatives{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextAlternatives *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextAlternatives */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextAlternatives */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextAlternatives */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextAlternatives */

// An array of alternative possible interpretations that the user might select.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextalternatives/alternativestrings
func (t_ TextAlternatives) AlternativeStrings() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("alternativeStrings"))
	return rv
}/* debug [instance_properties/getter]: alternativeStrings */


// An array of alternative possible interpretations that the user might select.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextalternatives/alternativestrings
func (t_ TextAlternatives) SetAlternativeStrings(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAlternativeStrings:"), value)
}/* debug [instance_properties/setter]: alternativeStrings */


// The text that was initially chosen as the input string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextalternatives/primarystring
func (t_ TextAlternatives) PrimaryString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("primaryString"))
	return rv
}/* debug [instance_properties/getter]: primaryString */


// The text that was initially chosen as the input string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextalternatives/primarystring
func (t_ TextAlternatives) SetPrimaryString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPrimaryString:"), value)
}/* debug [instance_properties/setter]: primaryString */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextAlternatives */



