// Code generated from Apple documentation for GameSave. DO NOT EDIT.

package gamesave

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [GSSyncedDirectory] class.
type IGSSyncedDirectory interface {
	objectivec.IObject
	FinishSyncingCompletionHandler(statusDisplay appkit.IWindow, completion unsafe.Pointer)
	FinishSyncingWithCompletionHandler(completion unsafe.Pointer)
	DirectoryState() unsafe.Pointer
}

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

// Alloc allocates a new instance without initialization.
func (gc _GSSyncedDirectoryClass) Alloc() GSSyncedDirectory {
	rv := objc.Send[GSSyncedDirectory](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Requests an instance of the game-save directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectory/open(forContainerIdentifier:)
func (gc _GSSyncedDirectoryClass) OpenDirectoryForContainerIdentifier(containerIdentifier string) GSSyncedDirectory {
	rv := objc.Send[GSSyncedDirectory](objc.ID(gc.class), objc.Sel("openDirectoryForContainerIdentifier:"), objc.String(containerIdentifier))
	return rv
}


// Waits for the directory sync to complete, showing the sync’s progress in a modal alert.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectory/finishSyncing(_:completionHandler:)
func (g_ GSSyncedDirectory) FinishSyncingCompletionHandler(statusDisplay appkit.IWindow, completion unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("finishSyncing:completionHandler:"), statusDisplay, completion)
}


// Waits for the directory sync to complete, without showing any user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectory/finishSyncing(completionHandler:)
func (g_ GSSyncedDirectory) FinishSyncingWithCompletionHandler(completion unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("finishSyncingWithCompletionHandler:"), completion)
}


// The state of the directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectory/directoryState
func (g_ GSSyncedDirectory) DirectoryState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("directoryState"))
	return rv
}




