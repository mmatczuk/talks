#include <stdio.h>

typedef struct str
{
    char *s;
    int len;
}str;


int main()
{
    str a;
    a.s = "abc";
    printf("%s", a); // "abc " - undefined behaviour! // HL
    return 0;
}
