package docs

func init() {

	Definition["err"] = "# _murex_ Language Guide\n\n## Command Reference: `err`\n\n> Print a line to the STDERR\n\n### Description\n\nWrite parameters to STDERR with a trailing new line character.\n\n### Usage\n\n    err: string to write -> <stderr>\n\n### Examples\n\n    » err Hello, World!\n    Hello, World!\n\n### Detail\n\n`err` outputs as `string` data-type. This can be changed by casting\n\n    err { \"Code\": 404, \"Message\": \"Page not found\" } ? cast json\n    \nHowever passing structured data-types along the STDERR stream is not recommended\nas any other function within your code might also pass error messages along the\nsame stream and thus taint your structured data. This is why _murex_ does not\nsupply a `tout` function for STDERR. The recommended solution for passing\nmessages like these which you want separate from your STDOUT stream is to create\na new _murex_ named pipe.\n\n    » pipe: --create messages\n    » bg { <messages> -> pretty }\n    » tout: <messages> json { \"Code\": 404, \"Message\": \"Page not found\" }\n    » pipe: --close messages\n    {\n        \"Code\": 404,\n        \"Message\": \"Page not found\"\n    }\n    \n#### ANSI Constants\n\n`err` supports ANSI constants.\n\n### See Also\n\n* [`(` (brace quote)](../commands/brace-quote.md):\n  Write a string to the STDOUT without new line\n* [`>>` (append file)](../commands/greater-than-greater-than.md):\n  Writes STDIN to disk - appending contents if file already exists\n* [`>` (truncate file)](../commands/greater-than.md):\n  Writes STDIN to disk - overwriting contents if file already exists\n* [`bg`](../commands/bg.md):\n  Run processes in the background\n* [`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* [`out`](../commands/out.md):\n  `echo` a string to the STDOUT with a trailing new line character\n* [`pretty`](../commands/pretty.md):\n  Prettifies JSON to make it human readable\n* [`pt`](../commands/pt.md):\n  Pipe telemetry. Writes data-types and bytes written\n* [`tout`](../commands/tout.md):\n  Print a string to the STDOUT and set it's data-type\n* [pipe](../commands/pipe.md):\n  \n* [sprintf](../commands/sprintf.md):\n  "

}
