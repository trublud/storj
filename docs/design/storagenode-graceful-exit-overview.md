# Storage Node Graceful Exit Overview

## Abstract

When a Storage Node wants to leave the network, we would like to ensure that the pieces get transferred to other nodes. This avoids additional repair cost and bandwidth. This design document describes the necessary components for the process.

## Background

Storage Nodes may want to leave the network. Taking a Storage Node directly offline would mean that the Satellite needs to repair the lost pieces. To avoid this we need a process to transfer the pieces to other nodes before taking the exiting node offline. We call this process Graceful Exit. Since this process can take time we need overview of the nodes exiting and the progress.

A Storage Node may also want to exit a single Satellite rather than the whole network. Example reasons could be that the Storage Node doesn't have as much bandwidth available or pricing **(TODO find better word)** for Satellites has changed.

The process must also consider that the Storage Nodes may have limited bandwidth available or patience. This means the process should prioritize pieces with smaller durability.

### Non-Goals

There are scenarios that this design document does not handle:

- Storage Node runs out of bandwidth during exiting.
- Storage Node may also want to transfer pieces partially due to decreased available capacity.
- Storage Node wants to rejoin a Satellite it previously exited.

## Design

The design is divided into four parts:

- [Process for gathering pieces that need to be transferred.](storagenode-graceful-exit-pieces.md)
- [Protocol for transferring pieces from one Storage Node to another.](storagenode-graceful-exit-protocol.md)
- [Reporting for graceful exit process.](storagenode-graceful-exit-reports.md)
- [User Interface for interacting with graceful exit.](storagenode-graceful-exit-ui.md)

Overall a good graceful exit process looks like:

1. Storage Node Operator initiaties the graceful exit process, which:
    - notifies the satellite about the graceful exit, and
    - adds entry about exiting to storagenode database.
2. Satellite receives graceful exit request, which:
    - adds entry about exiting to satellite database, and
    - starts gathering of pieces that need to be transferred.
3. Satellite finishes gathering pieces that need to be transferred.
4. Storage Node keeps polling Satellite for pieces to transfer.
5. When the Satellite doesn't have any more pieces to transfer, it will respond with a completion receipt.
6. Storage Node stores completion information in the database.
7. Satellite Operator creates a reports about exited storage nodes to release escrow.

For all of these steps we need to ensure that we have sufficient monitoring.

When a Graceful Exit has been started, it must either succeed or fail. The escrow, held from storage node, will be released only on success. We will call the failure scenario an ungraceful exit.

Ungraceful exit can happen when:

- Storage Node doesn't transfer pieces,
- Storage Node is too slow to transfer pieces,
- Storage Node decided to terminate the process.

## Open issues (if applicable)

- Should we have an `satellites` table in storage node to store trusted satellites and graceful exit information?