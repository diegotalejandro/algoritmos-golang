import math
from sympy import *
m = 48 #plazo en meses
Wo = 8000000 #Monto liquido del credito
Rm = 215886 #Cuota mensual
#Formula para sacar CAE------------------------------------------
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
#Estos son los gastos q teni q modificar, onda para recibir uno o mas gastos
GastosAdicionales=66339+138514
Wo += GastosAdicionales
#Formula para sacar interes mensual---------------------------------
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
#prints de la info total, nose q mas poner xD------------------------------------------
print("cuota: "+str(Rm))
print("monto del credito: "+str(Wo-GastosAdicionales))
print("plazo en meses: "+str(m))
print("Gastos adicionales: "+str(GastosAdicionales))
print("costo total del credito: "+str(Rm*m))
