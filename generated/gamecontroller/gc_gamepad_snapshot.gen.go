// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class GCGamepadSnapshot */


/* debug [class_header]: Header for GCGamepadSnapshot */
// The class instance for the [GCGamepadSnapshot] class.
var (
	GCGamepadSnapshotClass     _GCGamepadSnapshotClass
	GCGamepadSnapshotClassOnce sync.Once
)

func getGCGamepadSnapshotClass() _GCGamepadSnapshotClass {
	GCGamepadSnapshotClassOnce.Do(func() {
		GCGamepadSnapshotClass = _GCGamepadSnapshotClass{objc.GetClass("GCGamepadSnapshot")}
	})
	return GCGamepadSnapshotClass
}

type _GCGamepadSnapshotClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCGamepadSnapshot */
// An interface definition for the [GCGamepadSnapshot] class.
type IGCGamepadSnapshot interface {
	IGCGamepad
	
/* debug [class_interface_properties]: Properties for GCGamepadSnapshot */
	// properties:
	SnapshotData() objc.IObject /* cross-framework: NSData */
	SetSnapshotData(value objc.IObject /* cross-framework: NSData */)
	GCCurrentExtendedGamepadSnapshotDataVersion() GCExtendedGamepadSnapshotDataVersion
	GCCurrentMicroGamepadSnapshotDataVersion() GCMicroGamepadSnapshotDataVersion
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCGamepadSnapshot */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCGamepadSnapshot */
// Alloc allocates a new instance without initialization.
func (gc _GCGamepadSnapshotClass) Alloc() GCGamepadSnapshot {
	rv := objc.Send[GCGamepadSnapshot](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCGamepadSnapshotClass) New() GCGamepadSnapshot {
	rv := objc.Send[GCGamepadSnapshot](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCGamepadSnapshot) Init() GCGamepadSnapshot {
	rv := objc.Send[GCGamepadSnapshot](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCGamepadSnapshot) Autorelease() GCGamepadSnapshot {
	rv := objc.Send[GCGamepadSnapshot](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCGamepadSnapshot creates a new GCGamepadSnapshot instance.
func NewGCGamepadSnapshot() GCGamepadSnapshot {
	return getGCGamepadSnapshotClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCGamepadSnapshot */
// A recording of all of the values provided by a object.
//
// To create a gamepad snapshot, call the method on a object. The class is a subclass of the class, so you use the parent class’s properties to read the individual element values. The snapshot is stored in a device independent format. To get the flattened data representation of the snapshot data, read the property.


// A recording of all of the values provided by a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGamepadSnapshot
type GCGamepadSnapshot struct {
	GCGamepad
}

// GCGamepadSnapshotFrom constructs a [GCGamepadSnapshot] from an unsafe.Pointer.
//
// A recording of all of the values provided by a object.
func GCGamepadSnapshotFrom(ptr unsafe.Pointer) GCGamepadSnapshot {
	return GCGamepadSnapshot{
		GCGamepad: GCGamepadFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCGamepadSnapshot */

// Initializes a snapshot object associated with a specific controller using a flattened data representation obtained from another snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGamepadSnapshot/init(controller:snapshotData:)
func NewGCGamepadSnapshotWithControllerSnapshotData(controller IGCController, data objc.IObject /* cross-framework: NSData */) GCGamepadSnapshot {
	instance := getGCGamepadSnapshotClass().Alloc()
	rv := objc.Send[GCGamepadSnapshot](instance.ID, objc.Sel("initWithController:snapshotData:"), controller, data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGCGamepadSnapshotWithControllerSnapshotData */


// Initializes a snapshot object with the flattened data representation obtained from another snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGamepadSnapshot/init(snapshotData:)
func NewGCGamepadSnapshotWithSnapshotData(data objc.IObject /* cross-framework: NSData */) GCGamepadSnapshot {
	instance := getGCGamepadSnapshotClass().Alloc()
	rv := objc.Send[GCGamepadSnapshot](instance.ID, objc.Sel("initWithSnapshotData:"), data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGCGamepadSnapshotWithSnapshotData */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCGamepadSnapshot */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCGamepadSnapshot */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCGamepadSnapshot */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCGamepadSnapshot */

// The flattened control input values for the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGamepadSnapshot/snapshotData
func (g_ GCGamepadSnapshot) SnapshotData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](g_.ID, objc.Sel("snapshotData"))
	return rv
}/* debug [instance_properties/getter]: snapshotData */


// The flattened control input values for the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGamepadSnapshot/snapshotData
func (g_ GCGamepadSnapshot) SetSnapshotData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSnapshotData:"), value)
}/* debug [instance_properties/setter]: snapshotData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccurrentextendedgamepadsnapshotdataversion
func (g_ GCGamepadSnapshot) GCCurrentExtendedGamepadSnapshotDataVersion() GCExtendedGamepadSnapshotDataVersion {
	rv := objc.Send[GCExtendedGamepadSnapshotDataVersion](g_.ID, objc.Sel("GCCurrentExtendedGamepadSnapshotDataVersion"))
	return rv
}/* debug [instance_properties/getter]: GCCurrentExtendedGamepadSnapshotDataVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccurrentmicrogamepadsnapshotdataversion
func (g_ GCGamepadSnapshot) GCCurrentMicroGamepadSnapshotDataVersion() GCMicroGamepadSnapshotDataVersion {
	rv := objc.Send[GCMicroGamepadSnapshotDataVersion](g_.ID, objc.Sel("GCCurrentMicroGamepadSnapshotDataVersion"))
	return rv
}/* debug [instance_properties/getter]: GCCurrentMicroGamepadSnapshotDataVersion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCGamepadSnapshot */


