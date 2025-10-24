// Code generated from Apple documentation for GameSave. DO NOT EDIT.

package gamesave

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GSSyncedDirectoryVersion */


/* debug [class_header]: Header for GSSyncedDirectoryVersion */
// The class instance for the [GSSyncedDirectoryVersion] class.
var (
	GSSyncedDirectoryVersionClass     _GSSyncedDirectoryVersionClass
	GSSyncedDirectoryVersionClassOnce sync.Once
)

func getGSSyncedDirectoryVersionClass() _GSSyncedDirectoryVersionClass {
	GSSyncedDirectoryVersionClassOnce.Do(func() {
		GSSyncedDirectoryVersionClass = _GSSyncedDirectoryVersionClass{objc.GetClass("GSSyncedDirectoryVersion")}
	})
	return GSSyncedDirectoryVersionClass
}

type _GSSyncedDirectoryVersionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GSSyncedDirectoryVersion */
// An interface definition for the [GSSyncedDirectoryVersion] class.
type IGSSyncedDirectoryVersion interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GSSyncedDirectoryVersion */
	// properties:
	Description() objc.IObject /* cross-framework: NSString */
	IsLocal() bool
	LocalizedNameOfSavingComputer() objc.IObject /* cross-framework: NSString */
	ModifiedDate() objc.IObject /* cross-framework: NSDate */
	Url() objc.IObject /* cross-framework: NSURL */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GSSyncedDirectoryVersion */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GSSyncedDirectoryVersion */
// Alloc allocates a new instance without initialization.
func (gc _GSSyncedDirectoryVersionClass) Alloc() GSSyncedDirectoryVersion {
	rv := objc.Send[GSSyncedDirectoryVersion](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GSSyncedDirectoryVersionClass) New() GSSyncedDirectoryVersion {
	rv := objc.Send[GSSyncedDirectoryVersion](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GSSyncedDirectoryVersion) Init() GSSyncedDirectoryVersion {
	rv := objc.Send[GSSyncedDirectoryVersion](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GSSyncedDirectoryVersion) Autorelease() GSSyncedDirectoryVersion {
	rv := objc.Send[GSSyncedDirectoryVersion](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGSSyncedDirectoryVersion creates a new GSSyncedDirectoryVersion instance.
func NewGSSyncedDirectoryVersion() GSSyncedDirectoryVersion {
	return getGSSyncedDirectoryVersionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GSSyncedDirectoryVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectoryVersion
type GSSyncedDirectoryVersion struct {
	objectivec.Object
}

// GSSyncedDirectoryVersionFrom constructs a [GSSyncedDirectoryVersion] from an unsafe.Pointer.
func GSSyncedDirectoryVersionFrom(ptr unsafe.Pointer) GSSyncedDirectoryVersion {
	return GSSyncedDirectoryVersion{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GSSyncedDirectoryVersion *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GSSyncedDirectoryVersion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GSSyncedDirectoryVersion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GSSyncedDirectoryVersion */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GSSyncedDirectoryVersion */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectoryVersion/description
func (g_ GSSyncedDirectoryVersion) Description() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("description"))
	return rv
}/* debug [instance_properties/getter]: description */


// if the directory version is local; otherwise .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectoryVersion/isLocal
func (g_ GSSyncedDirectoryVersion) IsLocal() bool {
	rv := objc.Send[bool](g_.ID, objc.Sel("isLocal"))
	return rv
}/* debug [instance_properties/getter]: isLocal */


// The localized name of the device that saved this version.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectoryVersion/localizedNameOfSavingComputer
func (g_ GSSyncedDirectoryVersion) LocalizedNameOfSavingComputer() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("localizedNameOfSavingComputer"))
	return rv
}/* debug [instance_properties/getter]: localizedNameOfSavingComputer */


// The date that this version was last modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectoryVersion/modifiedDate
func (g_ GSSyncedDirectoryVersion) ModifiedDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](g_.ID, objc.Sel("modifiedDate"))
	return rv
}/* debug [instance_properties/getter]: modifiedDate */


// The URL of a directory where you read and write game-save data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectoryVersion/url
func (g_ GSSyncedDirectoryVersion) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](g_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GSSyncedDirectoryVersion */






