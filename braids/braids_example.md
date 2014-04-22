##Processor Pipeline Example:

###Braid
1. Strands
1. Permutations/Actions
1. Coordination
1. CDT Space

###Processor Pipeline
1. Instructions
1. Stages (fetch, decode, execute, write)
1. Hazard Avoidance (data, structural, control)
1. Set of all registers (register file)



The depth of the pipeline is equivalent to the number of strands in the braid. Thus, showing that the number of strands in a braid is directly linked to the concurrency of the system.


**Why such a low-level example?**

Simply, to show that braids work very well as a constructive abstraction throughout computational systems.




###Explanation

In the case of a processor pipeline the **unitary measure** would be a **clock cycle**. 

The group action would be the movement of each instruction to the next stage in the pipeline.

In the case of processors the action is purely sequential. In the arena of generic CDT space the group action could call for strands to jump from any CDT to any other CDT.







