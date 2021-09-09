### Ownership - Channels

- Owner of channel is a goroutine that instantiates, writes and closes a channel

- Channel utilizers only have a read-only view into the channel


Ownership of channels avoids:
- Deadlocking by writing to a nil channel
- Closing a nil channel, which will cause a panic
- Writing to a closed channel, which will cause a panic
- Closing a channel more than once

