// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AUParameterNode */


/* debug [class_header]: Header for AUParameterNode */
// The class instance for the [ParameterNode] class.
var (
	ParameterNodeClass     _ParameterNodeClass
	ParameterNodeClassOnce sync.Once
)

func getParameterNodeClass() _ParameterNodeClass {
	ParameterNodeClassOnce.Do(func() {
		ParameterNodeClass = _ParameterNodeClass{objc.GetClass("AUParameterNode")}
	})
	return ParameterNodeClass
}

type _ParameterNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ParameterNode */
// An interface definition for the [ParameterNode] class.
type IParameterNode interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ParameterNode */
	// properties:
	DisplayName() objc.IObject /* cross-framework: NSString */
	Identifier() objc.IObject /* cross-framework: NSString */
	ImplementorDisplayNameWithLengthCallback() ImplementorDisplayNameWithLengthCallback /* not a class type */
	SetImplementorDisplayNameWithLengthCallback(value ImplementorDisplayNameWithLengthCallback /* not a class type */)
	ImplementorStringFromValueCallback() ImplementorStringFromValueCallback /* not a class type */
	SetImplementorStringFromValueCallback(value ImplementorStringFromValueCallback /* not a class type */)
	ImplementorValueFromStringCallback() ImplementorValueFromStringCallback /* not a class type */
	SetImplementorValueFromStringCallback(value ImplementorValueFromStringCallback /* not a class type */)
	ImplementorValueObserver() ImplementorValueObserver /* not a class type */
	SetImplementorValueObserver(value ImplementorValueObserver /* not a class type */)
	ImplementorValueProvider() ImplementorValueProvider /* not a class type */
	SetImplementorValueProvider(value ImplementorValueProvider /* not a class type */)
	KeyPath() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ParameterNode */
	// methods:
	DisplayNameWithLength(maximumLength int) foundation.String
	RemoveParameterObserver(token ParameterObserverToken /* typedef */)
	TokenByAddingParameterAutomationObserver(observer ParameterAutomationObserver /* not a class type */) ParameterObserverToken /* typedef */
	TokenByAddingParameterObserver(observer ParameterObserver /* not a class type */) ParameterObserverToken /* typedef */
	TokenByAddingParameterRecordingObserver(observer ParameterRecordingObserver /* not a class type */) ParameterObserverToken /* typedef */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ParameterNode */
// Alloc allocates a new instance without initialization.
func (pc _ParameterNodeClass) Alloc() ParameterNode {
	rv := objc.Send[ParameterNode](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _ParameterNodeClass) New() ParameterNode {
	rv := objc.Send[ParameterNode](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ParameterNode) Init() ParameterNode {
	rv := objc.Send[ParameterNode](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ParameterNode) Autorelease() ParameterNode {
	rv := objc.Send[ParameterNode](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewParameterNode creates a new ParameterNode instance.
func NewParameterNode() ParameterNode {
	return getParameterNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ParameterNode */
// An object that represents a node in an audio unit’s parameter tree.
//
// Nodes are instances of either an or class.


// An object that represents a node in an audio unit’s parameter tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode
type ParameterNode struct {
	objectivec.Object
}

// ParameterNodeFrom constructs a [ParameterNode] from an unsafe.Pointer.
//
// An object that represents a node in an audio unit’s parameter tree.
func ParameterNodeFrom(ptr unsafe.Pointer) ParameterNode {
	return ParameterNode{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ParameterNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ParameterNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ParameterNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ParameterNode */

// Another version of the display name, possibly truncated to a desired length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/displayName(withLength:)
func (p_ ParameterNode) DisplayNameWithLength(maximumLength int) foundation.String {
	rv := objc.Send[foundation.String](p_.ID, objc.Sel("displayNameWithLength:"), maximumLength)
	return rv
}/* debug [instance_methods/method]: DisplayNameWithLength */


// Remove a specific parameter observer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/removeParameterObserver(_:)
func (p_ ParameterNode) RemoveParameterObserver(token ParameterObserverToken /* typedef */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeParameterObserver:"), token)
}/* debug [instance_methods/method]: RemoveParameterObserver */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/token(byAddingParameterAutomationObserver:)
func (p_ ParameterNode) TokenByAddingParameterAutomationObserver(observer ParameterAutomationObserver /* not a class type */) ParameterObserverToken /* typedef */ {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("tokenByAddingParameterAutomationObserver:"), observer)
	return rv
}/* debug [instance_methods/method]: TokenByAddingParameterAutomationObserver */


// Adds an observer for a single parameter or all parameters in a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/token(byAddingParameterObserver:)
func (p_ ParameterNode) TokenByAddingParameterObserver(observer ParameterObserver /* not a class type */) ParameterObserverToken /* typedef */ {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("tokenByAddingParameterObserver:"), observer)
	return rv
}/* debug [instance_methods/method]: TokenByAddingParameterObserver */


// Adds a recording observer for a single parameter or all parameters in a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/token(byAddingParameterRecordingObserver:)
func (p_ ParameterNode) TokenByAddingParameterRecordingObserver(observer ParameterRecordingObserver /* not a class type */) ParameterObserverToken /* typedef */ {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("tokenByAddingParameterRecordingObserver:"), observer)
	return rv
}/* debug [instance_methods/method]: TokenByAddingParameterRecordingObserver */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ParameterNode */

// A localized display name for the parameter node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/displayName
func (p_ ParameterNode) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("displayName"))
	return rv
}/* debug [instance_properties/getter]: displayName */


// A non-localized, permanent name for the parameter node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/identifier
func (p_ ParameterNode) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The callback for obtaining an abbreviated version of a parameter node display name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/implementorDisplayNameWithLengthCallback
func (p_ ParameterNode) ImplementorDisplayNameWithLengthCallback() ImplementorDisplayNameWithLengthCallback /* not a class type */ {
	rv := objc.Send[ImplementorDisplayNameWithLengthCallback](p_.ID, objc.Sel("implementorDisplayNameWithLengthCallback"))
	return rv
}/* debug [instance_properties/getter]: implementorDisplayNameWithLengthCallback */


// The callback for obtaining an abbreviated version of a parameter node display name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/implementorDisplayNameWithLengthCallback
func (p_ ParameterNode) SetImplementorDisplayNameWithLengthCallback(value ImplementorDisplayNameWithLengthCallback /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImplementorDisplayNameWithLengthCallback:"), value)
}/* debug [instance_properties/setter]: implementorDisplayNameWithLengthCallback */


// The callback for providing a string representation of a parameter value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/implementorStringFromValueCallback
func (p_ ParameterNode) ImplementorStringFromValueCallback() ImplementorStringFromValueCallback /* not a class type */ {
	rv := objc.Send[ImplementorStringFromValueCallback](p_.ID, objc.Sel("implementorStringFromValueCallback"))
	return rv
}/* debug [instance_properties/getter]: implementorStringFromValueCallback */


// The callback for providing a string representation of a parameter value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/implementorStringFromValueCallback
func (p_ ParameterNode) SetImplementorStringFromValueCallback(value ImplementorStringFromValueCallback /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImplementorStringFromValueCallback:"), value)
}/* debug [instance_properties/setter]: implementorStringFromValueCallback */


// The callback for converting a string to a parameter value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/implementorValueFromStringCallback
func (p_ ParameterNode) ImplementorValueFromStringCallback() ImplementorValueFromStringCallback /* not a class type */ {
	rv := objc.Send[ImplementorValueFromStringCallback](p_.ID, objc.Sel("implementorValueFromStringCallback"))
	return rv
}/* debug [instance_properties/getter]: implementorValueFromStringCallback */


// The callback for converting a string to a parameter value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/implementorValueFromStringCallback
func (p_ ParameterNode) SetImplementorValueFromStringCallback(value ImplementorValueFromStringCallback /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImplementorValueFromStringCallback:"), value)
}/* debug [instance_properties/setter]: implementorValueFromStringCallback */


// The callback for parameter value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/implementorValueObserver
func (p_ ParameterNode) ImplementorValueObserver() ImplementorValueObserver /* not a class type */ {
	rv := objc.Send[ImplementorValueObserver](p_.ID, objc.Sel("implementorValueObserver"))
	return rv
}/* debug [instance_properties/getter]: implementorValueObserver */


// The callback for parameter value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/implementorValueObserver
func (p_ ParameterNode) SetImplementorValueObserver(value ImplementorValueObserver /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImplementorValueObserver:"), value)
}/* debug [instance_properties/setter]: implementorValueObserver */


// The callback for refreshing known stale values in a parameter tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/implementorValueProvider
func (p_ ParameterNode) ImplementorValueProvider() ImplementorValueProvider /* not a class type */ {
	rv := objc.Send[ImplementorValueProvider](p_.ID, objc.Sel("implementorValueProvider"))
	return rv
}/* debug [instance_properties/getter]: implementorValueProvider */


// The callback for refreshing known stale values in a parameter tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/implementorValueProvider
func (p_ ParameterNode) SetImplementorValueProvider(value ImplementorValueProvider /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImplementorValueProvider:"), value)
}/* debug [instance_properties/setter]: implementorValueProvider */


// A key path generated by concatenating the identifiers of the parameter and its parents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/keyPath
func (p_ ParameterNode) KeyPath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("keyPath"))
	return rv
}/* debug [instance_properties/getter]: keyPath */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AUParameterNode */



