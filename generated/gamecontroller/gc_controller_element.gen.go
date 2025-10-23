// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GCControllerElement] class.
var (
	GCControllerElementClass     _GCControllerElementClass
	GCControllerElementClassOnce sync.Once
)

func getGCControllerElementClass() _GCControllerElementClass {
	GCControllerElementClassOnce.Do(func() {
		GCControllerElementClass = _GCControllerElementClass{objc.GetClass("GCControllerElement")}
	})
	return GCControllerElementClass
}

type _GCControllerElementClass struct {
	class objc.Class
}

// An interface definition for the [GCControllerElement] class.
type IGCControllerElement interface {
	objectivec.IObject
	Aliases() unsafe.Pointer
	Analog() bool
	Collection() GCControllerElement
	SetCollection(value IGCControllerElement)
	IsAnalog() bool
	SetIsAnalog(value bool)
	IsBoundToSystemGesture() bool
	SetIsBoundToSystemGesture(value bool)
	LocalizedName() string
	SetLocalizedName(value string)
	PreferredSystemGestureState() unsafe.Pointer
	SetPreferredSystemGestureState(value unsafe.Pointer)
	SfSymbolsName() string
	SetSfSymbolsName(value string)
	UnmappedLocalizedName() string
	SetUnmappedLocalizedName(value string)
	UnmappedSfSymbolsName() string
	SetUnmappedSfSymbolsName(value string)
}

// An input for a physical control, such as a button or thumbstick.
//
// is an abstract superclass for specific types of elements that represent controls on a game controller. Use the respective subclasses to either get the input of an element directly or set a handler that the element calls when the user changes a value. This class provides support for common features. For complex elements that have subelements, you can get the containing element using the property. For example, a direction pad ( ) has two axis control and four button subelements. If the user binds a controller element to a system gesture, the system sends the input to the system gesture recognizer first. If it doesn’t recognize a gesture, the system sends the input to your app but with a delay. If it does recognize a gesture, it doesn’t send any input to your app. To change this default behavior, you can set the property to to receive the input simultaneously without delay. Alternatively, set it to to disable the system gesture and receive the input exclusively. Use the property to check whether the user included an element in a system gesture. Use the property to determine whether an element’s input value is a range of values or a discrete digital value.


// An input for a physical control, such as a button or thumbstick.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement
type GCControllerElement struct {
	objectivec.Object
}

// GCControllerElementFrom constructs a [GCControllerElement] from an unsafe.Pointer.
//
// An input for a physical control, such as a button or thumbstick.
func GCControllerElementFrom(ptr unsafe.Pointer) GCControllerElement {
	return GCControllerElement{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GCControllerElementClass) Alloc() GCControllerElement {
	rv := objc.Send[GCControllerElement](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GCControllerElementClass) New() GCControllerElement {
	rv := objc.Send[GCControllerElement](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCControllerElement) Init() GCControllerElement {
	rv := objc.Send[GCControllerElement](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCControllerElement) Autorelease() GCControllerElement {
	rv := objc.Send[GCControllerElement](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCControllerElement creates a new GCControllerElement instance.
func NewGCControllerElement() GCControllerElement {
	return getGCControllerElementClass().New()
}



// The element’s aliases you use when accessing it with the subscript notation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/aliases
func (g_ GCControllerElement) Aliases() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("aliases"))
	return rv
}


// A Boolean value that indicates whether the element provides analog data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/isAnalog
func (g_ GCControllerElement) Analog() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("analog"))
	return rv
}


// The enclosing element for this element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/collection
func (g_ GCControllerElement) Collection() GCControllerElement {
	rv := objc.Send[GCControllerElement](g_.ID, objc.Sel("collection"))
	return rv
}


// The enclosing element for this element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/collection
func (g_ GCControllerElement) SetCollection(value IGCControllerElement) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setCollection:"), value)
}


// A Boolean value that indicates whether the element provides analog data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/isanalog
func (g_ GCControllerElement) IsAnalog() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isAnalog"))
	return rv
}


// A Boolean value that indicates whether the element provides analog data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/isanalog
func (g_ GCControllerElement) SetIsAnalog(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsAnalog:"), value)
}


// A Boolean value that indicates whether the user binds the element to a system gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/isboundtosystemgesture
func (g_ GCControllerElement) IsBoundToSystemGesture() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isBoundToSystemGesture"))
	return rv
}


// A Boolean value that indicates whether the user binds the element to a system gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/isboundtosystemgesture
func (g_ GCControllerElement) SetIsBoundToSystemGesture(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsBoundToSystemGesture:"), value)
}


// The localized name for the element or the remapped element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/localizedname
func (g_ GCControllerElement) LocalizedName() string {
	rv := objc.Send[string](g_.ID, objc.Sel("localizedName"))
	return rv
}


// The localized name for the element or the remapped element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/localizedname
func (g_ GCControllerElement) SetLocalizedName(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLocalizedName:"), objc.String(value))
}


// The preferred state for handling input when the user binds the element to a system gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/preferredsystemgesturestate
func (g_ GCControllerElement) PreferredSystemGestureState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("preferredSystemGestureState"))
	return rv
}


// The preferred state for handling input when the user binds the element to a system gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/preferredsystemgesturestate
func (g_ GCControllerElement) SetPreferredSystemGestureState(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPreferredSystemGestureState:"), value)
}


// A system symbol for the element or the remapped element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/sfsymbolsname
func (g_ GCControllerElement) SfSymbolsName() string {
	rv := objc.Send[string](g_.ID, objc.Sel("sfSymbolsName"))
	return rv
}


// A system symbol for the element or the remapped element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/sfsymbolsname
func (g_ GCControllerElement) SetSfSymbolsName(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSfSymbolsName:"), objc.String(value))
}


// The element’s localized name, not the remapped name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/unmappedlocalizedname
func (g_ GCControllerElement) UnmappedLocalizedName() string {
	rv := objc.Send[string](g_.ID, objc.Sel("unmappedLocalizedName"))
	return rv
}


// The element’s localized name, not the remapped name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/unmappedlocalizedname
func (g_ GCControllerElement) SetUnmappedLocalizedName(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setUnmappedLocalizedName:"), objc.String(value))
}


// The element’s system symbol, not the remapped symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/unmappedsfsymbolsname
func (g_ GCControllerElement) UnmappedSfSymbolsName() string {
	rv := objc.Send[string](g_.ID, objc.Sel("unmappedSfSymbolsName"))
	return rv
}


// The element’s system symbol, not the remapped symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/unmappedsfsymbolsname
func (g_ GCControllerElement) SetUnmappedSfSymbolsName(value string) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setUnmappedSfSymbolsName:"), objc.String(value))
}



