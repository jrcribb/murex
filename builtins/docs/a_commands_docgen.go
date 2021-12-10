package docs

func init() {

	Definition["a"] = "# _murex_ Shell Docs\n\n## Command Reference: `a` (mkarray)\n\n> A sophisticated yet simple way to build an array or list\n\n## Description\n\nPronounced \"make array\", like `mkdir` (etc), _murex_ has a pretty sophisticated\nbuiltin for generating arrays. Think like bash's `{1..9}` syntax:\n\n    a: [1..9]\n    \nExcept _murex_ also supports other sets of ranges like dates, days of the week,\nand alternative number bases.\n\n## Usage\n\n{{ include \"gen/includes/mkarray-range-usage.inc copy.md\" }}\n\n## Examples\n\n    » a: [1..3]\n    1\n    2\n    3\n    \n    » a: [3..1]\n    3\n    2\n    1\n    \n    » a: [01..03]\n    01\n    02\n    03\n\n## Detail\n\n### Advanced Array Syntax\n\nThe syntax for `a` is a comma separated list of parameters with expansions\nstored in square brackets. You can have an expansion embedded inside a\nparameter or as it's own parameter. Expansions can also have multiple\nparameters.\n\n    » a: 01,02,03,05,06,07\n    01\n    02\n    03\n    05\n    06\n    07\n    \n    » a: 0[1..3],0[5..7]\n    01\n    02\n    03\n    05\n    06\n    07\n    \n    » a: 0[1..3,5..7]\n    01\n    02\n    03\n    05\n    06\n    07\n    \n    » a: b[o,i]b\n    bob\n    bib\n    \nYou can also have multiple expansion blocks in a single parameter:\n\n    » a: a[1..3]b[5..7]\n    a1b5\n    a1b6\n    a1b7\n    a2b5\n    a2b6\n    a2b7\n    a3b5\n    a3b6\n    a3b7\n    \n`a` will cycle through each iteration of the last expansion, moving itself\nbackwards through the string; behaving like an normal counter.\n\n### Creating JSON arrays with `ja`\n\nAs you can see from the previous examples, `a` returns the array as a\nlist of strings. This is so you can stream excessively long arrays, for\nexample every IPv4 address: `a: [0..254].[0..254].[0..254].[0..254]`\n(this kind of array expansion would hang bash).\n\nHowever if you needed a JSON string then you can use all the same syntax\nas `a` but forgo the streaming capability:\n\n    » ja: [Monday..Sunday]\n    [\n        \"Monday\",\n        \"Tuesday\",\n        \"Wednesday\",\n        \"Thursday\",\n        \"Friday\",\n        \"Saturday\",\n        \"Sunday\"\n    ]\n    \nThis is particularly useful if you are adding formatting that might break\nunder `a`'s formatting (which uses the `str` data type).\n\n## See Also\n\n* [commands/`@[` (range) ](../commands/range.md):\n  Outputs a ranged subset of data from STDIN\n* [commands/`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [commands/`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [commands/`ja` (mkarray)](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [commands/`len` ](../commands/len.md):\n  Outputs the length of an array\n* [commands/`mtac`](../commands/mtac.md):\n  Reverse the order of an array\n* [types/`str` (string) ](../types/str.md):\n  string (primitive)\n* [commands/`ta` (mkarray)](../commands/ta.md):\n  A sophisticated yet simple way to build an array of a user defined data-type"

}
