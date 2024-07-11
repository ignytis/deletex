# deletex : Twitter history eraser

## Requirements

This application uses OAuth 1 for user authentication. To obtain a consumer key and a consumer key secret,
the application has to be registered at Twitter Developer Platform https://developer.twitter.com - see _Projects & apps_ -> _(your app)_ -> _keys and tokens_ -> _consumer key_

__Note:__ X had introduced paid access level, BUT some features, like deletion of posts by list, still work (as of 2024.07.10)

## Usage

### Help

```
deletex h
deletex -h
deletex help
deletex --help
```

__Common parameters__

- -c / --config - path to configuration file. See [config.example.yaml](config.example.yaml). Default value is `config.yaml`, so configuration file will be read from current directory.

### A typical workflow
- Register a client Twitter app at https://developer.twitter.com, obtain consumer API keys, save them to [config.yaml](config.example.yaml). Note: it's NOT needed if you would like just to generate a list of tweets to delete without actual deletion; Twitter credentials are needed for API calls.
- Generate and download a twitter archive: [Link](https://help.twitter.com/en/managing-your-account/how-to-download-your-twitter-archive)
- Convert the `data/tweet.json` file from arctive to JSON Lines format
- Generate a list of tweets to delete from JSON Lines file by filtering tweets
- Delete tweets using a list file

### Conversion of Javascript tweet dump file into JSON Lines

Twitter provides dumps as JavaScript files which are inapproptiate for analysis and filtering of records. This command converts the original JavaScript file into [JSON Lines format](https://jsonlines.org/).

```bash
deletex tweets:dump:to_jsonl \
    -i /home/john/somefolder/twitter_dump/data/tweet.js \
    -o /home/john/somefolder/twitter_dump_processed/tweets.js
```

### Producing a list of tweets to delete by filtering the JSON Lines file

```bash
deletex tweets:to_delete_list:from_jsonl \
    -i /home/john/somefolder/twitter_dump/tweets.json \
    -o /home/john/somefolder/twitter_dump/todo_delete.txt \
    -e "created_time >= '2022-01-01 00:00:00' && created_time < '2022-02-01 00:00:00'"
```

### Deletion of tweets using CSV file

__Parameters__

- -i / --input-file - path to CSV file where IDs of tweets to delete are listed. Each line contains a single column which is tweet ID. No CSV header. _Note: it was probably a bad idea to mention the CSV format in this command, because records are not actually comma-separated :D Probably this name will be changed later._

```bash
deletex tweets:delete:using_csv \
    -c $PWD/.local/config.yaml \
    -i /home/john/somefolder/tweets_to_delete.csv
```

## Filtering expressions

The [govaluate](https://github.com/Knetic/govaluate) library is used to parse expressions. Following tweet properties are defined for each record:
- created_time
- favorite_count
- full_text
- id
- retweet_count

### Examples of expressions

```
favorite_count + (retweet_count * 20) > 100
created_time >= '2022-01-01 00:00:00' && created_time < '2022-02-01 00:00:00'
```