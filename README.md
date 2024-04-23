# microq
A tiny event broker.

## Features
- [x] Websocket support
- [x] TCP support
- [x] Grouped subscriptions
- [ ] Offline messages
- [ ] Clustering
- [x] QoS (0, 1, 2)
- [x] Topics & Patterns

### Subscriptions
Topics and patterns are heavily inspired by [NATS's subject-based messaging](https://docs.nats.io/nats-concepts/subjects).

`Unsub` and `Sub` actions set on clients are applied in that order.

```markdown
# Sample algorithm only
SUB us.*
UNSUB us.payments
PUB us.accounts -> this will still be received by client
PUB us.payments -> this will not be received
SUB us.payments
PUB us.payments -> this will now be received
```
## Why does this exist?
Because.