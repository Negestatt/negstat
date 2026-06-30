---
status: accepted (deliberate v1 scope cut, revisit later)
---

# No Order cancellation or refunds, for now

Customers cannot cancel an Order or receive a refund once charged at placement — there is no cancellation/refund flow at all in v1. This is a deliberate scope decision rather than a permanent architectural stance: building refunds well requires deciding how they interact with payment gateway reversals, Print Batch production state, and designer Balance/Payout, and none of that is worth designing until there's a concrete need. Future readers should not assume the absence of refunds is an oversight.
