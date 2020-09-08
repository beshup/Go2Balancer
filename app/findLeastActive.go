package main

func FindLeastActive(library *Library) (least *Worker) {
	for index, worker := range library.workers {
		if index == 0 {
			least = worker
		} else {
			// avoiding race conditions here!
			if AccessBacklog(worker) < AccessBacklog(least) {
				least = worker
			}
		}
	}
	return
}
