// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [ParameterNode] class.
type IParameterNode interface {
	objectivec.IObject
	// properties:
	DisplayName() string /* primitive/slice/pointer. */
	Identifier() string /* primitive/slice/pointer. */
	ImplementorDisplayNameWithLengthCallback() ImplementorDisplayNameWithLengthCallback /* not a class type */
	SetImplementorDisplayNameWithLengthCallback(value ImplementorDisplayNameWithLengthCallback /* not a class type */)
	ImplementorStringFromValueCallback() ImplementorStringFromValueCallback /* not a class type */
	SetImplementorStringFromValueCallback(value ImplementorStringFromValueCallback /* not a class type */)
	ImplementorValueFromStringCallback() ImplementorValueFromStringCallback /* not a class type */
	SetImplementorValueFromStringCallback(value ImplementorValueFromStringCallback /* not a class type */)
	KeyPath() string /* primitive/slice/pointer. */
	ImplementorValueObserver() ImplementorValueObserver /* not a class type */
	SetImplementorValueObserver(value ImplementorValueObserver /* not a class type */)
	ImplementorValueProvider() ImplementorValueProvider /* not a class type */
	SetImplementorValueProvider(value ImplementorValueProvider /* not a class type */)
	// methods:
	DisplayNameWithLength(maximumLength int /* primitive/slice/pointer. */) objc.IObject /* cross-framework: String */
	RemoveParameterObserver(token objc.IObject /* cross-framework ParameterObserverToken */)
	TokenByAddingParameterAutomationObserver(observer ParameterAutomationObserver /* not a class type */) objc.IObject /* cross-framework: ParameterObserverToken */
	TokenByAddingParameterObserver(observer ParameterObserver /* not a class type */) objc.IObject /* cross-framework: ParameterObserverToken */
	TokenByAddingParameterRecordingObserver(observer ParameterRecordingObserver /* not a class type */) objc.IObject /* cross-framework: ParameterObserverToken */
}

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

// Alloc allocates a new instance without initialization.
func (pc _ParameterNodeClass) Alloc() ParameterNode {
	rv := objc.Send[ParameterNode](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Another version of the display name, possibly truncated to a desired length.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/displayName(withLength:)
func (p_ ParameterNode) DisplayNameWithLength(maximumLength int /* primitive/slice/pointer. */) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[String](p_.ID, objc.Sel("displayNameWithLength:"), maximumLength)
	return rv
}


// Remove a specific parameter observer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/removeParameterObserver(_:)
func (p_ ParameterNode) RemoveParameterObserver(token objc.IObject /* cross-framework ParameterObserverToken */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeParameterObserver:"), token)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/token(byAddingParameterAutomationObserver:)
func (p_ ParameterNode) TokenByAddingParameterAutomationObserver(observer ParameterAutomationObserver /* not a class type */) objc.IObject /* cross-framework: ParameterObserverToken */ {
	rv := objc.Send[ParameterObserverToken](p_.ID, objc.Sel("tokenByAddingParameterAutomationObserver:"), observer)
	return rv
}


// Adds an observer for a single parameter or all parameters in a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/token(byAddingParameterObserver:)
func (p_ ParameterNode) TokenByAddingParameterObserver(observer ParameterObserver /* not a class type */) objc.IObject /* cross-framework: ParameterObserverToken */ {
	rv := objc.Send[ParameterObserverToken](p_.ID, objc.Sel("tokenByAddingParameterObserver:"), observer)
	return rv
}


// Adds a recording observer for a single parameter or all parameters in a group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/token(byAddingParameterRecordingObserver:)
func (p_ ParameterNode) TokenByAddingParameterRecordingObserver(observer ParameterRecordingObserver /* not a class type */) objc.IObject /* cross-framework: ParameterObserverToken */ {
	rv := objc.Send[ParameterObserverToken](p_.ID, objc.Sel("tokenByAddingParameterRecordingObserver:"), observer)
	return rv
}


// A localized display name for the parameter node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/displayName
func (p_ ParameterNode) DisplayName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](p_.ID, objc.Sel("displayName"))
	return rv
}


// A non-localized, permanent name for the parameter node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/identifier
func (p_ ParameterNode) Identifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](p_.ID, objc.Sel("identifier"))
	return rv
}


// The callback for obtaining an abbreviated version of a parameter node display name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/implementorDisplayNameWithLengthCallback
func (p_ ParameterNode) ImplementorDisplayNameWithLengthCallback() ImplementorDisplayNameWithLengthCallback /* not a class type */ {
	rv := objc.Send[ImplementorDisplayNameWithLengthCallback](p_.ID, objc.Sel("implementorDisplayNameWithLengthCallback"))
	return rv
}


// The callback for obtaining an abbreviated version of a parameter node display name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/implementorDisplayNameWithLengthCallback
func (p_ ParameterNode) SetImplementorDisplayNameWithLengthCallback(value ImplementorDisplayNameWithLengthCallback /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImplementorDisplayNameWithLengthCallback:"), value)
}


// The callback for providing a string representation of a parameter value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/implementorStringFromValueCallback
func (p_ ParameterNode) ImplementorStringFromValueCallback() ImplementorStringFromValueCallback /* not a class type */ {
	rv := objc.Send[ImplementorStringFromValueCallback](p_.ID, objc.Sel("implementorStringFromValueCallback"))
	return rv
}


// The callback for providing a string representation of a parameter value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/implementorStringFromValueCallback
func (p_ ParameterNode) SetImplementorStringFromValueCallback(value ImplementorStringFromValueCallback /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImplementorStringFromValueCallback:"), value)
}


// The callback for converting a string to a parameter value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/implementorValueFromStringCallback
func (p_ ParameterNode) ImplementorValueFromStringCallback() ImplementorValueFromStringCallback /* not a class type */ {
	rv := objc.Send[ImplementorValueFromStringCallback](p_.ID, objc.Sel("implementorValueFromStringCallback"))
	return rv
}


// The callback for converting a string to a parameter value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/implementorValueFromStringCallback
func (p_ ParameterNode) SetImplementorValueFromStringCallback(value ImplementorValueFromStringCallback /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImplementorValueFromStringCallback:"), value)
}


// A key path generated by concatenating the identifiers of the parameter and its parents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/keyPath
func (p_ ParameterNode) KeyPath() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](p_.ID, objc.Sel("keyPath"))
	return rv
}


// The callback for parameter value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auparameternode/implementorvalueobserver
func (p_ ParameterNode) ImplementorValueObserver() ImplementorValueObserver /* not a class type */ {
	rv := objc.Send[ImplementorValueObserver](p_.ID, objc.Sel("implementorValueObserver"))
	return rv
}


// The callback for parameter value changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auparameternode/implementorvalueobserver
func (p_ ParameterNode) SetImplementorValueObserver(value ImplementorValueObserver /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImplementorValueObserver:"), value)
}


// The callback for refreshing known stale values in a parameter tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auparameternode/implementorvalueprovider
func (p_ ParameterNode) ImplementorValueProvider() ImplementorValueProvider /* not a class type */ {
	rv := objc.Send[ImplementorValueProvider](p_.ID, objc.Sel("implementorValueProvider"))
	return rv
}


// The callback for refreshing known stale values in a parameter tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/audiotoolbox/auparameternode/implementorvalueprovider
func (p_ ParameterNode) SetImplementorValueProvider(value ImplementorValueProvider /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImplementorValueProvider:"), value)
}



