package fwk

import (
	"runtime"
	"sync"
)

func ComputeParallel[I any, O any](input []I, compute func(arg0 I) O) chan O {
	var wg sync.WaitGroup
	in := make(chan I)
	out := make(chan O, len(input))

	for i := 0; i < runtime.NumCPU(); i++ {
		go func(id int) {
			wg.Add(1)

			for m := range in {
				out <- compute(m)
			}

			wg.Done()
		}(i)
	}

	for _, i := range input {
		in <- i
	}
	close(in)

	wg.Wait()
	close(out)

	return out
}
