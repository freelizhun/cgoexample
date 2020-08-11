#include "greeter.h"
#include <stdio.h>
#include <unistd.h>
int greet(const char *name, int year, char *out) {
	int n;
        sleep(60);
        n = sprintf(out, "Greetings, %s from %d! We come in peace :)", name, year);
        return n;
}
