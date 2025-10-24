// Code generated from Apple documentation for GameSave. DO NOT EDIT.

package gamesave

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GSSyncedDirectoryState */


/* debug [class_header]: Header for GSSyncedDirectoryState */
// The class instance for the [GSSyncedDirectoryState] class.
var (
	GSSyncedDirectoryStateClass     _GSSyncedDirectoryStateClass
	GSSyncedDirectoryStateClassOnce sync.Once
)

func getGSSyncedDirectoryStateClass() _GSSyncedDirectoryStateClass {
	GSSyncedDirectoryStateClassOnce.Do(func() {
		GSSyncedDirectoryStateClass = _GSSyncedDirectoryStateClass{objc.GetClass("GSSyncedDirectoryState")}
	})
	return GSSyncedDirectoryStateClass
}

type _GSSyncedDirectoryStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GSSyncedDirectoryState */
// An interface definition for the [GSSyncedDirectoryState] class.
type IGSSyncedDirectoryState interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GSSyncedDirectoryState */
	// properties:
	ConflictedVersions() []GSSyncedDirectoryVersion
	Error() objc.IObject /* cross-framework: Error */
	State() GSSyncState
	Url() objc.IObject /* cross-framework: NSURL */
	DirectoryState() IGSSyncedDirectoryState
	SetDirectoryState(value IGSSyncedDirectoryState)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GSSyncedDirectoryState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GSSyncedDirectoryState */
// Alloc allocates a new instance without initialization.
func (gc _GSSyncedDirectoryStateClass) Alloc() GSSyncedDirectoryState {
	rv := objc.Send[GSSyncedDirectoryState](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GSSyncedDirectoryStateClass) New() GSSyncedDirectoryState {
	rv := objc.Send[GSSyncedDirectoryState](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GSSyncedDirectoryState) Init() GSSyncedDirectoryState {
	rv := objc.Send[GSSyncedDirectoryState](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GSSyncedDirectoryState) Autorelease() GSSyncedDirectoryState {
	rv := objc.Send[GSSyncedDirectoryState](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGSSyncedDirectoryState creates a new GSSyncedDirectoryState instance.
func NewGSSyncedDirectoryState() GSSyncedDirectoryState {
	return getGSSyncedDirectoryStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GSSyncedDirectoryState */
// Represents the state and its associated properties of the directory
//
// Use the property to determine the validity of the other properties


// Represents the state and its associated properties of the directory
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectoryState
type GSSyncedDirectoryState struct {
	objectivec.Object
}

// GSSyncedDirectoryStateFrom constructs a [GSSyncedDirectoryState] from an unsafe.Pointer.
//
// Represents the state and its associated properties of the directory
func GSSyncedDirectoryStateFrom(ptr unsafe.Pointer) GSSyncedDirectoryState {
	return GSSyncedDirectoryState{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GSSyncedDirectoryState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GSSyncedDirectoryState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GSSyncedDirectoryState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GSSyncedDirectoryState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GSSyncedDirectoryState */

// The conflicting versions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectoryState/conflictedVersions
func (g_ GSSyncedDirectoryState) ConflictedVersions() []GSSyncedDirectoryVersion {
	rv := objc.Send[[]GSSyncedDirectoryVersion](g_.ID, objc.Sel("conflictedVersions"))
	return rv
}/* debug [instance_properties/getter]: conflictedVersions */


// The error preventing you from using the directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectoryState/error
func (g_ GSSyncedDirectoryState) Error() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](g_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */


// Specifies the current state of the directory
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectoryState/state
func (g_ GSSyncedDirectoryState) State() GSSyncState {
	rv := objc.Send[GSSyncState](g_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */


// The URL of a directory to read and write game-save data in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectoryState/url
func (g_ GSSyncedDirectoryState) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](g_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// The state of the directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamesave/gssynceddirectory/directorystate
func (g_ GSSyncedDirectoryState) DirectoryState() IGSSyncedDirectoryState {
	rv := objc.Send[GSSyncedDirectoryState](g_.ID, objc.Sel("directoryState"))
	return rv
}/* debug [instance_properties/getter]: directoryState */


// The state of the directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamesave/gssynceddirectory/directorystate
func (g_ GSSyncedDirectoryState) SetDirectoryState(value IGSSyncedDirectoryState) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDirectoryState:"), value)
}/* debug [instance_properties/setter]: directoryState */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GSSyncedDirectoryState */



