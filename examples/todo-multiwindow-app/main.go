// Multi-Window Todo App using only generated bindings (no darwinkit)
//
// This advanced example demonstrates:
// - Multiple windows (main window + preferences window)
// - Complex UI layout with multiple views
// - Text input and dynamic content
// - Window management and lifecycle
// - More sophisticated event handling
//
// It creates a simple todo list app with:
// - Main window: Add todos, view list, manage items
// - Preferences window: Configure app settings
package main

import (
	"fmt"
	"runtime"

	"github.com/ebitengine/purego"
	"github.com/ebitengine/purego/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

func init() {
	runtime.LockOSThread()

	// Load AppKit framework
	_, err := purego.Dlopen("/System/Library/Frameworks/AppKit.framework/AppKit", purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

// Foundation/AppKit types
type NSPoint struct {
	X, Y float64
}

type NSSize struct {
	Width, Height float64
}

type NSRect struct {
	Origin NSPoint
	Size   NSSize
}

// UI State
var (
	app              objc.ID
	mainWindow       appkit.Window
	prefsWindow      appkit.Window
	todoInput        appkit.TextField
	todoListView     appkit.View
	statusLabel      appkit.TextField
	maxItemsField    appkit.TextField
	todos            []string
	todoItemViews    []appkit.View
	maxTodoItems     = 10
	isPrefsOpen      = false
)

// Helper: Create NSString from Go string
func NewNSString(str string) objc.ID {
	return objc.ID(objc.GetClass("NSString")).Send(
		objc.RegisterName("stringWithUTF8String:"),
		str,
	)
}

// Helper: Get string value from NSString
func GetStringValue(nsstring objc.ID) string {
	if nsstring == 0 {
		return ""
	}
	cstr := nsstring.Send(objc.RegisterName("UTF8String"))
	if cstr == 0 {
		return ""
	}
	return objc.CString(uintptr(cstr))
}

// Helper: Create Button
func NewButton(frame NSRect) appkit.Button {
	button := appkit.ButtonClass.Alloc()
	objc.Send[objc.ID](button.ID(), objc.RegisterName("initWithFrame:"), frame)
	return button
}

// Helper: Create TextField
func NewTextField(frame NSRect) appkit.TextField {
	textField := appkit.TextFieldClass.Alloc()
	objc.Send[objc.ID](textField.ID(), objc.RegisterName("initWithFrame:"), frame)
	return textField
}

// Helper: Create View
func NewView(frame NSRect) appkit.View {
	view := appkit.ViewClass.Alloc()
	objc.Send[objc.ID](view.ID(), objc.RegisterName("initWithFrame:"), frame)
	return view
}

// Helper: Create NSColor (RGB values 0-1)
func NewNSColor(r, g, b, a float64) objc.ID {
	colorClass := objc.GetClass("NSColor")
	return objc.ID(colorClass).Send(
		objc.RegisterName("colorWithCalibratedRed:green:blue:alpha:"),
		r, g, b, a,
	)
}

// updateStatusLabel updates the status label with todo count
func updateStatusLabel() {
	statusLabel.SetStringValue(NewNSString(fmt.Sprintf("Todos: %d / %d", len(todos), maxTodoItems)))
}

// refreshTodoList rebuilds the todo list UI
func refreshTodoList() {
	// Remove all existing todo item views
	for _, itemView := range todoItemViews {
		objc.Send[objc.ID](itemView.ID(), objc.RegisterName("removeFromSuperview"))
	}
	todoItemViews = nil

	// Create new todo item views
	yPos := 10.0
	for i, todo := range todos {
		itemView := createTodoItemView(i, todo, yPos)
		objc.Send[objc.ID](todoListView.ID(), objc.RegisterName("addSubview:"), itemView.ID())
		todoItemViews = append(todoItemViews, itemView)
		yPos += 35
	}

	// Update frame of scrollable area if needed
	newHeight := float64(len(todos)*35 + 10)
	if newHeight < 200 {
		newHeight = 200
	}
	frame := NSRect{
		Origin: NSPoint{X: 0, Y: 0},
		Size:   NSSize{Width: 460, Height: newHeight},
	}
	objc.Send[objc.ID](todoListView.ID(), objc.RegisterName("setFrame:"), frame)

	updateStatusLabel()
}

// createTodoItemView creates a view for a single todo item
func createTodoItemView(index int, text string, yPos float64) appkit.View {
	itemView := NewView(NSRect{
		Origin: NSPoint{X: 5, Y: yPos},
		Size:   NSSize{Width: 450, Height: 30},
	})

	// Item number label
	numberLabel := NewTextField(NSRect{
		Origin: NSPoint{X: 0, Y: 5},
		Size:   NSSize{Width: 30, Height: 20},
	})
	numberLabel.SetStringValue(NewNSString(fmt.Sprintf("%d.", index+1)))
	numberLabel.SetEditable(false)
	numberLabel.SetBordered(false)
	numberLabel.SetBackgroundColor(objc.ID(0))
	objc.Send[objc.ID](itemView.ID(), objc.RegisterName("addSubview:"), numberLabel.ID())

	// Todo text label
	textLabel := NewTextField(NSRect{
		Origin: NSPoint{X: 35, Y: 5},
		Size:   NSSize{Width: 330, Height: 20},
	})
	textLabel.SetStringValue(NewNSString(text))
	textLabel.SetEditable(false)
	textLabel.SetBordered(false)
	textLabel.SetBackgroundColor(objc.ID(0))
	objc.Send[objc.ID](itemView.ID(), objc.RegisterName("addSubview:"), textLabel.ID())

	// Delete button
	deleteBtn := NewButton(NSRect{
		Origin: NSPoint{X: 370, Y: 2},
		Size:   NSSize{Width: 70, Height: 24},
	})
	deleteBtn.SetTitle(NewNSString("Remove"))
	deleteBtn.SetBezelStyle(1) // NSRoundedBezelStyle

	// Create delete handler for this specific index
	handler := createDeleteHandler(index)
	deleteBtn.SetTarget(handler)
	deleteBtn.SetAction(objc.RegisterName("deleteClicked:"))
	objc.Send[objc.ID](itemView.ID(), objc.RegisterName("addSubview:"), deleteBtn.ID())

	return itemView
}

// Event Handlers

func createAppDelegate() objc.ID {
	className := "AppDelegate"
	class := objc.GetClass(className)
	if class == 0 {
		superClass := objc.GetClass("NSObject")
		windowShouldClose := func(self objc.ID, _cmd objc.SEL, sender objc.ID) bool {
			objc.Send[objc.ID](app, objc.RegisterName("terminate:"), objc.ID(0))
			return true
		}
		class, _ = objc.RegisterClass(
			className,
			superClass,
			[]*objc.Protocol{},
			[]objc.FieldDef{},
			[]objc.MethodDef{
				{
					Cmd: objc.RegisterName("windowShouldClose:"),
					Fn:  windowShouldClose,
				},
			},
		)
	}
	delegate := objc.Send[objc.ID](objc.ID(class), objc.RegisterName("alloc"))
	delegate = objc.Send[objc.ID](delegate, objc.RegisterName("init"))
	return delegate
}

func createAddTodoHandler() objc.ID {
	className := "AddTodoHandler"
	class := objc.GetClass(className)
	if class == 0 {
		superClass := objc.GetClass("NSObject")
		addTodo := func(self objc.ID, _cmd objc.SEL, sender objc.ID) {
			// Get text from input field
			stringValue := objc.Send[objc.ID](todoInput.ID(), objc.RegisterName("stringValue"))
			text := GetStringValue(stringValue)

			if text == "" {
				fmt.Println("Cannot add empty todo")
				return
			}

			if len(todos) >= maxTodoItems {
				fmt.Printf("Maximum todos reached (%d)\n", maxTodoItems)
				return
			}

			todos = append(todos, text)
			fmt.Printf("Added todo: %s (total: %d)\n", text, len(todos))

			// Clear input field
			todoInput.SetStringValue(NewNSString(""))

			// Refresh the list
			refreshTodoList()
		}
		class, _ = objc.RegisterClass(
			className,
			superClass,
			[]*objc.Protocol{},
			[]objc.FieldDef{},
			[]objc.MethodDef{
				{
					Cmd: objc.RegisterName("addTodo:"),
					Fn:  addTodo,
				},
			},
		)
	}
	handler := objc.Send[objc.ID](objc.ID(class), objc.RegisterName("alloc"))
	handler = objc.Send[objc.ID](handler, objc.RegisterName("init"))
	return handler
}

func createClearAllHandler() objc.ID {
	className := "ClearAllHandler"
	class := objc.GetClass(className)
	if class == 0 {
		superClass := objc.GetClass("NSObject")
		clearAll := func(self objc.ID, _cmd objc.SEL, sender objc.ID) {
			fmt.Println("Clearing all todos")
			todos = nil
			refreshTodoList()
		}
		class, _ = objc.RegisterClass(
			className,
			superClass,
			[]*objc.Protocol{},
			[]objc.FieldDef{},
			[]objc.MethodDef{
				{
					Cmd: objc.RegisterName("clearAll:"),
					Fn:  clearAll,
				},
			},
		)
	}
	handler := objc.Send[objc.ID](objc.ID(class), objc.RegisterName("alloc"))
	handler = objc.Send[objc.ID](handler, objc.RegisterName("init"))
	return handler
}

// Dynamic delete handler creation
var deleteHandlerCounter = 0

func createDeleteHandler(index int) objc.ID {
	// Create a unique class name for each delete handler
	deleteHandlerCounter++
	className := fmt.Sprintf("DeleteHandler%d", deleteHandlerCounter)

	superClass := objc.GetClass("NSObject")
	deleteClicked := func(self objc.ID, _cmd objc.SEL, sender objc.ID) {
		if index >= 0 && index < len(todos) {
			fmt.Printf("Deleting todo at index %d: %s\n", index, todos[index])
			todos = append(todos[:index], todos[index+1:]...)
			refreshTodoList()
		}
	}

	class, _ := objc.RegisterClass(
		className,
		superClass,
		[]*objc.Protocol{},
		[]objc.FieldDef{},
		[]objc.MethodDef{
			{
				Cmd: objc.RegisterName("deleteClicked:"),
				Fn:  deleteClicked,
			},
		},
	)

	handler := objc.Send[objc.ID](objc.ID(class), objc.RegisterName("alloc"))
	handler = objc.Send[objc.ID](handler, objc.RegisterName("init"))
	return handler
}

func createPreferencesHandler() objc.ID {
	className := "PreferencesHandler"
	class := objc.GetClass(className)
	if class == 0 {
		superClass := objc.GetClass("NSObject")
		openPrefs := func(self objc.ID, _cmd objc.SEL, sender objc.ID) {
			if !isPrefsOpen {
				fmt.Println("Opening preferences window")
				createPreferencesWindow()
			} else {
				// Bring to front if already open
				prefsWindow.MakeKeyAndOrderFront(objc.ID(0))
			}
		}
		class, _ = objc.RegisterClass(
			className,
			superClass,
			[]*objc.Protocol{},
			[]objc.FieldDef{},
			[]objc.MethodDef{
				{
					Cmd: objc.RegisterName("openPrefs:"),
					Fn:  openPrefs,
				},
			},
		)
	}
	handler := objc.Send[objc.ID](objc.ID(class), objc.RegisterName("alloc"))
	handler = objc.Send[objc.ID](handler, objc.RegisterName("init"))
	return handler
}

func createPrefsDelegate() objc.ID {
	className := "PrefsDelegate"
	class := objc.GetClass(className)
	if class == 0 {
		superClass := objc.GetClass("NSObject")
		windowShouldClose := func(self objc.ID, _cmd objc.SEL, sender objc.ID) bool {
			fmt.Println("Closing preferences window")
			isPrefsOpen = false
			return true
		}
		class, _ = objc.RegisterClass(
			className,
			superClass,
			[]*objc.Protocol{},
			[]objc.FieldDef{},
			[]objc.MethodDef{
				{
					Cmd: objc.RegisterName("windowShouldClose:"),
					Fn:  windowShouldClose,
				},
			},
		)
	}
	delegate := objc.Send[objc.ID](objc.ID(class), objc.RegisterName("alloc"))
	delegate = objc.Send[objc.ID](delegate, objc.RegisterName("init"))
	return delegate
}

func createSavePrefsHandler() objc.ID {
	className := "SavePrefsHandler"
	class := objc.GetClass(className)
	if class == 0 {
		superClass := objc.GetClass("NSObject")
		savePrefs := func(self objc.ID, _cmd objc.SEL, sender objc.ID) {
			// Get max items value
			stringValue := objc.Send[objc.ID](maxItemsField.ID(), objc.RegisterName("stringValue"))
			text := GetStringValue(stringValue)

			var newMax int
			fmt.Sscanf(text, "%d", &newMax)

			if newMax > 0 && newMax <= 50 {
				maxTodoItems = newMax
				fmt.Printf("Updated max todo items to: %d\n", maxTodoItems)
				updateStatusLabel()
			} else {
				fmt.Println("Invalid max items value (must be 1-50)")
			}
		}
		class, _ = objc.RegisterClass(
			className,
			superClass,
			[]*objc.Protocol{},
			[]objc.FieldDef{},
			[]objc.MethodDef{
				{
					Cmd: objc.RegisterName("savePrefs:"),
					Fn:  savePrefs,
				},
			},
		)
	}
	handler := objc.Send[objc.ID](objc.ID(class), objc.RegisterName("alloc"))
	handler = objc.Send[objc.ID](handler, objc.RegisterName("init"))
	return handler
}

// createMainWindow creates the main todo list window
func createMainWindow() {
	mainWindow = appkit.WindowClass.Alloc()

	frame := NSRect{
		Origin: NSPoint{X: 100, Y: 100},
		Size:   NSSize{Width: 500, Height: 500},
	}

	objc.Send[objc.ID](
		mainWindow.ID(),
		objc.RegisterName("initWithContentRect:styleMask:backing:defer:"),
		frame,
		1|2|8, // Titled | Closable | Resizable
		2,     // NSBackingStoreBuffered
		false,
	)

	mainWindow.SetTitle(NewNSString("Todo List - Generated Bindings"))

	delegate := createAppDelegate()
	mainWindow.SetDelegate(delegate)

	contentView := objc.Send[objc.ID](mainWindow.ID(), objc.RegisterName("contentView"))

	// Title label
	titleLabel := NewTextField(NSRect{
		Origin: NSPoint{X: 20, Y: 450},
		Size:   NSSize{Width: 460, Height: 30},
	})
	titleLabel.SetStringValue(NewNSString("Multi-Window Todo App"))
	titleLabel.SetEditable(false)
	titleLabel.SetBordered(false)
	titleLabel.SetBackgroundColor(objc.ID(0))
	titleLabel.SetAlignment(2) // Center
	// Make it bold and larger
	font := objc.Send[objc.ID](objc.ID(objc.GetClass("NSFont")), objc.RegisterName("boldSystemFontOfSize:"), 18.0)
	objc.Send[objc.ID](titleLabel.ID(), objc.RegisterName("setFont:"), font)
	objc.Send[objc.ID](contentView, objc.RegisterName("addSubview:"), titleLabel.ID())

	// Input section
	inputLabel := NewTextField(NSRect{
		Origin: NSPoint{X: 20, Y: 410},
		Size:   NSSize{Width: 100, Height: 20},
	})
	inputLabel.SetStringValue(NewNSString("New Todo:"))
	inputLabel.SetEditable(false)
	inputLabel.SetBordered(false)
	inputLabel.SetBackgroundColor(objc.ID(0))
	objc.Send[objc.ID](contentView, objc.RegisterName("addSubview:"), inputLabel.ID())

	todoInput = NewTextField(NSRect{
		Origin: NSPoint{X: 120, Y: 410},
		Size:   NSSize{Width: 250, Height: 22},
	})
	todoInput.SetPlaceholderString(NewNSString("Enter todo item..."))
	objc.Send[objc.ID](contentView, objc.RegisterName("addSubview:"), todoInput.ID())

	// Add button
	addBtn := NewButton(NSRect{
		Origin: NSPoint{X: 380, Y: 408},
		Size:   NSSize{Width: 100, Height: 26},
	})
	addBtn.SetTitle(NewNSString("Add"))
	addBtn.SetBezelStyle(1)
	addHandler := createAddTodoHandler()
	addBtn.SetTarget(addHandler)
	addBtn.SetAction(objc.RegisterName("addTodo:"))
	objc.Send[objc.ID](contentView, objc.RegisterName("addSubview:"), addBtn.ID())

	// Status label
	statusLabel = NewTextField(NSRect{
		Origin: NSPoint{X: 20, Y: 380},
		Size:   NSSize{Width: 460, Height: 20},
	})
	statusLabel.SetStringValue(NewNSString("Todos: 0 / 10"))
	statusLabel.SetEditable(false)
	statusLabel.SetBordered(false)
	statusLabel.SetBackgroundColor(objc.ID(0))
	objc.Send[objc.ID](contentView, objc.RegisterName("addSubview:"), statusLabel.ID())

	// Scroll view for todo list
	scrollView := objc.Send[objc.ID](objc.ID(objc.GetClass("NSScrollView")), objc.RegisterName("alloc"))
	scrollFrame := NSRect{
		Origin: NSPoint{X: 20, Y: 60},
		Size:   NSSize{Width: 460, Height: 310},
	}
	objc.Send[objc.ID](scrollView, objc.RegisterName("initWithFrame:"), scrollFrame)
	objc.Send[objc.ID](scrollView, objc.RegisterName("setHasVerticalScroller:"), true)
	objc.Send[objc.ID](scrollView, objc.RegisterName("setHasHorizontalScroller:"), false)
	objc.Send[objc.ID](scrollView, objc.RegisterName("setBorderType:"), 1) // Line border

	// Document view (contains todos)
	todoListView = NewView(NSRect{
		Origin: NSPoint{X: 0, Y: 0},
		Size:   NSSize{Width: 460, Height: 310},
	})
	objc.Send[objc.ID](scrollView, objc.RegisterName("setDocumentView:"), todoListView.ID())
	objc.Send[objc.ID](contentView, objc.RegisterName("addSubview:"), scrollView)

	// Bottom buttons
	clearBtn := NewButton(NSRect{
		Origin: NSPoint{X: 20, Y: 20},
		Size:   NSSize{Width: 120, Height: 30},
	})
	clearBtn.SetTitle(NewNSString("Clear All"))
	clearBtn.SetBezelStyle(1)
	clearHandler := createClearAllHandler()
	clearBtn.SetTarget(clearHandler)
	clearBtn.SetAction(objc.RegisterName("clearAll:"))
	objc.Send[objc.ID](contentView, objc.RegisterName("addSubview:"), clearBtn.ID())

	prefsBtn := NewButton(NSRect{
		Origin: NSPoint{X: 360, Y: 20},
		Size:   NSSize{Width: 120, Height: 30},
	})
	prefsBtn.SetTitle(NewNSString("Preferences..."))
	prefsBtn.SetBezelStyle(1)
	prefsHandler := createPreferencesHandler()
	prefsBtn.SetTarget(prefsHandler)
	prefsBtn.SetAction(objc.RegisterName("openPrefs:"))
	objc.Send[objc.ID](contentView, objc.RegisterName("addSubview:"), prefsBtn.ID())
}

// createPreferencesWindow creates the preferences window
func createPreferencesWindow() {
	prefsWindow = appkit.WindowClass.Alloc()

	frame := NSRect{
		Origin: NSPoint{X: 300, Y: 300},
		Size:   NSSize{Width: 400, Height: 200},
	}

	objc.Send[objc.ID](
		prefsWindow.ID(),
		objc.RegisterName("initWithContentRect:styleMask:backing:defer:"),
		frame,
		1|2, // Titled | Closable (not resizable)
		2,   // NSBackingStoreBuffered
		false,
	)

	prefsWindow.SetTitle(NewNSString("Preferences"))

	delegate := createPrefsDelegate()
	prefsWindow.SetDelegate(delegate)

	contentView := objc.Send[objc.ID](prefsWindow.ID(), objc.RegisterName("contentView"))

	// Title
	titleLabel := NewTextField(NSRect{
		Origin: NSPoint{X: 20, Y: 150},
		Size:   NSSize{Width: 360, Height: 30},
	})
	titleLabel.SetStringValue(NewNSString("Application Settings"))
	titleLabel.SetEditable(false)
	titleLabel.SetBordered(false)
	titleLabel.SetBackgroundColor(objc.ID(0))
	titleLabel.SetAlignment(2)
	font := objc.Send[objc.ID](objc.ID(objc.GetClass("NSFont")), objc.RegisterName("boldSystemFontOfSize:"), 16.0)
	objc.Send[objc.ID](titleLabel.ID(), objc.RegisterName("setFont:"), font)
	objc.Send[objc.ID](contentView, objc.RegisterName("addSubview:"), titleLabel.ID())

	// Max items setting
	maxLabel := NewTextField(NSRect{
		Origin: NSPoint{X: 20, Y: 100},
		Size:   NSSize{Width: 200, Height: 20},
	})
	maxLabel.SetStringValue(NewNSString("Maximum todo items (1-50):"))
	maxLabel.SetEditable(false)
	maxLabel.SetBordered(false)
	maxLabel.SetBackgroundColor(objc.ID(0))
	objc.Send[objc.ID](contentView, objc.RegisterName("addSubview:"), maxLabel.ID())

	maxItemsField = NewTextField(NSRect{
		Origin: NSPoint{X: 230, Y: 98},
		Size:   NSSize{Width: 80, Height: 22},
	})
	maxItemsField.SetStringValue(NewNSString(fmt.Sprintf("%d", maxTodoItems)))
	objc.Send[objc.ID](contentView, objc.RegisterName("addSubview:"), maxItemsField.ID())

	// Save button
	saveBtn := NewButton(NSRect{
		Origin: NSPoint{X: 260, Y: 20},
		Size:   NSSize{Width: 120, Height: 30},
	})
	saveBtn.SetTitle(NewNSString("Save"))
	saveBtn.SetBezelStyle(1)
	saveHandler := createSavePrefsHandler()
	saveBtn.SetTarget(saveHandler)
	saveBtn.SetAction(objc.RegisterName("savePrefs:"))
	objc.Send[objc.ID](contentView, objc.RegisterName("addSubview:"), saveBtn.ID())

	isPrefsOpen = true
	prefsWindow.MakeKeyAndOrderFront(objc.ID(0))
}

func main() {
	fmt.Println("=== Multi-Window Todo App (Generated Bindings Only) ===\n")

	// Get NSApplication
	app = objc.Send[objc.ID](objc.ID(objc.GetClass("NSApplication")), objc.RegisterName("sharedApplication"))
	objc.Send[objc.ID](app, objc.RegisterName("setActivationPolicy:"), 0) // Regular app

	// Create main window
	createMainWindow()
	mainWindow.MakeKeyAndOrderFront(objc.ID(0))

	// Activate app
	objc.Send[objc.ID](app, objc.RegisterName("activateIgnoringOtherApps:"), true)

	fmt.Println("✅ Multi-window todo app created")
	fmt.Println("   Features:")
	fmt.Println("   - Add and remove todo items")
	fmt.Println("   - Dynamic UI updates")
	fmt.Println("   - Preferences window")
	fmt.Println("   - Configurable settings")
	fmt.Println("   - All using generated bindings!")
	fmt.Println()

	// Run event loop
	objc.Send[objc.ID](app, objc.RegisterName("run"))
}
