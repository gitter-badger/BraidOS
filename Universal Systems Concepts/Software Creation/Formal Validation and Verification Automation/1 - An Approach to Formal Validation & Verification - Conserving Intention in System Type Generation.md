#Conserving Intention in System Type Generation (An Approach to Formal Validation & Verification Part 1)

- **Abstract**
- **Introduction** (1 page)
- **The Problem** (1 page)
- **My Idea** (2 Pages)
- **The Details** (5 Pages)
- **Related Work** (1-2 Pages)
- **Conclusions and Further Work** (0.5 Pages)


##Introduction

In this first paper we would like to setup a *Theoretical Framework* for approaching the problem of Formal Validation & Verification in practical computation.

###The Problem of Formally Validating and Verifying Software

A modern problem in software systems is the common belief that 'if it compiles, it works'. This allows problems to persist in software implementations (potentially to the hazard of human lives) due to lack of formal *System* Validation and verification.

We define Formal Validation and Verification as: **The intention-conserving translation of Human Intent into System Concept followed by the concept-conserving translation of System Concept into implementation.**

- **Human Intention**(Definition)
- **System Concept** (Definition)
- **Conservation** (Definition)
- **Implementation** (Definition)

- Formal Validator - *application* specific
	- must be constructed per application
- Formal Verifier - *build-chain* specific
	- must be constructed per build-chain (set of System Concepts that are potential dependencies)

###Problems
- What kind of problems can arise?
	- Fault-Tolerance
	- Reliability Issues
- What should be done to avoid all problem types?

###Workflow Phase-based Validation and Verification

- *Design/Engineering Phase*
- **Validator** passes data from Design Phase to Implementation Phase
- *Implementation Phase*
- **Verifier** passes data from Implementation Phase to Product Phase
- *Product Phase*
- **Maintainer** (Debugger) and **Updater** (New Features/Modules) pass data from Product Phase to Design Phase

###Department-based (Human-Based) Validation and Verification

- **Implementation** Group
- **Quality Assurance** Group

##Using Principles of Symmetry as a Conservation Method

##Conserving Validity and Verification

##Conclusions



