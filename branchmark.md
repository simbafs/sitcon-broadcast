# Branchmark

Branchmark with [wrk v4.1.0](https://github.com/wg/wrk/tree/4.1.0). These APIs can be called via polling, so minimizing latency is important.

| api endpoint             |       latency       | latency max |    Req/Sec     | Req/Sec max |
| :----------------------- | :-----------------: | :---------: | :------------: | :---------: |
| `/api/now`               | 225.55us ± 522.97us |   14.92ms   | 34.34k ± 2.80k |   40.34k    |
| `/api/sessin/R0`         | 745.39us ± 401.08us |   6.35ms    | 6.90k ± 517.25 |    7.71k    |
| `/api/session/R0/all`    | 78.77ms ± 219.70ms  |    1.18s    | 2.86k ± 580.93 |    3.48k    |
| `/api/session/R0/2d8a5e` | 672.16us ± 531.32us |   14.76ms   | 7.90k ± 651.17 |    9.50k    |
