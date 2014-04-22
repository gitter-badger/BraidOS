##A Generalized Braid:
	- Strands (tasks)
	- Permutations (actions - scheduler/planner)
	- Coordinations (when two tasks 'intersect')
	- Space (Composite Data Type (CDT) Space)

	

###Strands (Tasks)
Strands are analagous to threads. Where a Braid is analagous to a multithreaded process as it is multi-stranded.

Each Braid (process) traverses a CDT space with multiple threads. Whether or not these threads cooperate with each other across the space is up to the internal coordination methods of the braid.


###Permutations

These are group actions on the the set of strands to determine where they will end up at the end of the "unitary measure".

The unitary measure can be anything that is consistent across the system (clock cycle, time.amount cycle, etc.).

The unitary measure can also be asynchronous perhaps event-driven or the like. But the measure must be consistent across the space.

A CDT space can be partitioned for modular design reasons. In this case, each space partition will have it's own assigned measure for overseeing how strands must traverse that partition.


###Coordination
	
A braid requires internal coordination as individual strands may compete over a common resource or more specifically: two strands may be trying to make changes to the same CDT within the same unitary measure.

The way that these strands coordinate can be specified for that braid. The specifications can also be laid out by the CDT.

It is possible that the CDT allows for only atomic operations on its internal data.


###CDT Space (all instances of CDTs)

We define a Composite Data Type to mean any general data structure (for everything from primitive singular types to database tables, files, etc.)

A CDT Space is a space modularly built out of sub-componets which are CDTs. These CDTs each have internal queues for input/ouput.

In other words, a program is a movement from CDT to CDT whereupon entering and exiting a CDT can cause change to that CDT.


###Conclusions

What should be(come) obvious is the fact that braids can actually be used to replace many of the concepts we are used to thinking about in basic computing including the design of modern day operating systems.

A braid covers persistent storage systems, process scheduling, multi-threaded programming, and even memory management (where pages are our CDTs).

Thus we hope to implement these concepts into the fully functioning and lightweight: BraidOS.
