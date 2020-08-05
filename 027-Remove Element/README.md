## 题目
给定一个数组和一个值，在原地删除与值相同的数字，返回新数组的长度。

元素的顺序可以改变，并且对新的数组不会有影响。

### 示例
```
Example 1:
	Input: [], value = 0
	Output: 0
```

```
Example 2:
	Input:  [0,4,4,0,0,2,4,4], value = 4
	Output: 4
	
	Explanation: 
	the array after remove is [0,0,0,2]
```

## 思路
思路和026 相同， 维护两个指针 `i` 和 `p`， `i` 随着遍历向前移动，只有遇到与 `val` 不同的值时指针`p`才向前移动一位，然后交换 `i` `p`位置的数据。