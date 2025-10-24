// Code generated from Apple documentation for GameSave. DO NOT EDIT.

package gamesave

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GSSyncedDirectory */


/* debug [class_header]: Header for GSSyncedDirectory */
// The class instance for the [GSSyncedDirectory] class.
var (
	GSSyncedDirectoryClass     _GSSyncedDirectoryClass
	GSSyncedDirectoryClassOnce sync.Once
)

func getGSSyncedDirectoryClass() _GSSyncedDirectoryClass {
	GSSyncedDirectoryClassOnce.Do(func() {
		GSSyncedDirectoryClass = _GSSyncedDirectoryClass{objc.GetClass("GSSyncedDirectory")}
	})
	return GSSyncedDirectoryClass
}

type _GSSyncedDirectoryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GSSyncedDirectory */
// An interface definition for the [GSSyncedDirectory] class.
type IGSSyncedDirectory interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GSSyncedDirectory */
	// properties:
	DirectoryState() IGSSyncedDirectoryState
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GSSyncedDirectory */
	// methods:
	Close()
	FinishSyncingCompletionHandler(statusDisplay appkit.Window, completion unsafe.Pointer)
	FinishSyncingWithCompletionHandler(completion unsafe.Pointer)
	ResolveConflictsWithVersion(version IGSSyncedDirectoryVersion)
	TriggerPendingUploadWithCompletionHandler(completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GSSyncedDirectory */
// Alloc allocates a new instance without initialization.
func (gc _GSSyncedDirectoryClass) Alloc() GSSyncedDirectory {
	rv := objc.Send[GSSyncedDirectory](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GSSyncedDirectoryClass) New() GSSyncedDirectory {
	rv := objc.Send[GSSyncedDirectory](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GSSyncedDirectory) Init() GSSyncedDirectory {
	rv := objc.Send[GSSyncedDirectory](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GSSyncedDirectory) Autorelease() GSSyncedDirectory {
	rv := objc.Send[GSSyncedDirectory](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGSSyncedDirectory creates a new GSSyncedDirectory instance.
func NewGSSyncedDirectory() GSSyncedDirectory {
	return getGSSyncedDirectoryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GSSyncedDirectory */
// A cloud-synced directory for game-save data.
//
// To get an instance of the directory, call , which returns the directory for the iCloud container associated with the specified identifier. Calling this method starts syncing the directory in the background on the specified container. When the game needs to access the contents of the directory, show a UI while the directory fully syncs using the method. If you’re showing your own UI, call the method to wait for the directory to finish syncing. After the directory is ready to use, syncing pauses until you close the directory object or the object is deallocated. To resume syncing during the game, close and re-open the directory by calling and then .


// A cloud-synced directory for game-save data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectory
type GSSyncedDirectory struct {
	objectivec.Object
}

// GSSyncedDirectoryFrom constructs a [GSSyncedDirectory] from an unsafe.Pointer.
//
// A cloud-synced directory for game-save data.
func GSSyncedDirectoryFrom(ptr unsafe.Pointer) GSSyncedDirectory {
	return GSSyncedDirectory{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GSSyncedDirectory *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GSSyncedDirectory */

// Requests an instance of the game-save directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectory/open(forContainerIdentifier:)
func (gc _GSSyncedDirectoryClass) OpenDirectoryForContainerIdentifier(containerIdentifier objc.IObject /* cross-framework: NSString */) GSSyncedDirectory {
	rv := objc.Send[GSSyncedDirectory](objc.ID(gc.class), objc.Sel("openDirectoryForContainerIdentifier:"), containerIdentifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OpenDirectoryForContainerIdentifier) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GSSyncedDirectory */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GSSyncedDirectory */

// Closes the directory, and resumes syncing the directory to the cloud.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectory/close()
func (g_ GSSyncedDirectory) Close() {
	objc.Send[objc.ID](g_.ID, objc.Sel("close"))
}/* debug [instance_methods/method]: Close */


// Waits for the directory sync to complete, showing the sync’s progress in a modal alert.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectory/finishSyncing(_:completionHandler:)
func (g_ GSSyncedDirectory) FinishSyncingCompletionHandler(statusDisplay appkit.Window, completion unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("finishSyncing:completionHandler:"), statusDisplay, completion)
}/* debug [instance_methods/method]: FinishSyncingCompletionHandler */


// Waits for the directory sync to complete, without showing any user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectory/finishSyncing(completionHandler:)
func (g_ GSSyncedDirectory) FinishSyncingWithCompletionHandler(completion unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("finishSyncingWithCompletionHandler:"), completion)
}/* debug [instance_methods/method]: FinishSyncingWithCompletionHandler */


// Indicates that you resolved a conflict.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectory/resolveConflicts(with:)
func (g_ GSSyncedDirectory) ResolveConflictsWithVersion(version IGSSyncedDirectoryVersion) {
	objc.Send[objc.ID](g_.ID, objc.Sel("resolveConflictsWithVersion:"), version)
}/* debug [instance_methods/method]: ResolveConflictsWithVersion */


// Triggers an upload of the directory for any changes that were pending.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectory/triggerPendingUpload(completionHandler:)
func (g_ GSSyncedDirectory) TriggerPendingUploadWithCompletionHandler(completion unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("triggerPendingUploadWithCompletionHandler:"), completion)
}/* debug [instance_methods/method]: TriggerPendingUploadWithCompletionHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GSSyncedDirectory */

// The state of the directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectory/directoryState
func (g_ GSSyncedDirectory) DirectoryState() IGSSyncedDirectoryState {
	rv := objc.Send[GSSyncedDirectoryState](g_.ID, objc.Sel("directoryState"))
	return rv
}/* debug [instance_properties/getter]: directoryState */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GSSyncedDirectory */



