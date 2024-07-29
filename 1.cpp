#include <stdio.h>
#include <string>


int main() {
    std::string a = "hello";
    char* b = "hello2";
    printf("%s", a.c_str());

    return 0;
}