# chainhub-backend

## Weekly pending fees

Read the current Curve fee week on Ethereum, Arbitrum and Base. The command checks each FeeCollector's installed burner at a pinned block. It runs the existing estimator for the known legacy burner and the Dutch auction estimator as soon as the collector installs the proposal burner. It writes one JSON report to stdout. An unrecognized burner stops the calculation with its address.

```sh
go run . weeklyPendingFees
go run . weeklyPendingFees -chain ethereum -flows
go run . weeklyPendingFees -chain ethereum -pools all -pool-limit 50
```

**During the migration.** Ethereum and Arbitrum retain the same live calculation as the scheduled `pendingPoolFees` job. Ethereum returns the legacy USD pool component and collected crvUSD separately. Arbitrum returns its original MIM amount, then converts it at the current MIM and crvUSD prices for the cross-chain total. The raw values are in `legacyComponents`. Arbitrum's legacy Tenderly bundle needs the existing `TENDERLY_NODE` configuration. The reporting command reads fee metadata without publishing or changing the job's caches.

Base enters the calculation when it activates the Dutch burner. While its legacy burner is installed, the row reports `outside_legacy_coverage` and is excluded from the total. A Base-only query in that state omits the estimate. After migration the same query includes Base's auction inventory.

The scheduled `pendingPoolFees` job uses the same on-chain switch. It keeps the existing distribution file for the dashboard and the legacy calculators on the other networks. For migrated networks that file represents the share attributable to veCRV holders. Ethereum uses the allocator share valued in USD and treasury-only L2 routes contribute zero. The detailed CLI report also shows protocol-wide and treasury amounts.

**Dutch auction estimates.** `estimatedPendingCrvUSD` is the value still in the pipeline at the observed blocks. It includes earlier inventory that can enter this week's distribution. The reporting week starts Thursday at 00:00 UTC. The estimate is a snapshot and does not forecast new fees from future swaps or loans.

| Phase | What contributes to the estimate |
| --- | --- |
| SLEEP and COLLECT | Observed crvUSD, eligible collector and burner inventory, and simulated receipts from listed pools. |
| EXCHANGE | Observed crvUSD and the remaining balance of lots staged for this week. Each lot is capped at its staged amount. The full-lot crvUSD quote is reported separately because execution prices decay and lots may remain unsold. |
| FORWARD | Observed crvUSD adjusted for the collector's buffer and the current forward keeper rate. Unsold tokens can be staged in a later week. |

The Ethereum estimate includes the hooker and allocator balances. The veCRV estimate uses the allocator's on-chain distributor weight. Arbitrum and Base route funds to the Ethereum DAO treasury. Each L2 reports the configured bridge minimum and whether the estimated balance meets it. The cross-chain total combines the pending amounts on the selected networks. It includes the treasury share.

The Dutch estimate is unavailable if the distribution route cannot be verified, forwarding is killed, or a required crvUSD conversion price is missing. Observed balances remain available as exact decimal strings. Market estimates use the current Curve prices for both the tokens and crvUSD. Before forwarding, actual sales, keeper incentives and execution time can change the result.

**Cost and coverage**

- `-pools auto` simulates listed pools before collection closes once the new burner is installed. `-pools off` reads existing inventory. `-pools all` scans in any phase and adds directly claimable crvUSD to the active estimate even after collection closes. `-pool-limit N` restricts simulations to the highest-TVL pools per chain. Zero scans all listed pools when simulations are enabled.
- Pool simulations measure token or LP balance increases at the collector. The pool section counts omitted pools, unsupported claims, unverified receipts and missing prices. Assets retained at legacy proxies, unwrapped native currency and assets outside the catalog or discovered lots are outside the priced estimate. `partial` flags known gaps within the selected scope.
- `-flows` additionally reads the week's executed native and CoW sales, distributions at the Ethereum allocator, and bridges initiated by the current L2 hooker. These fields describe successive stages and are reported separately from pending balances. The allocator can distribute funds from other sources, including during the migration.
- The default mode reads staging logs only during COLLECT and EXCHANGE. If a gateway cannot serve that history, the command reads current lot records for Curve-listed tokens and marks the estimate partial. The optional historical ledger needs a log-capable RPC. The report includes the block hash and the HTTP request count on its RPC client. Legacy simulations can also call external services.
- Override the CRVHub endpoints with `ETHEREUM_RPC_URL`, `ARBITRUM_RPC_URL` or `BASE_RPC_URL`. The Curve pool catalog and prices are fetched once per run when needed. `CURVE_POOLS_URL` accepts a self-hosted endpoint with the same response as Curve's `getPools/all`. Use `-timeout` to adjust the overall deadline.

The deployed addresses and calculation follow [Curve ownership proposal 1524](https://www.curve.finance/dao/ethereum/proposals/1524-ownership), the [Dutch burner contracts](https://github.com/curvefi/curve-burners/tree/2154e79e190193f36b48b23a80cbf4646e604d8d) and the [Ethereum fee allocator](https://eth.blockscout.com/address/0x22530D384cd9915e096ead2Db7F82EE81f8eB468?tab=contract).
