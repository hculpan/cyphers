# Cyphers
A quick little utility to randomly generate cyphers for the Cypher RPG game.  Note that, at present,
this utilitly only generates sublte cyphers.

Usage:

        cyphers [flags] [n]

Where "n" is the number of cyphers to generate.

      -no-clip             Do not put on clipboard
  -f, -full                Full (non-Roll20 formatted) output
  -u, -unique              Do not select the same cypher twice
  

TODO:
  * The unique functionality does not actually work
  * Fix formatting for Roll20 text output, if possible
  * Automatic rolling on subtables in cyphers
  * If conditions (e.g., if cypher level <3 output this, else output that)
