// Code generated from Apple documentation for CoreAudioKit. DO NOT EDIT.

package coreaudiokit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class UITableViewController */


/* debug [class_header]: Header for UITableViewController */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TableViewController */
// An interface definition for the [TableViewController] class.
type ITableViewController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TableViewController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TableViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TableViewController */
// Alloc allocates a new instance without initialization.
func (tc _TableViewControllerClass) Alloc() TableViewController {
	rv := objc.Send[TableViewController](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TableViewController */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TableViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TableViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TableViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TableViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TableViewController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class UITableViewController */



