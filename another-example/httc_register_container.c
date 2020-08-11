#include <stdio.h>
/* gcc -shared -fPIC -nostartfiles httc_register_container.c -o httc_register_container.so */
int httc_register_container(const char *s,const char * b)
{
    static int handle = 100;
    printf("the image need to tag is %s:%s\n", s,b);
    return handle++;
}

