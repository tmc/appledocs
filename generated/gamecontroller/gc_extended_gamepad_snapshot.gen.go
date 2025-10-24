// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class GCExtendedGamepadSnapshot */


/* debug [class_header]: Header for GCExtendedGamepadSnapshot */
// The class instance for the [GCExtendedGamepadSnapshot] class.
var (
	GCExtendedGamepadSnapshotClass     _GCExtendedGamepadSnapshotClass
	GCExtendedGamepadSnapshotClassOnce sync.Once
)

func getGCExtendedGamepadSnapshotClass() _GCExtendedGamepadSnapshotClass {
	GCExtendedGamepadSnapshotClassOnce.Do(func() {
		GCExtendedGamepadSnapshotClass = _GCExtendedGamepadSnapshotClass{objc.GetClass("GCExtendedGamepadSnapshot")}
	})
	return GCExtendedGamepadSnapshotClass
}

type _GCExtendedGamepadSnapshotClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCExtendedGamepadSnapshot */
// An interface definition for the [GCExtendedGamepadSnapshot] class.
type IGCExtendedGamepadSnapshot interface {
	IGCExtendedGamepad
	
/* debug [class_interface_properties]: Properties for GCExtendedGamepadSnapshot */
	// properties:
	SnapshotData() objc.IObject /* cross-framework: NSData */
	SetSnapshotData(value objc.IObject /* cross-framework: NSData */)
	GCCurrentExtendedGamepadSnapshotDataVersion() GCExtendedGamepadSnapshotDataVersion
	GCCurrentMicroGamepadSnapshotDataVersion() GCMicroGamepadSnapshotDataVersion
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCExtendedGamepadSnapshot */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCExtendedGamepadSnapshot */
// Alloc allocates a new instance without initialization.
func (gc _GCExtendedGamepadSnapshotClass) Alloc() GCExtendedGamepadSnapshot {
	rv := objc.Send[GCExtendedGamepadSnapshot](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCExtendedGamepadSnapshotClass) New() GCExtendedGamepadSnapshot {
	rv := objc.Send[GCExtendedGamepadSnapshot](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCExtendedGamepadSnapshot) Init() GCExtendedGamepadSnapshot {
	rv := objc.Send[GCExtendedGamepadSnapshot](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCExtendedGamepadSnapshot) Autorelease() GCExtendedGamepadSnapshot {
	rv := objc.Send[GCExtendedGamepadSnapshot](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCExtendedGamepadSnapshot creates a new GCExtendedGamepadSnapshot instance.
func NewGCExtendedGamepadSnapshot() GCExtendedGamepadSnapshot {
	return getGCExtendedGamepadSnapshotClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCExtendedGamepadSnapshot */
// A recording of all of the values provided by a object.
//
// To create a gamepad snapshot, call the method on a object. The class is a subclass of the class, so you use the parent class’s properties to read the individual element values. The snapshot is stored in a device independent format. To get the flattened data representation of the snapshot data, read the property.


// A recording of all of the values provided by a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepadSnapshot
type GCExtendedGamepadSnapshot struct {
	GCExtendedGamepad
}

// GCExtendedGamepadSnapshotFrom constructs a [GCExtendedGamepadSnapshot] from an unsafe.Pointer.
//
// A recording of all of the values provided by a object.
func GCExtendedGamepadSnapshotFrom(ptr unsafe.Pointer) GCExtendedGamepadSnapshot {
	return GCExtendedGamepadSnapshot{
		GCExtendedGamepad: GCExtendedGamepadFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCExtendedGamepadSnapshot */

// Initializes a snapshot object associated with a specific controller using a flattened data representation obtained from another snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepadSnapshot/init(controller:snapshotData:)
func NewGCExtendedGamepadSnapshotWithControllerSnapshotData(controller IGCController, data objc.IObject /* cross-framework: NSData */) GCExtendedGamepadSnapshot {
	instance := getGCExtendedGamepadSnapshotClass().Alloc()
	rv := objc.Send[GCExtendedGamepadSnapshot](instance.ID, objc.Sel("initWithController:snapshotData:"), controller, data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGCExtendedGamepadSnapshotWithControllerSnapshotData */


// Initializes a snapshot object with the flattened data representation obtained from another snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepadSnapshot/init(snapshotData:)
func NewGCExtendedGamepadSnapshotWithSnapshotData(data objc.IObject /* cross-framework: NSData */) GCExtendedGamepadSnapshot {
	instance := getGCExtendedGamepadSnapshotClass().Alloc()
	rv := objc.Send[GCExtendedGamepadSnapshot](instance.ID, objc.Sel("initWithSnapshotData:"), data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGCExtendedGamepadSnapshotWithSnapshotData */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCExtendedGamepadSnapshot */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCExtendedGamepadSnapshot */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCExtendedGamepadSnapshot */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCExtendedGamepadSnapshot */

// Flattens a snapshot into an archivable memory representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepadSnapshot/snapshotData
func (g_ GCExtendedGamepadSnapshot) SnapshotData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](g_.ID, objc.Sel("snapshotData"))
	return rv
}/* debug [instance_properties/getter]: snapshotData */


// Flattens a snapshot into an archivable memory representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepadSnapshot/snapshotData
func (g_ GCExtendedGamepadSnapshot) SetSnapshotData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSnapshotData:"), value)
}/* debug [instance_properties/setter]: snapshotData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccurrentextendedgamepadsnapshotdataversion
func (g_ GCExtendedGamepadSnapshot) GCCurrentExtendedGamepadSnapshotDataVersion() GCExtendedGamepadSnapshotDataVersion {
	rv := objc.Send[GCExtendedGamepadSnapshotDataVersion](g_.ID, objc.Sel("GCCurrentExtendedGamepadSnapshotDataVersion"))
	return rv
}/* debug [instance_properties/getter]: GCCurrentExtendedGamepadSnapshotDataVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccurrentmicrogamepadsnapshotdataversion
func (g_ GCExtendedGamepadSnapshot) GCCurrentMicroGamepadSnapshotDataVersion() GCMicroGamepadSnapshotDataVersion {
	rv := objc.Send[GCMicroGamepadSnapshotDataVersion](g_.ID, objc.Sel("GCCurrentMicroGamepadSnapshotDataVersion"))
	return rv
}/* debug [instance_properties/getter]: GCCurrentMicroGamepadSnapshotDataVersion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCExtendedGamepadSnapshot */


