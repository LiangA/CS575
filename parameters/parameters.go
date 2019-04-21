package parameters

type Clock struct {
	Now int // the clock now, start at 0
	Max int // the maximum of the clock, which comes from sum of burst time of all process
}

type Process struct {
	Pid      string //process id
	Priority int    // process priority, the larger the more important
	Arrive   int    // process arrive time
	Burst    int    // process burst time
	Remain   int    // process remaining burst time
	State    string // process have three possible states: ready, run, finished
}

// these could be a design, but in this case may not suit(after consideration) so give it up
// type Ready struct {
// 	Queue []Process
// }

// type Running struct {
// 	Queue []Process
// }
