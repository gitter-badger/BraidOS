##Braids - Coordination

###Intra-Coordination

This is the internal coordination procedures between the strands of the braid. These are analogous to atomic operations within the braid. These internal coordination procedures do not need to be known to any external systems the braid is running alongside. 

###Inter-Coordination

This is the coordination procedure between braids in a **multi-braided system**. It is possible that multiple braids traverse the same CDT space (as is the case with Operating Systems running multiple processes). In this situation the coordination procedures can be looked at as the internal coordination procedures of a **'Braid of Braids'**. In simple systems, it is possible that the coordination procedures are a parameter of the space itself (e.g. simple memory management - the self-partitioning of the CDT space in alignment with the number of processes).