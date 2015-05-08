#C.B.T. Notes

- C.B.T. - is for modeling anything that can be considered a "program" or a "system".
- The two codeword Turing Machine. C\_1 <=> C\_2
	- [http://en.wikipedia.org/wiki/Counter_machine#Two-counter_machines_are_Turing_equivalent_.28with_a_caveat.29](http://en.wikipedia.org/wiki/Counter_machine#Two-counter_machines_are_Turing_equivalent_.28with_a_caveat.29)
- A counter machine is one of four types of register machines:
	- [http://en.wikipedia.org/wiki/Register_machine#Overview](http://en.wikipedia.org/wiki/Register_machine#Overview)
- Turing Heads running over a tape produce strands over a space
- What is a computational braid?
	- A **computational braid** is a n-Proof System (n-set of proofs) over a Universe U\_i
		- such that the actions of the 'Braid Group' are function types over U\_i
- A good question is: what about Multi-Braided Systems?
	- What if each braid has its own Universe?
- **Braids = Systems** (not machines)
	- A 'Machine' is an implementation of a system
	- Multi-Braided System = System w/ Sub-Systems
		- How do they act as one system?
			- **A Braid of Braids** (??)
- Application of Interpretation Theory = Monads (??)
	- Adjoint Functors (F: U\_i -> U\_j) (G: U\_j -> U\_i)
	- G * F is a Monad functor (an endofunctor)

##Sequential Composition

- Sequential composition is the Time operation on a proof; it is the time that a certain computation runs on ... which means that it can vary from proof-to-proof.
- Thus, we say that Sequential Composition is **relative**
	- Discrete = Synchronous
	- Continuous = Asynchronous

##Systems Theory

- All modern day Operating Systems are based (not on a Memory Type or the bit) but are founded on a Single-Type Space: The File.
	- Files are the "Tape" for operating System Machines.
	- Programs "run over" files
- Systems Concepts - are system concepts applicable to every Universe?
	- for example, the concept of cache support can exist as a concept in multiple programming languages -- but the actual implementation, etc. can vary (wildly sometimes).
- Does C.B.T. tell us what we can do with a proof system over a Universe (??)
- **Action** - any proof that when run on a system can cause a change of state for that computer system.
	- Programming is a kind of action
- What is programming?
	- Any action on a computer system that causes change in the system is called **programming** the system.
- How does an action proof cause a proof system over a Universe to change?
	- These 'actions' are a major part of '**System Primitives**'
- What really is a **kernel**?? What really is an **operating System**?
	- A Filesystem as the **base Universe**
	- and some set of function types over that filesystem
- All programming is **meta-programming**
	- Proofs are sequences of sub-proofs
	- **Programming** is the art/science of putting sub-proofs together into one larger proof
- **Scheduling**
	- If enough processes are multiplexed/scheduled over a single resource such that the scheduler changes processes every cycle, then we have made the scheduler create a contiguous program of scheduling. This is a limit of a scheduler -- to be the only program running.
		- The scheduler is, itself, a program.
	- A scheduler can converge at being the only program running.
		- Thus, we might like to think that: the scheduler is the only program ever running. And that software management is about Dynamically Generating Sub-Proofs to run within the scheduler (??)

##Mathematics

- Braid Groups are generalizations of Symmetric Groups, since a Braid Group is about permutations and *how* the permutations are performed.
	- Thus, we say that Braid *Groups* must be generalized into Computational Braids, since Computational Braids do not have *Groups* **(??)**
	- Unless, in most computer systems we are constantly doing Group Homomorphisms **(??)**
- The Symmetric Group of a set of size n has a relationship to Galois Groups of polynomials of degree n
	- A Galois Group is a group of field automorphisms under composition.

##Modern Programming

- **Switch-Case Statements**
	- Function Types:
		- One Input
		- Multiple Outputs
			- Only one output is allowed to be non-null
- **If-Else Loops**
	- Function Types:
		- One Input
		- Multiple Outputs
			- Only One Output can be non-null

##Configuration Spaces

- Regardless of the CDT space, all configuration spaces can be assigned a metric
- Sequential Composition over the configuration space = "Actions" of the "Braid Group" on the CDT Space
- The **configuration dimension** of the **Configuration Universe** that is the Configuration Space of our proof system **can change**.
	- We can add or remove running proofs from the system
	- If we have a common Sequential Composition Unit e.g. a clock cycle
		- then we can use a consistent **Configuration Dimension** and use "multiplexing/concurrency" when we are adding or removing proofs.

##Braid Groups

- **Nielsen-Thurston Classification**
	- Periodic Braids
	- Reducible Braids
	- Pseudo-Anosov Braids
- The Braid Group is surjective into the symmetric group
	- The Braid Group takes into consideration how the Braids twist and cross
	- While the Symmetric Group only takes into consideration the possible permutations over n braids.

##Computational Braids

- A model of computation is any model that allows you to perform computations using a memory-based system
- The Turing Machine was revolutionary because it defined the basic principles of performing computations with a memory-based system
- Computational Braids are trying to generalize what a **memory-based computer system** is.
	- Can we generalize this **Memory Model** via a generalization of topological braids?
	- Memory is not restricted to being used by one proof system from one Universe.
		- We can have multiple proof systems over multiple Universe "running" in our Computer system, and we use Memory as a means to control it all. (It is the encoding Type coupled with some various other rules).
- CDT Space + Turing Heads
- The CDT Space (Memory) is holding the data of proofs from Universes & Turing Heads produce strands over some Universe as they traverse the Memory.
	- A Turing Machine is a model for traversing memory so as to produce proofs (strands over a Universe)
	- Is there a relationship between configuration spaces and memory?
	- We could just say that the configuration space is used as the memory
		- since every config space can be assigned a toplogy