#include <iostream>
#include <fstream>
using namespace std;
string codificar(){

}
int algoritmo_dinamico(int P,int N,int N_costo[],int N_puntaje[]){
bool A[P];
A[0]=true;
for(int i=1;i<P;i++){
  A[i]=false;
    for(int j=i;j<=N;j++){
      if(a[i-N_costo[j]]==true && i-N_costo[j]>=0){
        a[i]=true;
          }
    }
  }
  return A[P];
}



int main () {

 
 ifstream ficheroEntrada;
 string frase;
 
 ficheroEntrada.open ("ficheroTexto.txt");
 getline(ficheroEntrada, frase);
 ficheroEntrada.close();

 
 cout << "Frase leida: " << frase << endl;
 
 return 0;

}
