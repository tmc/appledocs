# Multi-Window Todo App Example

An advanced example demonstrating the generated AppKit bindings with a multi-window todo list application.

## Features

This example showcases:

- **Multiple Windows**: Main window + Preferences window
- **Complex UI Layout**: Dynamic todo list with scrolling, input fields, multiple buttons
- **Advanced Event Handling**: Add, remove, and manage todo items
- **Window Management**: Opening and closing secondary windows
- **Dynamic Content**: Real-time UI updates as todos are added/removed
- **Settings Management**: Configurable maximum number of todos
- **Pure Generated Bindings**: No DarwinKit dependency, only purego/objc + generated types

## Building and Running

```bash
cd examples/todo-multiwindow-app
go mod download
go build
./todo-multiwindow-app
```

## What You'll See

1. **Main Window**:
   - Text input field for new todos
   - "Add" button to add items to the list
   - Scrollable list of todo items with individual "Remove" buttons
   - Status label showing current/max todo count
   - "Clear All" button to remove all todos
   - "Preferences..." button to open settings

2. **Preferences Window**:
   - Setting to configure maximum number of todos (1-50)
   - "Save" button to apply changes

## Implementation Highlights

### Window Creation
Shows how to create multiple windows with different style masks and manage their lifecycle.

### Dynamic UI Updates
Demonstrates rebuilding the todo list UI dynamically as items are added or removed.

### Event Handler Patterns
Shows three patterns for event handling:
1. Static handlers (app delegate, clear all)
2. Singleton handlers (add todo, preferences)
3. Dynamic per-item handlers (delete buttons)

### Helper Functions
Well-organized helper functions for creating common UI elements:
- `NewButton()`, `NewTextField()`, `NewView()`
- `NewNSString()`, `GetStringValue()`
- `NewNSColor()`

### Scroll View
Demonstrates using NSScrollView for scrollable content areas.

## Code Organization

The example is structured with:
- Global state management
- Helper function library
- Event handler creators
- UI construction functions
- Clear separation of concerns

This demonstrates a practical pattern for building real AppKit applications using only the generated bindings.
