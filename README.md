# consul-kv-search

Search for data in Consul K/V store and display key and value in console. Currently only value content is searched.
You can provide a search query as glob pattern (e.g. "*data*") or regular expression pattern (e.g. ".*data.*").

## Install

Download a pre-compiled binary for your operating system from here: https://github.com/pteich/consul-kv-search/releases
You need just this binary. It works on macOS (OSX/Darwin), Linux and Windows.

## Usage

````bash
consul-kv-search "*data*"
consul-kv-search -a "127.0.0.1:8500" -p "/my/start/path" -g "*data*"
consul-kv-search -a "127.0.0.1:8500" -r ".*data.*" 
````

## CLI Options

| Flag           | Default               |                | 
|----------------|-----------------------|----------------|
| `-h --help`    |                       | show help      |
| `-v --version` |                       | show version   |
| `-a --address` | 127.0.0.1:8500        | Consul address | 
| `-p --path`    | /                     | path (prefix) in KV to start search |
| `-g --glob`    | true                  | interpret query as glob pattern |
| `-r --regex`   | false                 | interpret query as regex pattern |
| `--keys`       | false                 | search in keys only |
| `--values`       | false                 | search in values only |
 
