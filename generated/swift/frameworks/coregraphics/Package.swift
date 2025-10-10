// swift-tools-version: 5.9
import PackageDescription

let package = Package(
    name: "CoreGraphicsSwift",
    platforms: [
        .macOS(.v13)
    ],
    products: [
        .library(
            name: "CoreGraphicsSwift",
            type: .dynamic,
            targets: ["CoreGraphicsSwift"]
        ),
    ],
    targets: [
        .target(
            name: "CoreGraphicsSwift",
            path: ".",
            sources: ["CoreGraphics.swift"]
        ),
    ]
)
