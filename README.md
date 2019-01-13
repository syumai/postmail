# postmail

* Send email using `LOGIN` auth

## Install

`go get -u github.com/syumai/postmail`

## Configuration

* Set environtment variables.
* Using [direnv](https://direnv.net/) is recommended.

```sh
export POSTMAIL_SMTP_SERVER_NAME=smtp-mail.outlook.com # example
export POSTMAIL_SMTP_PORT=587                          # example
export POSTMAIL_SMTP_USER_NAME=username
export POSTMAIL_SMTP_PASSWORD=password
```

## Usage

```console
{
  echo "From: from@syumai.net"
  echo "To: to@syumai.net"
  echo "Subject: example mail"
  echo "Content-Type: text/plain; charset=UTF-8"
  echo "Content-Transfer-Encoding: 8bit"
  echo "MIME-Version: 1.0"
  echo
  echo "Hello, world" 
} | postmail -f from@syumai.net to@syumai.net
```
