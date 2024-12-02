TSS
---


0. build tss executable binary
```
git clone https://github.com/Charles-Leo-G/tss-lib
cd tss
go build
```

1. init 3 parties
```
./tss init --home ~/.test1 --vault_name "default" --moniker "test1" --password "123456789"
./tss init --home ~/.test2 --vault_name "default" --moniker "test2" --password "123456789"
./tss init --home ~/.test3 --vault_name "default" --moniker "test3" --password "123456789"
```

2. generate channel id
replace value of "--channel_id" for following commands with generated one
```
./tss channel --channel_expire 30
```

3. keygen 
```
./tss keygen --home ~/.test1 --vault_name "default" --parties 3 --threshold 1 --password "123456789" --channel_password "123456789" --channel_id "802671B1B19"
./tss keygen --home ~/.test2 --vault_name "default" --parties 3 --threshold 1 --password "123456789" --channel_password "123456789" --channel_id "802671B1B19"
./tss keygen --home ~/.test3 --vault_name "default" --parties 3 --threshold 1 --password "123456789" --channel_password "123456789" --channel_id "802671B1B19"
```

4. sign
```
./tss sign --home ~/.test1 --vault_name "default" --password "123456789" --channel_password "123456789" --channel_id "802671B1B19"
./tss sign --home ~/.test2 --vault_name "default" --password "123456789" --channel_password "123456789" --channel_id "802671B1B19"
```

5. regroup - replace existing 3 parties with 3 brand new parties
```
# start 2 old parties (answer Y for isOld and IsNew interactive questions)
./tss regroup --home ~/.test1 --vault_name "default" --password "123456789" --new_parties 3 --new_threshold 1 --channel_password "123456789" --channel_id "802671B1B19"
./tss regroup --home ~/.test2 --vault_name "default" --password "123456789" --new_parties 3 --new_threshold 1 --channel_password "123456789" --channel_id "802671B1B19"
# start the new parties (answer n for isIold and Y for IsNew interactive questions)
./tss regroup --home ~/.test3 --vault_name "default" --password "123456789" --new_parties 3 --new_threshold 1 --channel_password "123456789" --channel_id "802671B1B19"
```
