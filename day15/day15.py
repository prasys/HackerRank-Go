class Node:
    def __init__(self,data):
        self.data = data
        self.next = None 
class Solution: 
    def display(self,head):
        current = head
        while current:
            print(current.data,end=' ')
            current = current.next

    def insert(self,head,data):
        temp = head
        if temp is None: # if it's the first none, we create first
            head = Node(data)
            return head
        else: # else we will traverse untill we see a blank node and we will point to it :) 
            while temp.next is not None:
                temp = temp.next
            temp.next = Node(data)
            return head
            
             
            
    

mylist= Solution()
T=int(input())
head=None
for i in range(T):
    data=int(input())
    head=mylist.insert(head,data)    
mylist.display(head); 	  