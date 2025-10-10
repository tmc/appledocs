// AppKitUI.swift - AppKit UI layer for displaying CoreGraphics drawings

import Cocoa
import CoreGraphics
import ImageIO
import UniformTypeIdentifiers

// MARK: - Custom View for Drawing

class DrawingView: NSView {
    var cgContext: CGContext?

    override func draw(_ dirtyRect: NSRect) {
        super.draw(dirtyRect)

        guard let context = NSGraphicsContext.current?.cgContext else { return }

        // White background
        context.setFillColor(red: 1.0, green: 1.0, blue: 1.0, alpha: 1.0)
        context.fill(bounds)

        // Draw blue rectangle
        context.setFillColor(red: 0.2, green: 0.4, blue: 0.8, alpha: 1.0)
        context.fill(CGRect(x: 50, y: 50, width: 100, height: 100))

        // Draw red circle with stroke
        context.setFillColor(red: 0.8, green: 0.2, blue: 0.2, alpha: 1.0)
        context.fillEllipse(in: CGRect(x: 200, y: 50, width: 150, height: 150))

        context.setStrokeColor(red: 0.0, green: 0.0, blue: 0.0, alpha: 1.0)
        context.setLineWidth(3.0)
        context.strokeEllipse(in: CGRect(x: 200, y: 50, width: 150, height: 150))

        // Draw green triangle using path
        context.beginPath()
        context.move(to: CGPoint(x: 100, y: 300))
        context.addLine(to: CGPoint(x: 200, y: 250))
        context.addLine(to: CGPoint(x: 150, y: 350))
        context.closePath()

        context.setFillColor(red: 0.2, green: 0.8, blue: 0.2, alpha: 1.0)
        context.fillPath()
    }
}

// MARK: - Window Controller

class DrawingWindowController: NSWindowController {
    convenience init() {
        let window = NSWindow(
            contentRect: NSRect(x: 0, y: 0, width: 400, height: 400),
            styleMask: [.titled, .closable, .miniaturizable, .resizable],
            backing: .buffered,
            defer: false
        )

        window.title = "CoreGraphics Drawing - Swift + Go"
        window.center()

        // Create custom drawing view
        let drawingView = DrawingView(frame: window.contentView!.bounds)
        drawingView.autoresizingMask = [.width, .height]
        window.contentView = drawingView

        self.init(window: window)
    }
}

// MARK: - C API for Go

private var applicationDelegate: AppDelegate?
private var windowController: DrawingWindowController?

// Application delegate
class AppDelegate: NSObject, NSApplicationDelegate {
    func applicationDidFinishLaunching(_ notification: Notification) {
        print("Swift: Application finished launching")
    }

    func applicationShouldTerminateAfterLastWindowClosed(_ sender: NSApplication) -> Bool {
        return true
    }
}

/// Initialize the AppKit application
@_cdecl("ui_init_app")
public func uiInitApp() {
    // Initialize application
    let app = NSApplication.shared
    app.setActivationPolicy(.regular)

    // Set up application delegate
    applicationDelegate = AppDelegate()
    app.delegate = applicationDelegate

    print("Swift: Application initialized")
}

/// Create and show the main window
@_cdecl("ui_create_window")
public func uiCreateWindow() {
    windowController = DrawingWindowController()
    windowController?.showWindow(nil)
    windowController?.window?.makeKeyAndOrderFront(nil)

    // Activate the application
    NSApplication.shared.activate(ignoringOtherApps: true)

    print("Swift: Window created and shown")
}

/// Run the application event loop
/// This function blocks until the application quits
@_cdecl("ui_run_app")
public func uiRunApp() {
    print("Swift: Starting application event loop")
    NSApplication.shared.run()
}

/// Get the window's backing CGContext for custom drawing
/// Returns an opaque pointer to the CGContext
@_cdecl("ui_get_context")
public func uiGetContext() -> OpaquePointer? {
    guard let window = windowController?.window,
          let contentView = window.contentView,
          let context = NSGraphicsContext.current?.cgContext else {
        return nil
    }

    return OpaquePointer(Unmanaged.passRetained(context as AnyObject).toOpaque())
}

/// Save the current window content as PNG
@_cdecl("ui_save_png")
public func uiSavePNG(
    _ pathPtr: UnsafePointer<CChar>,
    _ outSuccess: UnsafeMutablePointer<Bool>
) {
    let path = String(cString: pathPtr)

    guard let window = windowController?.window,
          let contentView = window.contentView,
          let bitmapRep = contentView.bitmapImageRepForCachingDisplay(in: contentView.bounds) else {
        outSuccess.pointee = false
        return
    }

    contentView.cacheDisplay(in: contentView.bounds, to: bitmapRep)

    guard let data = bitmapRep.representation(using: .png, properties: [:]) else {
        outSuccess.pointee = false
        return
    }

    do {
        try data.write(to: URL(fileURLWithPath: path))
        outSuccess.pointee = true
        print("Swift: Saved PNG to \(path)")
    } catch {
        print("Swift: Failed to save PNG: \(error)")
        outSuccess.pointee = false
    }
}

/// Quit the application
@_cdecl("ui_quit")
public func uiQuit() {
    NSApplication.shared.terminate(nil)
}
