#include <stdio.h>
#include <zmq.h>

int main() {
    int
	x, y, z;

    zmq_version(&x, &y, &z);
    printf("%i.%i.%i\n", x, y, z);

    return 0;
}
