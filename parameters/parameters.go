package parameters

type Clock struct {
	Now int // the clock now, start at 0
	Max int // the maximum of the clock, which comes from sum of burst time of all process
}

type Process struct {
	Pid      string `json:"pid"`      //process id
	Priority int    `json:"priority"` // process priority, the larger the more important
	Arrive   int    `json:"arrive"`   // process arrive time
	Burst    int    `json:"burst"`    // process burst time
	Remain   int    `json:"remain"`   // process remaining burst time
	State    string `json:"state"`    // process have three possible states: ready, run, finished
}

// these could be a design, but in this case may not suit(after consideration) so give it up
// type Ready struct {
// 	Queue []Process
// }

// type Running struct {
// 	Queue []Process
// }
