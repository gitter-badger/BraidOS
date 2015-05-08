#Computational Braid Theory - Notes (Journals)

- Processors "traverse" a 2^64 code (space).
- Boolean Algebras -> Heyting Algebras -> Complete Heyting Algebras
	- An example of a Complete Heyting Algebra:
		- The system of all open sets of a given topological space, ordered by inclusion, is a complete heyting algebra.
- A **0-Strand** CDT Space has a **single-point configuration space**.
- A **1-Strand** CDT Space is **its own configuration space**.
	- This is the current method for processors, which at any given time is in a configuration over (F\_2)^64 space (for a 64-bit machine).
	- This tells us that a Configuration Space is the **state space** for a single program or a system (multiple programs).
- A **n-strand** CDT space has a **n-tuple configuration space** (where n>=2).
- A braid works on a cycle, each cycle allows for a configuration change.
- There are **3 ways to change configuration (and/or the configuration space):**
	1. A strand changes value
	2. Creation/Destruction of a strand
	3. CDT Space Change
- Movements across the configuration space (**configuration changes**) are (**compile** to) **canonical actions** (or reducible actions -- if a sequence of canonical actions) of the **Braid Group B\_n**.
	- An implementation of this is physical **Turing Operations**
		- e.g. pointer management, register management, etc. (Memory Space Traversal Operations)
- Using C.B.T. to construct Calculus
	- using strands over a space to derive the basis of calculus (differentiability and integrability), given a Type: Points, a Line := a traversal over terms in Type: point (dxdydz).
	- An integral has a 'starting' and 'ending' value, similiar to how a strand has a starting and ending value.
- **Graded Configuration Space** - A configuration space that is based on multiple braids.
- C.B.T. is about handling the realm of:
	- Multiple Programs
	- Multiple Universes
- In other words, CBT is a Model for creating **Computer Systems** (e.g. Proof Systems), sets of running programs (proofs) each with their own Type Universes.
- Thus, CBT is a model that allows us to create **System Concepts**
	- an example of a computer system could be: Robotic Swarms
- **Turing Machine**
	- Finite set of possible States
	- Finite Alphabet
	- Initial state
	- set of final or accepting states
	- Transition function (left/right)
- **Computational Braids**
	- Proof System
	- Universe
- **Virtual Memory** (or just memory generally) is a Type in a CDT space, that we can then define **Coordinate Types** (a System Type) to "handle" i.e. define interpretation functors that are structurally dependent on the Memory Type.
- There is a strong correlation between System Coordinate Types and interpretation functors.
