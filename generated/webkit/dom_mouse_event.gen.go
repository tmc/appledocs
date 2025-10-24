// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DOMMouseEvent */


/* debug [class_header]: Header for DOMMouseEvent */
// The class instance for the [DOMMouseEvent] class.
var (
	DOMMouseEventClass     _DOMMouseEventClass
	DOMMouseEventClassOnce sync.Once
)

func getDOMMouseEventClass() _DOMMouseEventClass {
	DOMMouseEventClassOnce.Do(func() {
		DOMMouseEventClass = _DOMMouseEventClass{objc.GetClass("DOMMouseEvent")}
	})
	return DOMMouseEventClass
}

type _DOMMouseEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMMouseEvent */
// An interface definition for the [DOMMouseEvent] class.
type IDOMMouseEvent interface {
	IDOMUIEvent
	
/* debug [class_interface_properties]: Properties for DOMMouseEvent */
	// properties:
	AltKey() bool
	Button() objectivec.IObject
	ClientX() int
	ClientY() int
	CtrlKey() bool
	FromElement() IDOMNode
	MetaKey() bool
	OffsetX() int
	OffsetY() int
	RelatedTarget() unsafe.Pointer
	ScreenX() int
	ScreenY() int
	ShiftKey() bool
	ToElement() IDOMNode
	X() int
	Y() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMMouseEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMMouseEvent */
// Alloc allocates a new instance without initialization.
func (dc _DOMMouseEventClass) Alloc() DOMMouseEvent {
	rv := objc.Send[DOMMouseEvent](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMMouseEventClass) New() DOMMouseEvent {
	rv := objc.Send[DOMMouseEvent](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMMouseEvent) Init() DOMMouseEvent {
	rv := objc.Send[DOMMouseEvent](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMMouseEvent) Autorelease() DOMMouseEvent {
	rv := objc.Send[DOMMouseEvent](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMMouseEvent creates a new DOMMouseEvent instance.
func NewDOMMouseEvent() DOMMouseEvent {
	return getDOMMouseEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMMouseEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMouseEvent
type DOMMouseEvent struct {
	DOMUIEvent
}

// DOMMouseEventFrom constructs a [DOMMouseEvent] from an unsafe.Pointer.
func DOMMouseEventFrom(ptr unsafe.Pointer) DOMMouseEvent {
	return DOMMouseEvent{
		DOMUIEvent: DOMUIEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMMouseEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMouseEvent/initMouseEvent:::::::::::::::
func NewDOMMouseEventMouseEvent(type_ objc.IObject /* cross-framework: NSString */, canBubble bool, cancelable bool, view IDOMAbstractView, detail int, screenX int, screenY int, clientX int, clientY int, ctrlKey bool, altKey bool, shiftKey bool, metaKey bool, button objectivec.IObject, relatedTarget unsafe.Pointer) DOMMouseEvent {
	instance := getDOMMouseEventClass().Alloc()
	rv := objc.Send[DOMMouseEvent](instance.ID, objc.Sel("initMouseEvent:::::::::::::::"), type_, canBubble, cancelable, view, detail, screenX, screenY, clientX, clientY, ctrlKey, altKey, shiftKey, metaKey, button, relatedTarget)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDOMMouseEventMouseEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMouseEvent/initMouseEvent(_:canBubble:cancelable:view:detail:screenX:screenY:clientX:clientY:ctrlKey:altKey:shiftKey:metaKey:button:relatedTarget:)
func NewDOMMouseEventMouseEventCanBubbleCancelableViewDetailScreenXScreenYClientXClientYCtrlKeyAltKeyShiftKeyMetaKeyButtonRelatedTarget(type_ objc.IObject /* cross-framework: NSString */, canBubble bool, cancelable bool, view IDOMAbstractView, detail int, screenX int, screenY int, clientX int, clientY int, ctrlKey bool, altKey bool, shiftKey bool, metaKey bool, button objectivec.IObject, relatedTarget unsafe.Pointer) DOMMouseEvent {
	instance := getDOMMouseEventClass().Alloc()
	rv := objc.Send[DOMMouseEvent](instance.ID, objc.Sel("initMouseEvent:canBubble:cancelable:view:detail:screenX:screenY:clientX:clientY:ctrlKey:altKey:shiftKey:metaKey:button:relatedTarget:"), type_, canBubble, cancelable, view, detail, screenX, screenY, clientX, clientY, ctrlKey, altKey, shiftKey, metaKey, button, relatedTarget)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDOMMouseEventMouseEventCanBubbleCancelableViewDetailScreenXScreenYClientXClientYCtrlKeyAltKeyShiftKeyMetaKeyButtonRelatedTarget */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMMouseEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMMouseEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMMouseEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMMouseEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMouseEvent/altKey
func (d_ DOMMouseEvent) AltKey() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("altKey"))
	return rv
}/* debug [instance_properties/getter]: altKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMouseEvent/button
func (d_ DOMMouseEvent) Button() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("button"))
	return rv
}/* debug [instance_properties/getter]: button */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMouseEvent/clientX
func (d_ DOMMouseEvent) ClientX() int {
	rv := objc.Send[int](d_.ID, objc.Sel("clientX"))
	return rv
}/* debug [instance_properties/getter]: clientX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMouseEvent/clientY
func (d_ DOMMouseEvent) ClientY() int {
	rv := objc.Send[int](d_.ID, objc.Sel("clientY"))
	return rv
}/* debug [instance_properties/getter]: clientY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMouseEvent/ctrlKey
func (d_ DOMMouseEvent) CtrlKey() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("ctrlKey"))
	return rv
}/* debug [instance_properties/getter]: ctrlKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMouseEvent/fromElement
func (d_ DOMMouseEvent) FromElement() IDOMNode {
	rv := objc.Send[DOMNode](d_.ID, objc.Sel("fromElement"))
	return rv
}/* debug [instance_properties/getter]: fromElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMouseEvent/metaKey
func (d_ DOMMouseEvent) MetaKey() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("metaKey"))
	return rv
}/* debug [instance_properties/getter]: metaKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMouseEvent/offsetX
func (d_ DOMMouseEvent) OffsetX() int {
	rv := objc.Send[int](d_.ID, objc.Sel("offsetX"))
	return rv
}/* debug [instance_properties/getter]: offsetX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMouseEvent/offsetY
func (d_ DOMMouseEvent) OffsetY() int {
	rv := objc.Send[int](d_.ID, objc.Sel("offsetY"))
	return rv
}/* debug [instance_properties/getter]: offsetY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMouseEvent/relatedTarget
func (d_ DOMMouseEvent) RelatedTarget() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("relatedTarget"))
	return rv
}/* debug [instance_properties/getter]: relatedTarget */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMouseEvent/screenX
func (d_ DOMMouseEvent) ScreenX() int {
	rv := objc.Send[int](d_.ID, objc.Sel("screenX"))
	return rv
}/* debug [instance_properties/getter]: screenX */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMouseEvent/screenY
func (d_ DOMMouseEvent) ScreenY() int {
	rv := objc.Send[int](d_.ID, objc.Sel("screenY"))
	return rv
}/* debug [instance_properties/getter]: screenY */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMouseEvent/shiftKey
func (d_ DOMMouseEvent) ShiftKey() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("shiftKey"))
	return rv
}/* debug [instance_properties/getter]: shiftKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMouseEvent/toElement
func (d_ DOMMouseEvent) ToElement() IDOMNode {
	rv := objc.Send[DOMNode](d_.ID, objc.Sel("toElement"))
	return rv
}/* debug [instance_properties/getter]: toElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMouseEvent/x
func (d_ DOMMouseEvent) X() int {
	rv := objc.Send[int](d_.ID, objc.Sel("x"))
	return rv
}/* debug [instance_properties/getter]: x */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMouseEvent/y
func (d_ DOMMouseEvent) Y() int {
	rv := objc.Send[int](d_.ID, objc.Sel("y"))
	return rv
}/* debug [instance_properties/getter]: y */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMMouseEvent */


