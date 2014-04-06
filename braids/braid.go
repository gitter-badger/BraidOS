package braid

import (
	
	"fmt"
	"time"
	"os"
	
)

//  I wish that go had generics --> may be reverting over to C++ for this concept simulation

type braid struct {

	tasks int;

}

type Group struct {

	// define a group and group action
}

func permutation (Group G) {

	braid B;

	// Action of G on the set X of strands (tasks) in braid B => group homomorphism from G -> SYM(X)

	for i := 1; i < B.tasks; i++ {

	}
}

type task struct {
	id int;
	counter byte;
	r_set //CDT;
	stack //CDT;
}

type intersection struct {
		task t1;
		task t2;
}

func coordinations (intersection i) {

	// t1 will coordinate with t2 according to this function
	// whenever the two are addressing the same resource at the same time
}

type space struct {
	// can be graph or can be dimensional, etc.
	//node n1;
	// or ...
	//int x, y, z;

}

