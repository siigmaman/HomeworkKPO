#include <stdlib.h>
#include <time.h>

int CheckHealthC() {
    static int initialized = 0;
    if (!initialized) {
        srand(time(NULL));
        initialized = 1;
    }
    return rand() % 2;
}