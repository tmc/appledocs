// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNScriptChanges */


/* debug [class_header]: Header for CNScriptChanges */
// The class instance for the [CNScriptChanges] class.
var (
	CNScriptChangesClass     _CNScriptChangesClass
	CNScriptChangesClassOnce sync.Once
)

func getCNScriptChangesClass() _CNScriptChangesClass {
	CNScriptChangesClassOnce.Do(func() {
		CNScriptChangesClass = _CNScriptChangesClass{objc.GetClass("CNScriptChanges")}
	})
	return CNScriptChangesClass
}

type _CNScriptChangesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNScriptChanges */
// An interface definition for the [CNScriptChanges] class.
type ICNScriptChanges interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNScriptChanges */
	// properties:
	AddedDetectionTracks() []CNDetectionTrack
	DataRepresentation() objc.IObject /* cross-framework: NSData */
	FNumber() float32
	UserDecisions() []CNDecision
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNScriptChanges */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNScriptChanges */
// Alloc allocates a new instance without initialization.
func (cc _CNScriptChangesClass) Alloc() CNScriptChanges {
	rv := objc.Send[CNScriptChanges](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNScriptChangesClass) New() CNScriptChanges {
	rv := objc.Send[CNScriptChanges](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNScriptChanges) Init() CNScriptChanges {
	rv := objc.Send[CNScriptChanges](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNScriptChanges) Autorelease() CNScriptChanges {
	rv := objc.Send[CNScriptChanges](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNScriptChanges creates a new CNScriptChanges instance.
func NewCNScriptChanges() CNScriptChanges {
	return getCNScriptChangesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNScriptChanges */
// An object that represents a snapshot of the changes made to a movie script, including the added user decisions and detection tracks.
//
// Use as a snapshot to quickly revert to previously saved edits.


// An object that represents a snapshot of the changes made to a movie script, including the added user decisions and detection tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScriptChanges
type CNScriptChanges struct {
	objectivec.Object
}

// CNScriptChangesFrom constructs a [CNScriptChanges] from an unsafe.Pointer.
//
// An object that represents a snapshot of the changes made to a movie script, including the added user decisions and detection tracks.
func CNScriptChangesFrom(ptr unsafe.Pointer) CNScriptChanges {
	return CNScriptChanges{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNScriptChanges */

// Creates a previously saved data representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScriptChanges/initWithDataRepresentation:
func NewCNScriptChangesWithDataRepresentation(dataRepresentation objc.IObject /* cross-framework: NSData */) CNScriptChanges {
	instance := getCNScriptChangesClass().Alloc()
	rv := objc.Send[CNScriptChanges](instance.ID, objc.Sel("initWithDataRepresentation:"), dataRepresentation)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNScriptChangesWithDataRepresentation */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNScriptChanges */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNScriptChanges */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNScriptChanges */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNScriptChanges */

// All detection tracks added since recording the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScriptChanges/addedDetectionTracks
func (c_ CNScriptChanges) AddedDetectionTracks() []CNDetectionTrack {
	rv := objc.Send[[]CNDetectionTrack](c_.ID, objc.Sel("addedDetectionTracks"))
	return rv
}/* debug [instance_properties/getter]: addedDetectionTracks */


// Persistent data representation of changes for later restoration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScriptChanges/dataRepresentation
func (c_ CNScriptChanges) DataRepresentation() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("dataRepresentation"))
	return rv
}/* debug [instance_properties/getter]: dataRepresentation */


// The f-stop number to apply to the entire movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScriptChanges/fNumber
func (c_ CNScriptChanges) FNumber() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("fNumber"))
	return rv
}/* debug [instance_properties/getter]: fNumber */


// All active user decisions, including those made at recording time, unless removed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScriptChanges/userDecisions
func (c_ CNScriptChanges) UserDecisions() []CNDecision {
	rv := objc.Send[[]CNDecision](c_.ID, objc.Sel("userDecisions"))
	return rv
}/* debug [instance_properties/getter]: userDecisions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNScriptChanges */


