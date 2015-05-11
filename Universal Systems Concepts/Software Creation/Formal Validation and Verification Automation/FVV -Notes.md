#Formal Validation and Verification Automation 
##Modern Problems and Solutions

A modern problem in software systems is the common belief that 'if it compiles, it works'. This allows problems to persist in software implementations (potentially to the hazard of human lives) due to lack of formal *System* Validation and verification.

###Problems
- What kind of problems can arise?
	- Fault-Tolerance
	- Reliability Issues
- What should be done to avoid all problem types?

###Phase-based Validation and Verification

- *Design/Engineering Phase*
- **Validator** passes data from Design Phase to Implementation Phase
- *Implementation Phase*
- **Verifier** passes data from Implementation Phase to Product Phase
- *Product Phase*
- **Maintainer** (Debugger) and **Updater** (New Features/Modules) pass data from Product Phase to Design Phase

###Department-based (Human-Based) Validation and Verification

- **Implementation** Group
- **Quality Assurance** Group

###Software Design Methods

- T.D.D.
- Testing Suites/Frameworks
- Performance/Load Testing
- Penetration Testing
- etc.

##System Architectures and F.V.V.

Formal Validation and Verification is limited to the Universe that the programs are dependent on. Sometimes a particular systems architecture will not offer the option to truly allow a program over it to guarantee a verification of some set of system requirements. The ability to dynamically generate architectures would be a useful architecture itself. {Future research (in progress)}.

###FVV - What is it commonly?

One way to look at FVV is: **The intention-conserving translation of Human Intent into System Concept followed by the concept-conserving translation of System Concept into implementation.**

- **Human Intention**(Definition)
- **System Concept** (Definition)
- **Conservation** (Definition)
- **Implementation** (Definition)

- Formal Validator - *application* specific
	- must be constructed per application
- Formal Verifier - *build-chain* specific
	- must be constructed per build-chain (set of System Concepts that are potential dependencies)


##The Paradox

**How can we formally validate and verify a formal validator and verifier?**

##Symmetry

**Symmetry implies conservation**

###Symmetry in Mathematics

Symmetry is a natural bijection between two sets (or categorically -- a structure preserving morphism between two objects inside the Category -- as in the case of the Category of Groups, a bijection is a group isomorphism).

###Symmetry in Computing
**In computing, Symmetry implies immutability**

The 'set' is a set of types (CDTs) that is shared (bijected) between componets in a system.

- Networking (Bijection = Protocol i.e. a set of shared types)
- Software Tests (Conservation of Validity & Verification)
- Consensus
- Atomic Operations (are Symmetric Functions)
- Copy Functions for arbitrary Data