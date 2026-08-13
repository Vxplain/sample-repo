# System design

The sample service keeps its state in memory so tests remain deterministic.

1. The HTTP handler validates JSON input.
2. The task store validates titles and assigns IDs.
3. The completion flow updates a task under the store lock.
4. Responses are encoded as JSON.

This document was renamed and expanded in a pull request to exercise rename metadata.
