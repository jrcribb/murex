package docs

func init() {

	Definition["tread"] = "# _murex_ Language Guide\n\n## Command Reference: `tread`\n\n> `read` a line of input from the user and store as a user defined *typed* variable    \n\n### Description\n\nA readline function to allow a line of data inputed from the terminal and then\nstore that as a typed variable.\n\n### Usage\n\n    tread: data-type \"prompt\" var_name\n    \n    <stdin> -> tread: data-type var_name\n\n### Examples\n\n    tread: qs \"Please paste a URL: \" url\n    out: \"The query string values included were:\"\n    $url -> format json\n    \n    out: Please paste a URL: -> tread: qs url\n    out: \"The query string values included were:\"\n    $url -> format json\n\n### Detail\n\nIf `tread` is called as a method then the prompt string is taken from STDIN.\nOtherwise the prompt string will be the first parameter. However if no prompt\nstring is given then `tread` will not write a prompt.\n\nThe last parameter will be the variable name to store the string read by `tread`.\nThis variable cannot be prefixed by dollar, `$`, otherwise the shell will write\nthe output of that variable as the last parameter rather than the name of the\nvariable.\n\n### See Also\n\n* [`(` (brace quote)](../commands/brace-quote.md):\n  Write a string to the STDOUT without new line\n* [`err`](../commands/err.md):\n  Print a line to the STDERR\n* [`out`](../commands/out.md):\n  `echo` a string to the STDOUT with a trailing new line character\n* [`read`](../commands/read.md):\n  `read` a line of input from the user and store as a variable\n* [`tout`](../commands/tout.md):\n  Print a string to the STDOUT and set it's data-type\n* [cast](../commands/cast.md):\n  \n* [format](../commands/format.md):\n  \n* [pretty](../commands/pretty.md):\n  \n* [sprintf](../commands/sprintf.md):\n  "

}
