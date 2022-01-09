class Difference:
    def __init__(self, a):
        self.__elements = a
    
    maximumDifference = 0
    def computeDifference(self):
        
        for i in range(0,len(self.__elements)):
            for j in range(i+1,len(self.__elements)):
                diff = abs(self.__elements[i] - self.__elements[j])
                if (diff > self.maximumDifference):
                    self.maximumDifference = diff 
                

# End of Difference class

_ = input()
a = [int(e) for e in input().split(' ')]

d = Difference(a)
d.computeDifference()

print(d.maximumDifference)