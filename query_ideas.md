total amount traded in a particular asset on a particular day
compute account balance in all assets as of a particular time
for a bunch of accounts (e.g. 50 accounts belonging to a party), daily summary of activity

for an account, show daily balances for a year
   - timescale with precomputed aggregate: 50ms
   - timescale just querying: 5s




-- sqlite:
2022/01/17 14:58:34 create parties took 5m28.257045319s



-- timescale 15*2000 30k inserts per second, consistent
-- vanilla   10*2000 20k inserts per second, down to ~3k/sec after ~100m rows

-- sqlite = 3k inserts per second, pretty consistent
-- badgerdb 2*1000 = 2k inserts per second. once up to ~60m, down to between 50/1000 per second with huge variance (sometimes taking a couple of minutes to sort itself out)
64gb memory used

count 200m transfers:
 - sqlite 353
 - vanilla pg 100
 - timescale 36



database thoughts

- Append only
- Delete on snapshot restore?
- Do things like, e.g. account names, asset names ever change?%


- Can we embed a postgres container image into datanode binary?
https://github.com/mudler/poco


Compression:

select hypertable_size('transfers')
155960582144


select chunks_detailed_size('transfers')
(_timescaledb_internal,_hyper_2_1_chunk,2145435648,234897408,8192,2380341248,)