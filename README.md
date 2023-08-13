# check_aws_ses

## Overview

check_aws_ses is a nagios plugins to check reputation status or send quota usage of your Amazon SES account.

## Prerequisites

IAM Policy `sesv2:GetAccount` action is required to perform this command.

## Usage

```
check_aws_ses is a nagios plugins to check reputation status or send quota usage of your Amazon SES account.

Usage:
  check_aws_ses [command]

Available Commands:
  help        Help about any command
  reputation  Check the reputation status of your Amazon SES account
  sendquota   Check the send quota usage of your Amazon SES account in the current region
  version     Print the version information

Flags:
  -R, --region string      AWS region (required)
  -A, --accesskey string   AWS ACCESS KEY ID (required if secretkey is set)
  -S, --secretkey string   AWS SECRET ACCESS KEY (required if accesskey is set)
  -h, --help               help for check_aws_ses
  -v, --version            version for check_aws_ses

Use "check_aws_ses [command] --help" for more information about a command.
```

### reputation

```
Check the reputation status of your Amazon SES account.

The reputation status can be one of the following:
  - HEALTHY (check status is OK)
    There are no reputation-related issues that currently impact your account.
  - PROBATION (check status is WARNING)
    We've identified potential issues with your Amazon SES account.
    We're placing your account under review while you work on correcting these issues.
  - SHUTDOWN (check status is CRITICAL)
    Your account's ability to send email is currently paused because
    of an issue with the email sent from your account. When you correct the issue,
    you can contact us and request that your account's ability to send email is resumed.

Usage:
  check_aws_ses reputation [flags]

Flags:
  -h, --help   help for reputation

Global Flags:
  -A, --accesskey string   AWS ACCESS KEY ID (required if secretkey is set)
  -R, --region string      AWS region (required)
  -S, --secretkey string   AWS SECRET ACCESS KEY (required if accesskey is set)
```

### sendquota

```
Check the send quota usage of your Amazon SES account in the current region.

send quota is an object that contains information about the per-day and per-second sending
limits for your Amazon SES account in the current AWS Region.

- SentLast24Hours
  The number of emails sent from your Amazon SES account in the current region over the past 24 hours.
- Max24HourSend
  The maximum number of emails that you can send in the current region over a 24-hour period.
  A value of "-1" signifies an unlimited quota.
- SendQuotaUsage (%)
  SentLast24Hours / Max24HourSend * 100

Usage:
  check_aws_ses sendquota [flags]

Flags:
  -w, --warning float    send quota usage result in warning status (%)
  -c, --critical float   send quota usage result in critical status (%)
  -h, --help             help for sendquota

Global Flags:
  -A, --accesskey string   AWS ACCESS KEY ID (required if secretkey is set)
  -R, --region string      AWS region (required)
  -S, --secretkey string   AWS SECRET ACCESS KEY (required if accesskey is set)
```

## Examples

### check reputation status

- using IAM Role or environment variables (`AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`)

```sh
$ check_aws_ses reputation -R ap-northeast-1
 OK: reputation status is HEALTHY
```

- with accesskey/secretkey flags

```sh
$ check_aws_ses reputation -R ap-northeast-1 -A someaccesskey -S somesecretkey
 OK: reputation status is HEALTHY
```

### check send quota usage

- using IAM Role or environment variables (`AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`)

```sh
$ check_aws_ses sendquota -R ap-northeast-1 -w 10 -c 20
 OK: send quota usage is 0.4% (SentLast24Hours: 200, Max24HourSend: 50000)
```

- with accesskey/secretkey flags

```sh
$ check_aws_ses sendquota -R ap-northeast-1 -A someaccesskey -S somesecretkey -w 10 -c 20
 OK: send quota usage is 0.4% (SentLast24Hours: 200, Max24HourSend: 50000)
```
