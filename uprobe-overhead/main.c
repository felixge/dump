#include <stdio.h>

int add(int a, int b) {
  return a + b;
}

void main() {
  long int sum = 0;
  int i;
  for (i = 0; i < 1000*1000; i++) {
    sum = add(sum, i);
  }
  fprintf(stdout, "sum: %ld\n", sum);
}
