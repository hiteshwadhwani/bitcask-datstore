
## Introduction

This repository contains a Go implementation of the Bitcask key-value store as described in the Riak Bitcask paper. Bitcask is a high-performance, persistent key-value storage system optimized for fast reads and writes. It is designed to handle large datasets with low-latency access, making it suitable for applications that require efficient data storage and retrieval.

ref - https://riak.com/assets/bitcask-intro.pdf

## Features

- **High Performance:** Optimized for fast read and write operations.
- **Persistence:** Ensures data durability by storing data on disk.
- **Concurrency Control:** Utilizes read-write locks to manage concurrent access.
- **Data Integrity:** Implements mechanisms to maintain data consistency.
- **Extensible Options:** Supports configurable options for customization.


## Architecture

The Bitcask implementation follows the architecture outlined in the Bitcask paper, focusing on simplicity and performance. Key components include:

- **Data Files:** Stores key-value pairs on disk.
- **Indexing Mechanism:** Maintains an in-memory index to map keys to their locations in data files.
- **Concurrency Management:** Uses synchronization primitives to handle concurrent operations safely.
