In wombat the % character goes up a level. At the top (file) level % get into WCTL (Wombat Compile Time Language).
There are 2 operations currently:
 * %operator -- maps operator (prefix and/or suffix) to standard expression
 * %import file -- imports operators and names from file
 
 Higher still is lexical. Currently two ops:
  * %%token someName "regularExpression" -- each token-pos must match only one
  * %%include file
