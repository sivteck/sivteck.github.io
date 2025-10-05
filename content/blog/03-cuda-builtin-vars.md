---
title: CUDA Builtin variables and Indexing [WIP]
description: Understanding threadIdx and blockIdx
date: 2025-10-02
tags: dev
draft: true
---

Basic CUDA examples to understand indexing

One Block, Ten Threads:

```c++
#include <iostream>
__device__ void foo(int i, int j, int* rGPU) {
    rGPU[i + j] = i + j;
}

__global__ void kernelVarHunter(int* rGPU) {
    int i = blockIdx.x;
    int j = threadIdx.x;
    foo(i, j, rGPU);
}

int main() {
  int *result = (int *)malloc(10 * sizeof(int));

  int *rGPU = NULL;
  cudaMalloc((void **)&rGPU, 10 * sizeof(int));

  kernelVarHunter<<<1, 10>>>(rGPU);
  cudaMemcpy(result, rGPU, 10 * sizeof(int), cudaMemcpyDeviceToHost);

  for (int i = 0; i < 10; i++) {
    std::cout << result[i] << " ";
  }
}
```

Compile with nvcc, `nvcc foo.c -o foo` and run it `./foo`.


```
# output
0 1 2 3 4 5 6 7 8 9
```

100 Blocks, each having 128 threads

```c++
#include <iostream>

__device__ void foo(int i, int j, int* rGPU) {
    // j - 0 to 127 | i - 0 to 99
    rGPU[j + 128*i] = j + 128 * i;
}

__global__ void kernelVarHunter(int* rGPU) {
    int i = blockIdx.x;
    int j = threadIdx.x;
    foo(i, j, rGPU);
}

int main() {
    int* result = (int *)malloc(12800*sizeof(int));
    int* rGPU = NULL;
    cudaMalloc((void **)&rGPU, 12800 * sizeof(int));
    kernelVarHunter<<<100,128>>>(rGPU);

    cudaMemcpy(result, rGPU, 12800 * sizeof(int), cudaMemcpyDeviceToHost);

    for (int i = 0; i < 100; i++) {
        for (int j = 0; j < 128; j++) {
            std::cout << result[j + i*128] << " ";
        }
        std::cout << std::endl;
    }
}
```

```
# output
0 1 2 3 4 5 6 7 8 9 ... 127
...
...
12672 12673 12674 ... 12799
```

Grid -> Block -> Thread

blockDim represents dimensions of a block of threads.

Can be three dimensional,
```
blockDim.x
blockDim.y
blockDim.z
```

blockIdx represents each unique block in a grid.
```
blockIdx.x
blockIdx.y
blockIdx.z
```

threadIdx gives a unique thread coordinate inside a block
```
threadIdx.x
threadIdx.y
threadIdx.z
```

Work Per Thread = Total Work/(Number Of Blocks * Threads Per Block)

gridDim
blockDim
dim3
blockIdx
threadIdx
divup
