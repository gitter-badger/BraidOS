#BraidOS


**Welcome to BraidOS! The world's first 'Braided System'!**


##Generalized Braids

Unlike the mathematical framework behind Braid Theory, we will be approaching the concepts derived here from a purely abstract sense and not in a geometric or topological sense.

Generalized braids will be the computational subject that we use (this concept is not found in the mathematical literature). However, being familiar with braid theory will be invaluable in assisting a newcomer to quickly assessing the details of the 'what' and 'why' a particular idea or implementation was chosen for this project.

Basic Braid Theory is not too fundamentally hard to learn. The biggest hurdle may be familiarity with mathematical notation. But if you can grasp basic concepts from set theory and generally intuitive topological concepts then you will be able to easily parse Braid-Theoretic concepts.

##Configuration Spaces

The best way to think of these guys is as State Spaces. Or to map it directly to systems concepts, its not necessarily incorrect to think of them as strictly some sort of 'Systems Configuration Space' rather than purely as a 'Mathematical Configuration Space'.

The relationship between braid groups and topological spaces is straightforward and limited. However, using the abstract notion of 'A Generalized Braid' we should be able to allow ourselves to ignore any and all topological constrictions when strictly working in the computational realm.


##Robotics

The first goal of this new OS design will be in robotics applications.

The reason for this initial goal is the based on the mathematical relationship between braid groups and configuration spaces of n-tuples of topological spaces.

In other words, braids naturally form a simple abstraction for a robots system configuration i.e. traversing a state space whose dimensional architecture is given by the D.O.F of the machine and/or its components.

In the realm of embedded systems, one of the most common things to do is control servos, motors, hydraulics, etc. that mechanically control a specific D.O.F. within the system.

The total collection of these components can be taken as a n-tuple State within the 'System Configuration Space'.



##Systems Programming

We can use BraidOS to be a System Kernel that essentially allows all control of the entire machine to have a reliable and concurrent software foundation that developers can easily map concepts too.

Instead of multi-threaded programming, concurrent data structures, non-blocking algorithms, etc. our programmer has to strictly spin up new Braids to complete a task.

The abstraction that the 'Braided System' concept brings to the table allows us to much more easily and reliably create complex highly-concurrent, state-changing tasks to the machine.

The goal of BraidOS is not performance. Our goal is strictly reliability and ease of development of complex systems such as the ones just described.

We do however consider performance to be critical (especially in embedded applications) and are thus trying to keep these powerful concepts as low-level as possible to afford their benefits to developers of all kinds.





