import math
func1 = 0;
m = 12;
Wo = 1000000;
Rm = 89697;

def funcionDada(x):
for i in range(1,m):
func1 = func1+(1/(1+x)^i);
return func1-(Wo/Rm);

func2(x) = 0;

def derivadaFuncionDada(x):
for i in range(1,m):
        func2(x) = func2(x)+(1+x)^(-i-1);
    return func2(x);

i = int(0);
x = float(input("ingresar el valor de x; " ));
tempo = 0;
while(x != tempo and i<100):
    tempo = x;
    x = x-funcionDada(x)/derivadaFuncionDada(x);
    e = abs((x-tempo)/x);

    print("x"+str(i)+"="+str(x)+"error= "+str(e)+"\n");
    i = i+1;

    if(i == 100):
        print("\n\n no converge");
    else
        print("\n\n Solucion x="+str(x));
