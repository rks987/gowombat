In wombat the % character goes up a level. At the top (file) level % get into 
WCTL (Wombat Compile Time Language).
There are 2 operations currently:
 * %operator -- maps operator (prefix and/or suffix) to standard expression
 * %import file -- imports (all) operators and (selected) names from file 
(containing %export and %noexport lines)
 * also: %( ...wctl... )% returns a wctl Expr -- used to create primitive stuff
 
Higher still is lexical. Currently two ops:
 * %%token someName "regularExpression" -- each token-pos must match only one
 * %%include file.wh

There is an implicit %%include WOMBATPATH/include/wombat.wh

The lexical scanner sends a sequence of tokens on a channel, each being a struct:
 * file/line/char pos
 * token type name
 * token text

This is then processed by the active operators to create an AST. Each operator 
will map to a standard form:
 * "call": 2 parameters
 * "tuple": list of parameters
 * "closure": 1 parameter
 * "disjoint union": 2 parameters (atom(.left or .right), value)
 * "Tuple": 1 parameter List(Type)
 * etc

We end up with an AST of entries consisting of:
 * file/line/char pos of the operator
 * 2 chans (parent up/down)
 * standard form type (call/tuple/...)
 * list of pointers to parameters as appropriate

Each node listens on its parents down channel and all its paremters up channels 
until all those channels are closed. Whenever it learns something to change its 
type it lets its parent and children know on their up channel and childrens down 
channel. When its incoming channel is closed it closes all its out channels.

Whenever it gets worked out that a particular closure will be called at a 
particular place then a specialization of the closure replaces it there. This 
establishes additional channels between the original closure and its 
specializations and between the specializations and the relevant call expression.

Is that all??
