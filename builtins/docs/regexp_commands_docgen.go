package docs

func init() {

	Definition["regexp"] = "# _murex_ Shell Guide\n\n## Command Reference: `regexp`\n\n> Regexp tools for arrays / lists of strings\n\n### Description\n\n`regexp` provides a few tools for text matching and manipulation against an\narray or list of strings - thus `regexp` is _murex_ data-type aware.\n\n### Usage\n\n    <stdin> -> regexp expression -> <stdout>\n\n### Examples\n\n#### Find elements:\n\n    » ja: [monday..sunday] -> regexp 'f/^([a-z]{3})day/'\n    [\n        \"mon\",\n        \"fri\",\n        \"sun\"\n    ]\n    \nThis returns only 3 days because only 3 days match the expression (where\nthe days have to be 6 characters long) and then it only returns the first 3\ncharacters because those are inside the parenthesis.\n\n#### Match elements:\n\n    » ja: [monday..sunday] -> regexp 'm/(mon|fri|sun)day/'\n    [\n        \"monday\",\n        \"friday\",\n        \"sunday\"\n    ]\n    \n#### Substitute expression:\n\n    » ja: [monday..sunday] -> regexp 's/day/night/'\n    [\n        \"monnight\",\n        \"tuesnight\",\n        \"wednesnight\",\n        \"thursnight\",\n        \"frinight\",\n        \"saturnight\",\n        \"sunnight\"\n    ]\n\n### Flags\n\n* `f/`\n    output found expressions\n* `m/`\n    output elements that match expression\n* `s/`\n    output all elements - substituting elements that match expression\n\n### Detail\n\n`regexp` is data-type aware so will work against lists or arrays of whichever\n_murex_ data-type is passed to it via STDIN and return the output in the\nsame data-type.\n\n### See Also\n\n* [commands/`2darray` ](../commands/2darray.md):\n  Create a 2D JSON array from multiple input sources\n* [commands/`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [commands/`append`](../commands/append.md):\n  Add data to the end of an array\n* [commands/`ja`](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [commands/`jsplit` ](../commands/jsplit.md):\n  Splits STDIN into a JSON array based on a regex parameter\n* [commands/`len` ](../commands/len.md):\n  Outputs the length of an array\n* [commands/`map` ](../commands/map.md):\n  Creates a map from two data sources\n* [commands/`match`](../commands/match.md):\n  Match an exact value in an array\n* [commands/`msort` ](../commands/msort.md):\n  Sorts an array - data type agnostic\n* [commands/`prepend` ](../commands/prepend.md):\n  Add data to the start of an array\n* [commands/`pretty`](../commands/pretty.md):\n  Prettifies JSON to make it human readable\n* [commands/`ta`](../commands/ta.md):\n  A sophisticated yet simple way to build an array of a user defined data-type\n* [commands/prefix](../commands/prefix.md):\n  \n* [commands/suffix](../commands/suffix.md):\n  "

}