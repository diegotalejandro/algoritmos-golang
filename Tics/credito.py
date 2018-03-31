import math
from sympy import *
m = 12;
Wo = 1000000;
Rm = 89697;

y, i = symbols("y i")
func1=Function('func1')
func1=summation((1/(1+y)**i),(i,1,m))*Rm-Wo
func2=func1.diff(y)
#---------------------------------------------
"""def funcionDada(x):

    #func1=float(0)
    #result=0

    #for a in range(1,m):
    #    func1 += (1/(1+x)**a)
    #    break
    #    func1 = func1*Wo - Rm
    #print(func1,"\n\n\n")
    return summation((1/(1+x)**i),(i,1,m))*Rm-Wo
    funcionDada(0.01)
#--------------------------------------------

def derivadaFuncionDada(y):
    #func2=float(0)
    #result=0

    #for b in range(1,m):
    #    func2 += 1/((1+x)**(b+1))
    #    break
    #    func2 = func2*Rm
    #print(derivada)
    return func1.diff(x)
    derivadaFuncionDada(0.01)
    print("hola")"""
#--------------------------------------------
i = int(0);
#x = float(input("ingresar el valor de x; " ));
x=0.0000000000001
tempo = 0;
#while(x != tempo and i<100):

while(i<100):
    #tempo = x;
    x = x-func1.subs(y,x)/func2.subs(y,x);
    e = abs((x-tempo)/x);

    print("x"+str(i)+"="+str(x)+"error= "+str(e)+"\n");
    i = i+1;
    if i == 100:
        print("\n\n no converge");
        break
    else:
        print("\n\n Solucion x="+str(x));
