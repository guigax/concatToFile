# concatToFile

A CLI to read, concatenate and write multiple values into files.

The original intent of this project was to make multiple files of a multi line `INSERT INTO`, based on another file content. Feel free to fit it for your usage.

## Table of Contents

- [concatToFile](#concattofile)
  * [Requirements](#requirements)
  * [Run](#run)
  * [Features](#features)
    + [source](#source)
    + [destination](#destination)
    + [before](#before)
    + [beforeR](#beforer)
    + [after](#after)
    + [afterR](#afterr)
    + [split](#split)
    + [name](#name)
    + [format](#format)
    + [remove](#remove)
  * [Usage](#usage)

## Requirements

[Go 1.18+](https://go.dev/dl/)

## Run

Execute the `go run` command

    go run concat.go

You can also build and execute the code with:

OS | Command
--- | --- 
Windows | `go build concat.go && concat.exe`
Linux | `go build concat.go && ./concat`

It accepts command-line arguments, listed by: 

    go run concat.go --help

## Features

Available flags to use it on the CLI.

### source

Defines a path to the file that will be parsed. The resulting file(s) will be put on the same directory as the original file.

    -source="generateFiles/text_created.txt"

### destination

Path where the files will be generated (default `./`)

    -destination="generateFiles"

### before

Content that will be concatenated before everything.

    -before="INSERT INTO table01 (column1, column2, column3) VALUES "

### beforeR

Short to "before repeat". Content that will be concatenated before every repetition.

    -beforeR="(01, '"

### after

Content that will be concatenated after everything.

    -after=";"

### afterR

Short to "after repeat". Content that will be concatenated after every repetition.

    -afterR="', CURRENT_TIMESTAMP),"

### split

At which line of the original file it will split, to form multiple resulting files (default `100000`).

    -split=5000

### name

Resulting file name, it will be concatenated with `_XX`, which is the number of the file generated (default `generatedFile`).

    -name="table01_20220101"

### format

Resulting file format (default `txt`).

    -format="txt"

### remove

If passed, it removes the last character of the resulting file(s) (default `false`). It does not remove the contents of the `-after` flag.

    -remove

## Usage

Using the file `generateFiles/text_created.txt` as the base file, I want to output multiple files on the following format, splitting 5000 lines per resulting file, in a sql file:

```sql
INSERT INTO table01 (column1, column2, column3) VALUES 
(01, 'line 1', CURRENT_TIMESTAMP),
(01, 'line 2', CURRENT_TIMESTAMP),
(01, 'line 3', CURRENT_TIMESTAMP),
(01, 'line 4', CURRENT_TIMESTAMP),
(01, 'line 5', CURRENT_TIMESTAMP)
;
```

Then the command-line arguments will be like this:

    go run concat.go -source="generateFiles/text_created.txt" -destination="generateFiles" -before="INSERT INTO table01 (column1, column2, column3) VALUES " -beforeR="(01, '" -afterR="', CURRENT_TIMESTAMP)," -after=";" -split=5000 -name="table01_20220101" -format="sql" -remove

The result will be 151 files named `table01_20220101_01.sql`, `table01_20220101_02.sql`, `table01_20220101_03.sql`, etc... Each with 5000 records in them at the destination.