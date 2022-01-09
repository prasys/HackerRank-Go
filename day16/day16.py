#!/bin/python3

import math
import os
import random
import re
import sys


#the main line had to be removed otherwise this program doesn't even work/fails the test cases
#Refer to https://www.hackerrank.com/challenges/30-exceptions-string-to-integer/forum
S = input()
try:
    print(int(S))
except ValueError:
    print("Bad String")