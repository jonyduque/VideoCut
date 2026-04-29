# Implementation Plan: VideoCut

## Phase 1: Setup & Scaffolding

- [ ] **Task: Initialize Wails Project and Clean Scaffolding**
    - [ ] Deep Planning: Analyze Wails React/TS template structure and plan cleanup of example files.
    - [ ] Write Technical Specification: Detail the project folder structure and initial configurations.
    - [ ] Setup: Run `wails init` with React/TS template.
    - [ ] Review: Verify project builds and clean state is achieved.
- [ ] **Task: Configure Project Dependencies (Frontend & Backend)**
    - [ ] Deep Planning: Select specific versions for Zustand and other utility libraries.
    - [ ] Write Technical Specification: List all dependencies and their roles in the project.
    - [ ] Implement: Install Zustand and configure Go modules.
    - [ ] Review: Verify all dependencies are correctly resolved and integrated.
- [ ] **Task: Conductor - User Manual Verification 'Setup & Scaffolding' (Protocol in workflow.md)**

## Phase 2: Backend Development (Go Logic)

- [ ] **Task: Implement Path Resolution and Environment Verification**
    - [ ] Deep Planning: Explore `os.Executable` vs `os.Getwd` for portable path resolution.
    - [ ] Write Technical Specification: Define the `App` struct fields for path management and binary location logic.
    - [ ] Write Tests: Create tests for relative path resolution and FFmpeg binary existence check.
    - [ ] Implement: Create routines in `app.go` to locate binaries and manage directories.
    - [ ] Review: Verify logic handles different execution paths and restricted permissions.
- [ ] **Task: Implement Metadata Extraction (ffprobe Integration)**
    - [ ] Deep Planning: Analyze `ffprobe` output formats (JSON vs CSV) for efficient parsing in Go.
    - [ ] Write Technical Specification: Define the `GetDuration` binding and its error handling.
    - [ ] Write Tests: Mock `ffprobe` calls to verify duration extraction and error reporting.
    - [ ] Implement: Write the Go function to execute `ffprobe` and return the video duration.
    - [ ] Review: Check for resource leaks in process execution and efficiency of parsing.
- [ ] **Task: Implement Video Splitting Logic (FFmpeg Integration)**
    - [ ] Deep Planning: Determine optimal FFmpeg flags for frame-accurate re-encoding vs speed.
    - [ ] Write Technical Specification: Define the `SplitVideo` binding, including overlap math and output naming.
    - [ ] Write Tests: Create unit tests for the overlap calculation logic and command string construction.
    - [ ] Implement: Develop the FFmpeg invocation routine with support for the "Sobra" (overlap) feature.
    - [ ] Review: Verify the generated commands are correct and process management is resilient.
- [ ] **Task: Conductor - User Manual Verification 'Backend Development' (Protocol in workflow.md)**

## Phase 3: Frontend Development (React & UI)

- [ ] **Task: Implement UI Shell and Drag-and-Drop Zone**
    - [ ] Deep Planning: Design the React component tree and plan the D&D event handling.
    - [ ] Write Technical Specification: Detail the CSS layout for the reception area and D&D states.
    - [ ] Write Tests: Implement tests for the D&D listener and path capture logic.
    - [ ] Implement: Build the main interface with a modern/sleek look and interactive D&D area.
    - [ ] Review: Check for responsive behavior on the fixed-size window and visual feedback quality.
- [ ] **Task: Implement Timeline Slider and Metadata Preview**
    - [ ] Deep Planning: Choose a slider implementation (native vs library) for best UX and precision.
    - [ ] Write Technical Specification: Define the state sync between the slider and the duration metadata.
    - [ ] Write Tests: Test real-time updates of the selected time display and slider constraints.
    - [ ] Implement: Create the metadata display and the interactive timeline control.
    - [ ] Review: Verify synchronization between the UI state and the backend-provided duration.
- [ ] **Task: Implement Export Settings and Action Controls**
    - [ ] Deep Planning: Plan the conditional rendering for the custom destination selector.
    - [ ] Write Technical Specification: Detail the Zustand store structure for the export configuration.
    - [ ] Write Tests: Test the form validation and the "Processing" state transitions.
    - [ ] Implement: Build the settings form (Sobra, Naming, Destination) and the main action button.
    - [ ] Review: Verify UI locking during processing and clear error/success notifications.
- [ ] **Task: Conductor - User Manual Verification 'Frontend Development' (Protocol in workflow.md)**

## Phase 4: Final Validation & Resilience

- [ ] **Task: End-to-End Integration Testing**
    - [ ] Deep Planning: Define the critical user path from file drop to successful multi-segment cut.
    - [ ] Write Technical Specification: List the integration test scenarios (Standard cut, Multi-part cut, Error states).
    - [ ] Write Tests: Implement E2E tests simulating the full application lifecycle.
    - [ ] Implement: Refine IPC calls and state synchronization based on integration results.
    - [ ] Review: Analyze overall program efficiency and look for final bug fixes.
- [ ] **Task: Portability and Permission Stress Tests**
    - [ ] Deep Planning: Plan tests for restricted folder access and missing FFmpeg scenarios.
    - [ ] Write Technical Specification: Detail the resilience requirements and expected fallback behaviors.
    - [ ] Implement: Verify the app's behavior on a clean Windows environment without admin rights.
    - [ ] Review: Ensure all project guidelines and "Definition of Done" are met.
- [ ] **Task: Conductor - User Manual Verification 'Final Validation' (Protocol in workflow.md)**
