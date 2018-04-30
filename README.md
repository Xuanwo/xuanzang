# xuanzang

xuanzang is a fulltext search backend build upon [wukong](https://github.com/huichen/wukong).

It will load content form [sitemap](https://en.wikipedia.org/wiki/Site_map), index automatically and provide a fulltext search engine by RESTful API.

## Features

- Support chinese word segmentation
- Load content from sitemap

## Quick start

Create a config file like following:

```yaml
port: 8080
db_path: /project/xuanzang/database
index_path: /project/xuanzang/index

source:
  type: sitemap
  url: https://xuanwo.org/sitemap.xml
  duration: 3600

dictionary: /project/xuanzang/dictionary.txt
stop_tokens: /project/xuanzang/stop_tokens.txt

logger:
  level: debug
```

Use xuanzang:

```bash
:) xuanzang -c /path/to/config.yaml
```

Try to do a fulltext search by curl:

```bash
:) curl 127.0.0.1:8080?text=aspire
{"tokens":["aspire"],"docs":[{"title":"友情链接 // Xuanwo's Blog","url":"https://xuanwo.org/blogroll/","content_text":""}],"total":1}
```

## Installation

Get the latest xuanzang for Linux, macOS and Windows from [releases](https://github.com/Xuanwo/xuanzang/releases)

## Configuration

xuanzang has following config options:

```yaml
# port will controls which port xuanzang will used to provide service.
port: 8080
# db_path is the path to xuanzang's database file.
db_path: /path/to/database
# index_path is the path to xuanzang's index folder.
index_path: /path/to/index

# source is where xuanzang to fetch content.
source:
  # Currently, we only support sitemap.
  type: sitemap
  # url is the URL to sitemap
  url: https://example.com/sitemap.xml
  # xuanzang will index content for every duration seconds.
  duration: 3600

# dictionary is the path to dictionary data.
# https://github.com/Xuanwo/xuanzang/data/dictionary.txt is suitable for new comer.
dictionary: /path/to/dictionary
# stop_tokens is the path to stop_tokens data.
# https://github.com/Xuanwo/xuanzang/data/stop_tokens.txt is suitable for new comer.
stop_tokens: /path/to/stop_tokens

# logger will control the logger behavior.
logger:
  # Available value: fatal, panic, error, warn, info, debug.
  level: debug
  # if output set to "stdout", xuanzang will use stdout to do output.
  output: /path/to/log
```

## Contributing

Please see [_`Contributing Guidelines`_](./CONTRIBUTING.md) of this project before submitting patches.

## LICENSE

The Apache License (Version 2.0, January 2004).
