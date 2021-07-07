// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void burn_cgo() {
	while (1) asm("");
}

static void *burn_pthread(void *arg) {
	while (1) asm("");
}

static void *burn_pthread_go(void *arg) {
	extern void burn_callback(void);
	burn_callback();
}

void prep_cgo() {
	burn_cgo();
}

void prep_pthread() {
	int res;
	pthread_t tid;

	res = pthread_create(&tid, NULL, burn_pthread, NULL);
	if (res != 0) {
		fprintf(stderr, "pthread_create: %s\n", strerror(res));
		exit(EXIT_FAILURE);
	}
}

void prep_pthread_go() {
	int res;
	pthread_t tid;

	res = pthread_create(&tid, NULL, burn_pthread_go, NULL);
	if (res != 0) {
		fprintf(stderr, "pthread_create: %s\n", strerror(res));
		exit(EXIT_FAILURE);
	}
}
