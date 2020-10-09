# consul-kv-search

Search for data in Consul K/V store and display key and value in console. You can search in keys, values or both (default).
You can provide a search query as glob pattern (e.g. "*data*" this is the default) or regular expression pattern (e.g. ".*data.*").

## Install

Download a pre-compiled binary for your operating system from here: https://github.com/pteich/consul-kv-search/releases
You need just this binary. It works on macOS (OSX/Darwin), Linux and Windows.

## Usage

````bash
consul-kv-search "*data*"
consul-kv-search -a "127.0.0.1:8500" -p "/my/start/path" "*data*"
consul-kv-search -a "127.0.0.1:8500" -r ".*data.*" 
````

## CLI Options

| Flag           | Default               |                | 
|----------------|-----------------------|----------------|
| `-h --help`    |                       | show help      |
| `-v --version` |                       | show version   |
| `-a --address` | 127.0.0.1:8500        | Consul address | 
| `-t --token`   |                       | Consul access token | 
| `-w --wrap`    |  false                | Wrap text output | 
| `-p --path`    | /                     | path (prefix) in KV to start search |
| `-r --regex`   | false                 | interpret query as regex pattern |
| `--keys`       | false                 | search in keys only |
| `--values`     | false                 | search in values only |
 
