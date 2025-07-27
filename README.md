# `netchat` TCP Connection Checker

## Install

Build the binary:

```bash
go build -o netchat
```

Install the binary globally:

```bash
go install github.com/yourusername/netchat@latest
```

## Usage

Start a server and a client to test the connection with a message:

```bash
netchat server -p 1234
```

```bash
netchat client -a localhost -p 1234 "Hello, World!"
```

Check what protocol is used by the server:

```bash
netchat proto -a localhost -p 1234
```
