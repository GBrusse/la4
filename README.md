Language Assignment #4: Go
Issued: Tuesday, April 18
Due: Thursday, April 27

Purpose:
This assignment asks you to begin using a garbage-collected, concurrent, object-
oriented, and imperative programming language named Go. Go has been con-
sidered a competitor of the venerable systems-programming duo of C and C++,
but it has garbage collection. Go was designed by Robert Griesemer, Rob Pike,
and Ken Thompson, at Google, in 2009.

Documentation:
Go lecture slides are at:
pub/slides/slides-go.pdf
Go is demonstrated by:
pub/sum/go
pub/etc/sleepsort
Go is too new to be described in our textbook.
Links to programming-language documentation can be found at:
http://csweb.boisestate.edu/~buff/pl.html


Assignment:
Begin, by porting the simple banking application from Java to Go.
Model your Go solution on the Java solution. Thus, you will
have multiple Go packages.

Then, reimplement the Accrue functions, to use goroutines. This is a bit silly,
since interest-accrual is too simple to deserve separate threads. In any event,
use a channel, so the bank’s Accrue function can sum the interest, and print
the total.

Hints and Advice:
• An abstract class/method can be approximated with an interface that
declares a function that has no definition.
• Your bank can represent its set of accounts as a map from interface point-
ers to interface values. An interface value is essentially an object reference:
accounts map[*IAccount]IAccount
and iterate over them like this:

for _,a:=range b.accounts {
· · ·
}