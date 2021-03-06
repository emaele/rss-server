# RSS Server and Telegram notifier

Receive RSS feed elements directly in a telegram chat!

## Configuration

| **Var** 	                 | **Description**                 	                                                    |
|----------------------------|--------------------------------------------------------------------------------------|
| `RSS_SERVER_BIND_ADDRESS`  | The bind address in the form `address:port` (optional, default is localhost:26009)   |
| `TELEGRAM_TOKEN`           | This is your telegram token, grab it one from [@botfather](https://t.me/botfather)   |
| `TELEGRAM_CHAT`            | This is the telegram chat_id where the posts will be sent                            |
| `AUTHORIZATION_TOKEN`      | Authorization token for all the http calls                                           |

## Setup

### Build

Clone the repo, move into the folder and type

```Bash
go build
```

### Run it

Execute the binary
```Bash
./rss-server-notifier
```

## How to use it

### Authorization

Set the `Authorization` header with the previous set token.

### Add RSS feed

Adding a feed it's pretty easy, just call the `/add` endpoint with a JSON struct like that:

```JSON
{
    "URL":"https://feed_url",
    "Filter": "([0-9].*|(regex))" 
}
```
