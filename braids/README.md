##A Generalized Braid:
	- Strands (tasks)
	- Permutations (actions - scheduler/planner)
	- Coordinations (when two tasks 'intersect')
	- Space (Composite Data Type (CDT) Space)

###Strands (Tasks)
Strands are analagous to threads. Where a Braid is analagous to a multithreaded process as it is multi-stranded.

Each Braid (process) traverses a CDT space with multiple threads. Whether or not these threads cooperate with each other across the space is up to the internal coordination methods of the braid.

###Permutations

These are group actions on the the set of strands to determine where they wil end up at the end of the "unitary measure".

The unitary measure can be anything that is consistent across the system (clock cycle, time.amount cycle, etc.).

In the case of a processor pipeline this unitary measure would be a clock cycle. The group action in that case would be the new position of each instruction in it's overall lifecycle in each new measure (e.g. the stage in the pipeline the instruction would be at).

In the case of processors the action is purely sequential. In the arena of generic CDT space the group action could call for strands to jump from any CDT to any other CDT every unitary measure.

**The unitary measure can also be asynchronous perhaps event-driven or the like. But the measure must be consistent across the space.**

A CDT space can be partitioned (e.g. a multi-clock design). In this case each space partition must have it's own assigned measure for overseeing how strands must traverse the space. (A good way to think about it is in terms of a space metric).





###Coordinations
	
A braid requires internal coordinations as individual strands may compete over a common resource or more specifically: 
two strands may be trying to make changes to the same CDT within the same unitary measure.

The way that these strands coordinate can be specified for that braid. The specifications can also be laid out by the CDT.
It is possible that the CDT allows for only atomic operations on its internal data.


###CDT Space (all instances of CDTs)

We define a Composite Data Type to mean any general data structure (for everything from primitive singular types to database tables, files, etc.)

A CDT space is a space modularly built out of sub-componets which are CDTs. These CDTs each have internal queues for input/ouput.
In other words, a program is a movement from CDT to CDT whereupon entering and exiting a CDT can cause change to that CDT.

In the example of the processor pipeline: the CDT space would be the register file (or set of all registers). Whereupon each instruction was a strand
and that instruction would traverse the pipeline via it's atomic operations: fetch, decode, execute, write. 

The permutation would be the action known to happen pre-hand such as it appears written in assembly: ADD A, B, C (add registers A and B and store it at C)
The permutation shows that the next instruction should end up adding CDT A and CDT B and storing it at CDT C.

This involves two reads and a write operation and the manipulation of three CDTs. Thus the instruction will traverse a CDT each clock cycle.

First, the instruction register (also a CDT) will be read.
Second, the instruction will be decoded.
Third, the instruction will be executed at the ALU (reading from multiple registers CDTs).
Fourth, the instruction will write its results to a register (a CDT).

The processor pipeline has instruction level concurrency, though. And sometimes even execution parallelism, such as when the processor has multiple ALUs.


What the coordinations are working to prevent are things like pipeline stall.


###Conclusions

What should be(come) obvious is the fact that braids can actually be used to replace many of the concepts we are used to thinking about in basic computing including the design of modern day operating systems.

A braid covers persistent storage systems, process scheduling, multi-threaded programming, and even memory management (where pages are our CDTs).

One of the far reaches of this effort is to eventually implement the concepts into a fully functioning lightweight BraidOS.
