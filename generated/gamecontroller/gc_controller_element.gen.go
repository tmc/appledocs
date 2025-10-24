// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GCControllerElement */


/* debug [class_header]: Header for GCControllerElement */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCControllerElement */
// An interface definition for the [GCControllerElement] class.
type IGCControllerElement interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GCControllerElement */
	// properties:
	Aliases() unsafe.Pointer
	Collection() IGCControllerElement
	Analog() bool
	BoundToSystemGesture() bool
	LocalizedName() objc.IObject /* cross-framework: NSString */
	SetLocalizedName(value objc.IObject /* cross-framework: NSString */)
	PreferredSystemGestureState() GCSystemGestureState
	SetPreferredSystemGestureState(value GCSystemGestureState)
	SfSymbolsName() objc.IObject /* cross-framework: NSString */
	SetSfSymbolsName(value objc.IObject /* cross-framework: NSString */)
	UnmappedLocalizedName() objc.IObject /* cross-framework: NSString */
	SetUnmappedLocalizedName(value objc.IObject /* cross-framework: NSString */)
	UnmappedSfSymbolsName() objc.IObject /* cross-framework: NSString */
	SetUnmappedSfSymbolsName(value objc.IObject /* cross-framework: NSString */)
	IsAnalog() bool
	SetIsAnalog(value bool)
	IsBoundToSystemGesture() bool
	SetIsBoundToSystemGesture(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCControllerElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCControllerElement */
// Alloc allocates a new instance without initialization.
func (gc _GCControllerElementClass) Alloc() GCControllerElement {
	rv := objc.Send[GCControllerElement](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCControllerElement */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCControllerElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCControllerElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCControllerElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCControllerElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCControllerElement */

// The element’s aliases you use when accessing it with the subscript notation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/aliases
func (g_ GCControllerElement) Aliases() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("aliases"))
	return rv
}/* debug [instance_properties/getter]: aliases */


// The enclosing element for this element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/collection
func (g_ GCControllerElement) Collection() IGCControllerElement {
	rv := objc.Send[GCControllerElement](g_.ID, objc.Sel("collection"))
	return rv
}/* debug [instance_properties/getter]: collection */


// A Boolean value that indicates whether the element provides analog data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/isAnalog
func (g_ GCControllerElement) Analog() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("analog"))
	return rv
}/* debug [instance_properties/getter]: analog */


// A Boolean value that indicates whether the user binds the element to a system gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/isBoundToSystemGesture
func (g_ GCControllerElement) BoundToSystemGesture() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("boundToSystemGesture"))
	return rv
}/* debug [instance_properties/getter]: boundToSystemGesture */


// The localized name for the element or the remapped element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/localizedName
func (g_ GCControllerElement) LocalizedName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("localizedName"))
	return rv
}/* debug [instance_properties/getter]: localizedName */


// The localized name for the element or the remapped element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/localizedName
func (g_ GCControllerElement) SetLocalizedName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLocalizedName:"), value)
}/* debug [instance_properties/setter]: localizedName */


// The preferred state for handling input when the user binds the element to a system gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/preferredSystemGestureState
func (g_ GCControllerElement) PreferredSystemGestureState() GCSystemGestureState {
	rv := objc.Send[GCSystemGestureState](g_.ID, objc.Sel("preferredSystemGestureState"))
	return rv
}/* debug [instance_properties/getter]: preferredSystemGestureState */


// The preferred state for handling input when the user binds the element to a system gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/preferredSystemGestureState
func (g_ GCControllerElement) SetPreferredSystemGestureState(value GCSystemGestureState) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setPreferredSystemGestureState:"), value)
}/* debug [instance_properties/setter]: preferredSystemGestureState */


// A system symbol for the element or the remapped element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/sfSymbolsName
func (g_ GCControllerElement) SfSymbolsName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("sfSymbolsName"))
	return rv
}/* debug [instance_properties/getter]: sfSymbolsName */


// A system symbol for the element or the remapped element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/sfSymbolsName
func (g_ GCControllerElement) SetSfSymbolsName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSfSymbolsName:"), value)
}/* debug [instance_properties/setter]: sfSymbolsName */


// The element’s localized name, not the remapped name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/unmappedLocalizedName
func (g_ GCControllerElement) UnmappedLocalizedName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("unmappedLocalizedName"))
	return rv
}/* debug [instance_properties/getter]: unmappedLocalizedName */


// The element’s localized name, not the remapped name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/unmappedLocalizedName
func (g_ GCControllerElement) SetUnmappedLocalizedName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setUnmappedLocalizedName:"), value)
}/* debug [instance_properties/setter]: unmappedLocalizedName */


// The element’s system symbol, not the remapped symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/unmappedSfSymbolsName
func (g_ GCControllerElement) UnmappedSfSymbolsName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("unmappedSfSymbolsName"))
	return rv
}/* debug [instance_properties/getter]: unmappedSfSymbolsName */


// The element’s system symbol, not the remapped symbol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/unmappedSfSymbolsName
func (g_ GCControllerElement) SetUnmappedSfSymbolsName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setUnmappedSfSymbolsName:"), value)
}/* debug [instance_properties/setter]: unmappedSfSymbolsName */


// A Boolean value that indicates whether the element provides analog data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/isanalog
func (g_ GCControllerElement) IsAnalog() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isAnalog"))
	return rv
}/* debug [instance_properties/getter]: isAnalog */


// A Boolean value that indicates whether the element provides analog data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/isanalog
func (g_ GCControllerElement) SetIsAnalog(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsAnalog:"), value)
}/* debug [instance_properties/setter]: isAnalog */


// A Boolean value that indicates whether the user binds the element to a system gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/isboundtosystemgesture
func (g_ GCControllerElement) IsBoundToSystemGesture() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isBoundToSystemGesture"))
	return rv
}/* debug [instance_properties/getter]: isBoundToSystemGesture */


// A Boolean value that indicates whether the user binds the element to a system gesture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccontrollerelement/isboundtosystemgesture
func (g_ GCControllerElement) SetIsBoundToSystemGesture(value bool) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setIsBoundToSystemGesture:"), value)
}/* debug [instance_properties/setter]: isBoundToSystemGesture */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCControllerElement */



