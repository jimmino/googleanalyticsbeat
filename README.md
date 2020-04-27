# GoogleAnalytics Beat

Welcome to GoogleAnalytics Beat.

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/jimmino/googleanalyticsbeat`

## Getting Started with GoogleAnalytics Beat

This Beat was forked from https://github.com/GeneralElectric/GABeat and updated accordingly

### Requirements

* [Golang](https://golang.org/dl/) 1.13

### Init Project
To get running with GoogleAnalytics Beat 

```
- install Go
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push GoogleAnalytics Beat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/jimmino/googleanalyticsbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for GoogleAnalytics Beat run the command below. This will generate a binary
in the same directory with the name googleanalyticsbeat.

```
go build
```


### Run

To run GoogleAnalytics Beat with debugging output enabled, run:

```
./googleanalyticsbeat -c googleanalyticsbeat.yml -e -d "*"
```


### Clone

To clone GoogleAnalytics Beat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/jimmino/googleanalyticsbeat
git clone https://github.com/jimmino/googleanalyticsbeat ${GOPATH}/src/github.com/jimmino/googleanalyticsbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

