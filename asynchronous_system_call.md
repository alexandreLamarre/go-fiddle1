### Example

- File descriptor is set to non-blocking mode
- If file descriptor is not ready for I/O operation, system call does not block, but returns an error
- Asynchronous IO increases the application complexity
- Setup event loops using callback functions

### How does Go handle this? netpoller

- uses Netpoller to convert asynchronous system call to blocking system call

- When a goroutine makes a asynchronous systeml call and file descriptor is not ready, goroutine is parked at netpoller OS thread

- Netpoller uses interface provided by OS to do polling on file descriptors

- Netpoller gets notification from OS, when file descriptor is ready for I/O operation

- Netpoller notifies goroutine to retry I/O operation.

- Complexity of managing asynchronous system call is move from Application to Go runtime, which manages it efficiently

### Summary 

- Go uses netpoller to handle asynchronous system call

- netpoller uses interface provided by OS to do polling on file descriptors and notifies the goroutine to try I/O operation when it is ready