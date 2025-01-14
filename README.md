# Distributed CAS Storage

A distributed content-addressed storage (CAS) system designed to enable users to store and retrieve files across multiple peer nodes. This system leverages a decentralized architecture, providing features such as secure file storage, encryption, and peer-to-peer (P2P) communication. It allows multiple peers (servers) to share their storage resources, ensuring that data is replicated and accessible without relying on a central authority.

## About

Distributed CAS Storage offers a decentralized file storage solution that allows users to upload, store, and retrieve files securely across multiple peers. The system works by utilizing content-based addressing (hashing), which ensures that files can be reliably identified and retrieved without the need for a central server. Data is encrypted both during transfer and while stored in the network, providing enhanced security.

Key Features of this project:

- **Decentralized File Storage**: Files are split across multiple peer nodes, so there’s no single point of failure. The storage network becomes more robust as more peers join.
- **Content-Addressable Storage**: Each file is identified by a unique hash of its contents, making it immutable and ensuring data consistency.
- **Peer-to-Peer Communication**: Nodes (servers) directly communicate with each other using a TCP protocol to share files and metadata.
- **Replication and Fault Tolerance**: Files are replicated across multiple peers to improve availability and resilience against node failures.
- **Encryption**: AES encryption is used to ensure that file data remains confidential both in transit and at rest.

This system is ideal for scenarios requiring secure, distributed, and redundant file storage without a centralized server infrastructure.

## Features

- **Decentralized File Storage**: Files are distributed across a network of peers, and data is accessible from any peer that has it stored, reducing reliance on a single server.
- **Data Encryption**: Files are encrypted using AES (Advanced Encryption Standard) to ensure secure storage and transfer. This ensures that only authorized peers can read the data.
- **Content-Addressable Storage**: Each file is identified by a hash of its content, which ensures data integrity and makes the system immune to accidental or malicious modification of files.
- **Peer-to-Peer Communication**: Nodes communicate directly with each other using a peer-to-peer network. This allows for efficient file retrieval and reduces reliance on centralized infrastructure.
- **Automatic Peer Discovery**: New peers automatically discover and connect with existing peers in the network, forming a distributed mesh network of nodes.
- **Data Replication**: Files are replicated across multiple nodes in the network to ensure high availability, resilience, and fault tolerance.

## Architecture

The system follows a decentralized, distributed architecture. It consists of multiple peer nodes that communicate with each other to store and retrieve files. Each node in the network:

- **Acts as a File Storage Node**: Stores a part of the overall file storage and can retrieve files stored on other peers.
- **Participates in P2P File Sharing**: Nodes use TCP to directly communicate with each other, offering and requesting files based on their unique content-based hash keys.
- **File Hashing**: When a file is uploaded to the system, its content is hashed (e.g., using SHA-256), and the hash serves as a unique identifier. This hash is used to locate the file across the network.
- **Replication**: Files are replicated across multiple peers to ensure data availability even if one or more peers go offline.

Key components of the system:

- **Nodes**: Each peer in the network that both stores files and serves requests for file retrieval.
- **Peer Discovery**: New nodes are introduced to the network through a bootstrap process. Nodes can discover each other and form a decentralized network automatically.
- **File Storage**: Files are stored using their content hash, ensuring they are immutable and directly retrievable.
- **Encryption**: File data is encrypted during both storage and transfer to prevent unauthorized access.

## Installation

### Prerequisites

Before you can run this system, ensure that you have the following tools installed:

- **Go (v1.18 or higher)**: The programming language used to develop this system. Download and install Go from [the official website](https://golang.org/dl/).
- **Git**: A version control system used to clone the repository. Install Git from [here](https://git-scm.com/).
- **A Go-Compatible Editor**: Editors like Visual Studio Code or GoLand are recommended for Go development. You can also use any text editor with Go support.

### Clone the repository

First, clone the repository from GitHub to your local machine using Git:

```bash
git clone https://github.com/Hamzenium/Distributed-CAS-Storage.git
