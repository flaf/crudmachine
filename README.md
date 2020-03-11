```sh
# Run a subscriber.
go nats-sub.go -s 127.0.0.1 msg.test

# Run a publisher.
for i in XXX YYY ZZZ; do go run nats-pub.go -s 127.0.0.1 msg.test "Hello World! $i"; done
```


