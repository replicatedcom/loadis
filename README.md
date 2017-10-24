## loadis - the ledis loader

## Runnins

```bash
make docker shell
make build
./bin/loadis <ledis dump filename>
```

## Commands

Supported commands are:

 * `keys` -- list all keys in the database
 * `hgetall <hkey>` -- list all keys and values in the hash
 * `smembers <skey>` -- list all members in the set
 * `get <key>` -- get the value of the string key
