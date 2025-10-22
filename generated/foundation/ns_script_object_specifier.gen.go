// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScriptObjectSpecifier] class.
var (
	ScriptObjectSpecifierClass     _ScriptObjectSpecifierClass
	ScriptObjectSpecifierClassOnce sync.Once
)

func getScriptObjectSpecifierClass() _ScriptObjectSpecifierClass {
	ScriptObjectSpecifierClassOnce.Do(func() {
		ScriptObjectSpecifierClass = _ScriptObjectSpecifierClass{objc.GetClass("NSScriptObjectSpecifier")}
	})
	return ScriptObjectSpecifierClass
}

type _ScriptObjectSpecifierClass struct {
	class objc.Class
}

// An interface definition for the [ScriptObjectSpecifier] class.
type IScriptObjectSpecifier interface {
	objectivec.IObject
}

// An abstract class used to represent natural language expressions.
//
// is the abstract superclass for classes that instantiate objects called “object specifiers.” An object specifier represents an AppleScript reference form, which is a natural-language expression such as or or . The scripting system maps these words or phrases to attributes and relationships of scriptable objects. A reference form rarely occurs in isolation; usually a script statement consists of a series of reference forms preceded by a command and typically connected to each other by , such as: The expression specifies a location in the application’s AppleScript object model—the objects the application makes available to scripters. The classes of objects in the object model often closely match the classes of actual objects in the application, but they are not required to. An object specifier locates objects in the running application that correspond to the specified object model objects. Your application typically creates object specifiers when it implements the method for its scriptable classes. That method is defined by the NSScriptObjectSpecifiers protocol. It is unlikely that you would ever need to create your own subclass of ; the set of valid AppleScript reference forms is determined by Apple Computer and object specifier classes are already implemented for this set. If for some reason you do need to create a subclass, you must override the primitive method to return indices to the elements within the container whose values are matched with the child specifier’s key. In addition, you probably need to declare any special instance variables and implement an initializer that invokes super’s designated initializer, , and initializes these variables. For a comprehensive treatment of object specifiers, including sample code, see in .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier
type ScriptObjectSpecifier struct {
	objectivec.Object
}

// ScriptObjectSpecifierFrom constructs a [ScriptObjectSpecifier] from an unsafe.Pointer.
//
// An abstract class used to represent natural language expressions.
func ScriptObjectSpecifierFrom(ptr unsafe.Pointer) ScriptObjectSpecifier {
	return ScriptObjectSpecifier{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _ScriptObjectSpecifierClass) Alloc() ScriptObjectSpecifier {
	rv := objc.Send[ScriptObjectSpecifier](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScriptObjectSpecifierClass) New() ScriptObjectSpecifier {
	rv := objc.Send[ScriptObjectSpecifier](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScriptObjectSpecifier) Init() ScriptObjectSpecifier {
	rv := objc.Send[ScriptObjectSpecifier](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScriptObjectSpecifier) Autorelease() ScriptObjectSpecifier {
	rv := objc.Send[ScriptObjectSpecifier](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScriptObjectSpecifier creates a new ScriptObjectSpecifier instance.
func NewScriptObjectSpecifier() ScriptObjectSpecifier {
	return getScriptObjectSpecifierClass().New()
}




// Returns an object initialized with the given attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/init(containerClassDescription:containerSpecifier:key:)
func NewScriptObjectSpecifierWithContainerClassDescriptionContainerSpecifierKey(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property string) ScriptObjectSpecifier {
	instance := getScriptObjectSpecifierClass().Alloc()
	rv := objc.Send[ScriptObjectSpecifier](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:"), classDesc, container, objc.String(property))
	rv.Autorelease()
	return rv
}



// Returns an object initialized with a given container specifier and key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptObjectSpecifier/init(containerSpecifier:key:)
func NewScriptObjectSpecifierWithContainerSpecifierKey(container IScriptObjectSpecifier, property string) ScriptObjectSpecifier {
	instance := getScriptObjectSpecifierClass().Alloc()
	rv := objc.Send[ScriptObjectSpecifier](instance.ID, objc.Sel("initWithContainerSpecifier:key:"), container, objc.String(property))
	rv.Autorelease()
	return rv
}


// Sets the receiver’s child reference.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/child
func (s_ ScriptObjectSpecifier) Child() NSScriptObjectSpecifier {
	rv := objc.Send[NSScriptObjectSpecifier](s_.ID, objc.Sel("child"))
	return rv
}


// SetChild sets the value of the child property.
// Sets the receiver’s child reference.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/child
func (s_ ScriptObjectSpecifier) SetChild(value IScriptObjectSpecifier) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setChild:"), value)
}

// Sets the container specifier of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/container
func (s_ ScriptObjectSpecifier) Container() NSScriptObjectSpecifier {
	rv := objc.Send[NSScriptObjectSpecifier](s_.ID, objc.Sel("container"))
	return rv
}


// SetContainer sets the value of the container property.
// Sets the container specifier of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/container
func (s_ ScriptObjectSpecifier) SetContainer(value IScriptObjectSpecifier) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setContainer:"), value)
}

// Sets the class description of the receiver’s container specifier to a given specifier.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/containerclassdescription
func (s_ ScriptObjectSpecifier) ContainerClassDescription() NSScriptClassDescription {
	rv := objc.Send[NSScriptClassDescription](s_.ID, objc.Sel("containerClassDescription"))
	return rv
}


// SetContainerClassDescription sets the value of the containerClassDescription property.
// Sets the class description of the receiver’s container specifier to a given specifier.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/containerclassdescription
func (s_ ScriptObjectSpecifier) SetContainerClassDescription(value IScriptClassDescription) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setContainerClassDescription:"), value)
}

// Sets whether the receiver’s container should be an object involved in a filter reference or the top-level object.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/containerisobjectbeingtested
func (s_ ScriptObjectSpecifier) ContainerIsObjectBeingTested() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("containerIsObjectBeingTested"))
	return rv
}


// SetContainerIsObjectBeingTested sets the value of the containerIsObjectBeingTested property.
// Sets whether the receiver’s container should be an object involved in a filter reference or the top-level object.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/containerisobjectbeingtested
func (s_ ScriptObjectSpecifier) SetContainerIsObjectBeingTested(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setContainerIsObjectBeingTested:"), value)
}

// Sets whether the receiver’s container is to be the container for a range specifier or a top-level object.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/containerisrangecontainerobject
func (s_ ScriptObjectSpecifier) ContainerIsRangeContainerObject() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("containerIsRangeContainerObject"))
	return rv
}


// SetContainerIsRangeContainerObject sets the value of the containerIsRangeContainerObject property.
// Sets whether the receiver’s container is to be the container for a range specifier or a top-level object.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/containerisrangecontainerobject
func (s_ ScriptObjectSpecifier) SetContainerIsRangeContainerObject(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setContainerIsRangeContainerObject:"), value)
}

// Returns an Apple event descriptor that represents the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/descriptor
func (s_ ScriptObjectSpecifier) Descriptor() NSAppleEventDescriptor {
	rv := objc.Send[NSAppleEventDescriptor](s_.ID, objc.Sel("descriptor"))
	return rv
}


// SetDescriptor sets the value of the descriptor property.
// Returns an Apple event descriptor that represents the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/descriptor
func (s_ ScriptObjectSpecifier) SetDescriptor(value IAppleEventDescriptor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDescriptor:"), value)
}

// Returns the object specifier in which an evaluation error occurred.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/evaluationerror
func (s_ ScriptObjectSpecifier) EvaluationError() NSScriptObjectSpecifier {
	rv := objc.Send[NSScriptObjectSpecifier](s_.ID, objc.Sel("evaluationError"))
	return rv
}


// SetEvaluationError sets the value of the evaluationError property.
// Returns the object specifier in which an evaluation error occurred.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/evaluationerror
func (s_ ScriptObjectSpecifier) SetEvaluationError(value IScriptObjectSpecifier) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEvaluationError:"), value)
}

// Sets the value of the evaluation error.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/evaluationerrornumber
func (s_ ScriptObjectSpecifier) EvaluationErrorNumber() int {
	rv := objc.Send[int](s_.ID, objc.Sel("evaluationErrorNumber"))
	return rv
}


// SetEvaluationErrorNumber sets the value of the evaluationErrorNumber property.
// Sets the value of the evaluation error.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/evaluationerrornumber
func (s_ ScriptObjectSpecifier) SetEvaluationErrorNumber(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEvaluationErrorNumber:"), value)
}

// Sets the key of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/key
func (s_ ScriptObjectSpecifier) Key() string {
	rv := objc.Send[string](s_.ID, objc.Sel("key"))
	return rv
}


// SetKey sets the value of the key property.
// Sets the key of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/key
func (s_ ScriptObjectSpecifier) SetKey(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setKey:"), objc.String(value))
}

// Returns the class description of the objects specified by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/keyclassdescription
func (s_ ScriptObjectSpecifier) KeyClassDescription() NSScriptClassDescription {
	rv := objc.Send[NSScriptClassDescription](s_.ID, objc.Sel("keyClassDescription"))
	return rv
}


// SetKeyClassDescription sets the value of the keyClassDescription property.
// Returns the class description of the objects specified by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/keyclassdescription
func (s_ ScriptObjectSpecifier) SetKeyClassDescription(value IScriptClassDescription) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setKeyClassDescription:"), value)
}

// Returns the actual object represented by the nested series of object specifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/objectsbyevaluatingspecifier
func (s_ ScriptObjectSpecifier) ObjectsByEvaluatingSpecifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("objectsByEvaluatingSpecifier"))
	return rv
}


// SetObjectsByEvaluatingSpecifier sets the value of the objectsByEvaluatingSpecifier property.
// Returns the actual object represented by the nested series of object specifiers.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptobjectspecifier/objectsbyevaluatingspecifier
func (s_ ScriptObjectSpecifier) SetObjectsByEvaluatingSpecifier(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setObjectsByEvaluatingSpecifier:"), value)
}


