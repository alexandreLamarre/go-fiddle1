## M:N Scheduler

The Go scheduler is part of the Go runtime.

Go scheduler runs in user space

Go scheduler uses OS threads to schedule go routines
for execution

Go runtime create number of worker OS threads equal to GOMAXPROCS

GOMAXPROCS - default value is number of processors on machine

Go scheduler distributes runnable goroutines over multiple worker OS threads

At any time, N goroutines could be scheduled on M OS threads that run on at most GOMAXPROCS numbers of processors

Go 1.14+, Go scheduler implements asynchronous preemption

-> prevents long running goroutines from hogging onto CPU that could block other goroutine

### Goroutine states

- Runnable : waiting to be run
- Executing : being executed
- Preempted : if it exceeds its time slice, put back in to runnable queue
- Waiting : I/O or event wait, when I/O completes moved back to runnable state