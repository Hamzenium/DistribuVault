def f(x):
    ''' Pre: x is a natural number '''
    a = x
    y = 10
    while a > 0:
        a = a - y
        y = y - 1
    return a * y
print(f(100))