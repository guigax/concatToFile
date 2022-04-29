# concatToFile

A CLI to read, concatenate and write multiple values into files.

This project was conceived to made files of a multi line `INSERT INTO`, based on another file with the contents of the SQL.

# Requirements

[Go 1.18+](https://go.dev/dl/)

# Run

Execute the `go run` command

    go run ./

You can also build and execute the code with:

    go build && concat.exe

It accepts command-line arguments, listed by: 

    go run ./ --help

# Features

### path

Defines a path to the file that will use on the concatenate proccess.

    -path="generateFiles/text_created.txt"

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

    -afterR="', 'John Doe'),"

### remove

If passed, it removes the last character of the resulting file(s) (default `false`).

    -remove

### split

At which line of the original file it will split, to form multiple resulting files (default `100000`).

    -split=5000

### name

Resulting file name, it will be concatenated with `_XX`, which is the number of the file generated (default `generatedFile`).

    -name="table01_20220101"

### format

Resulting file format (default `sql`).

    -format="txt"

# Usage

Using the file `generateFiles/text_created.txt` as the base file, I want to output multiple files o the following format, splitting 5000 lines per resulting file, in a txt file:

```sql
INSERT INTO table01 (column1, column2, column3) VALUES 
(01, 'line 1', 'John Doe'),
(01, 'line 2', 'John Doe'),
(01, 'line 3', 'John Doe'),
(01, 'line 4', 'John Doe'),
(01, 'line 5', 'John Doe')
;
```

Then the command-line arguments will be passed in this way:

    go run ./ -path="generateFiles/text_created.txt" -before="INSERT INTO table01 (column1, column2, column3) VALUES " -beforeR="(01, '" -afterR="', 'John Doe')," -after=";" -split=5000 -name="table01_20220101" -format="txt" -remove

The result will be 151 files named `table01_20220101_01.txt`, `table01_20220101_02.txt`, `table01_20220101_03.txt`, etc... Each with 5000 records in them in the same path of the original file.