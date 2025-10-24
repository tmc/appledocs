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

// An interface definition for the [GSSyncedDirectoryState] class.
type IGSSyncedDirectoryState interface {
	objectivec.IObject
	// properties:
	ConflictedVersions() []GSSyncedDirectoryVersion /* not a class type */
	Error() objc.IObject /* cross-framework: Error */
	State() GSSyncState
	Url() objc.IObject /* cross-framework: NSURL */
	DirectoryState() IGSSyncedDirectoryState
	SetDirectoryState(value IGSSyncedDirectoryState)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (gc _GSSyncedDirectoryStateClass) Alloc() GSSyncedDirectoryState {
	rv := objc.Send[GSSyncedDirectoryState](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The conflicting versions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectoryState/conflictedVersions
func (g_ GSSyncedDirectoryState) ConflictedVersions() []GSSyncedDirectoryVersion /* not a class type */ {
	rv := objc.Send[[]GSSyncedDirectoryVersion](g_.ID, objc.Sel("conflictedVersions"))
	return rv
}


// The error preventing you from using the directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectoryState/error
func (g_ GSSyncedDirectoryState) Error() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](g_.ID, objc.Sel("error"))
	return rv
}


// Specifies the current state of the directory
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectoryState/state
func (g_ GSSyncedDirectoryState) State() GSSyncState {
	rv := objc.Send[GSSyncState](g_.ID, objc.Sel("state"))
	return rv
}


// The URL of a directory to read and write game-save data in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameSave/GSSyncedDirectoryState/url
func (g_ GSSyncedDirectoryState) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](g_.ID, objc.Sel("url"))
	return rv
}


// The state of the directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamesave/gssynceddirectory/directorystate
func (g_ GSSyncedDirectoryState) DirectoryState() IGSSyncedDirectoryState {
	rv := objc.Send[GSSyncedDirectoryState](g_.ID, objc.Sel("directoryState"))
	return rv
}


// The state of the directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamesave/gssynceddirectory/directorystate
func (g_ GSSyncedDirectoryState) SetDirectoryState(value IGSSyncedDirectoryState) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDirectoryState:"), value)
}




