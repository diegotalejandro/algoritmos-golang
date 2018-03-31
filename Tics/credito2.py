import math
from sympy import *
m = 48
Wo = 8000000
Rm = 215886
#no tocar------------------------------------------
y, i = symbols("y i")
func1=Function('func1')
func1=summation((1/(1+y)**i),(i,1,m))*Rm-Wo
func2=func1.diff(y)
i = int(0);
x=0.0000000000001
tempo = 0;
while(x != tempo and i<100):
    tempo = x;
    x = x-func1.subs(y,x)/func2.subs(y,x);
    e = abs((x-tempo)/x);
    i = i+1;
    if i == 3:
        print("CAE: "+str(12*x*100)+"%");
        break
GastosAdicionales=66339+138514
Wo += GastosAdicionales
y, i = symbols("y i")
func1=summation((1/(1+y)**i),(i,1,m))*Rm-Wo
func2=func1.diff(y)
i = 0
x=0.0000000000001
tempo = 0;
while(x != tempo and i<100):
    tempo = x;
    x = x-func1.subs(y,x)/func2.subs(y,x);
    e = abs((x-tempo)/x);
    i = i+1;
    if i == 3:
        print("interes mensual: "+str(x*100)+"%");
        break
#no tocar------------------------------------------
print("cuota: "+str(Rm))
print("monto del credito: "+str(Wo-GastosAdicionales))
print("plazo en meses: "+str(m))
print("Gastos adicionales: "+str(GastosAdicionales))
print("costo total del credito: "+str(Rm*m))
