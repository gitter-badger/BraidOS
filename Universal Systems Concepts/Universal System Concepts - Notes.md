#Universal System Concepts - Notes

- A **System Concept** is just that, a *concept*. It can be implemented via a variety of different proofs (and/or proof systems).

##Software Creation (PLs & Alternatives)

###Compilation Support

- Compilation Optimization is taking a Proof, P\_1, over some universe and compiling it another proof, P\_2 over some (other) Universe, given that:
	- m(p\_1) > m(p\_2)
	- where m(p\_x) is the Modulus of the Proof p\_x

####Optimization Support

- Clarity => Single Optimization Interface (Language ??)
- Local Optimization Support => CDT Optimization
- Global Optimization Support => Strand Optimization
- Inter-Procedural Optimization (Braid Optimization)
- Inter-Machine Optimization (Multi-Braid Optimization)

###Formal Validation and Verification Automation Support

This is about relationships between proofs. Specifically, it is about **Verification Interpretations** between two proofs, such that a proof becomes the **Verifier** of another proof under a **Verification Functor**.

###Types of Verification

- Cyclic Verification
- Dual Verification (Symmetric Verification)
- Idempotent Verification

####Symmetry

- Symmetric Functions (Actions)
- Symmetric Functions on n-tuple Configuration Spaces

###System Build Chains

- Package Manager (system)
- The OS ("kernel") is an API -- to call what your software needs

##Software Management (Computer System Management)

###Caching Support

- Data Layout/Sorting 
- Cache Oblivious Modeling as an Abstract Machine
- Unique Objectification (via Linear Type Theory / Unique Typing)
- Caching is a compilation between one Memory Type and another Memory Type (??)
	- e.g. a compilation between Virtual Memory and LX Caches, or registers (??)
- Caching should cover:
	- Networking Support
	- I/O Support


###Concurrency Support

Concurrency is a natural bi-product of Computational Braid Theory.

- Isolation & Communication support of programs within that space
- **Example:** Networking Nodes share a CDT space (e.g. IP nodes, server-client nodes, etc.)
	- The basic notion of a 'protocol' is a being a common CDT space
- Multi-Braided Programming = Parallel Programming
- Category of Concurrent Programs = Braided Monoidal Category (???)


####Transactions -- Concurrency Support (as a Universal System Concept)

We can use the theory of Braid Groups to create a **braid-theoretic interpretation of transactions**. Since (sigma\_i)(sigma\_j) = (sigma\_j)(sigma\_i)

This may have a relationship to a System Type (Coordinations) (??)

###Memory Management

- **Data Storage**
	- Data Stores
	- Query Languages (Query IO Systems in general)
	- Indexing Support
- **Main Memory**
	- Memory Allocation
		- Unified Theory of GCs
			- Re-express in terms of ETT and/or CBT (??) 


###Scaling Support

- Multiple CDT spaces (acting as one space)
	- **Trans-Nativity:** Combining & Separating these spaces (but acting as **ONE OVERALL SPACE (Machine)**)


###Scheduling Support

- **Temporal Uniqueness sharing a single-Turing-head**
	- The head "generates" strands as it traverses the space.

- We can mimic a Computational Braid (machine) *within* a single program using scheduling, as well. This is usually tackled via multi-threading (which handles "scope", etc.).

Using System Types (Coordinates and Actions) for custom preemptive scheduling support.

- Braid Group-Theoretic Schedulers
- Artin Group Schedulers
- Group Presentation Scheduler Generation
- Network Topology Semigroup Scheduling


####Variable-Length Mulitplexing -- Scheduling Support

We would like to discuss **variable-length (variable-Measure) multiplexing** over a resource (limited by a **generalized shannon limit** -- Shannon's Limit talks about the bandwidth of a channel -- thus, by a generalized Shannon Limit, we generalize "channel" to "resource"/'resource pool' -- in networking, a common channel is multiplexed over for varying signals to travel across it -- in the case of a resource, we would like to multiplex access to the resource using various Variable-Length Multiplexing techniques).

