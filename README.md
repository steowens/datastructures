# Data Structures package.

I know that Golang has this philosophy that you can do anythyhing with maps and slices,
but good lord, when you are trying to read the code written using these raw types it is 
quite an eye chart.

My personal philosophy is create sensible interfaces that lend themselves to writing code
that people more familiar with other languages are liable to recongnize and be able to 
make sense of.  Hence this library of basic data structures so that not everyting has to 
be a brain teaser.  We have enough to think about already.

As I run across the need for data structures I will add them, right now the following
data structures (or my take on them anyhow) are supported.

1. Set - a generic set data structure with union, intersection and compliment operations.
2. Bag - implements an unordered bag of values with union, intersection, difference operations. 
   
---
github.com/steowens/datastructures (c) by Stephen Owens>

github.com/steowens/datastructures is licensed under a
Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License.

You should have received a copy of the license (LICENSE) along with this
work. If not, see <http://creativecommons.org/licenses/by-nc-sa/4.0/>.

