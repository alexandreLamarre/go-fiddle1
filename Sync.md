### Go Sync package

#### When to use channels and when to use mutex

Channels:
- Passing copy of data
- Distributing units of work
- Communicating asynchronous results

Mutex:
- Caches
- State

Specifically, Mutex is used for:

- Protecting shared resources
  - sync.Mutex : Provide exclusive access to a shared resource
  - sync.RWMutex : Allows multiple readers. Writers get exclusive lock


#### sync.Atomic

- Low level atomic operations on memory
- Lockless operation
- Used for atomic operations on counters

```
atomic.AddUint64(&ops, 1)
value := atomic.LoadUint64(&ops)
```

#### Sync.Cond

- Condition variable is one of the synchronization mechanisms
- A condition variable is basically a container of goroutines that are waiting for a certain condition

How do we make a gouroutine wait until a certain condition occurs
1. Wait in a loop for the condition

```
var sharedRsc = make(map[string]string)
go func() {
    defer wg.Done()
    mu.Lock()
    for len(sharedRsc) == 0 {
        mu.Unlock()
        time.Sleep(100*time.Millisecond)
        mu.Lock()
    }
    // Do processing
    fmt.Println(sharedRsc["rsc])
    mu.Unlock()
}()
```

- We need some way to make goroutine suspend while waiting
- We need some way to signal the suspended goroutine when a particular event has occurred

Channels?
- we could use channels to block a goroutine

But what if we want multiple conditions? <span style="color:red"> sync.Cond </span>
- Conditional Variable are type
```
    var c *sync.Cond
```

- We use constructor method sync.NewCond() to create a conditional variable, it takes sync.Locker interface as input, which is ususally sync.Mutex
    m := sync.Mutex{}
    c := sync.NewCond(&m)

sync.Cond has 3 methods.
    - c.Wait()
    - c.Signal()
    - c.Broadcast()

c.Wait()
- suspends execution of the calling goroutine
- automatically unlocks c.L
- Wait cannot return unless awoken by Broadcast or Signal
- Wait locks c.L before returning
- Because c.L is not locked when Wait first resumes, the caller typically cannot assume that the condition is true when Wait returns . Instead, the caller should Wait in a loop

c.Signal()
- Signal wakes one goroutine waiting on c, if there is any
- Signal finds goroutine that has been waiting the longest and notifies that.
- It is allowed but not required for the caller to hold c.L during the call.

c.Broadcast()
- Broadcast wakes all goroutines waiting on c.
- It is allowed but not required for the caller to hold c.L during the call

#### Summary of conditional
- Conditional variable is used to synchronise execution of goroutines
- Wait suspends the execution of goroutine
- Signal wakes on goroutine waiting on c.
- Broadcast wakes all goroutines waiting on that condition
