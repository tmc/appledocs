// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
	DisplayNameWithLength(maximumLength int) string
	RemoveParameterObserver(token unsafe.Pointer)
	TokenByAddingParameterAutomationObserver(observer unsafe.Pointer) unsafe.Pointer
	TokenByAddingParameterObserver(observer unsafe.Pointer) unsafe.Pointer
	TokenByAddingParameterRecordingObserver(observer unsafe.Pointer) unsafe.Pointer
}

// An object that represents a node in an audio unit’s parameter tree.
//
// Nodes are instances of either an or class.
//
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
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/displayName(withLength:)
func (p_ ParameterNode) DisplayNameWithLength(maximumLength int) string {
	rv := objc.Send[string](p_.ID, objc.Sel("displayNameWithLength:"), maximumLength)
	return rv
}

// Remove a specific parameter observer.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/removeParameterObserver(_:)
func (p_ ParameterNode) RemoveParameterObserver(token unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removeParameterObserver:"), token)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/token(byAddingParameterAutomationObserver:)
func (p_ ParameterNode) TokenByAddingParameterAutomationObserver(observer unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("tokenByAddingParameterAutomationObserver:"), observer)
	return rv
}

// Adds an observer for a single parameter or all parameters in a group.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/token(byAddingParameterObserver:)
func (p_ ParameterNode) TokenByAddingParameterObserver(observer unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("tokenByAddingParameterObserver:"), observer)
	return rv
}

// Adds a recording observer for a single parameter or all parameters in a group.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/token(byAddingParameterRecordingObserver:)
func (p_ ParameterNode) TokenByAddingParameterRecordingObserver(observer unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("tokenByAddingParameterRecordingObserver:"), observer)
	return rv
}

// A localized display name for the parameter node.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/displayName
func (p_ ParameterNode) DisplayName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("displayName"))
	return rv
}

// A non-localized, permanent name for the parameter node.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/identifier
func (p_ ParameterNode) Identifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("identifier"))
	return rv
}

// The callback for obtaining an abbreviated version of a parameter node display name.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/implementorDisplayNameWithLengthCallback
func (p_ ParameterNode) ImplementorDisplayNameWithLengthCallback() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("implementorDisplayNameWithLengthCallback"))
	return rv
}


// SetImplementorDisplayNameWithLengthCallback sets the value of the implementorDisplayNameWithLengthCallback property.
// The callback for obtaining an abbreviated version of a parameter node display name.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/implementorDisplayNameWithLengthCallback
func (p_ ParameterNode) SetImplementorDisplayNameWithLengthCallback(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImplementorDisplayNameWithLengthCallback:"), value)
}
// The callback for converting a string to a parameter value.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/implementorValueFromStringCallback
func (p_ ParameterNode) ImplementorValueFromStringCallback() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("implementorValueFromStringCallback"))
	return rv
}


// SetImplementorValueFromStringCallback sets the value of the implementorValueFromStringCallback property.
// The callback for converting a string to a parameter value.

//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/implementorValueFromStringCallback
func (p_ ParameterNode) SetImplementorValueFromStringCallback(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImplementorValueFromStringCallback:"), value)
}
// A key path generated by concatenating the identifiers of the parameter and its parents.
//
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/AUParameterNode/keyPath
func (p_ ParameterNode) KeyPath() string {
	rv := objc.Send[string](p_.ID, objc.Sel("keyPath"))
	return rv
}



