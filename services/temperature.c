#include <stdlib.h>
#include <time.h>

float CheckTempC() {
    static int initialized = 0;
    if (!initialized) {
        srand(time(NULL));
        initialized = 1;
    }
    return 35.0f + ((float)rand() / RAND_MAX) * 6.0f;
}