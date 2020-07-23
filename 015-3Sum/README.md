## 题目
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:
The solution set must not contain duplicate triplets.

Example:
```
Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
```

## 题目大意

给定一个数组，要求在这个数组中找出 3 个数之和为 0 的所有组合, 不能有重复的。

## 思路
![](https://i.loli.net/2019/04/26/5cc2b50f091b7.jpg)

三个数字肯定至少一个为负（全0除外），即 a + b = -c

- 首先对切片排序，重复的数字会挨在一起。
- 两层循环遍历：
  - 第一层: 指针 `p1` 从头部向尾部移动，移动过程中若发现 该位置数字 > 0 即可停止循环返回结果；
  - 第二层：指针 `p2` 从 `p1 + 1` 位置开始向后移动，`p3` 从尾部向前移动，移动过程中注意重复结果；
    - 若发现满足`a + b = -c` 保存 `a b c` 三个数字， 然后 `p2++` `p3--` 继续寻找；
    - 若发现满足 `a + b < -c` 则说明在 `a`（p1位置）不变的情况下 `b` 数字小，然后 `p2++` 增大 `b` 继续寻找；
    - 若发现满足 `a + b > -c` 则说明在 `a`（p1位置）不变的情况下 `c` 数字大，然后 `p3--` 减小 `c` 继续寻找；
