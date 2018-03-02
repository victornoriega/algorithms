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
        convert << (i+1);
        number = convert.str();
        fileName = "randomList" +number+"k.txt";
        archivoSalida.open(fileName);
        for(int j = 0 ; j < (i+1)*1000; j++){
            archivoSalida << rand() % 500;
            archivoSalida << "\n";
        }
        archivoSalida.close();
    }

    return 0;
}
