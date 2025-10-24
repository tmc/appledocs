// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSStoryboardSegue */


/* debug [class_header]: Header for NSStoryboardSegue */
// The class instance for the [StoryboardSegue] class.
var (
	StoryboardSegueClass     _StoryboardSegueClass
	StoryboardSegueClassOnce sync.Once
)

func getStoryboardSegueClass() _StoryboardSegueClass {
	StoryboardSegueClassOnce.Do(func() {
		StoryboardSegueClass = _StoryboardSegueClass{objc.GetClass("NSStoryboardSegue")}
	})
	return StoryboardSegueClass
}

type _StoryboardSegueClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for StoryboardSegue */
// An interface definition for the [StoryboardSegue] class.
type IStoryboardSegue interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for StoryboardSegue */
	// properties:
	DestinationController() objc.ID
	Identifier() StoryboardSegueIdentifier /* typedef */
	SourceController() objc.ID
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for StoryboardSegue */
	// methods:
	Perform()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for StoryboardSegue */
// Alloc allocates a new instance without initialization.
func (sc _StoryboardSegueClass) Alloc() StoryboardSegue {
	rv := objc.Send[StoryboardSegue](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StoryboardSegueClass) New() StoryboardSegue {
	rv := objc.Send[StoryboardSegue](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StoryboardSegue) Init() StoryboardSegue {
	rv := objc.Send[StoryboardSegue](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StoryboardSegue) Autorelease() StoryboardSegue {
	rv := objc.Send[StoryboardSegue](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStoryboardSegue creates a new StoryboardSegue instance.
func NewStoryboardSegue() StoryboardSegue {
	return getStoryboardSegueClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for StoryboardSegue */
// A transition or containment relationship between two scenes in a storyboard.
//
// In this context, a is a view controller or a window controller and a is an instance of the class. A storyboard segue has a procedural notion of being invoked, known in the API as being . You can take advantage of hooks into the segue performance process by way of the protocol. You do not create storyboard segue objects directly. Instead, the system creates them as needed as segues are invoked. To run code during initialization and performance of a segue, override the and methods. You can initiate a segue programmatically with the method of the protocol. For example, you might do this to transition from a scene in one storyboard file to a scene in another.


// A transition or containment relationship between two scenes in a storyboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue
type StoryboardSegue struct {
	objectivec.Object
}

// StoryboardSegueFrom constructs a [StoryboardSegue] from an unsafe.Pointer.
//
// A transition or containment relationship between two scenes in a storyboard.
func StoryboardSegueFrom(ptr unsafe.Pointer) StoryboardSegue {
	return StoryboardSegue{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for StoryboardSegue */

// The designated initializer for a storyboard segue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue/init(identifier:source:destination:)
func NewStoryboardSegueWithIdentifierSourceDestination(identifier StoryboardSegueIdentifier /* typedef */, sourceController objc.IObject, destinationController objc.IObject) StoryboardSegue {
	instance := getStoryboardSegueClass().Alloc()
	rv := objc.Send[StoryboardSegue](instance.ID, objc.Sel("initWithIdentifier:source:destination:"), identifier, sourceController, destinationController)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewStoryboardSegueWithIdentifierSourceDestination */


// Creates a storyboard segue and a block used when the segue is performed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue/init(identifier:source:destination:performHandler:)
func NewStoryboardSegueWithIdentifierSourceDestinationPerformHandler(identifier StoryboardSegueIdentifier /* typedef */, sourceController objc.IObject, destinationController objc.IObject, performHandler unsafe.Pointer) StoryboardSegue {
	rv := objc.Send[StoryboardSegue](objc.ID(getStoryboardSegueClass().class), objc.Sel("segueWithIdentifier:source:destination:performHandler:"), identifier, sourceController, destinationController, performHandler)
	return rv
}/* debug [class_init_methods/constructor]: NewStoryboardSegueWithIdentifierSourceDestinationPerformHandler */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for StoryboardSegue */

// Creates a storyboard segue and a block used when the segue is performed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue/init(identifier:source:destination:performHandler:)
func (sc _StoryboardSegueClass) SegueWithIdentifierSourceDestinationPerformHandler(identifier StoryboardSegueIdentifier /* typedef */, sourceController objc.IObject, destinationController objc.IObject, performHandler unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("segueWithIdentifier:source:destination:performHandler:"), identifier, sourceController, destinationController, performHandler)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SegueWithIdentifierSourceDestinationPerformHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for StoryboardSegue */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for StoryboardSegue */

// Performs a visual transition from one controller to another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue/perform()
func (s_ StoryboardSegue) Perform() {
	objc.Send[objc.ID](s_.ID, objc.Sel("perform"))
}/* debug [instance_methods/method]: Perform */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for StoryboardSegue */

// The ending/contained view controller or window controller for the storyboard segue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue/destinationController
func (s_ StoryboardSegue) DestinationController() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("destinationController"))
	return rv
}/* debug [instance_properties/getter]: destinationController */


// An optional, unique identifier for the storyboard segue that you can specify using the Identity inspector in Interface Builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue/identifier-swift.property
func (s_ StoryboardSegue) Identifier() StoryboardSegueIdentifier /* typedef */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The starting/containing view controller or window controller for the storyboard segue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStoryboardSegue/sourceController
func (s_ StoryboardSegue) SourceController() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("sourceController"))
	return rv
}/* debug [instance_properties/getter]: sourceController */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSStoryboardSegue */


