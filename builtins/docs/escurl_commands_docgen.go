package docs

func init() {

	Definition["escurl"] = "# _murex_ Shell Docs\n\n## Command Reference: `escurl`\n\n> Encode or decodes text for the URL\n\n## Description\n\n`escurl` takes input from either STDIN or the parameters and returns the same\ndata, escaped for the URL.\n\n`!eschtml` does the same process in reverse, where it takes URL escaped data\nand returns its unescaped counterpart.\n\n## Usage\n\nEscape\n\n    <stdin> -> escurl -> <stdout>\n    \n    escurl string to escape -> <stdout>\n    \nUnescape\n\n    <stdin> -> !escurl -> <stdout>\n    \n    !escurl string to unescape -> <stdout>\n\n## Examples\n\nEscape\n\n    » out: \"!? <>\" -> escurl\n    %21%3F%20%3C%3E%0A \n    \nUnescape\n\n    out: '%21%3F%20%3C%3E%0A' -> !escurl\n    !? <>\n\n## Synonyms\n\n* `escurl`\n* `!escurl`\n\n\n## See Also\n\n* [commands/`escape`](../commands/escape.md):\n  Escape or unescape input \n* [commands/`esccli`](../commands/esccli.md):\n  Escapes an array so output is valid shell code\n* [commands/`eschtml`](../commands/eschtml.md):\n  Encode or decodes text for HTML\n* [commands/`get`](../commands/get.md):\n  Makes a standard HTTP request and returns the result as a JSON object\n* [commands/`getfile`](../commands/getfile.md):\n  Makes a standard HTTP request and return the contents as _murex_-aware data type for passing along _murex_ pipelines.\n* [commands/`post`](../commands/post.md):\n  HTTP POST request with a JSON-parsable return"

}
