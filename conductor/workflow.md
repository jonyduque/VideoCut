# Project Workflow

## Guiding Principles

1. **The Plan is the Source of Truth:** All work must be tracked in `plan.md`
2. **Deep Planning:** Every task must be planned deeply, exploring different implementation possibilities and ensuring cohesion with the entire project.
3. **The Tech Stack is Deliberate:** Changes to the tech stack must be documented in `tech-stack.md` *before* implementation
4. **Test-Driven Development:** Write unit tests before implementing functionality.
5. **High Code Coverage:** Aim for >80% code coverage for all modules.
6. **User Experience First:** Every decision should prioritize user experience.
7. **Rigorous Review:** Every completed task must undergo a thorough review for bugs, quality, and efficiency.
8. **Non-Interactive & CI-Aware:** Prefer non-interactive commands. Use `CI=true` for watch-mode tools (tests, linters) to ensure single execution.

## Task Workflow

All tasks follow a strict lifecycle:

### Standard Task Workflow

1. **Select Task:** Choose the next available task from `plan.md` in sequential order.

2. **Mark In Progress:** Before beginning work, edit `plan.md` and change the task from `[ ]` to `[~]`.

3. **Deep Planning & Technical Specification:**
   - **Research:** Analyze the task requirements and explore different technical possibilities for implementation.
   - **Cohesion Check:** Evaluate how the proposed solution fits with the project's existing architecture and goals.
   - **Define Specifications:** Establish the intended results and write a concise technical specification of the implementation strategy.

4. **Write Failing Tests (Red Phase):**
   - **Requirement:** If tests for the specific functionality do not exist, they must be written *before* any application code.
   - Create or update the test file for the feature or bug fix.
   - Write unit tests that clearly define the expected behavior based on the technical specification.
   - **CRITICAL:** Run the tests and confirm that they fail as expected. Do not proceed until you have failing tests.

5. **Implement to Pass Tests (Green Phase):**
   - Write the minimum amount of application code necessary to make the failing tests pass.
   - Run the test suite again and confirm that all tests now pass.

6. **Rigorous Code Review & Refactoring:**
   - **Quality Check:** Review the code for clarity, bugs, and strict adherence to project guidelines (`code_styleguides/`).
   - **Efficiency Analysis:** Look for opportunities to improve code quality and program efficiency.
   - **Refactor:** Improve the implementation while ensuring all tests continue to pass.

7. **Verify Coverage:** Run coverage reports using the project's chosen tools.
   - Target: >80% coverage for new code.

8. **Document Deviations:** If implementation differs from tech stack:
   - **STOP** implementation
   - Update `tech-stack.md` with new design
   - Add dated note explaining the change
   - Resume implementation

9. **Commit Code Changes:**
   - Stage all code changes related to the task.
   - Propose a clear, concise commit message (Conventional Commits format).
   - Perform the commit.

10. **Attach Task Summary with Git Notes:**
    - **Step 10.1: Get Commit Hash:** Obtain the hash of the *just-completed commit*.
    - **Step 10.2: Draft Note Content:** Create a detailed summary including the task name, planning rationale, changes made, and verification results.
    - **Step 10.3: Attach Note:** Use `git notes add -m "<note content>" <commit_hash>`.

11. **Get and Record Task Commit SHA:**
    - **Step 11.1: Update Plan:** Read `plan.md`, update status from `[~]` to `[x]`, and append the 7-character commit hash.
    - **Step 11.2: Write Plan:** Write the updated content back to `plan.md`.

12. **Commit Plan Update:**
    - Stage the modified `plan.md`.
    - Commit with `conductor(plan): Mark task '<TASK NAME>' as complete`.

### Phase Completion Verification and Checkpointing Protocol

**Trigger:** Executed immediately after a task is completed that concludes a phase in `plan.md`.

1.  **Announce Protocol Start:** Inform the user that the phase is complete and verification has begun.

2.  **Ensure Test Coverage for Phase Changes:**
    -   **Step 2.1: Determine Phase Scope:** Identify files changed since the last checkpoint.
    -   **Step 2.2: List Changed Files:** `git diff --name-only <previous_checkpoint_sha> HEAD`.
    -   **Step 2.3: Verify and Create Tests:** Ensure every code file has a corresponding test file validating the phase's requirements.

3.  **Execute Automated Tests with Proactive Debugging:**
    -   Announce the test command (e.g., `CI=true go test ./...`).
    -   Execute and verify. If failure occurs, attempt up to two proposed fixes before stopping for guidance.

4.  **Propose a Detailed Manual Verification Plan:**
    -   Analyze `product.md` and `plan.md` to create a step-by-step verification guide for the user.
    -   Format:
        ```
        **Manual Verification Steps:**
        1. [Command/Action]
        2. [Expected Outcome]
        ```

5.  **Await Explicit User Feedback:**
    -   **PAUSE** and await explicit user confirmation ("yes" or feedback) before proceeding.

6.  **Create Checkpoint Commit:**
    -   Stage all changes and commit: `conductor(checkpoint): Checkpoint end of Phase X`.

7.  **Attach Auditable Verification Report:**
    -   Attach a git note with the full verification report (tests, manual steps, user feedback).

8.  **Get and Record Phase Checkpoint SHA:**
    -   Update `plan.md` header with `[checkpoint: <sha>]`.

9. **Commit Plan Update:**
    - Commit the updated `plan.md`: `conductor(plan): Mark phase '<PHASE NAME>' as complete`.

10.  **Announce Completion:** Inform the user the checkpoint is created and the phase is fully verified.

### Quality Gates

Before marking any task complete, verify:

- [ ] Planning logic covers all edge cases and project cohesion.
- [ ] Technical specifications are established before coding.
- [ ] All tests pass.
- [ ] Code coverage meets requirements (>80%).
- [ ] Code follows project's code style guidelines.
- [ ] All public functions/methods are documented (e.g., GoDoc, JSDoc).
- [ ] Type safety is strictly enforced.
- [ ] No linting or static analysis errors.
- [ ] Code has been reviewed for bugs and efficiency improvements.
- [ ] Documentation updated if needed.

## Development Commands

### Setup
```bash
# Go dependencies
go mod tidy
# Frontend dependencies
cd frontend && npm install
```

### Daily Development
```bash
# Start Wails dev mode
wails dev
# Run Go tests
go test ./...
# Run Frontend tests
cd frontend && npm test
```

### Before Committing
```bash
# Full check (tests + lint)
# [Define project specific check command here]
```

## Testing Requirements

### Unit Testing
- Every module must have corresponding tests.
- Write tests *before* implementation for new functionality.
- Mock external dependencies (e.g., FFmpeg calls).
- Test both success and failure cases.

### Integration Testing
- Test complete user flows (Drag-and-drop to Cut).
- Verify file system operations and path resolutions.

## Code Review Process

### Self-Review Checklist
Before marking task as complete:

1. **Planning & Specs**
   - Does the implementation match the pre-coding specification?
   - Were all explored possibilities considered?

2. **Functionality & Efficiency**
   - Feature works as specified.
   - Algorithms and I/O operations are optimized for efficiency.
   - Edge cases (e.g., zero overlap, very short videos) handled.

3. **Code Quality**
   - Follows style guide.
   - Variable/function names are clear and descriptive.
   - Appropriate comments for complex logic.

4. **Testing & Security**
   - Unit tests are comprehensive and passed.
   - Coverage is >80%.
   - No hardcoded paths; everything is relative to the executable.
   - No security vulnerabilities in OS command execution.

## Commit Guidelines

### Message Format
```
<type>(<scope>): <description>

[optional body]

[optional footer]
```

### Types
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation only
- `style`: Formatting, missing semicolons, etc.
- `refactor`: Code change that neither fixes a bug nor adds a feature
- `test`: Adding missing tests
- `chore`: Maintenance tasks

## Definition of Done

A task is complete when:

1. Deep planning and technical specs are documented.
2. Unit tests are written and passing.
3. Code is implemented and passes tests.
4. Code review for bugs, style, and efficiency is complete.
5. Code coverage meets project requirements (>80%).
6. Changes committed with proper Conventional Commit message.
7. Git note with task summary attached to the commit.
8. Implementation notes and SHA added to `plan.md`.
