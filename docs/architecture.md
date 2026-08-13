# Architecture

The sample service keeps its state in memory so tests remain deterministic.

1. The HTTP handler validates JSON input.
2. The task store validates titles and assigns IDs.
3. Responses are encoded as JSON.

This document was merged through a pull request to preserve a realistic timeline event.
