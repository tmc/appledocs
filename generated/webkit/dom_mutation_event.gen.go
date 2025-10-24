// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DOMMutationEvent */


/* debug [class_header]: Header for DOMMutationEvent */
// The class instance for the [DOMMutationEvent] class.
var (
	DOMMutationEventClass     _DOMMutationEventClass
	DOMMutationEventClassOnce sync.Once
)

func getDOMMutationEventClass() _DOMMutationEventClass {
	DOMMutationEventClassOnce.Do(func() {
		DOMMutationEventClass = _DOMMutationEventClass{objc.GetClass("DOMMutationEvent")}
	})
	return DOMMutationEventClass
}

type _DOMMutationEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMMutationEvent */
// An interface definition for the [DOMMutationEvent] class.
type IDOMMutationEvent interface {
	IDOMEvent
	
/* debug [class_interface_properties]: Properties for DOMMutationEvent */
	// properties:
	AttrChange() objectivec.IObject
	AttrName() objc.IObject /* cross-framework: NSString */
	NewValue() objc.IObject /* cross-framework: NSString */
	PrevValue() objc.IObject /* cross-framework: NSString */
	RelatedNode() IDOMNode
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMMutationEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMMutationEvent */
// Alloc allocates a new instance without initialization.
func (dc _DOMMutationEventClass) Alloc() DOMMutationEvent {
	rv := objc.Send[DOMMutationEvent](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMMutationEventClass) New() DOMMutationEvent {
	rv := objc.Send[DOMMutationEvent](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMMutationEvent) Init() DOMMutationEvent {
	rv := objc.Send[DOMMutationEvent](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMMutationEvent) Autorelease() DOMMutationEvent {
	rv := objc.Send[DOMMutationEvent](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMMutationEvent creates a new DOMMutationEvent instance.
func NewDOMMutationEvent() DOMMutationEvent {
	return getDOMMutationEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMMutationEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMutationEvent
type DOMMutationEvent struct {
	DOMEvent
}

// DOMMutationEventFrom constructs a [DOMMutationEvent] from an unsafe.Pointer.
func DOMMutationEventFrom(ptr unsafe.Pointer) DOMMutationEvent {
	return DOMMutationEvent{
		DOMEvent: DOMEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMMutationEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMutationEvent/initMutationEvent::::::::
func NewDOMMutationEventMutationEvent(type_ objc.IObject /* cross-framework: NSString */, canBubble bool, cancelable bool, relatedNode IDOMNode, prevValue objc.IObject /* cross-framework: NSString */, newValue objc.IObject /* cross-framework: NSString */, attrName objc.IObject /* cross-framework: NSString */, attrChange objectivec.IObject) DOMMutationEvent {
	instance := getDOMMutationEventClass().Alloc()
	rv := objc.Send[DOMMutationEvent](instance.ID, objc.Sel("initMutationEvent::::::::"), type_, canBubble, cancelable, relatedNode, prevValue, newValue, attrName, attrChange)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDOMMutationEventMutationEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMutationEvent/initMutationEvent(_:canBubble:cancelable:relatedNode:prevValue:newValue:attrName:attrChange:)
func NewDOMMutationEventMutationEventCanBubbleCancelableRelatedNodePrevValueNewValueAttrNameAttrChange(type_ objc.IObject /* cross-framework: NSString */, canBubble bool, cancelable bool, relatedNode IDOMNode, prevValue objc.IObject /* cross-framework: NSString */, newValue objc.IObject /* cross-framework: NSString */, attrName objc.IObject /* cross-framework: NSString */, attrChange objectivec.IObject) DOMMutationEvent {
	instance := getDOMMutationEventClass().Alloc()
	rv := objc.Send[DOMMutationEvent](instance.ID, objc.Sel("initMutationEvent:canBubble:cancelable:relatedNode:prevValue:newValue:attrName:attrChange:"), type_, canBubble, cancelable, relatedNode, prevValue, newValue, attrName, attrChange)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDOMMutationEventMutationEventCanBubbleCancelableRelatedNodePrevValueNewValueAttrNameAttrChange */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMMutationEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMMutationEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMMutationEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMMutationEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMutationEvent/attrChange
func (d_ DOMMutationEvent) AttrChange() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("attrChange"))
	return rv
}/* debug [instance_properties/getter]: attrChange */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMutationEvent/attrName
func (d_ DOMMutationEvent) AttrName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("attrName"))
	return rv
}/* debug [instance_properties/getter]: attrName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMutationEvent/newValue
func (d_ DOMMutationEvent) NewValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("newValue"))
	return rv
}/* debug [instance_properties/getter]: newValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMutationEvent/prevValue
func (d_ DOMMutationEvent) PrevValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("prevValue"))
	return rv
}/* debug [instance_properties/getter]: prevValue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMMutationEvent/relatedNode
func (d_ DOMMutationEvent) RelatedNode() IDOMNode {
	rv := objc.Send[DOMNode](d_.ID, objc.Sel("relatedNode"))
	return rv
}/* debug [instance_properties/getter]: relatedNode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMMutationEvent */


