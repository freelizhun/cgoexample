#include <stdio.h>
/* gcc -shared -fPIC -nostartfiles openimage.c -o openimage.so */
int openimage(const char *s,const char * b)
{
    static int handle = 100;
    printf("the foyouage need to tagtagtagtaged is %s:%s\n", s,b);
    return handle++;
}

