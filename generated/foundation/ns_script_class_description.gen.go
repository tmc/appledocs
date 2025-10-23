// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScriptClassDescription] class.
var (
	ScriptClassDescriptionClass     _ScriptClassDescriptionClass
	ScriptClassDescriptionClassOnce sync.Once
)

func getScriptClassDescriptionClass() _ScriptClassDescriptionClass {
	ScriptClassDescriptionClassOnce.Do(func() {
		ScriptClassDescriptionClass = _ScriptClassDescriptionClass{objc.GetClass("NSScriptClassDescription")}
	})
	return ScriptClassDescriptionClass
}

type _ScriptClassDescriptionClass struct {
	class objc.Class
}

// An interface definition for the [ScriptClassDescription] class.
type IScriptClassDescription interface {
	IClassDescription
	AppleEventCodeForKey(key string) unsafe.Pointer
	ClassDescriptionForKey(key string) ScriptClassDescription
	HasOrderedToManyRelationshipForKey(key string) bool
	HasPropertyForKey(key string) bool
	HasReadablePropertyForKey(key string) bool
	HasWritablePropertyForKey(key string) bool
	IsLocationRequiredToCreateForKey(toManyRelationshipKey string) bool
	IsReadOnlyKey(key string) bool
	KeyWithAppleEventCode(appleEventCode unsafe.Pointer) String
	MatchesAppleEventCode(appleEventCode unsafe.Pointer) bool
	SelectorForCommand(commandDescription IScriptCommandDescription) objc.SEL
	SupportsCommand(commandDescription IScriptCommandDescription) bool
	TypeForKey(key string) String
	AppleEventCode() unsafe.Pointer
	ClassName() string
	DefaultSubcontainerAttributeKey() string
	ImplementationClassName() string
	SuiteName() string
	SuperclassDescription() NSScriptClassDescription
	Superclass() NSScriptClassDescription
	SetSuperclass(value IScriptClassDescription)
}

// A scriptable class that a macOS app supports.
//
// A scriptable application provides scriptability information that describes the commands and objects scripters can use in scripts that target the application. That includes information about the classes those scriptable objects are created from. An application’s scriptability information is collected automatically by an instance of . The registry object creates an for each class it finds and caches these objects in memory. Cocoa scripting uses registry information in handling scripting requests that target the application. A class description instance stores the name, attributes, relationships, and supported commands for a class. For example, a scriptable class for a drawing application might support attributes such as and , relationships such as collections of , , and , and commands such as and . As with many of the classes in Cocoa’s built-in scripting support, your application may never need to directly work with instances of . However, one case where you might need access to a class description is if you override in a scriptable class. For information on how to do this, see in . Another case where your application may need access to class description information is if you override in a specifier class. Although you can subclass , it is unlikely that you would need to do so, or even to create instances of it.


// A scriptable class that a macOS app supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription
type ScriptClassDescription struct {
	ClassDescription
}

// ScriptClassDescriptionFrom constructs a [ScriptClassDescription] from an unsafe.Pointer.
//
// A scriptable class that a macOS app supports.
func ScriptClassDescriptionFrom(ptr unsafe.Pointer) ScriptClassDescription {
	return ScriptClassDescription{
		ClassDescription: ClassDescriptionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _ScriptClassDescriptionClass) Alloc() ScriptClassDescription {
	rv := objc.Send[ScriptClassDescription](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScriptClassDescriptionClass) New() ScriptClassDescription {
	rv := objc.Send[ScriptClassDescription](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScriptClassDescription) Init() ScriptClassDescription {
	rv := objc.Send[ScriptClassDescription](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScriptClassDescription) Autorelease() ScriptClassDescription {
	rv := objc.Send[ScriptClassDescription](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScriptClassDescription creates a new ScriptClassDescription instance.
func NewScriptClassDescription() ScriptClassDescription {
	return getScriptClassDescriptionClass().New()
}



// Returns the class description for the specified class or, if it is not scriptable, for the first superclass that is.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/init(for:)
func NewScriptClassDescriptionForClass(aClass objc.Class) ScriptClassDescription {
	rv := objc.Send[ScriptClassDescription](objc.ID(getScriptClassDescriptionClass().class), objc.Sel("classDescriptionForClass:"), aClass)
	return rv
}


// Initializes and returns a newly allocated instance of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/init(suiteName:className:dictionary:)
func NewScriptClassDescriptionWithSuiteNameClassNameDictionary(suiteName string, className string, classDeclaration objectivec.IObject) ScriptClassDescription {
	instance := getScriptClassDescriptionClass().Alloc()
	rv := objc.Send[ScriptClassDescription](instance.ID, objc.Sel("initWithSuiteName:className:dictionary:"), objc.String(suiteName), objc.String(className), classDeclaration)
	rv.Autorelease()
	return rv
}



// Returns the class description for the specified class or, if it is not scriptable, for the first superclass that is.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/init(for:)
func (sc _ScriptClassDescriptionClass) ClassDescriptionForClass(aClass objc.Class) ScriptClassDescription {
	rv := objc.Send[ScriptClassDescription](objc.ID(sc.class), objc.Sel("classDescriptionForClass:"), aClass)
	return rv
}


// Returns the Apple event code for the specified attribute or relationship in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/appleEventCode(forKey:)
func (s_ ScriptClassDescription) AppleEventCodeForKey(key string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("appleEventCodeForKey:"), objc.String(key))
	return rv
}


// Returns the class description instance for the class type of the specified attribute or relationship.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/forKey(_:)
func (s_ ScriptClassDescription) ClassDescriptionForKey(key string) ScriptClassDescription {
	rv := objc.Send[ScriptClassDescription](s_.ID, objc.Sel("classDescriptionForKey:"), objc.String(key))
	return rv
}


// Returns a Boolean value indicating whether the described class has an ordered to-many relationship identified by the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/hasOrderedToManyRelationship(forKey:)
func (s_ ScriptClassDescription) HasOrderedToManyRelationshipForKey(key string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasOrderedToManyRelationshipForKey:"), objc.String(key))
	return rv
}


// Returns a Boolean value indicating whether the described class has a property identified by the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/hasProperty(forKey:)
func (s_ ScriptClassDescription) HasPropertyForKey(key string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasPropertyForKey:"), objc.String(key))
	return rv
}


// Returns a Boolean value indicating whether the described class has a readable property identified by the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/hasReadableProperty(forKey:)
func (s_ ScriptClassDescription) HasReadablePropertyForKey(key string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasReadablePropertyForKey:"), objc.String(key))
	return rv
}


// Returns a Boolean value indicating whether the described class has a writable property identified by the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/hasWritableProperty(forKey:)
func (s_ ScriptClassDescription) HasWritablePropertyForKey(key string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasWritablePropertyForKey:"), objc.String(key))
	return rv
}


// Returns a Boolean value indicating whether an insertion location must be specified when creating a new object in the specified to-many relationship of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/isLocationRequiredToCreate(forKey:)
func (s_ ScriptClassDescription) IsLocationRequiredToCreateForKey(toManyRelationshipKey string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isLocationRequiredToCreateForKey:"), objc.String(toManyRelationshipKey))
	return rv
}


// Returns a Boolean value indicating whether a specified property in the receiver is read-only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/isReadOnlyKey:
func (s_ ScriptClassDescription) IsReadOnlyKey(key string) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isReadOnlyKey:"), objc.String(key))
	return rv
}


// Given an Apple event code that identifies a property or element class, returns the key for the corresponding attribute, one-to-one relationship, or one-to-many relationship.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/key(withAppleEventCode:)
func (s_ ScriptClassDescription) KeyWithAppleEventCode(appleEventCode unsafe.Pointer) String {
	rv := objc.Send[String](s_.ID, objc.Sel("keyWithAppleEventCode:"), appleEventCode)
	return rv
}


// Returns a Boolean value indicating whether a primary or secondary Apple event code in the receiver matches the passed code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/matchesAppleEventCode(_:)
func (s_ ScriptClassDescription) MatchesAppleEventCode(appleEventCode unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("matchesAppleEventCode:"), appleEventCode)
	return rv
}


// Returns the selector associated with the receiver for the specified command description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/selector(forCommand:)
func (s_ ScriptClassDescription) SelectorForCommand(commandDescription IScriptCommandDescription) objc.SEL {
	rv := objc.Send[objc.SEL](s_.ID, objc.Sel("selectorForCommand:"), commandDescription)
	return rv
}


// Returns a Boolean value indicating whether the receiver or any superclass supports the specified command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/supportsCommand(_:)
func (s_ ScriptClassDescription) SupportsCommand(commandDescription IScriptCommandDescription) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("supportsCommand:"), commandDescription)
	return rv
}


// Returns the name of the declared type of the attribute or relationship identified by the passed key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/type(forKey:)
func (s_ ScriptClassDescription) TypeForKey(key string) String {
	rv := objc.Send[String](s_.ID, objc.Sel("typeForKey:"), objc.String(key))
	return rv
}


// Returns the Apple event code associated with the receiver’s class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/appleEventCode
func (s_ ScriptClassDescription) AppleEventCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("appleEventCode"))
	return rv
}


// Returns the name of the class the receiver describes, as provided at initialization time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/className
func (s_ ScriptClassDescription) ClassName() string {
	rv := objc.Send[string](s_.ID, objc.Sel("className"))
	return rv
}


// Returns the value of the entry of the class dictionary from which the receiver was instantiated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/defaultSubcontainerAttributeKey
func (s_ ScriptClassDescription) DefaultSubcontainerAttributeKey() string {
	rv := objc.Send[string](s_.ID, objc.Sel("defaultSubcontainerAttributeKey"))
	return rv
}


// Returns the name of the Objective-C class instantiated to implement the scripting class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/implementationClassName
func (s_ ScriptClassDescription) ImplementationClassName() string {
	rv := objc.Send[string](s_.ID, objc.Sel("implementationClassName"))
	return rv
}


// Returns the name of the receiver’s suite.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/suiteName
func (s_ ScriptClassDescription) SuiteName() string {
	rv := objc.Send[string](s_.ID, objc.Sel("suiteName"))
	return rv
}


// Returns the class description instance for the superclass of the receiver’s class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScriptClassDescription/superclass
func (s_ ScriptClassDescription) SuperclassDescription() NSScriptClassDescription {
	rv := objc.Send[NSScriptClassDescription](s_.ID, objc.Sel("superclassDescription"))
	return rv
}


// Returns the class description instance for the superclass of the receiver’s class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/superclass
func (s_ ScriptClassDescription) Superclass() NSScriptClassDescription {
	rv := objc.Send[NSScriptClassDescription](s_.ID, objc.Sel("superclass"))
	return rv
}


// Returns the class description instance for the superclass of the receiver’s class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptclassdescription/superclass
func (s_ ScriptClassDescription) SetSuperclass(value IScriptClassDescription) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSuperclass:"), value)
}


