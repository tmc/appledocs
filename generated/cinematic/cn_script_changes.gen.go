// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CNScriptChanges] class.
type ICNScriptChanges interface {
	objectivec.IObject
	// properties:
	AddedDetectionTracks() []CNDetectionTrack /* primitive/slice/pointer. */
	DataRepresentation() foundation.objc.IObject /* cross-framework: NSData */
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (cc _CNScriptChangesClass) Alloc() CNScriptChanges {
	rv := objc.Send[CNScriptChanges](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// All detection tracks added since recording the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScriptChanges/addedDetectionTracks
func (c_ CNScriptChanges) AddedDetectionTracks() []CNDetectionTrack /* primitive/slice/pointer. */ {
	rv := objc.Send[[]CNDetectionTrack](c_.ID, objc.Sel("addedDetectionTracks"))
	return rv
}


// Persistent data representation of changes for later restoration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScriptChanges/dataRepresentation
func (c_ CNScriptChanges) DataRepresentation() foundation.objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("dataRepresentation"))
	return rv
}



