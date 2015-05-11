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

Formal Validation and Verification is limited on Instruction Set Architectures. We can only get maximal information-theoretic compression without having static dependencies (such as a hardware ISA). {This gives rise to the need for the **RAPS computer Architecture**}

##Systems Concept Implementation vs. Physical Implementation

- Systems Concepts are methods in Explicit Type Theory
- Physical Devices/Architectures (Implementations) are the "covering space" for Generalized Braids (Machines) - they are propositions (set of all proofs - i.e. set of all objects that can be formally constructed via a concept-preserving translation from a Systems Concept).