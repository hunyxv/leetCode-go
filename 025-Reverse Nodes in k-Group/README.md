## 题目
按照每 K 个元素翻转的方式翻转链表。如果不满足 K 个元素的就不翻转。

### 示例
```
Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5
```

### 思路

024 加强版本，其实就是每个 k 长度子列进行逆序操作。