// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DOMKeyboardEvent */


/* debug [class_header]: Header for DOMKeyboardEvent */
// The class instance for the [DOMKeyboardEvent] class.
var (
	DOMKeyboardEventClass     _DOMKeyboardEventClass
	DOMKeyboardEventClassOnce sync.Once
)

func getDOMKeyboardEventClass() _DOMKeyboardEventClass {
	DOMKeyboardEventClassOnce.Do(func() {
		DOMKeyboardEventClass = _DOMKeyboardEventClass{objc.GetClass("DOMKeyboardEvent")}
	})
	return DOMKeyboardEventClass
}

type _DOMKeyboardEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMKeyboardEvent */
// An interface definition for the [DOMKeyboardEvent] class.
type IDOMKeyboardEvent interface {
	IDOMUIEvent
	
/* debug [class_interface_properties]: Properties for DOMKeyboardEvent */
	// properties:
	AltGraphKey() bool
	AltKey() bool
	CharCode() int
	CtrlKey() bool
	KeyCode() int
	KeyIdentifier() objc.IObject /* cross-framework: NSString */
	KeyLocation() objectivec.IObject
	Location() objectivec.IObject
	MetaKey() bool
	ShiftKey() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMKeyboardEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMKeyboardEvent */
// Alloc allocates a new instance without initialization.
func (dc _DOMKeyboardEventClass) Alloc() DOMKeyboardEvent {
	rv := objc.Send[DOMKeyboardEvent](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMKeyboardEventClass) New() DOMKeyboardEvent {
	rv := objc.Send[DOMKeyboardEvent](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMKeyboardEvent) Init() DOMKeyboardEvent {
	rv := objc.Send[DOMKeyboardEvent](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMKeyboardEvent) Autorelease() DOMKeyboardEvent {
	rv := objc.Send[DOMKeyboardEvent](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMKeyboardEvent creates a new DOMKeyboardEvent instance.
func NewDOMKeyboardEvent() DOMKeyboardEvent {
	return getDOMKeyboardEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMKeyboardEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMKeyboardEvent
type DOMKeyboardEvent struct {
	DOMUIEvent
}

// DOMKeyboardEventFrom constructs a [DOMKeyboardEvent] from an unsafe.Pointer.
func DOMKeyboardEventFrom(ptr unsafe.Pointer) DOMKeyboardEvent {
	return DOMKeyboardEvent{
		DOMUIEvent: DOMUIEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMKeyboardEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMKeyboardEvent/initKeyboardEvent:canBubble:cancelable:view:keyIdentifier:keyLocation:ctrlKey:altKey:shiftKey:metaKey:
func NewDOMKeyboardEventKeyboardEventCanBubbleCancelableViewKeyIdentifierKeyLocationCtrlKeyAltKeyShiftKeyMetaKey(type_ objc.IObject /* cross-framework: NSString */, canBubble bool, cancelable bool, view IDOMAbstractView, keyIdentifier objc.IObject /* cross-framework: NSString */, keyLocation objectivec.IObject, ctrlKey bool, altKey bool, shiftKey bool, metaKey bool) DOMKeyboardEvent {
	instance := getDOMKeyboardEventClass().Alloc()
	rv := objc.Send[DOMKeyboardEvent](instance.ID, objc.Sel("initKeyboardEvent:canBubble:cancelable:view:keyIdentifier:keyLocation:ctrlKey:altKey:shiftKey:metaKey:"), type_, canBubble, cancelable, view, keyIdentifier, keyLocation, ctrlKey, altKey, shiftKey, metaKey)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDOMKeyboardEventKeyboardEventCanBubbleCancelableViewKeyIdentifierKeyLocationCtrlKeyAltKeyShiftKeyMetaKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMKeyboardEvent/initKeyboardEvent:canBubble:cancelable:view:keyIdentifier:keyLocation:ctrlKey:altKey:shiftKey:metaKey:altGraphKey:
func NewDOMKeyboardEventKeyboardEventCanBubbleCancelableViewKeyIdentifierKeyLocationCtrlKeyAltKeyShiftKeyMetaKeyAltGraphKey(type_ objc.IObject /* cross-framework: NSString */, canBubble bool, cancelable bool, view IDOMAbstractView, keyIdentifier objc.IObject /* cross-framework: NSString */, keyLocation objectivec.IObject, ctrlKey bool, altKey bool, shiftKey bool, metaKey bool, altGraphKey bool) DOMKeyboardEvent {
	instance := getDOMKeyboardEventClass().Alloc()
	rv := objc.Send[DOMKeyboardEvent](instance.ID, objc.Sel("initKeyboardEvent:canBubble:cancelable:view:keyIdentifier:keyLocation:ctrlKey:altKey:shiftKey:metaKey:altGraphKey:"), type_, canBubble, cancelable, view, keyIdentifier, keyLocation, ctrlKey, altKey, shiftKey, metaKey, altGraphKey)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDOMKeyboardEventKeyboardEventCanBubbleCancelableViewKeyIdentifierKeyLocationCtrlKeyAltKeyShiftKeyMetaKeyAltGraphKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMKeyboardEvent/initKeyboardEvent(_:canBubble:cancelable:view:keyIdentifier:location:ctrlKey:altKey:shiftKey:metaKey:)
func NewDOMKeyboardEventKeyboardEventCanBubbleCancelableViewKeyIdentifierLocationCtrlKeyAltKeyShiftKeyMetaKey(type_ objc.IObject /* cross-framework: NSString */, canBubble bool, cancelable bool, view IDOMAbstractView, keyIdentifier objc.IObject /* cross-framework: NSString */, location objectivec.IObject, ctrlKey bool, altKey bool, shiftKey bool, metaKey bool) DOMKeyboardEvent {
	instance := getDOMKeyboardEventClass().Alloc()
	rv := objc.Send[DOMKeyboardEvent](instance.ID, objc.Sel("initKeyboardEvent:canBubble:cancelable:view:keyIdentifier:location:ctrlKey:altKey:shiftKey:metaKey:"), type_, canBubble, cancelable, view, keyIdentifier, location, ctrlKey, altKey, shiftKey, metaKey)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDOMKeyboardEventKeyboardEventCanBubbleCancelableViewKeyIdentifierLocationCtrlKeyAltKeyShiftKeyMetaKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMKeyboardEvent/initKeyboardEvent(_:canBubble:cancelable:view:keyIdentifier:location:ctrlKey:altKey:shiftKey:metaKey:altGraphKey:)
func NewDOMKeyboardEventKeyboardEventCanBubbleCancelableViewKeyIdentifierLocationCtrlKeyAltKeyShiftKeyMetaKeyAltGraphKey(type_ objc.IObject /* cross-framework: NSString */, canBubble bool, cancelable bool, view IDOMAbstractView, keyIdentifier objc.IObject /* cross-framework: NSString */, location objectivec.IObject, ctrlKey bool, altKey bool, shiftKey bool, metaKey bool, altGraphKey bool) DOMKeyboardEvent {
	instance := getDOMKeyboardEventClass().Alloc()
	rv := objc.Send[DOMKeyboardEvent](instance.ID, objc.Sel("initKeyboardEvent:canBubble:cancelable:view:keyIdentifier:location:ctrlKey:altKey:shiftKey:metaKey:altGraphKey:"), type_, canBubble, cancelable, view, keyIdentifier, location, ctrlKey, altKey, shiftKey, metaKey, altGraphKey)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDOMKeyboardEventKeyboardEventCanBubbleCancelableViewKeyIdentifierLocationCtrlKeyAltKeyShiftKeyMetaKeyAltGraphKey */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMKeyboardEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMKeyboardEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMKeyboardEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMKeyboardEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMKeyboardEvent/altGraphKey
func (d_ DOMKeyboardEvent) AltGraphKey() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("altGraphKey"))
	return rv
}/* debug [instance_properties/getter]: altGraphKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMKeyboardEvent/altKey
func (d_ DOMKeyboardEvent) AltKey() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("altKey"))
	return rv
}/* debug [instance_properties/getter]: altKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMKeyboardEvent/charCode
func (d_ DOMKeyboardEvent) CharCode() int {
	rv := objc.Send[int](d_.ID, objc.Sel("charCode"))
	return rv
}/* debug [instance_properties/getter]: charCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMKeyboardEvent/ctrlKey
func (d_ DOMKeyboardEvent) CtrlKey() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("ctrlKey"))
	return rv
}/* debug [instance_properties/getter]: ctrlKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMKeyboardEvent/keyCode
func (d_ DOMKeyboardEvent) KeyCode() int {
	rv := objc.Send[int](d_.ID, objc.Sel("keyCode"))
	return rv
}/* debug [instance_properties/getter]: keyCode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMKeyboardEvent/keyIdentifier
func (d_ DOMKeyboardEvent) KeyIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("keyIdentifier"))
	return rv
}/* debug [instance_properties/getter]: keyIdentifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMKeyboardEvent/keyLocation
func (d_ DOMKeyboardEvent) KeyLocation() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("keyLocation"))
	return rv
}/* debug [instance_properties/getter]: keyLocation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMKeyboardEvent/location
func (d_ DOMKeyboardEvent) Location() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("location"))
	return rv
}/* debug [instance_properties/getter]: location */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMKeyboardEvent/metaKey
func (d_ DOMKeyboardEvent) MetaKey() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("metaKey"))
	return rv
}/* debug [instance_properties/getter]: metaKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMKeyboardEvent/shiftKey
func (d_ DOMKeyboardEvent) ShiftKey() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("shiftKey"))
	return rv
}/* debug [instance_properties/getter]: shiftKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMKeyboardEvent */


