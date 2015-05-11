##C10M Scalability

####Hacker News Comment 1

	"> The kernel is the problem
	I had an epiphany one day when I realized that the kernel is nothing but a library with an expensive 
	calling convention.
	The only reason we bother calling the kernel at all is because it has privileges that userspace 
	programs don't have, and it uses those privileges to control/mediate access to shared resources.
	The downside of the kernel is that there's no way around it. You can't "opt out" of any of its decisions. 
	With any normal library, if had an O(n^2) algorithm somewhere, or wasn't optimized for your use case, 
	or just generally got in the way, you could choose not to link against it. User-space libraries are 
	democratic in this respect; you vote with your linker line. But with the kernel, 
	it's "my way or the highway." The kernel is the only way to the hardware.
	Here's an unfortunate example: O_DIRECT is one of those few ways that you can sidestep the kernel. 
	With O_DIRECT you can bypass the kernel's page cache, which there are very good reasons for wanting to do. 
	But Linus's opinion is that "I should have fought back harder": https://lkml.org/lkml/2007/1/10/233 
	He thinks it's unfortunate that you can sidestep the kernel, because 
	"You need a buffer whatever IO you do, and it might as well be the page cache."
	But what if you want better memory accounting, or better isolation between users, or a different data 
	structure, or any number of other tweaks to the page cache implementation? Well thankfully we have 
	O_DIRECT. Otherwise, your only choice would be to try to convince the Linux kernel maintainers to 
	integrate your change, after you've tweaked it so that it's suitable for absolutely everyone else that 
	uses Linux, and given it an interface that Linux is willing to support forever. Good luck with that.
	The kernel has always been the problem. User-space is where it's at."

###Hacker News Comment 2
	
	""The kernel is just a library" isn't exactly the same sentiment as "the kernel should be as small 
	as possible" -- I believed the latter before I fully understood the former. "The kernel is just a 
	library" means that all of the experience we have designing and factoring userspace APIs carries 
	over into kernel design. Furthermore it means that the kernel is a strictly less flexible library 
	than userspace libraries, with a strictly more expensive calling convention, and that its only 
	advantage is that it can protect and mediate access to hardware."

###Hacker News Comment 3
	
	"It has another advantage, which is that it acts as a trusted third party that can allow two 
	mutually untrusting users/program to share a software-implemented resource - for example, a disk 
	cache or a TCP stack."


###Mulit-Core Scalability

- Keep data structures per core. Then on aggregation read all the counters.
- Atomics. Instructions supported by the CPU that can called from C. Guaranteed to be atomic, never conflict. Expensive, so don’t want to use for everything.
- Lock-free data structures. Accessible by threads that never stop and wait for each other. Don’t do it yourself. It’s very complex to work across different architectures.
- Threading model. Pipelined vs worker thread model. It’s not just synchronization that’s the problem, but how your threads are architected.
- Processor affinity. Tell the OS to use the first two cores. Then set where your threads run on which cores. You can also do the same thing with interrupts. So you own these CPUs and Linux doesn’t.


###Memory Scalability

**Co-locate Data**

- Don’t scribble data all over memory via pointers. Each time you follow a pointer it will be a cache miss: [hash pointer] -> [Task Control Block] -> [Socket] -> [App]. That’s four cache misses.

- Keep all the data together in one chunk of memory: [TCB | Socket | App]. Prereserve memory by preallocating all the blocks. This reduces cache misses from 4 to 1.

**Paging**

- The paging table for 32gigs require 64MB of paging tables which doesn’t fit in cache. So you have two caches misses, one for the paging table and one for what it points to. This is detail we can’t ignore for scalable software.

- **Solutions**: compress data; use cache efficient structures instead of binary search tree that has a lot of memory accesses 

- NUMA architectures double the main memory access time. Memory may not be on a local socket but is on another socket.

**Memory pools**

- Preallocate all memory all at once on startup.

- Allocate on a per object, per thread, and per socket basis.

**Hyper-threading**

- Network processors can run up to 4 threads per processor, Intel only has 2.

- This masks the latency, for example, from memory accesses because when one thread waits the other goes at full speed.

**Hugepages**

- Reduces page table size. Reserve memory from the start and then your application manages the memory.


###Summary

**NIC**

- **Problem**: going through the kernel doesn’t work well.

- ***Solution***: take the adapter away from the OS by using your own driver and manage them yourself

**CPU**

- **Problem**: if you use traditional kernel methods to coordinate your application it doesn’t work well.

- ***Solution***: Give Linux the first two CPUs and you application manages the remaining CPUs. No interrupts will happen on those CPUs that you don’t allow.

**Memory**

- **Problem**: Takes special care to make work well.

- ***Solution***: At system startup allocate most of the memory in hugepages that you manage.


###Article Comment

	"This breakdown of bottlenecks is impressive, it is refreshing to read something this correct.
	
	Now I want to play devils' advocate (mostly because I thoroughly agree w/ this guy). The solutions 
	proposed sound like customized hardware specific solutions, sound like a move back to the old days, 
	when you could not just put some fairly random hardware together, slap linux on top and go ... 
	that will be the biggest backlash to this, people fear appliance/vendor/driver lock-in, and the fear 
	is a rational one.
	
	What are the plans to make these very correct architectural practices available to the layman. 
	Some sort of API is needed, so individual hardware-stacks can code to it and this API must not be a 
	heavy-weight abstraction. This is a tough challenge.
	
	Best of luck to the C10M movement, it is brilliant, and I would be a happy programmer if I can slap 
	together a system that does C10M sometime in the next few years"

###Conclusions

