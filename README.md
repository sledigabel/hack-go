# Hack projects in GOLANG

This project regroups my hacky projects in GOLANG in my quest of learning that awesome language

## COLLATZ

Solution to the following problem:

The following iterative sequence is defined for the set of positive integers.

n => n/2 (n is even)
n => 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
13 40 20 10 5 16 8 4 2 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has
not been proved yet (search for the “Collatz Problem” if you're interested), it is thought that all
starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?
*NOTE*: Once the chain starts the terms are allowed to go above one million.
