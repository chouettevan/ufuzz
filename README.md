# Ufuzz
---
Ufuzz is a highly flexible web fuzzer meant for complex cases where other alternatives (like gobuster) might not be sufficient.
---
The tool does not provide filtering ,saving or data formatting features,as these can be acheived using pipes and the GNU Coreutils.

## Installation
clone the repository,
install any dependencies with

```
go get
```
then build the project with

```
go build
```

you should find an executable named `ufuzz`
## Usage

in order to fuzz,you must first write a config file containing an HTTP request template,in the following form:

```
GET /S1 HTTP/1.1
Host:localhost
Connection:keep-alive

```
you then run the follwing command:
```
ufuzz --host localhost --port 80/443 --config /path/to/config -w /path/to/wordlists
```
ufuzz will replace the S1 placeholder with input from the wordlist


