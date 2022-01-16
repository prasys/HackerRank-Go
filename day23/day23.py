import sys

class Node:
    def __init__(self,data):
        self.right=self.left=None
        self.data = data
class Solution:
    def insert(self,root,data):
        if root==None:
            return Node(data)
        else:
            if data<=root.data:
                cur=self.insert(root.left,data)
                root.left=cur
            else:
                cur=self.insert(root.right,data)
                root.right=cur
        return root

    def levelOrder(self,root):
        try:
            treeQueue = [] # queue for it
            temp_node = root
            while temp_node is not None:
                print(temp_node.data,end=" ")
                if temp_node.left is not None:
                    treeQueue.append(temp_node.left) # left first
                if temp_node.right is not None:
                    treeQueue.append(temp_node.right)
                    
                temp_node = treeQueue.pop(0)
        except Exception:
            pass
                    
                
            

T=int(input())
myTree=Solution()
root=None
for i in range(T):
    data=int(input())
    root=myTree.insert(root,data)
myTree.levelOrder(root)
