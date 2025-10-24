// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TableViewController] class.
var (
	TableViewControllerClass     _TableViewControllerClass
	TableViewControllerClassOnce sync.Once
)

func getTableViewControllerClass() _TableViewControllerClass {
	TableViewControllerClassOnce.Do(func() {
		TableViewControllerClass = _TableViewControllerClass{objc.GetClass("UITableViewController")}
	})
	return TableViewControllerClass
}

type _TableViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [TableViewController] class.
type ITableViewController interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other CoreAudioKit classes.


// A parent class referenced by other CoreAudioKit classes. [Full Topic]
type TableViewController struct {
	objectivec.Object
}

// TableViewControllerFrom constructs a [TableViewController] from an unsafe.Pointer.
//
// A parent class referenced by other CoreAudioKit classes.
func TableViewControllerFrom(ptr unsafe.Pointer) TableViewController {
	return TableViewController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TableViewControllerClass) Alloc() TableViewController {
	rv := objc.Send[TableViewController](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TableViewControllerClass) New() TableViewController {
	rv := objc.Send[TableViewController](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TableViewController) Init() TableViewController {
	rv := objc.Send[TableViewController](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TableViewController) Autorelease() TableViewController {
	rv := objc.Send[TableViewController](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTableViewController creates a new TableViewController instance.
func NewTableViewController() TableViewController {
	return getTableViewControllerClass().New()
}




