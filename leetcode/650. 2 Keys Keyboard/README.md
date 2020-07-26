Saying n being a number of input

First approach 
1. Factorization
1.1 examine whether input n is prime number or not
1.1.1 method for finding prime number in golang is ProbablyPrime()
1.1.2 It use MillerRabin's algorithm
1.2 if not prime, begin for loop
1.2.1 conduct recursive division n with arbitary number starting with 3 to specific number whose power of 2 is equal to n or greater than
1.2.2 queueing every primenumbers and its exponent 
2. Find exponent
3. Find Output