# Domain Documentation Configuration

This repository uses a **single-context** layout for its domain documentation.

This means that:

*   **`CONTEXT.md`**: A single `CONTEXT.md` file located at the repository root (`NEGESTAT/CONTEXT.md`) serves as the primary source of domain language and project overview for all agent skills.
*   **Architectural Decision Records (ADRs)**: Architectural Decision Records are located in a single `docs/adr/` directory at the repository root (`NEGESTAT/docs/adr/`).

Agent skills like `improve-codebase-architecture`, `diagnosing-bugs`, and `tdd` will read these files to understand the project's domain language and past architectural decisions.
