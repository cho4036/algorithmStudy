from typing import List


# class Solution:
#     def generateParenthesis(self, n: int) -> List[str]:
#         if n == 1:
#             return ["()"]
#         tmp = []
#         tmpList = self.generateParenthesis(n-1)
#         for s in tmpList:
#             tmp.append("()"+s)
#             tmp.append("("+s+")")
#             tmp.append(s+"()")
#         my_set = set(tmp)
#         tmp = list(my_set)
#         tmp.sort()
#         return tmp

class Node:
    def __init__(self, data = "", sum = 0, count = 0) -> None:
        if data=="":
            print("initialize")
        self.data = data
        self.sum = sum
        self.count = count
        self.left = None
        self.right = None
    def insert(self, data, sum, count, left, right) :
        if left :
            self.left = Node(data, sum, count)
        if right :
            self.right = Node(data, sum, count)

class Solution:
    def __init__(self) -> None:
        self.root = Node()

    def generateParenthesis(self, n: int) -> List[str]:
        l = []
        self.dfs(self.root, 0, 0, 2*n, l)
        return l

    def dfs(self, node: Node, left:int, current:int, depth:int, l: List[str]):
        if current > depth :
            return 
        if left < depth/2 :
            if node.left == None:
                node.left = Node(node.data+"(", node.sum+1, node.count+1)
            self.dfs(node.left, left+1, current+1, depth, l)
        if node.sum > 0 :
            if node.right == None:
                node.right = Node(node.data+")", node.sum-1, node.count+1)
            self.dfs(node.right, left, current+1, depth, l)
        if node.count == depth:
            l.append(node.data)
        

s = Solution()
print(s.generateParenthesis(4))
t = ["(((())))","((()()))","((())())","((()))()","(()(()))","(()()())","(()())()","(())(())","(())()()","()((()))","()(()())","()(())()","()()(())","()()()()"]
t.sort()
print(t)

print(s.generateParenthesis(3))
t = ["((()))","(()())","(())()","()(())","()()()"]
t.sort()
print(t)

print(s.generateParenthesis(2))
t = ["()()","(())"]
t.sort()
print(t)

print(s.generateParenthesis(1))
t = ["()"]
t.sort()
print(t)