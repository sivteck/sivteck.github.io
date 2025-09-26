---
title: Clangd setup for CUDA.
description: Installing and Setting up clangd in Ubuntu 
date: 2025-09-26
tags: devtools
---

To use clangd for CUDA 12.8 and higher, we need to install the latest clangd version.

1. Install [clangd-20](https://apt.llvm.org/).
2. Symlink clangd-20 to clangd,

```bash
sudo ln -s /usr/bin/clangd-20 /usr/local/bin/clangd 
```

