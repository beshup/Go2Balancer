package main

// taking care of race conditions
func AccessBacklog(w *Worker) (num int) {
	w.mux.RLock()
	num = w.backlog
	w.mux.Unlock()
	return
}
