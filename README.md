# TSATU

Book recommendations from India's best podcast show - [The Seen And The Unseen](https://seenunseen.in/) hosted by [Amit Varma](https://amitvarma.com/)

___To Be Updated Soon___

## Running

```bash
$ go run main.go | tee logs
```

If you want to parse the log data by printing the book and author information as json, try this 

```bash
$ go run main.go | tee logs
$ cat logs | head -1 | jq -s '.[] | sort_by(.count)' # This is for books
```
