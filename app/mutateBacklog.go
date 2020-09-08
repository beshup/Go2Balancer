package main

// taking care of race conditions
func MutateBacklog(w *Worker, add bool) {
	w.mux.RLock()
	if add == true {
		w.backlog += 1
	} else {
		w.backlog -= 1
	}
	w.mux.Unlock()
	return
}
