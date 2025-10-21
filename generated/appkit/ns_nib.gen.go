// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Nib] class.
var (
	NibClass     _NibClass
	NibClassOnce sync.Once
)

func getNibClass() _NibClass {
	NibClassOnce.Do(func() {
		NibClass = _NibClass{objc.GetClass("NSNib")}
	})
	return NibClass
}

type _NibClass struct {
	class objc.Class
}

// An interface definition for the [Nib] class.
type INib interface {
	objectivec.IObject
	InstantiateWithOwnerTopLevelObjects(owner objectivec.IObject, topLevelObjects objectivec.IObject) bool
	InstantiateNibWithExternalNameTable(externalNameTable objectivec.IObject) bool
	InstantiateNibWithOwnerTopLevelObjects(owner objectivec.IObject, topLevelObjects objectivec.IObject) bool
}

// An object wrapper, or container, for an Interface Builder nib file.
//
// An object keeps the contents of a nib file resident in memory, ready for unarchiving and instantiation. When you create a nib object using the contents of a nib file, the object loads the contents of the referenced nib bundle—the object graph as well as any images and sounds—into memory; but it does not yet unarchive it. To unarchive all of the nib data and thus truly instantiate the nib you must call one of the methods of . During the instantiation process, each object in the archive is unarchived and then initialized using the method befitting its type. View classes are initialized using their method. Custom objects are initialized using their method. In the case of Cocoa views (and custom views that have options on an associated Interface Builder palette) the initialization process also reads in any values set by the user in Interface Builder. Once all objects have been instantiated and initialized from the archive, the nib loading code attempts to reestablish the connections between each object’s outlets and the corresponding target objects. If your custom objects have outlets, the object attempts to reestablish any connections you created in Interface Builder. It starts by trying to establish the connections using your object’s own methods first. For each outlet that needs a connection, the object looks for a method of the form in your object. If that method exists, the nib object calls it, passing the target object as a parameter. If you did not define a setter method with that exact name, the object searches the object for an instance variable (of type ) with the corresponding outlet name and tries to set its value directly. If an instance variable with the correct name cannot be found, initialization of that connection does not occur. After all objects have been initialized and their connections reestablished, each object receives an message. You can override this method in your custom objects to perform any additional initialization.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNib
type Nib struct {
	objectivec.Object
}

// NibFrom constructs a [Nib] from an unsafe.Pointer.
//
// An object wrapper, or container, for an Interface Builder nib file.
func NibFrom(ptr unsafe.Pointer) Nib {
	return Nib{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NibClass) Alloc() Nib {
	rv := objc.Send[Nib](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NibClass) New() Nib {
	rv := objc.Send[Nib](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ Nib) Init() Nib {
	rv := objc.Send[Nib](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ Nib) Autorelease() Nib {
	rv := objc.Send[Nib](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNib creates a new Nib instance.
func NewNib() Nib {
	return getNibClass().New()
}




// Returns an object initialized to the nib file at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNib/initWithContentsOfURL:
func NewNibWithContentsOfURL(nibFileURL foundation.IURL) Nib {
	instance := getNibClass().Alloc()
	rv := objc.Send[Nib](instance.ID, objc.Sel("initWithContentsOfURL:"), nibFileURL)
	rv.Autorelease()
	return rv
}



// Initializes an instance with nib data and specified bundle for locating resources.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNib/init(nibData:bundle:)
func NewNibWithNibDataBundle(nibData foundation.IData, bundle foundation.IBundle) Nib {
	instance := getNibClass().Alloc()
	rv := objc.Send[Nib](instance.ID, objc.Sel("initWithNibData:bundle:"), nibData, bundle)
	rv.Autorelease()
	return rv
}



// Returns an object initialized to the nib file in the specified bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNib/init(nibNamed:bundle:)
func NewNibWithNibNamedBundle(nibName INibName, bundle foundation.IBundle) Nib {
	instance := getNibClass().Alloc()
	rv := objc.Send[Nib](instance.ID, objc.Sel("initWithNibNamed:bundle:"), nibName, bundle)
	rv.Autorelease()
	return rv
}


// Instantiates objects in the nib file with the specified owner.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNib/instantiate(withOwner:topLevelObjects:)
func (n_ Nib) InstantiateWithOwnerTopLevelObjects(owner objectivec.IObject, topLevelObjects objectivec.IObject) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("instantiateWithOwner:topLevelObjects:"), owner, topLevelObjects)
	return rv
}

// Unarchives and instantiates the in-memory contents of the receiver’s nib file, creating a distinct object tree and top level objects.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNib/instantiateNibWithExternalNameTable:
func (n_ Nib) InstantiateNibWithExternalNameTable(externalNameTable objectivec.IObject) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("instantiateNibWithExternalNameTable:"), externalNameTable)
	return rv
}

// Unarchives and instantiates the in-memory contents of the receiver’s nib file, creating a distinct object tree and set of top level objects.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNib/instantiateNibWithOwner:topLevelObjects:
func (n_ Nib) InstantiateNibWithOwnerTopLevelObjects(owner objectivec.IObject, topLevelObjects objectivec.IObject) bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("instantiateNibWithOwner:topLevelObjects:"), owner, topLevelObjects)
	return rv
}


