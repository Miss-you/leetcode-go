#leetcode-258.Add Digits

Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.

For example:

Given num = 38, the process is like: 3 + 8 = 11, 1 + 1 = 2. Since 2 has only one digit, return it.

Follow up:
Could you do it without any loop/recursion in O(1) runtime?

如果纯粹按照上面的想法，显然是想不出有什么优化方式。但是本题要求我们使用O（1）的时间复杂度，那么应该换一种表达方式才能够解决。

##数根

在数学中，数根(又称数字根Digital root)是自然数的一种性质，换句话说，每个自然数都有一个数根。

数根是将一正整数的各个位数相加（即横向相加），若加完后的值大于10的话，则继续将各位数进行横向相加直到其值小于十为止[1]，或是，将一数字重复做数字和，直到其值小于十为止，则所得的值为该数的数根。

例如54817的数根为7，因为5+4+8+1+7=25，25大于10则再加一次，2+5=7，7小于十，则7为54817的数根。

参考计算地址

https://en.wikipedia.org/wiki/Digital_root

##用途

数根可以计算模运算的同余，对于非常大的数字的情况下可以节省很多时间。

数字根可作为一种检验计算正确性的方法。例如，两数字的和的数根等于两数字分别的数根的和。

另外，数根也可以用来判断数字的整除性，如果数根能被3或9整除，则原来的数也能被3或9整除。

##实现

dr(n) = 0 if n = 0,

dr(n) = 9 if n%9 = 0,

dr(n) = n mod 9 if n%9 != 0

```
func addDigits(num int) int {
    if (num == 0){
        return 0
    }
    
    finalDigit := num%9
    if (finalDigit == 0){
        return 9
    } else{
        return finalDigit
    }
}
```