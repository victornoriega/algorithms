#include <iostream>
#include <cstdlib>
#include <fstream>
using namespace std;

int main(){
    ofstream archivoSalida;
    archivoSalida.open("orderedList1kk.txt");

    for(int i = 0 ; i < 1000000; i++){
        archivoSalida << i;
        archivoSalida << "\n";
    }
    return 0;
}
