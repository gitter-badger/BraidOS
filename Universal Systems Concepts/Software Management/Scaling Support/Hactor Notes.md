##Hactor Notes


if you'll be creating a kernel of some sort in 2014, at least make it **trans-machine**, as in it can transcend one machine, so you **run one kernel on a cluster**.

I do. You see, it's easy to communicate inside a process (use the language native methods: functions, variables, structs). It becomes harder interprocess. And then it becomes a whole another layer intermachine. That HAS to go in a modern kernel. **Intermachine should be the same as intraprocess.**


this is why I have this data format. To remove exactly these boundaries

however without slowing things down (this is the trick)

one example. **Can I use your local filesystem API to work on a file on another server without any weird gotchas and limitations?**

if not, I think it should be designed so you can

the next problem is that you can **no longer have resource access be identity based** like all other Operating Systems right now, because **identity becomes fuzzy between machines.**

that's where **Object Capability Security** comes. It's a model, you can look it up.

[http://en.wikipedia.org/wiki/Object-capability_model](http://en.wikipedia.org/wiki/Object-capability_model)

this way any machine can provide full direct access to the resources of another, yet do it securely

it's quite powerful