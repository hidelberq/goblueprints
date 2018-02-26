# twittervotes

## Run

```
$ nsqlookupd
[nsqlookupd] 2018/02/27 01:49:36.928872 nsqlookupd v1.0.0-compat (built w/go1.9.1)
[nsqlookupd] 2018/02/27 01:49:36.929268 TCP: listening on [::]:4160
[nsqlookupd] 2018/02/27 01:49:36.930345 HTTP: listening on [::]:4161

$ nsqd --lookupd-tcp-address=localhost:4160
[nsqd] 2018/02/27 01:53:08.624348 TCP: listening on [::]:4150
[nsqd] 2018/02/27 01:53:08.624403 HTTP: listening on [::]:4151

$ nsq_tail --topic="votes" --lookupd-http-address=localhost:4161
2018/02/27 01:54:17 INF    1 [votes/tail448510#ephemeral] querying nsqlookupd http://localhost:4161/lookup?topic=votes
2018/02/27 01:54:17 INF    1 [votes/tail448510#ephemeral] (hidelberq:4150) connecting to nsqd
```
