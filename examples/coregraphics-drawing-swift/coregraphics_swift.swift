// coregraphics_swift.swift - Swift wrapper for CoreGraphics drawing operations
// Export C-compatible functions for Go to call via purego

import Foundation
import CoreGraphics
import ImageIO
import UniformTypeIdentifiers

// MARK: - Context Creation

/// Create a bitmap context for drawing
/// Returns an opaque pointer to CGContext (managed by Swift)
@_cdecl("cg_create_bitmap_context")
public func cgCreateBitmapContext(
    _ width: Int32,
    _ height: Int32
) -> UnsafeMutableRawPointer? {
    let colorSpace = CGColorSpaceCreateDeviceRGB()
    let bitmapInfo = CGImageAlphaInfo.premultipliedLast.rawValue

    guard let context = CGContext(
        data: nil,
        width: Int(width),
        height: Int(height),
        bitsPerComponent: 8,
        bytesPerRow: 0,
        space: colorSpace,
        bitmapInfo: bitmapInfo
    ) else {
        return nil
    }

    // Return retained reference for Go to manage
    return Unmanaged.passRetained(context as AnyObject).toOpaque()
}

/// Release a bitmap context
@_cdecl("cg_release_context")
public func cgReleaseContext(_ ctx: UnsafeMutableRawPointer) {
    Unmanaged<AnyObject>.fromOpaque(ctx).release()
}

// MARK: - Drawing Operations

/// Set the fill color (RGBA, 0.0-1.0)
@_cdecl("cg_set_fill_color")
public func cgSetFillColor(
    _ ctx: UnsafeMutableRawPointer,
    _ red: Double,
    _ green: Double,
    _ blue: Double,
    _ alpha: Double
) {
    let context = Unmanaged<AnyObject>.fromOpaque(ctx).takeUnretainedValue() as! CGContext
    context.setFillColor(red: red, green: green, blue: blue, alpha: alpha)
}

/// Set the stroke color (RGBA, 0.0-1.0)
@_cdecl("cg_set_stroke_color")
public func cgSetStrokeColor(
    _ ctx: UnsafeMutableRawPointer,
    _ red: Double,
    _ green: Double,
    _ blue: Double,
    _ alpha: Double
) {
    let context = Unmanaged<AnyObject>.fromOpaque(ctx).takeUnretainedValue() as! CGContext
    context.setStrokeColor(red: red, green: green, blue: blue, alpha: alpha)
}

/// Set the line width
@_cdecl("cg_set_line_width")
public func cgSetLineWidth(_ ctx: UnsafeMutableRawPointer, _ width: Double) {
    let context = Unmanaged<AnyObject>.fromOpaque(ctx).takeUnretainedValue() as! CGContext
    context.setLineWidth(width)
}

/// Fill a rectangle
@_cdecl("cg_fill_rect")
public func cgFillRect(
    _ ctx: UnsafeMutableRawPointer,
    _ x: Double,
    _ y: Double,
    _ width: Double,
    _ height: Double
) {
    let context = Unmanaged<AnyObject>.fromOpaque(ctx).takeUnretainedValue() as! CGContext
    let rect = CGRect(x: x, y: y, width: width, height: height)
    context.fill(rect)
}

/// Stroke a rectangle
@_cdecl("cg_stroke_rect")
public func cgStrokeRect(
    _ ctx: UnsafeMutableRawPointer,
    _ x: Double,
    _ y: Double,
    _ width: Double,
    _ height: Double
) {
    let context = Unmanaged<AnyObject>.fromOpaque(ctx).takeUnretainedValue() as! CGContext
    let rect = CGRect(x: x, y: y, width: width, height: height)
    context.stroke(rect)
}

/// Fill an ellipse
@_cdecl("cg_fill_ellipse")
public func cgFillEllipse(
    _ ctx: UnsafeMutableRawPointer,
    _ x: Double,
    _ y: Double,
    _ width: Double,
    _ height: Double
) {
    let context = Unmanaged<AnyObject>.fromOpaque(ctx).takeUnretainedValue() as! CGContext
    let rect = CGRect(x: x, y: y, width: width, height: height)
    context.fillEllipse(in: rect)
}

/// Stroke an ellipse
@_cdecl("cg_stroke_ellipse")
public func cgStrokeEllipse(
    _ ctx: UnsafeMutableRawPointer,
    _ x: Double,
    _ y: Double,
    _ width: Double,
    _ height: Double
) {
    let context = Unmanaged<AnyObject>.fromOpaque(ctx).takeUnretainedValue() as! CGContext
    let rect = CGRect(x: x, y: y, width: width, height: height)
    context.strokeEllipse(in: rect)
}

// MARK: - Path Operations

/// Begin a new path
@_cdecl("cg_begin_path")
public func cgBeginPath(_ ctx: UnsafeMutableRawPointer) {
    let context = Unmanaged<AnyObject>.fromOpaque(ctx).takeUnretainedValue() as! CGContext
    context.beginPath()
}

/// Move to point
@_cdecl("cg_move_to")
public func cgMoveTo(_ ctx: UnsafeMutableRawPointer, _ x: Double, _ y: Double) {
    let context = Unmanaged<AnyObject>.fromOpaque(ctx).takeUnretainedValue() as! CGContext
    context.move(to: CGPoint(x: x, y: y))
}

/// Add line to point
@_cdecl("cg_add_line_to")
public func cgAddLineTo(_ ctx: UnsafeMutableRawPointer, _ x: Double, _ y: Double) {
    let context = Unmanaged<AnyObject>.fromOpaque(ctx).takeUnretainedValue() as! CGContext
    context.addLine(to: CGPoint(x: x, y: y))
}

/// Close the current path
@_cdecl("cg_close_path")
public func cgClosePath(_ ctx: UnsafeMutableRawPointer) {
    let context = Unmanaged<AnyObject>.fromOpaque(ctx).takeUnretainedValue() as! CGContext
    context.closePath()
}

/// Stroke the current path
@_cdecl("cg_stroke_path")
public func cgStrokePath(_ ctx: UnsafeMutableRawPointer) {
    let context = Unmanaged<AnyObject>.fromOpaque(ctx).takeUnretainedValue() as! CGContext
    context.strokePath()
}

/// Fill the current path
@_cdecl("cg_fill_path")
public func cgFillPath(_ ctx: UnsafeMutableRawPointer) {
    let context = Unmanaged<AnyObject>.fromOpaque(ctx).takeUnretainedValue() as! CGContext
    context.fillPath()
}

// MARK: - Image Export

/// Create a PNG image from the context
/// Returns the number of bytes written, or -1 on error
/// The data pointer must be freed by the caller
@_cdecl("cg_create_png_data")
public func cgCreatePNGData(
    _ ctx: UnsafeMutableRawPointer,
    _ outDataPtr: UnsafeMutablePointer<UnsafeMutablePointer<UInt8>?>,
    _ outLength: UnsafeMutablePointer<Int>
) -> Int32 {
    let context = Unmanaged<AnyObject>.fromOpaque(ctx).takeUnretainedValue() as! CGContext

    guard let image = context.makeImage() else {
        return -1
    }

    let mutableData = NSMutableData()
    guard let destination = CGImageDestinationCreateWithData(
        mutableData as CFMutableData,
        UTType.png.identifier as CFString,
        1,
        nil as CFDictionary?
    ) else {
        return -1
    }

    CGImageDestinationAddImage(destination, image, nil as CFDictionary?)

    guard CGImageDestinationFinalize(destination) else {
        return -1
    }

    let length = mutableData.length
    let buffer = UnsafeMutablePointer<UInt8>.allocate(capacity: length)
    mutableData.getBytes(buffer, length: length)

    outDataPtr.pointee = buffer
    outLength.pointee = length

    return Int32(length)
}

/// Free PNG data allocated by cg_create_png_data
@_cdecl("cg_free_png_data")
public func cgFreePNGData(_ data: UnsafeMutablePointer<UInt8>) {
    data.deallocate()
}
