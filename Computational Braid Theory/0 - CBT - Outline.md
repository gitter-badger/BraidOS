#Computational Braid Theory -- Outline

>It sometimes happens that by ignoring a certain amount of data or structure one obtains a simpler, more flexible theory which, almost paradoxically, can give results not readily obtainable in the original setting.
- **Allen Hatcher** (*Algebraic Topology*)

##Foundations

###Turing Theory

 - Turing Machines (general description)
	 - Alternative descriptions
 - Turing Completeness
 - Universal Turing Machines
 - Computable Functions

###Topological Braid Theory

 - Algebraic and Geometric views on mathematical braids
	 - Geometric Braids
	 - Polygonal Braids
	 - Braid Groups (and their algebraic properties)
		 - Pure Braid Groups
		 - Relationship to Configuration Spaces
 - Some homological views on braids
	 - 1 -> P\_n -> B\_n -> S\_n -> 1

##Computational Braids

- Generalization of a Turing Machine using 'braid group actions over a config space'
- Sometimes spaces are not necessarily topological or have a metric
	- Thus, we **generalize "Braid Groups over Config Spaces" with "Proof Sheafs Over Universes"**

##Systems Theory

- **Computer System**
	- Memory + Memory Traversal Rules is a Universe U\_memory (what the Turing Machine actually models)
	- Encoding a proof sheaf over U\_i -> U\_memory is a compilation functor. The Universe (U\_i) that the proof sheaf is originally over is called the "Language" of the sheaf i.e. the programming language used to specify the qualifier and derivations for the sheaf.
		- The language used to specify a sheaf, type, etc. can only use terminology available in some Universe.
- Computational Braids as **Proof Systems over Universes**
	- Thus, in Computer Systems we deal with Computational Braids as Proof Systems over the Universe U\_memory.
	- The number of strands in the braid is equivalent to the number of proofs in the proof system. 
	- A proof over the Universe is analogous to a 'process' (or thread) in current OS theory. (??)