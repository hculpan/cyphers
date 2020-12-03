# Cyphers
A quick little utility to randomly generate cyphers for the Cypher RPG game.  Note that, at present,
this utilitly only generates sublte cyphers.

# Adding new cyphers
This project embeds the Cypher data into the executable.  It does this by converting text files into Go source files.  When the command "go generate" is used, It will look at any text file in the cypher-data subdirectory and attempt to convert this to a similarly named Go source code file, with xxxx.txt being named xxxxCypher.go in the types package.  So to add new cyphers, "go generate" will have to be executed and the project then rebuilt.

# Building
In addition to the normal `go build/install` commands, this project will require two additional steps to build from scratch.
1. To update the list of cyphers, use `go generate` to convert from the text data to the Go files (only needed for Numenera cyphers at this point)
2. When including binary resources (e.g., clip.png), use `fyne bundle <resource> > bundle.go`, then copy the resultant static resource into another file (e.g., main.go)

#TODO:
  * Convert subtle cyphers from Go source into a text file (which would then generate a Go file), hopefully making this easier to add to
  * The unique functionality does not actually work
  * Fix formatting for Roll20 text output, if possible
  * Automatic rolling on subtables in cyphers
  * If conditions (e.g., if cypher level <4 output this, else output that)
