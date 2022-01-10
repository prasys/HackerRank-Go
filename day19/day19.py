class AdvancedArithmetic(object):
    def divisorSum(n):
        raise NotImplementedError

class Calculator(AdvancedArithmetic):
    
    def __init__(self):
        super().__init__()

    def divisorSum(self, n):
        total = 0
        if n < 0:
            return 0
        for i in range(n+1):
            if i == 0:
                continue
            if n%i == 0:
                #print(i)
                total += i
        return total
                


n = int(input())
my_calculator = Calculator()
s = my_calculator.divisorSum(n)
print("I implemented: " + type(my_calculator).__bases__[0].__name__)
print(s)