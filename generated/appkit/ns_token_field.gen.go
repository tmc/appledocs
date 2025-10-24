// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NSTokenField */


/* debug [class_header]: Header for NSTokenField */
// The class instance for the [TokenField] class.
var (
	TokenFieldClass     _TokenFieldClass
	TokenFieldClassOnce sync.Once
)

func getTokenFieldClass() _TokenFieldClass {
	TokenFieldClassOnce.Do(func() {
		TokenFieldClass = _TokenFieldClass{objc.GetClass("NSTokenField")}
	})
	return TokenFieldClass
}

type _TokenFieldClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TokenField */
// An interface definition for the [TokenField] class.
type ITokenField interface {
	ITextField
	
/* debug [class_interface_properties]: Properties for TokenField */
	// properties:
	CompletionDelay() float64
	SetCompletionDelay(value float64)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	TokenStyle() TokenStyle
	SetTokenStyle(value TokenStyle)
	TokenizingCharacterSet() foundation.CharacterSet
	SetTokenizingCharacterSet(value foundation.CharacterSet)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TokenField */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TokenField */
// Alloc allocates a new instance without initialization.
func (tc _TokenFieldClass) Alloc() TokenField {
	rv := objc.Send[TokenField](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TokenFieldClass) New() TokenField {
	rv := objc.Send[TokenField](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TokenField) Init() TokenField {
	rv := objc.Send[TokenField](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TokenField) Autorelease() TokenField {
	rv := objc.Send[TokenField](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTokenField creates a new TokenField instance.
func NewTokenField() TokenField {
	return getTokenFieldClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TokenField */
// A text field that converts text into visually distinct tokens.
//
// Use a token field when you want typed text to be transformed into “tokens”, which are visually distinct elements in the text field interface. For example, you might use a token field in a mail app to display email addresses for individual users. The distinct appearance of tokens makes them easy for users to distinguish from surrounding text. uses an to implement much of the control’s functionality. provides cover methods for most methods of , which invoke the corresponding cell method.


// A text field that converts text into visually distinct tokens.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField
type TokenField struct {
	TextField
}

// TokenFieldFrom constructs a [TokenField] from an unsafe.Pointer.
//
// A text field that converts text into visually distinct tokens.
func TokenFieldFrom(ptr unsafe.Pointer) TokenField {
	return TokenField{
		TextField: TextFieldFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TokenField *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TokenField */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TokenField */

// Returns the default completion delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/defaultCompletionDelay
func (tc _TokenFieldClass) DefaultCompletionDelay() float64 {
	rv := objc.Send[float64](objc.ID(tc.class), objc.Sel("defaultCompletionDelay"))
	return rv
}/* debug [class_properties_class/property]: defaultCompletionDelay */

// Returns the default tokenizing character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/defaultTokenizingCharacterSet
func (tc _TokenFieldClass) DefaultTokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.Send[foundation.CharacterSet](objc.ID(tc.class), objc.Sel("defaultTokenizingCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: defaultTokenizingCharacterSet */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TokenField */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TokenField */

// The receiver’s completion delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/completionDelay
func (t_ TokenField) CompletionDelay() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("completionDelay"))
	return rv
}/* debug [instance_properties/getter]: completionDelay */


// The receiver’s completion delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/completionDelay
func (t_ TokenField) SetCompletionDelay(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCompletionDelay:"), value)
}/* debug [instance_properties/setter]: completionDelay */


// Returns the default completion delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/defaultCompletionDelay
func (t_ TokenField) DefaultCompletionDelay() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("defaultCompletionDelay"))
	return rv
}/* debug [instance_properties/getter]: defaultCompletionDelay */


// Returns the default tokenizing character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/defaultTokenizingCharacterSet
func (t_ TokenField) DefaultTokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.Send[foundation.CharacterSet](t_.ID, objc.Sel("defaultTokenizingCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: defaultTokenizingCharacterSet */


// Returns the token field’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/delegate
func (t_ TokenField) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// Returns the token field’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/delegate
func (t_ TokenField) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The token style of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/tokenStyle-swift.property
func (t_ TokenField) TokenStyle() TokenStyle {
	rv := objc.Send[TokenStyle](t_.ID, objc.Sel("tokenStyle"))
	return rv
}/* debug [instance_properties/getter]: tokenStyle */


// The token style of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/tokenStyle-swift.property
func (t_ TokenField) SetTokenStyle(value TokenStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTokenStyle:"), value)
}/* debug [instance_properties/setter]: tokenStyle */


// The recevier’s tokenizing character set to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/tokenizingCharacterSet
func (t_ TokenField) TokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.Send[foundation.CharacterSet](t_.ID, objc.Sel("tokenizingCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: tokenizingCharacterSet */


// The recevier’s tokenizing character set to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/tokenizingCharacterSet
func (t_ TokenField) SetTokenizingCharacterSet(value foundation.CharacterSet) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTokenizingCharacterSet:"), value)
}/* debug [instance_properties/setter]: tokenizingCharacterSet */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTokenField */



