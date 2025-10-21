// Code generated from Apple documentation for PreferencePanes. DO NOT EDIT.

package preferencepanes

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PreferencePane] class.
var (
	PreferencePaneClass     _PreferencePaneClass
	PreferencePaneClassOnce sync.Once
)

func getPreferencePaneClass() _PreferencePaneClass {
	PreferencePaneClassOnce.Do(func() {
		PreferencePaneClass = _PreferencePaneClass{objc.GetClass("NSPreferencePane")}
	})
	return PreferencePaneClass
}

type _PreferencePaneClass struct {
	class objc.Class
}

// An interface definition for the [PreferencePane] class.
type IPreferencePane interface {
	objectivec.IObject
	AssignMainView()
	DidSelect()
	DidUnselect()
	LoadMainView() unsafe.Pointer
	MainViewDidLoad()
	ReplyToShouldUnselect(shouldUnselect bool)
	UpdateHelpMenuWithArray(inArrayOfMenuItems unsafe.Pointer)
	WillSelect()
	WillUnselect()
}

// The interface for providing preference panes to System Preferences or other apps.
//
// Preference panes are subclasses of , packaged up in bundles and loaded by a preference application, such as System Preferences. These bundles have a suffix of . Bundles intended for use by System Preferences are located in the directories of the various file system domains. See the chapter in for information about domains. The preference pane bundle normally contains a nib file with the user interface for modifying user preferences. The nib file contains a window assigned to the _window outlet of the preference pane instance (the nib’s File’s Owner). The implementation of , invoked by the preference application, loads the nib file and uses the content view of _window as the preference pane’s main view. Override this method if you need a different technique for creating the user interface. The subclass is responsible for initializing the user interface with the current preference settings and recording any modifications the user makes. Through a series of , , and methods, the preference application notifies the preference pane when the pane is selected (displayed) and deselected, allowing the pane to perform the necessary actions at the appropriate times. Implement these methods (and any additional target-action methods connected to the interface) as needed to produce the desired behavior for your preference pane. Preference panes support Help menu items. Specify global help menu items in your bundle’s file under the key. To add dynamic help items, implement the method.
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane
type PreferencePane struct {
	objectivec.Object
}

// PreferencePaneFrom constructs a [PreferencePane] from an unsafe.Pointer.
//
// The interface for providing preference panes to System Preferences or other apps.
func PreferencePaneFrom(ptr unsafe.Pointer) PreferencePane {
	return PreferencePane{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PreferencePaneClass) Alloc() PreferencePane {
	rv := objc.Send[PreferencePane](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PreferencePaneClass) New() PreferencePane {
	rv := objc.Send[PreferencePane](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreferencePane) Init() PreferencePane {
	rv := objc.Send[PreferencePane](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreferencePane) Autorelease() PreferencePane {
	rv := objc.Send[PreferencePane](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreferencePane creates a new PreferencePane instance.
func NewPreferencePane() PreferencePane {
	return getPreferencePaneClass().New()
}


// Initializes a preference pane with the specified bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/init(bundle:)
func NewPreferencePaneWithBundle(bundle unsafe.Pointer) PreferencePane {
	instance := getPreferencePaneClass().Alloc()
	rv := objc.Send[PreferencePane](instance.ID, objc.Sel("initWithBundle:"), bundle)
	rv.Autorelease()
	return rv
}


// Locates and assigns the preference pane’s main view from the nib file loaded by .
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/assignMainView()
func (p_ PreferencePane) AssignMainView() {
	objc.Send[objc.ID](p_.ID, objc.Sel("assignMainView"))
}

// Notifies the preference pane that the main app has just displayed the preference pane’s main view.
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/didSelect()
func (p_ PreferencePane) DidSelect() {
	objc.Send[objc.ID](p_.ID, objc.Sel("didSelect"))
}

// Notifies the preference pane that the main app has just stopped displaying the preference pane’s main view.
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/didUnselect()
func (p_ PreferencePane) DidUnselect() {
	objc.Send[objc.ID](p_.ID, objc.Sel("didUnselect"))
}

// Loads the preference pane’s user interface into its main view.
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/loadMainView()
func (p_ PreferencePane) LoadMainView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("loadMainView"))
	return rv
}

// Notifies the preference pane that the main view is set up and prepared to be displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/mainViewDidLoad()
func (p_ PreferencePane) MainViewDidLoad() {
	objc.Send[objc.ID](p_.ID, objc.Sel("mainViewDidLoad"))
}

// Notifies the main application of the preference pane’s ability to be deselected.
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/reply(toShouldUnselect:)
func (p_ PreferencePane) ReplyToShouldUnselect(shouldUnselect bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("replyToShouldUnselect:"), shouldUnselect)
}

// Updates the help menu.
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/updateHelpMenu(with:)
func (p_ PreferencePane) UpdateHelpMenuWithArray(inArrayOfMenuItems unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("updateHelpMenuWithArray:"), inArrayOfMenuItems)
}

// Notifies the preference pane that the main app is about to display the preference pane’s main view.
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/willSelect()
func (p_ PreferencePane) WillSelect() {
	objc.Send[objc.ID](p_.ID, objc.Sel("willSelect"))
}

// Notifies the preference pane that the main app is about to stop displaying the preference pane’s main view.
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/willUnselect()
func (p_ PreferencePane) WillUnselect() {
	objc.Send[objc.ID](p_.ID, objc.Sel("willUnselect"))
}

// A Boolean value that indicates whether text fields save their values before changing preference panes.
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/autoSaveTextFields
func (p_ PreferencePane) AutoSaveTextFields() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("autoSaveTextFields"))
	return rv
}

// The preference pane’s bundle.
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/bundle
func (p_ PreferencePane) Bundle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("bundle"))
	return rv
}

// The first view in the keyboard focus chain.
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/firstKeyView
func (p_ PreferencePane) FirstKeyView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("firstKeyView"))
	return rv
}


// SetFirstKeyView sets the value of the firstKeyView property.
// The first view in the keyboard focus chain.

//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/firstKeyView
func (p_ PreferencePane) SetFirstKeyView(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFirstKeyView:"), value)
}
// The view that should have keyboard focus when the pane is selected.
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/initialKeyView
func (p_ PreferencePane) InitialKeyView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("initialKeyView"))
	return rv
}


// SetInitialKeyView sets the value of the initialKeyView property.
// The view that should have keyboard focus when the pane is selected.

//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/initialKeyView
func (p_ PreferencePane) SetInitialKeyView(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInitialKeyView:"), value)
}
// A Boolean value that indicates whether the preference pane is currently selected.
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/isSelected
func (p_ PreferencePane) Selected() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("selected"))
	return rv
}

// The last view in the keyboard focus chain.
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/lastKeyView
func (p_ PreferencePane) LastKeyView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("lastKeyView"))
	return rv
}


// SetLastKeyView sets the value of the lastKeyView property.
// The last view in the keyboard focus chain.

//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/lastKeyView
func (p_ PreferencePane) SetLastKeyView(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLastKeyView:"), value)
}
// The name of the preference pane’s nib file.
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/mainNibName
func (p_ PreferencePane) MainNibName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("mainNibName"))
	return rv
}

// The main view of the preference pane.
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/mainView
func (p_ PreferencePane) MainView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("mainView"))
	return rv
}


// SetMainView sets the value of the mainView property.
// The main view of the preference pane.

//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/mainView
func (p_ PreferencePane) SetMainView(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMainView:"), value)
}
// A Boolean value that indicates whether the preference pane is able to be deselected.
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePane/shouldUnselect
func (p_ PreferencePane) ShouldUnselect() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("shouldUnselect"))
	return rv
}


