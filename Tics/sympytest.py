from sympy import *
m = 12;
Wo = 1000000;
Rm = 89697;

x, i = symbols("x i")
func1=Function('func1')
func1=summation((1/(1+x)**i),(i,1,m))*Rm-Wo
derivada=func1.diff(x)
print(derivada.subs(x,0.01))
def f(x):
    return func1.diff(x)
#print(f(0))
