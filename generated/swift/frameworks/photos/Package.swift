// swift-tools-version: 5.9
import PackageDescription

let package = Package(
    name: "PhotosSwift",
    platforms: [
        .macOS(.v13)
    ],
    products: [
        .library(
            name: "PhotosSwift",
            type: .dynamic,
            targets: ["PhotosSwift"]
        )
    ],
    targets: [
        .target(
            name: "PhotosSwift",
            path: ".",
            sources: ["PhotosSwift.swift"]
        )
    ]
)
