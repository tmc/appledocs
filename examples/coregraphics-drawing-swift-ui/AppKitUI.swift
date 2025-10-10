// AppKitUI.swift - Minimal AppKit UI for CoreGraphics drawing

import Cocoa

// Custom view with CoreGraphics drawing
class DrawingView: NSView {
    override func draw(_ dirtyRect: NSRect) {
        guard let context = NSGraphicsContext.current?.cgContext else { return }

        // White background
        context.setFillColor(red: 1.0, green: 1.0, blue: 1.0, alpha: 1.0)
        context.fill(bounds)

        // Blue rectangle
        context.setFillColor(red: 0.2, green: 0.4, blue: 0.8, alpha: 1.0)
        context.fill(CGRect(x: 50, y: 50, width: 100, height: 100))

        // Red circle with stroke
        context.setFillColor(red: 0.8, green: 0.2, blue: 0.2, alpha: 1.0)
        context.fillEllipse(in: CGRect(x: 200, y: 50, width: 150, height: 150))
        context.setStrokeColor(red: 0.0, green: 0.0, blue: 0.0, alpha: 1.0)
        context.setLineWidth(3.0)
        context.strokeEllipse(in: CGRect(x: 200, y: 50, width: 150, height: 150))

        // Green triangle
        context.beginPath()
        context.move(to: CGPoint(x: 100, y: 300))
        context.addLine(to: CGPoint(x: 200, y: 250))
        context.addLine(to: CGPoint(x: 150, y: 350))
        context.closePath()
        context.setFillColor(red: 0.2, green: 0.8, blue: 0.2, alpha: 1.0)
        context.fillPath()
    }
}

// App delegate for auto-quit on window close
class AppDelegate: NSObject, NSApplicationDelegate {
    func applicationShouldTerminateAfterLastWindowClosed(_ sender: NSApplication) -> Bool {
        return true
    }
}

// Global state
private var appDelegate: AppDelegate?
private var window: NSWindow?

/// Run the application
@_cdecl("ui_run")
public func uiRun() {
    let app = NSApplication.shared
    app.setActivationPolicy(.regular)

    appDelegate = AppDelegate()
    app.delegate = appDelegate

    // Create window
    window = NSWindow(
        contentRect: NSRect(x: 0, y: 0, width: 400, height: 400),
        styleMask: [.titled, .closable, .miniaturizable],
        backing: .buffered,
        defer: false
    )
    window?.title = "CoreGraphics - Swift + Go"
    window?.center()
    window?.contentView = DrawingView()
    window?.makeKeyAndOrderFront(nil)

    app.activate(ignoringOtherApps: true)
    app.run()
}
