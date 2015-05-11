#Compilation

- In E.T.T. Function Types = Compilation, and vice versa.
	- F: U\_1 -> U\_2 (Universe Compilation)
	- F: T\_1 -> T\_2 (Type Conversion)
	- F: T\_1 -> T\_1 (Type Operations)

- For example, if U\_1 is the Universe of some proof representation, and U\_2 is the Universe of some CPU architecture (ISA), then a function type F: U\_1 -> U\_2 maps the proof representation (program) to the CPU arch. --- this is exactly what a compiler does for a programming syntax.
- Lambda Calculi are only concerned with function types, and this may confusingly seem perfect for all of Computational theory as it can be used to help describe a PL compilers job quite easily. However, programming languages are not the entirety of computational theory. In fact, compilation itself can play a larger role in practical computation than languages (this may seem contradictory at first, until you realize that PL compilation is only one kind of compilation).

##Just-In-Time Compilation

##Optimization Support

- If every program is already maximally optimal, what are we really doing when we "optimize" source code compilation?