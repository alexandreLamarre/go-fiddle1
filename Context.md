### Context Package

- We need a way to propagate request-scoped data down the call-graph.

- We need a way to propagate cancellation signal down the call-graph.

Context package servers two primary purpose:

- Provides API's for cancelling branches of call-graph
- Provides a data-bag for transporting request-scoped data through the call graph

TODO()
```
func fun() {
    ctx := context.TODO()
}
```


- TODO() returns an Empty Context
- TODOs intended purpose is to server as a placeholder

Context package can be used to send,
- Request-scoped values
- Cancellation signals

Across API boundaries to all goroutines involved in handling a request

context.Background() -> returns an empty context, it is the root of any Context tree