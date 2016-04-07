#Power of Three

Given an integer, write a function to determine if it is a power of three.

##Follow up:

Could you do it without using any loop / recursion?

给一个整数，判断其是否是3的幂，要求不可以有迭代和循环

从数学的角度想，可以先求这个数的3的对数，取整数，判断是否相等

```
class Solution(object):
    def isPowerOfThree(self, n):
        """
        :type n: int
        :rtype: bool
        """
        return n > 0 and 3 ** round(math.log(n, 3)) == n
```

对具体计算看看能否优化

```
3^x=n
log(3^x) = log(n)
x log(3) = log(n)
x = log(n) / log(3)
```

感觉这样好用一些。