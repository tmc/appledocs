// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NSTokenFieldCell */


/* debug [class_header]: Header for NSTokenFieldCell */
// The class instance for the [TokenFieldCell] class.
var (
	TokenFieldCellClass     _TokenFieldCellClass
	TokenFieldCellClassOnce sync.Once
)

func getTokenFieldCellClass() _TokenFieldCellClass {
	TokenFieldCellClassOnce.Do(func() {
		TokenFieldCellClass = _TokenFieldCellClass{objc.GetClass("NSTokenFieldCell")}
	})
	return TokenFieldCellClass
}

type _TokenFieldCellClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TokenFieldCell */
// An interface definition for the [TokenFieldCell] class.
type ITokenFieldCell interface {
	ITextFieldCell
	
/* debug [class_interface_properties]: Properties for TokenFieldCell */
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

	
/* debug [class_interface_methods]: Methods for TokenFieldCell */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TokenFieldCell */
// Alloc allocates a new instance without initialization.
func (tc _TokenFieldCellClass) Alloc() TokenFieldCell {
	rv := objc.Send[TokenFieldCell](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TokenFieldCellClass) New() TokenFieldCell {
	rv := objc.Send[TokenFieldCell](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TokenFieldCell) Init() TokenFieldCell {
	rv := objc.Send[TokenFieldCell](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TokenFieldCell) Autorelease() TokenFieldCell {
	rv := objc.Send[TokenFieldCell](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTokenFieldCell creates a new TokenFieldCell instance.
func NewTokenFieldCell() TokenFieldCell {
	return getTokenFieldCellClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TokenFieldCell */
// A text field cell subclass that enables tokenized editing of an array of objects.
//
// is a subclass of that provides tokenized editing of an array of objects similar to the address field in the Mail app. The objects may be strings or objects that can be represented as strings. A single token field cell can be presented in an control.


// A text field cell subclass that enables tokenized editing of an array of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell
type TokenFieldCell struct {
	TextFieldCell
}

// TokenFieldCellFrom constructs a [TokenFieldCell] from an unsafe.Pointer.
//
// A text field cell subclass that enables tokenized editing of an array of objects.
func TokenFieldCellFrom(ptr unsafe.Pointer) TokenFieldCell {
	return TokenFieldCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TokenFieldCell *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TokenFieldCell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TokenFieldCell */

// Returns the default completion delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/defaultCompletionDelay
func (tc _TokenFieldCellClass) DefaultCompletionDelay() float64 {
	rv := objc.Send[float64](objc.ID(tc.class), objc.Sel("defaultCompletionDelay"))
	return rv
}/* debug [class_properties_class/property]: defaultCompletionDelay */

// Returns the default tokenizing character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/defaultTokenizingCharacterSet
func (tc _TokenFieldCellClass) DefaultTokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.Send[foundation.CharacterSet](objc.ID(tc.class), objc.Sel("defaultTokenizingCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: defaultTokenizingCharacterSet */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TokenFieldCell */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TokenFieldCell */

// The receiver’s completion delay to a given delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/completionDelay
func (t_ TokenFieldCell) CompletionDelay() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("completionDelay"))
	return rv
}/* debug [instance_properties/getter]: completionDelay */


// The receiver’s completion delay to a given delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/completionDelay
func (t_ TokenFieldCell) SetCompletionDelay(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCompletionDelay:"), value)
}/* debug [instance_properties/setter]: completionDelay */


// Returns the default completion delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/defaultCompletionDelay
func (t_ TokenFieldCell) DefaultCompletionDelay() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("defaultCompletionDelay"))
	return rv
}/* debug [instance_properties/getter]: defaultCompletionDelay */


// Returns the default tokenizing character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/defaultTokenizingCharacterSet
func (t_ TokenFieldCell) DefaultTokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.Send[foundation.CharacterSet](t_.ID, objc.Sel("defaultTokenizingCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: defaultTokenizingCharacterSet */


// The receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/delegate
func (t_ TokenFieldCell) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The receiver’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/delegate
func (t_ TokenFieldCell) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The token style of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/tokenStyle
func (t_ TokenFieldCell) TokenStyle() TokenStyle {
	rv := objc.Send[TokenStyle](t_.ID, objc.Sel("tokenStyle"))
	return rv
}/* debug [instance_properties/getter]: tokenStyle */


// The token style of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/tokenStyle
func (t_ TokenFieldCell) SetTokenStyle(value TokenStyle) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTokenStyle:"), value)
}/* debug [instance_properties/setter]: tokenStyle */


// The receiver’s tokenizing character set to a given character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/tokenizingCharacterSet
func (t_ TokenFieldCell) TokenizingCharacterSet() foundation.CharacterSet {
	rv := objc.Send[foundation.CharacterSet](t_.ID, objc.Sel("tokenizingCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: tokenizingCharacterSet */


// The receiver’s tokenizing character set to a given character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/tokenizingCharacterSet
func (t_ TokenFieldCell) SetTokenizingCharacterSet(value foundation.CharacterSet) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTokenizingCharacterSet:"), value)
}/* debug [instance_properties/setter]: tokenizingCharacterSet */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTokenFieldCell */



