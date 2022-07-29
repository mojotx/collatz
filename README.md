# Collatz Conjecture

## Overview

From Wikipedia: [Collatz Conjecture](https://en.wikipedia.org/wiki/Collatz_conjecture)

> The Collatz conjecture is a famous unsolved mathematical problem. Ihe conjecture
asks whether repeating two simple arithmetic operations will eventually transform
every positive integer into 1. It concerns sequences of integers in which each
term is obtained from the previous term as follows: if the previous term is
even, the next term is one half of the previous term. If the previous term is
odd, the next term is 3 times the previous term plus 1. The conjecture is that
these sequences always reach 1, no matter which positive integer is chosen to
start the sequence.
>
> It is named after mathematician Lothar Collatz, who introduced the idea in 1937,
two years after receiving his doctorate. It is also known as the 3n + 1 problem,
the 3n + 1 conjecture, the Ulam conjecture (after Stanisław Ulam), Kakutani's
problem (after Shizuo Kakutani), the Thwaites conjecture (after Sir Bryan
Thwaites), Hasse's algorithm (after Helmut Hasse), or the Syracuse problem.
The sequence of numbers involved is sometimes referred to as the hailstone
sequence, hailstone numbers or hailstone numerals (because the values are
usually subject to multiple descents and ascents like hailstones in a cloud),
or as wondrous numbers.

This is my toy code for playing with the problem, written in
[Go](https://go.dev/).

## Installation

```shell

$ go install -v github.com/mojotx/collatz@latest

```

## Example

```shell
$ collatz 97
97 292 146 73 220 110 55 166 83 250 125 376 188 94 47 142 71 214 107 322 161 484 242 121 364 182 91 274 137 412 206 103 310 155 466 233 700 350 175 526 263 790 395 1186 593 1780 890 445 1336 668 334 167 502 251 754 377 1132 566 283 850 425 1276 638 319 958 479 1438 719 2158 1079 3238 1619 4858 2429 7288 3644 1822 911 2734 1367 4102 2051 6154 3077 9232 4616 2308 1154 577 1732 866 433 1300 650 325 976 488 244 122 61 184 92 46 23 70 35 106 53 160 80 40 20 10 5 16 8 4 2 1
steps: 118
```

## License

Licensed under the [MIT License](https://opensource.org/licenses/MIT). If you
find this useful or interesting, great!
