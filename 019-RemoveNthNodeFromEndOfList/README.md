## 题目
给定一个链表，删除倒数第 n 个节点，然后返回链表（头节点）。
### 例子
```
Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
```

## 思路

双指针法
- 首先 `p1` 指向 `head`， `p2` 移动到距离 `p1` 为 `n` 的位置；
- 然后同时向后移动 `p1` 和 `p2`， 直到 `p2` 指向末尾，此时 `p1` 位置就是应该删除的节点。