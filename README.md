licensechecker (lc)
-------------------
`lc` is a command line tool that recursively iterates over a supplied directory or file 
attempting to identify what software license each file is under using the list
of licenses supplied by the SPDX (Software Package Data Exchange) Project. It will pick up 
license files named appropiately or inline licenses such as the below in source files

`SPDX-License-Identifier: GPL-3.0-only`

In a nutshell this project is a reimplementation of http://www.boyter.org/2017/05/identify-software-licenses-python-vector-space-search-ngram-keywords/ using Go while I attempt to nut out the nuances of the language. 

It can produce report outputs as valid [SPDX](https://spdx.org/), CSV, JSON and CLI formatted. It has been designed to work inside CI systems that capture either stdout or file artifacts.

[![Build Status](https://travis-ci.org/boyter/lc.svg?branch=master)](https://travis-ci.org/boyter/lc)

Dual-licensed under MIT or the [UNLICENSE](http://unlicense.org).

### Why

In short taken from, http://ben.balter.com/licensee/

 * You've got an open source project. How do you know what you can and can't do with the software?
 * You've got a bunch of open source projects, how do you know what their licenses are?
 * You've got a project with a license file, but which license is it? Has it been modified?

Why should you care about what licenses your code runs under? See 

 * http://www.openlogic.com/resources/enterprise-blog/archive/use-spdx-for-open-source-license-compliance 
 * https://thenewstack.io/spdx-open-source-cheap-compliance-license-can-expensive/
 * https://www.infoworld.com/article/2839560/open-source-software/sticking-a-license-on-everything.html

### Installation

The binary name for `licencechecker` is `lc`.

For binary files see releases https://github.com/boyter/lc/releases To build from source you need to have Go setup with your GOPATH working and your go binary path exported like so,

```
export PATH=$PATH:$(go env GOPATH)/bin
```

then to install

```
$ go install
```


### Usage

Command line usage of `licensechecker` is designed to be as simple as possible.
Full details can be found in `lc --help`.

More information about [what licensechecker looks at and how it works](what-we-look-at.md)

Probably the most useful functionality is the `-f` modifier which specifies the output format.
By default `licencechecker` will print out results as it processes files. However as it was designed
to run at the end of CI tasks you may want to get a nicer output which can be done like so.

```
$ lc -f tabular .
```

The above will process starting in the current directory and print out a formatted list of results when finished.

To view all command line options

```
$ lc --help
```

Example output of `licencechecker` running against itself in tabular format while ignoring the .git, licenses and vendor directories

```
$ lc -pbl .git,vendor,licenses -f tabular .
Directory            File                    License                            Confidence  Size
.                    .gitignore              (MIT OR Unlicense)                 100.00%     275B
.                    .travis.yml             (MIT OR Unlicense)                 100.00%     74B
.                    CONTRIBUTING.md         (MIT OR Unlicense)                 100.00%     1.2K
.                    Gopkg.lock              (MIT OR Unlicense)                 100.00%     1.4K
.                    Gopkg.toml              (MIT OR Unlicense)                 100.00%     972B
.                    LICENSE                 Unlicense AND MIT                  94.83%      1.1K
.                    README.md               (MIT OR Unlicense)                 100.00%     7.3K
.                    UNLICENSE               MIT AND Unlicense                  95.16%      1.2K
.                    database_keywords.json  (MIT OR Unlicense)                 100.00%     3.6M
.                    main.go                 (MIT OR Unlicense)                 100.00%     3.4K
examples/identifier  LICENSE                 GPL-3.0+ AND MIT                   95.40%      1K
examples/identifier  LICENSE2                MIT AND GPL-3.0+                   99.65%      35K
examples/identifier  has_identifier.py       (MIT OR GPL-3.0+) AND GPL-2.0      100.00%     428B
parsers              constants.go            (MIT OR Unlicense)                 100.00%     4.8M
parsers              formatter.go            (MIT OR Unlicense)                 100.00%     8.1K
parsers              formatter_test.go       (MIT OR Unlicense)                 100.00%     976B
parsers              guesser.go              (MIT OR Unlicense)                 100.00%     10.2K
parsers              guesser_test.go         (MIT OR Unlicense)                 100.00%     1.6K
parsers              helpers.go              (MIT OR Unlicense) AND Apache-2.0  100.00%     2.5K
parsers              structs.go              (MIT OR Unlicense)                 100.00%     710B
scripts              build_database.py       (MIT OR Unlicense)                 100.00%     4.7K
scripts              include.go              (MIT OR Unlicense)                 100.00%     999B
```

Or to write out the results to a CSV file

```
$ lc --format csv -output licences.csv --pathblacklist .git,licenses,vendor .
```

Or to a SPDX 2.1 file

```
$lc -f spdx -o spdx_example.spdx --pbl .git,vendor,licenses -dn licensechecker -pn licensechecker .
```

You can specify multiple directories as additional arguments and all results will be merged into a single output

```
$ lc -f tabular ./examples/identifier ./scripts
```

You can also specify files and directories as additional arguments 

```
$ lc -f tabular README.md LICENSE ./examples/identifier
Directory              File               License                        Confidence  Size
                       README.md          NOASSERTION                    100.00%     7.5K
                       LICENSE            MIT                            94.83%      1.1K
./examples/identifier  LICENSE            GPL-3.0+ AND MIT               95.40%      1K
./examples/identifier  LICENSE2           MIT AND GPL-3.0+               99.65%      35K
./examples/identifier  has_identifier.py  (MIT OR GPL-3.0+) AND GPL-2.0  100.00%     428B
```

### SPDX

Running against itself to produce a SPDX file using tools from https://github.com/spdx/tools

```
$ go run main.go  -f spdx -o spdx_example.spdx --pbl .git,vendor,licenses -dn licensechecker -pn licensechecker . && java -jar ./spdx-tools-2.1.12-SNAPSHOT-jar-with-dependencies.jar Verify ./spdx_example.spdx
ERROR StatusLogger No log4j2 configuration file found. Using default configuration: logging only errors to the console. Set system property 'log4j2.debug' to show Log4j2 internal initialization logging.
03:49:29.479 [main] ERROR org.apache.jena.rdf.model.impl.RDFReaderFImpl - Rewired RDFReaderFImpl - configuration changes have no effect on reading
03:49:29.482 [main] ERROR org.apache.jena.rdf.model.impl.RDFReaderFImpl - Rewired RDFReaderFImpl - configuration changes have no effect on reading
This SPDX Document is valid.
```

### Package

Run go build for windows and linux then the following in linux, keep in mind need to update the version

```
zip -r9 lc-1.0.0-x86_64-pc-windows.zip lc.exe && zip -r9 lc-1.0.0-x86_64-unknown-linux.zip lc
```


### TODO

* Add error handling for all the file operations and just in general. Most are currently ignored
* Add logic to guess the file type for SPDX value FileType
* Add addtional unit and integration tests
* Investigate using "github.com/gosuri/uitable" for formatting https://github.com/gosuri/uitable