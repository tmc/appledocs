// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSSecureTextFieldCell */


/* debug [class_header]: Header for NSSecureTextFieldCell */
// The class instance for the [SecureTextFieldCell] class.
var (
	SecureTextFieldCellClass     _SecureTextFieldCellClass
	SecureTextFieldCellClassOnce sync.Once
)

func getSecureTextFieldCellClass() _SecureTextFieldCellClass {
	SecureTextFieldCellClassOnce.Do(func() {
		SecureTextFieldCellClass = _SecureTextFieldCellClass{objc.GetClass("NSSecureTextFieldCell")}
	})
	return SecureTextFieldCellClass
}

type _SecureTextFieldCellClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SecureTextFieldCell */
// An interface definition for the [SecureTextFieldCell] class.
type ISecureTextFieldCell interface {
	ITextFieldCell
	
/* debug [class_interface_properties]: Properties for SecureTextFieldCell */
	// properties:
	EchosBullets() bool
	SetEchosBullets(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SecureTextFieldCell */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SecureTextFieldCell */
// Alloc allocates a new instance without initialization.
func (sc _SecureTextFieldCellClass) Alloc() SecureTextFieldCell {
	rv := objc.Send[SecureTextFieldCell](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SecureTextFieldCellClass) New() SecureTextFieldCell {
	rv := objc.Send[SecureTextFieldCell](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SecureTextFieldCell) Init() SecureTextFieldCell {
	rv := objc.Send[SecureTextFieldCell](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SecureTextFieldCell) Autorelease() SecureTextFieldCell {
	rv := objc.Send[SecureTextFieldCell](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSecureTextFieldCell creates a new SecureTextFieldCell instance.
func NewSecureTextFieldCell() SecureTextFieldCell {
	return getSecureTextFieldCellClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SecureTextFieldCell */
// A text field whose value is hidden from the user.
//
// works with and overrides the general cell use of the field editor to provide its own field editor, which doesn’t display text or allow the user to cut or copy its value.


// A text field whose value is hidden from the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSecureTextFieldCell
type SecureTextFieldCell struct {
	TextFieldCell
}

// SecureTextFieldCellFrom constructs a [SecureTextFieldCell] from an unsafe.Pointer.
//
// A text field whose value is hidden from the user.
func SecureTextFieldCellFrom(ptr unsafe.Pointer) SecureTextFieldCell {
	return SecureTextFieldCell{
		TextFieldCell: TextFieldCellFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SecureTextFieldCell *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SecureTextFieldCell */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SecureTextFieldCell */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SecureTextFieldCell */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SecureTextFieldCell */

// A Boolean that indicates whether the receiver echoes a bullet character rather than each character typed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSecureTextFieldCell/echosBullets
func (s_ SecureTextFieldCell) EchosBullets() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("echosBullets"))
	return rv
}/* debug [instance_properties/getter]: echosBullets */


// A Boolean that indicates whether the receiver echoes a bullet character rather than each character typed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSecureTextFieldCell/echosBullets
func (s_ SecureTextFieldCell) SetEchosBullets(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEchosBullets:"), value)
}/* debug [instance_properties/setter]: echosBullets */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSecureTextFieldCell */



