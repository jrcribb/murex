package docs

func init() {
	Summary = map[string]string{

		"!":                     "Reads the STDIN and exit number from previous process and not's it's condition",
		"(":                     "Write a string to the STDOUT without new line",
		"2darray":               "Create a 2D JSON array from multiple input sources",
		"(murex named pipe)":    "Reads from a _murex_ named pipe",
		"<stdin>":               "Read the STDIN belonging to the parent code block",
		"=":                     "Evaluate a mathematical function",
		">>":                    "Writes STDIN to disk - appending contents if file already exists",
		">":                     "Writes STDIN to disk - overwriting contents if file already exists",
		"@[":                    "Outputs a ranged subset of data from STDIN",
		"[[":                    "Outputs an element from a nested structure",
		"[":                     "Outputs an element from an array, map or table",
		"a":                     "A sophisticated yet simple way to build an array or list",
		"alias":                 "Create an alias for a command",
		"alter":                 "Change a value within a structured data-type and pass that change along the pipeline without altering the original source input",
		"and":                   "Returns `true` or `false` depending on whether multiple conditions are met",
		"append":                "Add data to the end of an array",
		"autocomplete":          "Set definitions for tab-completion in the command line",
		"bg":                    "Run processes in the background",
		"cast":                  "Alters the data type of the previous function without altering it's output",
		"catch":                 "Handles the exception code raised by `try` or `trypipe",
		"cd":                    "Change (working) directory",
		"config":                "Query or define _murex_ runtime settings",
		"cpuarch":               "Output the hosts CPU architecture",
		"cpucount":              "Output the number of CPU cores available on your host",
		"debug":                 "Debugging information",
		"die":                   "Terminate murex with an exit number of 1",
		"err":                   "Print a line to the STDERR",
		"escape":                "Escape or unescapes input",
		"esccli":                "Escapes an array so output is valid shell code",
		"eschtml":               "Encode or decodes text for HTML",
		"escurl":                "Encode or decodes text for the URL",
		"event":                 "Event driven programming for shell scripts",
		"exec":                  "Runs an executable",
		"exit":                  "Exit murex",
		"exitnum":               "Output the exit number of the previous process",
		"export":                "Define a local variable and set it's value",
		"f":                     "Lists objects (eg files) in the current working directory",
		"false":                 "Returns a `false` value",
		"fg":                    "Sends a background process into the foreground",
		"fid-kill":              "Terminate a running _murex_ function",
		"fid-killall":           "Terminate _all_ running _murex_ functions",
		"fid-list":              "Lists all running functions within the current _murex_ session",
		"for":                   "A more familiar iteration loop to existing developers",
		"foreach":               "Iterate through an array",
		"format":                "Reformat one data-type into another data-type",
		"function":              "Define a function block",
		"g":                     "Glob pattern matching for file system objects (eg *.txt)",
		"get-type":              "Returns the data-type of a variable or pipe",
		"get":                   "Makes a standard HTTP request and returns the result as a JSON object",
		"getfile":               "Makes a standard HTTP request and return the contents as _murex_-aware data type for passing along _murex_ pipelines.",
		"global":                "Define a global variable and set it's value",
		"history":               "Outputs murex's command history",
		"if":                    "Conditional statement to execute different blocks of code depending on the result of the condition",
		"ja":                    "A sophisticated yet simply way to build a JSON array",
		"jsplit":                "Splits STDIN into a JSON array based on a regex parameter",
		"left":                  "Left substring every item in a list",
		"len":                   "Outputs the length of an array",
		"let":                   "Evaluate a mathematical function and assign to variable",
		"lockfile":              "Create and manage lock files",
		"man-summary":           "Outputs a man page summary of a command",
		"map":                   "Creates a map from two data sources",
		"match":                 "Match an exact value in an array",
		"msort":                 "Sorts an array - data type agnostic",
		"mtac":                  "Reverse the order of an array",
		"murex-docs":            "Displays the man pages for _murex_ builtins",
		"murex-package":         "_murex_'s package manager",
		"murex-parser":          "Runs the _murex_ parser against a block of code",
		"murex-update-exe-list": "Forces _murex_ to rescan $PATH looking for exectables",
		"null":                  "null function. Similar to /dev/null",
		"open-image":            "Renders bitmap image data on your terminal",
		"or":                    "Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.",
		"os":                    "Output the auto-detected OS name",
		"out":                   "`echo` a string to the STDOUT with a trailing new line character",
		"pipe":                  "Manage _murex_ named pipes",
		"post":                  "HTTP POST request with a JSON-parsable return",
		"prefix":                "Prefix a string to every item in a list",
		"prepend":               "Add data to the start of an array",
		"pretty":                "Prettifies JSON to make it human readable",
		"private":               "Define a private function block",
		"pt":                    "Pipe telemetry. Writes data-types and bytes written",
		"rand":                  "Random field generator",
		"read":                  "`read` a line of input from the user and store as a variable",
		"regexp":                "Regexp tools for arrays / lists of strings",
		"right":                 "Right substring every item in a list",
		"runtime":               "Returns runtime information on the internal state of _murex_",
		"rx":                    "Regexp pattern matching for file system objects (eg '.*\\.txt')",
		"set":                   "Define a local variable and set it's value",
		"source":                "Import _murex_ code from another file of code block",
		"struct-keys":           "Outputs all the keys in a structure as a file path",
		"suffix":                "Prefix a string to every item in a list",
		"summary":               "Defines a summary help text for a command",
		"switch":                "Blocks of cascading conditionals",
		"swivel-datatype":       "Converts tabulated data into a map of values for serialised data-types such as JSON and YAML",
		"swivel-table":          "Rotates a table by 90 degrees",
		"ta":                    "A sophisticated yet simple way to build an array of a user defined data-type",
		"test":                  "_murex_'s test framework - define tests, run tests and debug shell scripts",
		"time":                  "Returns the execution run time of a command or block",
		"tmp":                   "Create a temporary file and write to it",
		"tout":                  "Print a string to the STDOUT and set it's data-type",
		"tread":                 "`read` a line of input from the user and store as a user defined *typed* variable",
		"true":                  "Returns a `true` value",
		"try":                   "Handles errors inside a block of code",
		"trypipe":               "Checks state of each function in a pipeline and exits block on error",
		"version":               "Get _murex_ version",
		"while":                 "Loop until condition false",
	}

	Synonym = map[string]string{

		"!":                     "!",
		"(":                     "(",
		"2darray":               "2darray",
		"(murex named pipe)":    "(murex named pipe)",
		"<>":                    "(murex named pipe)",
		"<stdin>":               "<stdin>",
		"=":                     "=",
		">>":                    ">>",
		"fappend":               ">>",
		">":                     ">",
		"fwrite":                ">",
		"@[":                    "@[",
		"[[":                    "[[",
		"[":                     "[",
		"![":                    "[",
		"a":                     "a",
		"alias":                 "alias",
		"!alias":                "alias",
		"alter":                 "alter",
		"and":                   "and",
		"!and":                  "and",
		"append":                "append",
		"autocomplete":          "autocomplete",
		"bg":                    "bg",
		"cast":                  "cast",
		"catch":                 "catch",
		"!catch":                "catch",
		"cd":                    "cd",
		"config":                "config",
		"!config":               "config",
		"cpuarch":               "cpuarch",
		"cpucount":              "cpucount",
		"debug":                 "debug",
		"die":                   "die",
		"err":                   "err",
		"escape":                "escape",
		"!escape":               "escape",
		"esccli":                "esccli",
		"eschtml":               "eschtml",
		"!eschtml":              "eschtml",
		"escurl":                "escurl",
		"!escurl":               "escurl",
		"event":                 "event",
		"!event":                "event",
		"exec":                  "exec",
		"exit":                  "exit",
		"exitnum":               "exitnum",
		"export":                "export",
		"!export":               "export",
		"unset":                 "export",
		"f":                     "f",
		"false":                 "false",
		"fg":                    "fg",
		"fid-kill":              "fid-kill",
		"fid-killall":           "fid-killall",
		"fid-list":              "fid-list",
		"jobs":                  "fid-list",
		"for":                   "for",
		"foreach":               "foreach",
		"format":                "format",
		"function":              "function",
		"!function":             "function",
		"g":                     "g",
		"@g":                    "g",
		"get-type":              "get-type",
		"get":                   "get",
		"getfile":               "getfile",
		"global":                "global",
		"!global":               "global",
		"history":               "history",
		"if":                    "if",
		"!if":                   "if",
		"ja":                    "ja",
		"jsplit":                "jsplit",
		"left":                  "left",
		"len":                   "len",
		"let":                   "let",
		"lockfile":              "lockfile",
		"man-summary":           "man-summary",
		"map":                   "map",
		"match":                 "match",
		"!match":                "match",
		"msort":                 "msort",
		"mtac":                  "mtac",
		"murex-docs":            "murex-docs",
		"murex-package":         "murex-package",
		"murex-parser":          "murex-parser",
		"murex-update-exe-list": "murex-update-exe-list",
		"null":                  "null",
		"open-image":            "open-image",
		"or":                    "or",
		"!or":                   "or",
		"os":                    "os",
		"out":                   "out",
		"echo":                  "out",
		"pipe":                  "pipe",
		"!pipe":                 "pipe",
		"post":                  "post",
		"prefix":                "prefix",
		"prepend":               "prepend",
		"pretty":                "pretty",
		"private":               "private",
		"pt":                    "pt",
		"rand":                  "rand",
		"read":                  "read",
		"regexp":                "regexp",
		"!regexp":               "regexp",
		"right":                 "right",
		"runtime":               "runtime",
		"builtins":              "runtime",
		"rx":                    "rx",
		"set":                   "set",
		"!set":                  "set",
		"source":                "source",
		".":                     "source",
		"struct-keys":           "struct-keys",
		"suffix":                "suffix",
		"summary":               "summary",
		"!summary":              "summary",
		"switch":                "switch",
		"swivel-datatype":       "swivel-datatype",
		"swivel-table":          "swivel-table",
		"ta":                    "ta",
		"test":                  "test",
		"!test":                 "test",
		"time":                  "time",
		"tmp":                   "tmp",
		"tout":                  "tout",
		"tread":                 "tread",
		"true":                  "true",
		"try":                   "try",
		"trypipe":               "trypipe",
		"version":               "version",
		"while":                 "while",
		"!while":                "while",
	}
}
