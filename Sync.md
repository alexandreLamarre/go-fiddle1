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