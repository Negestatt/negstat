# Rush Orders bypass the Print Batch entirely

Customers can pay an extra fee to skip the Print Batch threshold (10 Orders or 14-day time-box) and trigger an individual print job immediately. We chose to bypass batching entirely, rather than only expediting shipping after a batch closes, because "fast delivery" customers care about production lead time, not just transit time — only bypassing batching actually shortens the up-to-14-day wait established in ADR 0001.

This also changes who absorbs under-fill cost (ADR 0002) for these orders specifically: a Rush Order is the most extreme possible under-fill (a batch of one), but here the customer's surcharge — not the platform — is expected to cover that cost. ADR 0002 still governs ordinary batched Orders.

Rush Orders are fully isolated from a design's regular Print Batch progress: they do not count toward the 10-Order threshold. A Rush Order's production already happened individually, so counting it toward a batch meant to be produced *together* would misrepresent what actually happened, and could let a batch reach "10/10" without any of those units ever being produced together.
