package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`echo`:            `out`,
	`!and`:            `and`,
	`!or`:             `or`,
	`!catch`:          `catch`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!global`:         `global`,
	`(`:               `brace-quote`,
	`!set`:            `set`,
	`!event`:          `event`,
	`!if`:             `if`,
	`>`:               `>`,
	`rx`:              `rx`,
	`getfile`:         `getfile`,
	`brace-quote`:     `brace-quote`,
	`err`:             `err`,
	`f`:               `f`,
	`g`:               `g`,
	`swivel-datatype`: `swivel-datatype`,
	`swivel-table`:    `swivel-table`,
	`post`:            `post`,
	`global`:          `global`,
	`set`:             `set`,
	`append`:          `append`,
	`read`:            `read`,
	`try`:             `try`,
	`event`:           `event`,
	`murex-docs`:      `murex-docs`,
	`>>`:              `>>`,
	`prepend`:         `prepend`,
	`pt`:              `pt`,
	`ttyfd`:           `ttyfd`,
	`trypipe`:         `trypipe`,
	`export`:          `export`,
	`get`:             `get`,
	`tread`:           `tread`,
	`catch`:           `catch`,
	`alter`:           `alter`,
	`out`:             `out`,
	`tout`:            `tout`,
	`and`:             `and`,
	`or`:              `or`,
	`if`:              `if`,
}
