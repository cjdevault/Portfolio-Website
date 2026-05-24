# Oracle Database Engineering Project

A two-tier database application for CSC 460 (Database Design) — an Oracle-backed pet café with a Java JDBC front-end. The system supports member accounts, reservations, visits, orders, pet care records, adoptions, events, and reporting through a menu-driven CLI.
  
**Authors:** Nathan Tebbs, Krish Patel, CJ De Vault, Shayden Lowry  
**Completed:** December 2025

---

## Overview

This project implements a full client–server stack:

1. **Database back-end** — Relational schema on Oracle (hosted on `aloe.cs.arizona.edu`), designed with an E-R model, functional dependency analysis, and normalization to 3NF/BCNF where applicable.
2. **JDBC front-end** — Text-based Java application that connects to Oracle and exposes CRUD operations and reports through nested menus.

![Final ER Diagram](docs/Final%20ER%20Diagram.png)

---

## Features

| Module | Capabilities |
|--------|--------------|
| **Member Management** | Register, look up, update, delete members; view visit history and adoption applications |
| **Reservation Management** | Create and manage café reservations |
| **Visit Management** | Record visits with membership tier snapshot logic |
| **Order Management** | Create and manage customer orders |
| **Pet Management** | Pets, health records, and behavioral assessments |
| **Adoption Management** | Track adoption applications and status |
| **Events** | Event scheduling and registration |
| **Reports** | SQL-backed reports (e.g., pet popularity) |

---

## Tech Stack

- **Language:** Java 16
- **Database:** Oracle DBMS
- **Connectivity:** JDBC (`ojdbc8.jar`)
- **Build:** GNU Make
- **Documentation:** LaTeX (`design.tex`), Markdown schema docs

---

## Project Structure

```
CSC460Prog4/
├── src/                          # Java source files
│   ├── PetCafeApplication.java   # Main entry point & menu system
│   ├── DatabaseConfig.java       # JDBC connection configuration
│   ├── MemberManager.java
│   ├── ReservationManager.java
│   ├── VisitManager.java
│   ├── OrderManager.java
│   ├── PetManager.java
│   ├── HealthRecordManager.java
│   ├── BehavioralAssessmentManager.java
│   ├── AdoptionManager.java
│   ├── EventManager.java
│   └── ReportManager.java
├── sql/
│   ├── create_schema.sql         # DDL — tables, constraints, sequences
│   ├── insert_sample_data.sql    # Sample data population
│   └── verify_data.sql           # Data verification queries
├── docs/
│   ├── Final ER Diagram.png
│   ├── design.tex                # LaTeX design document source
│   ├── SCHEMA_DOCUMENTATION.md
│   ├── DATA_DICTIONARY.md
│   ├── ER_DIAGRAM_TXT.md
│   └── FD_NORMALIZATION_ANALYSIS.md
├── Makefile                      # Build, run, and LaTeX automation
└── ReadMe.txt                    # Team contribution log
```

Compiled output (`bin/`, `design.pdf`, `docs/latex/`) is generated locally and excluded via `.gitignore`.

---

## Getting Started

### Prerequisites

- Java 16+
- Oracle JDBC driver at `/usr/lib/oracle/19.8/client64/lib/ojdbc8.jar` (available on UA **lectura** servers)
- An Oracle account on lectura with the schema loaded

### Database Setup

Run these scripts against your Oracle account **before** starting the application:

```bash
sql/create_schema.sql
sql/insert_sample_data.sql
```

Optional verification:

```bash
sql/verify_data.sql
```

### Build & Run

Intended to run on **lectura** while connected to the Oracle instance on aloe.

```bash
# Compile all Java sources
make build

# Run the application
make run USER=your_lectura_username PASS=your_lectura_password
```

Example:

```bash
make run USER=nate PASS=secret123
```

The entry point is `PetCafeApplication`, which accepts Oracle credentials and launches the menu-driven system.

### Build Design Document

```bash
make design        # Compiles docs/design.tex → design.pdf
make clean-docs    # Remove LaTeX auxiliary files
make clean-all     # Remove all build output (Java + LaTeX)
```

---

## My Contributions

**CJ De Vault**

- Implemented `MemberManager` and `ReservationManager`
- Created `DatabaseConfig` for JDBC connection handling
- Drafted the preliminary `PetCafeApplication` class
- Wired menu logic for member management, reservation management, order management, and adoption management, with revisions to other submenus
- Integrated `HealthRecordManager`, `PetManager`, and `BehavioralAssessmentManager` into the pet management submenu
- Authored functional dependency and normalization analysis ([`docs/FD_NORMALIZATION_ANALYSIS.md`](docs/FD_NORMALIZATION_ANALYSIS.md))
- Drafted [`sql/create_schema.sql`](sql/create_schema.sql)

---

## Team Contributions

<details>
<summary><strong>Krish Patel</strong></summary>

- Full implementation of `AdoptionManager.java`, `BehavioralAssessmentManager.java`, `PetManager.java`, `HealthRecordManager.java`
- Full testing of the above manager classes
- Full rough ER diagram sketch
- Final ER diagram sketch based on TA feedback and later additions
- Partial integration of `AdoptionManager` and `PetManager` into `PetCafeApplication`

</details>

<details>
<summary><strong>Nathan Tebbs</strong></summary>

- Full construction of the Makefile (Java compilation, run support, LaTeX build system)
- Implemented `OrderManager`, `VisitManager`, and related CRUD operations
- Added membership tier snapshot logic to visit creation
- Wrote and wired menu logic for order and visit management
- Documented the Pet Popularity Report (Query 4)
- Updated the Data Dictionary and Schema Documentation
- Drafted and assembled the complete `design.tex`
- Created the LaTeX output pipeline (`docs/latex/`)
- Updated documentation and planning to reflect TA feedback

</details>

<details>
<summary><strong>Shayden Lowry</strong></summary>

- Polished `create_schema.sql`
- Wrote SQL for populating the DB with test data
- Wrote SQL for verifying data
- Wrote `EventManager` JDBC class
- Error testing for various JDBC classes

</details>

---

## Documentation

| Document | Description |
|----------|-------------|
| [`docs/SCHEMA_DOCUMENTATION.md`](docs/SCHEMA_DOCUMENTATION.md) | Table purposes, attributes, relationships, and constraints |
| [`docs/DATA_DICTIONARY.md`](docs/DATA_DICTIONARY.md) | Full data dictionary |
| [`docs/FD_NORMALIZATION_ANALYSIS.md`](docs/FD_NORMALIZATION_ANALYSIS.md) | Functional dependencies and 3NF/BCNF analysis |
| [`docs/ER_DIAGRAM_TXT.md`](docs/ER_DIAGRAM_TXT.md) | Text description of the E-R model |
| [`docs/Final ER Diagram.png`](docs/Final%20ER%20Diagram.png) | Final entity–relationship diagram |

---

## License

Academic project completed for CSC 460 at the University of Arizona. Not licensed for commercial use.
