package main

import "flag"

func parseFlags() parameters {
	var (
		source      = flag.String("source", "", "path of file that will be parsed")
		destination = flag.String("destination", "./", "path where the file (s) will be generated")
		before      = flag.String("before", "", "a string that will be concatenated before all of the repetitions")
		beforeR     = flag.String("beforeR", "", "a string that will be concatenated before the start of each repetition")
		after       = flag.String("after", "", "a string that will be concatenated after all of the repetitions")
		afterR      = flag.String("afterR", "", "a string that will be concatenated after the end of each repetition")
		name        = flag.String("name", "generatedFile", "resulting file name (without file type)")
		format      = flag.String("format", "txt", "resulting file format")
		splitAt     = flag.Int("split", 100000, "at which line it will split the resulted files")
		remove      = flag.Bool("remove", false, "if passed, it removes the last character of the file before the contents of the \"after\" flag")
	)

	flag.Parse()

	return parameters{
		*source,
		*destination,
		*before,
		*beforeR,
		*after,
		*afterR,
		*name,
		*format,
		*splitAt,
		*remove,
	}
}
