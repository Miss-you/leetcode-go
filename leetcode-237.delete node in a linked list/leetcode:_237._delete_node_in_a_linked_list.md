#leetcode: 237. Delete Node in a Linked List

Write a function to delete a node (except the tail) in a singly linked list, given only access to that node.

Supposed the linked list is 1 -> 2 -> 3 -> 4 and you are given the third node with value 3, the linked list should become 1 -> 2 -> 4 after calling your function.

从链表中，删除指定的节点，但又没有给出前指针，所以先交换当前节点和下一节点的值，然后删除下一节点即可。

> 主要是这里说了删除的节点不包括最末尾的节点。。

```
# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def deleteNode(self, node):
        """
        :type node: ListNode
        :rtype: void Do not return anything, modify node in-place instead.
        """
        if node == None:
            return 
        node.val = node.next.val
        node.next = node.next.next
```