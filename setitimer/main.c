// Example program showing that signal delivery from setitimer ends up in
// "random" threads of the executing program.

#include <sys/time.h>
#include <unistd.h>
#include <signal.h>
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>

void sighandler(int);
void thread_loop();

int main(int argc, char *argv[]) {
  // Setup SIGPROF handler
  if (signal(SIGPROF, sighandler) == SIG_ERR) {
    perror("Unable to install sighandler");
    exit(1);
  }

  // Spawn a thread
  pthread_t thread;
  if(pthread_create(&thread, NULL, (void *)thread_loop, NULL)) {
    perror("Unable to create thread");
    return 1;
  }

  // Setup itimer to fire every 1s
  struct itimerval it_val;
  it_val.it_value.tv_sec = 1;
  it_val.it_value.tv_usec = 0;
  it_val.it_interval = it_val.it_value;
  if (setitimer(ITIMER_PROF, &it_val, NULL) == -1) {
    perror("error calling setitimer()");
    exit(1);
  }

  pthread_t tid = pthread_self();
  printf("hello from main: %ld\n", tid);

  while (1) {}
}

void thread_loop() {
  printf("hello\n");
  // Setup itimer to fire every 1s
  struct itimerval it_val;
  it_val.it_value.tv_sec = 1;
  it_val.it_value.tv_usec = 0;
  it_val.it_interval = it_val.it_value;
  if (setitimer(ITIMER_PROF, &it_val, NULL) == -1) {
    perror("error calling setitimer()");
    exit(1);
  }

  pthread_t tid = pthread_self();
  printf("hello from thread: %ld\n", tid);
  while (1) {}
}

void sighandler(int signo) {
  struct timeval t;
  gettimeofday(&t, NULL);

  pthread_t tid = pthread_self();
  printf("%ld.%ld received signal: %d in thread: %ld\n", t.tv_sec, t.tv_usec, signo, tid);
}
