# Cyphers
A quick little utility to randomly generate cyphers for the Cypher RPG game.  Note that, at present,
this utilitly only generates sublte cyphers.

Usage:

        cyphers [flags] [n]

Where "n" is the number of cyphers to generate.

      -no-clip             Do not put on clipboard
  -f, -full                Full (non-Roll20 formatted) output
  -u, -unique              Do not select the same cypher twice

Cypher lists
  You can select which list of cyphers to use.
  -n                       Use cypher list from Numenera (DEFAULT)
  -s                       Use the subtle cyphers from the Cypher System core rules

# Adding new cyphers
This project embeds the Cypher data into the executable.  It does this by converting text files into Go source files.  When the command "go generate" is used, It will look at any text file in the cypher-data subdirectory and attempt to convert this to a similarly named Go source code file, with xxxx.txt being named xxxxCypher.go in the types package.  So to add new cyphers, "go generate" will have to be executed and the project then rebuilt.

TODO:
  * Convert subtle cyphers from Go source into a text file (which would then generate a Go file), hopefully making this easier to add to
  * The unique functionality does not actually work
  * Fix formatting for Roll20 text output, if possible
  * Automatic rolling on subtables in cyphers
  * If conditions (e.g., if cypher level <3 output this, else output that)
