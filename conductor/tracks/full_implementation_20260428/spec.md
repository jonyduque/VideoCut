# Specification: Full Implementation of VideoCut

## 1. Goal
The objective of this track is to implement the VideoCut (Fatiador de Vídeos) desktop application, a portable utility for splitting videos with overlap, using Wails, Go, React, and FFmpeg.

## 2. Technical Scope
- **Backend:** Go-based logic for file I/O, OS path resolution, and FFmpeg process management.
- **Frontend:** React (TypeScript) SPA with Vanilla CSS for a minimalist, modern UI.
- **IPC:** Wails bindings for asynchronous communication between layers.
- **Binaries:** Bundled `ffmpeg.exe` and `ffprobe.exe` for portable execution.

## 3. Functional Requirements
- **Drag-and-Drop:** Intuitive file import.
- **Metadata Analysis:** Real-time duration extraction via `ffprobe`.
- **Precise Cutting:** Timeline-based split point selection.
- **Overlap Logic:** Configurable intersection time for context preservation.
- **Smart Naming:** Automatic output naming (`OriginalName_PartN.mp4`).
- **Destination Control:** Choice between source folder or custom path.

## 4. Architecture & Design
- **State Management:** Zustand for lightweight state handling.
- **Styling:** Vanilla CSS following a "Modern/Sleek" aesthetic.
- **Portability:** Zero-installation, relative path resolution for all assets.

## 5. Implementation Phases
1. **Setup & Scaffolding:** Wails initialization and basic structure.
2. **Backend Development:** Go routines for FFmpeg/ffprobe integration.
3. **Frontend Development:** UI components and IPC wiring.
4. **Resilience & Testing:** Final validation of the portable, non-admin workflow.
