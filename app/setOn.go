package main

// taking care of race conditions
func SetOn(w *Worker, isOn bool) {
	w.mux.RLock()
	w.on = isOn
	w.mux.Unlock()
	return
}
