#What *Really* Is A Computer System?

**Abstract:** We discuss modern Operating System models & various System Concepts. We try to generalize on these models and derive possible new directions for 'System Concepts'.

##An OS is Just Packaging

An "Operating System" is just the word given to any "container" that packages various "System Concepts" for how to use a machine.

###Packages

[http://en.wikipedia.org/wiki/List_of_operating_systems](http://en.wikipedia.org/wiki/List_of_operating_systems)


##Current System Concepts

**Quotations taken from David Dalrymple's article here: [http://davidad.github.io/blog/2014/03/12/the-operating-system-is-out-of-date/](http://davidad.github.io/blog/2014/03/12/the-operating-system-is-out-of-date/)**

**Thanks to David for his article which helped inspired me to pursue the ideas that I will collectively be calling: BraidOS**

###Programming Languages

Programming Languages are for Compiling/Assembling machine execution.

>"Every Language we have is a response to a set of problems at its time" -Bjarne Stroustrup

We won't even make an attempt to discuss the world of Programming Languages in its entirety. But let's discuss some of the more widely adopted programming language concepts, from FORTRAN (okay, even earlier) to present day.

Programming Languages were invented as a Systems Concept. However, most modern programming languages are a composite of various smaller (sub-)system concepts. So, a programming language is not to different from an 'OS' in that it too is a container for various system concepts.

####In the beginning: **Machine Code & Assemblers**
	
- **ISA Instruction Types**
	- Data handling and memory operations
	- Arithmetic and logic operations
	- Control flow operations 


####Fortran

**While Fortran may not be a 'System Concept' in of itself, it certainly has had enough influence to be addressed directly.**

Deserves it's own section

####Array Programming

APL - Created by Kenneth Iverson in the early 1960s

>"most programming languages are decidedly inferior to mathematical notation and are little used as tools of thought in ways that would be considered significant by, say, an applied mathematician." - Iverson


####The Rise of Structured Programming

 **Structured Program Theorem** - Any algorithm can be expressed using only three system concepts(control structures): Sequencing, Selection, Iteration.

This is a continuum of the Turing Machine concepts put into practice.

>From Wikipedia: 
>1. Executing one subprogram, and then another subprogram (sequence)
>2. Executing one of two subprograms according to the value of a boolean expression (selection)
>3. Executing a subprogram until a boolean expression is true (iteration)

The guy who sits in a chair and does work on a computational machine based on this theorem day-in and day-out is commonly referred to as a: **Software Engineer**.


####Object-Oriented Programming
####Imperative Programming
####Logic Programming
####Functional Programming


####Programming Languages - What For?

Humans don't *need* programming languages to get a computational machine to achieve an end goal.

Getting a computational machine to do what you want is achieved by applying *a set of System Concepts* to the computational machine.

As stated earlier, Programming Languages are a 'System Concept' composed of smaller system concepts.

Applications are written using a subset of the 'System Concepts' internal to a programming language (e.g. not every program will use a 'for loop').

Thus, the argument is presented that **computational applications are constructed using a set of System Concepts brought together to achieve a particular end result.** 

A programming language just happens to be a (conceptually-friendly) packaging of various smaller system concepts (similarly to modern day Operating Systems) used to achieve this. They are easily approachable by machine operators thanks to their human readable interface e.g. syntax (but actually implemented as a compiler/interpreter).

**We conclude that programming languages may not be the necessity that they are commonly believed to be.**


###The Unix(-like) Kernel

**While Unix may not be a 'System Concept' in of itself, it certainly has had enough influence to be addressed directly.**

>An operating system is a piece of software that facilitates the execution of multiple independent programs on one computer, using standard input and output routines.

####Signals

####Unix Pipes

####C Standard Library

####I/O Management 

Device Drivers

###Interactivity

>An interactive program is one which consumes input after producing output. Prior to SAGE, once a program produced its output, it was done, and the machine would halt or move on to the next job. What distinguishes an interactive system is that it will produce some output and then wait until more input is available.

###Transactions (Concurrency Primitives)

>Transactions are operations each guaranteed either to fail without any effect, or to run in a definite, strict order.

>- This one core idea enabled the development of systems called databases, which can reliably maintain the state of complex data structures across incessant read and write operations as well as some level of hardware failures.
- Modern filesystems are “journaled”, which means that they implement transactions.
- Transactions are also the key idea behind version control systems, which are increasingly adopted in all corners of the software world. In that context, they are called “commits”.
- Most recently, the core of crypto-currencies is a crude but clever solution to a distributed transaction processing problem. (In this context, transactions are in fact called transactions.)

###Automated Memory Reclamation (Garbage Collection)

>A garbage collector (GC) is a piece of software which maintains a data structure representing available memory, and marks a given memory location as available whenever it is no longer being referred to.

**Problems:** Bottleneck for performance in larger and more complex applications. Makes building real-time systems difficult (will need non-general memory reclamation instead). And there is no One-Size-Fits-All answer to handling memory as a resource within an application that will be maximally optimized on a per-application basis.

###Virtualization

>Virtualization is a general term for software facilities (possibly supported by hardware acceleration) to run programs as if they each have a computer all to themselves. 

**Problems:** *Thrashing* leading to numerous *Page Faults*.

###Hypermedia

>Hypermedia refers to any communications medium which comprises interactive systems. The most popular forms of hypermedia are those employing **hyperlinks**: certain elements of a viewed object which can be activated through interaction and whose activation triggers the display of a different object, which is determined by the hyperlink and possibly also by the interaction.

###Networking

>An internetwork is a set of communications channels between computers, where each computer is running a service that routes incoming messages to some other communications channel, so that each message eventually reaches its addressee.


##Suggested System Concepts

###Compilation Support

####Optimization Support

Optimization should always be an intermediate between a User-made something and a System-made something. Whether that be User-Code, or User-Input it will always need to be properly optimized *for* the system itself auto-magically.

###Concurrency Support

Isolation is easy, communication is hard.

###Caching Support

We define Caching as: Moving data from one memory node to another memory node (usually for execution scheduling purposes at a specific processing node).

In the case of hierarchical memory architecture Caching support would be used to move data from a lower memory tier (e.g. RAM) to a higher memory tier (e.g. L(n) Cache).

###Formal Validation & Verification Automation Support

This must be the **first thing implemented in the System**. Everything else must be built using a Formal Validation and Verification Scheme.

We will probably build a concept off of **Symmetry**.


- Formal Validator - *application* specific
	- must be constructed per application
- Formal Verifier - *toolchain* specific
	- must be constructed per toolchain


####Symmetry

- Symmetry is all about having a consistent 'framework': 
	- invariance 
	- and conservation.
- Symmetry can be used for things such as: 
	- immutability, 
	- testing, 
	- reflection, 
	- networking,
	- caching, 
	- etc.

###Scalability Support (Distributed System Support)
 
- Keeping the execution pipeline full
- Trans-Machine Nativity (making a cluster easily act as "one" machine)

###Scheduling Support

Allowing an application to arrange it's own execution scheduling.

 
###System Package (Toolchain) Options

Being able to *choose* what set of System Concepts you want to use for your 'application' or whatever.

####Operating Systems should never be 'set'
Operating Systems should change -- runtime reconfiguration should be an inherent property.
Operating Systems should  be flexible (choose which system concepts to package at what times, essentially).

##The OS is the Machine
You cannot access the actual hardware components without at least consulting it.
But the operating system is more than an extension of the machine. A good example to think about would be: 

	Imagine if we took the Linux Kernel code base and compiled it to
	Verilog. Then we compiled the Verilog code onto an FPGA as a
	processing soft-core. Our hardware would now BE an OS. 

It is really no different if you start with a RISC architecture and then control/extend it with a "System". (Sometimes we forget that software IS hardware as well).



  