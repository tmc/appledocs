// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class GCMicroGamepadSnapshot */


/* debug [class_header]: Header for GCMicroGamepadSnapshot */
// The class instance for the [GCMicroGamepadSnapshot] class.
var (
	GCMicroGamepadSnapshotClass     _GCMicroGamepadSnapshotClass
	GCMicroGamepadSnapshotClassOnce sync.Once
)

func getGCMicroGamepadSnapshotClass() _GCMicroGamepadSnapshotClass {
	GCMicroGamepadSnapshotClassOnce.Do(func() {
		GCMicroGamepadSnapshotClass = _GCMicroGamepadSnapshotClass{objc.GetClass("GCMicroGamepadSnapshot")}
	})
	return GCMicroGamepadSnapshotClass
}

type _GCMicroGamepadSnapshotClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GCMicroGamepadSnapshot */
// An interface definition for the [GCMicroGamepadSnapshot] class.
type IGCMicroGamepadSnapshot interface {
	IGCMicroGamepad
	
/* debug [class_interface_properties]: Properties for GCMicroGamepadSnapshot */
	// properties:
	SnapshotData() objc.IObject /* cross-framework: NSData */
	SetSnapshotData(value objc.IObject /* cross-framework: NSData */)
	GCCurrentExtendedGamepadSnapshotDataVersion() GCExtendedGamepadSnapshotDataVersion
	GCCurrentMicroGamepadSnapshotDataVersion() GCMicroGamepadSnapshotDataVersion
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GCMicroGamepadSnapshot */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GCMicroGamepadSnapshot */
// Alloc allocates a new instance without initialization.
func (gc _GCMicroGamepadSnapshotClass) Alloc() GCMicroGamepadSnapshot {
	rv := objc.Send[GCMicroGamepadSnapshot](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GCMicroGamepadSnapshotClass) New() GCMicroGamepadSnapshot {
	rv := objc.Send[GCMicroGamepadSnapshot](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GCMicroGamepadSnapshot) Init() GCMicroGamepadSnapshot {
	rv := objc.Send[GCMicroGamepadSnapshot](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GCMicroGamepadSnapshot) Autorelease() GCMicroGamepadSnapshot {
	rv := objc.Send[GCMicroGamepadSnapshot](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCMicroGamepadSnapshot creates a new GCMicroGamepadSnapshot instance.
func NewGCMicroGamepadSnapshot() GCMicroGamepadSnapshot {
	return getGCMicroGamepadSnapshotClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GCMicroGamepadSnapshot */
// A recording of all of the values provided by a object.
//
// To create a gamepad snapshot, call the method on a object. The class is a subclass of the class, so you use the parent class’s properties to read the individual element values. The snapshot is stored in a device independent format. To get the flattened data representation of the snapshot data, read the property.


// A recording of all of the values provided by a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepadSnapshot
type GCMicroGamepadSnapshot struct {
	GCMicroGamepad
}

// GCMicroGamepadSnapshotFrom constructs a [GCMicroGamepadSnapshot] from an unsafe.Pointer.
//
// A recording of all of the values provided by a object.
func GCMicroGamepadSnapshotFrom(ptr unsafe.Pointer) GCMicroGamepadSnapshot {
	return GCMicroGamepadSnapshot{
		GCMicroGamepad: GCMicroGamepadFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GCMicroGamepadSnapshot */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepadSnapshot/init(controller:snapshotData:)
func NewGCMicroGamepadSnapshotWithControllerSnapshotData(controller IGCController, data objc.IObject /* cross-framework: NSData */) GCMicroGamepadSnapshot {
	instance := getGCMicroGamepadSnapshotClass().Alloc()
	rv := objc.Send[GCMicroGamepadSnapshot](instance.ID, objc.Sel("initWithController:snapshotData:"), controller, data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGCMicroGamepadSnapshotWithControllerSnapshotData */


// Initializes a snapshot object with the flattened data representation obtained from another snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepadSnapshot/init(snapshotData:)
func NewGCMicroGamepadSnapshotWithSnapshotData(data objc.IObject /* cross-framework: NSData */) GCMicroGamepadSnapshot {
	instance := getGCMicroGamepadSnapshotClass().Alloc()
	rv := objc.Send[GCMicroGamepadSnapshot](instance.ID, objc.Sel("initWithSnapshotData:"), data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGCMicroGamepadSnapshotWithSnapshotData */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GCMicroGamepadSnapshot */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GCMicroGamepadSnapshot */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GCMicroGamepadSnapshot */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GCMicroGamepadSnapshot */

// The flattened control input values for the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepadSnapshot/snapshotData
func (g_ GCMicroGamepadSnapshot) SnapshotData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](g_.ID, objc.Sel("snapshotData"))
	return rv
}/* debug [instance_properties/getter]: snapshotData */


// The flattened control input values for the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepadSnapshot/snapshotData
func (g_ GCMicroGamepadSnapshot) SetSnapshotData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setSnapshotData:"), value)
}/* debug [instance_properties/setter]: snapshotData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccurrentextendedgamepadsnapshotdataversion
func (g_ GCMicroGamepadSnapshot) GCCurrentExtendedGamepadSnapshotDataVersion() GCExtendedGamepadSnapshotDataVersion {
	rv := objc.Send[GCExtendedGamepadSnapshotDataVersion](g_.ID, objc.Sel("GCCurrentExtendedGamepadSnapshotDataVersion"))
	return rv
}/* debug [instance_properties/getter]: GCCurrentExtendedGamepadSnapshotDataVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamecontroller/gccurrentmicrogamepadsnapshotdataversion
func (g_ GCMicroGamepadSnapshot) GCCurrentMicroGamepadSnapshotDataVersion() GCMicroGamepadSnapshotDataVersion {
	rv := objc.Send[GCMicroGamepadSnapshotDataVersion](g_.ID, objc.Sel("GCCurrentMicroGamepadSnapshotDataVersion"))
	return rv
}/* debug [instance_properties/getter]: GCCurrentMicroGamepadSnapshotDataVersion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GCMicroGamepadSnapshot */


