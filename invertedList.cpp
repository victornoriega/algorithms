#include <iostream>
#include <cstdlib>
#include <fstream>
#include <cstring>
#include <sstream>

#define N_FILES 10

using namespace std;

int main(){

    ofstream archivoSalida;
    string fileName;
    string number;
    ostringstream convert;


    for(int i = 0 ; i < N_FILES; i++){
        convert.str("");
        convert.clear();
        convert << ((i+1)*100);
        number = convert.str();
        fileName = "invertedList" +number+"k.txt";
        archivoSalida.open(fileName);
        for(int j = (i+1)*100000 ; j >0 ; j--){
            archivoSalida << j;
            archivoSalida << "\n";
        }
        archivoSalida.close();
    }

    return 0;
}
