package docs

func init() {

	Definition["msort"] = "# `msort` \n\n> Sorts an array - data type agnostic\n\n## Description\n\nThis builtin takes input from stdin, sorts it and the outputs it to stdout. \n\nThe code behind `msort` is significantly more lightweight than UNIX sort.\nIt doesn't work with numeric types (eg sorting floating point numbers),\nreversed order nor multi-column data. It is specifically designed to work\nwith lists of data. For example arrays in data formats like JSON (`json`),\nYAML (`yaml`) or S-Expressions (`sexp`); or lists of strings (`str`). The\nintention is to cover use cases not already covered by UNIX sort while also\nproviding something rudimentary for Murex scripts to function on Windows\nwithout having to write lots of ugly platform-specific code. This is also\nthe reason this builtin is called `msort` rather than conflicting with the\nexisting UNIX name, `sort`.\n\n## Usage\n\n```\n<stdin> -> msort -> <stdout>\n```\n\n## Examples\n\n```\n» tout: json ([\"c\", \"b\", \"a\"]) -> msort   \n[\n    \"a\",\n    \"b\",\n    \"c\"\n]\n```\n\nSince `msort` does not support reversed order, you will need to pipe the\noutput of `msort` into another builtin:\n\n```\n» tout: json ([\"c\", \"b\", \"a\"]) -> msort -> mtac \n[\n    \"c\",\n    \"b\",\n    \"a\"\n]\n```\n\n## Synonyms\n\n* `msort`\n* `list.sort`\n\n\n## See Also\n\n* [`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [`[` (range) ](../commands/range.md):\n  Outputs a ranged subset of data from STDIN\n* [`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [`alter`](../commands/alter.md):\n  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input\n* [`append`](../commands/append.md):\n  Add data to the end of an array\n* [`count`](../commands/count.md):\n  Count items in a map, list or array\n* [`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [`jsplit` ](../commands/jsplit.md):\n  Splits STDIN into a JSON array based on a regex parameter\n* [`mtac`](../commands/mtac.md):\n  Reverse the order of an array\n* [`prepend` ](../commands/prepend.md):\n  Add data to the start of an array\n* [index](../commands/item-index.md):\n  Outputs an element from an array, map or table"

}
