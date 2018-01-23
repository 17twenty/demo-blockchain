# Demo Blockchain

Simple demo on creating a genesis block and signing transactions for an
immutable ledger.

## Getting Started

Simplest way to do it is to just run the test and then take a look at the code:

```bash
$ go test -v
=== RUN   TestBlockchain
Previous hash:
Transaction: { from: 00000000-0000-0000-0000-000000000000, to: 00000000-0000-0000-0000-000000000000, description: genesis, amount: 0.0000
Hash: d862651e1cfe0e55e62e34bfa1422eb49da5865b8364e082a5fa9e8c92c622b4

Previous hash: d862651e1cfe0e55e62e34bfa1422eb49da5865b8364e082a5fa9e8c92c622b4
Transaction: { from: 7838a0d7-f1b7-4cd4-9f32-8aa75ed96ef4, to: de22b22f-91d2-4429-9994-d305d2254ed6, description: , amount: 1000.0000
Hash: 53439a31b1d09ebb6a7d086599caac0fd468db9478069b1929a53157dc32e10f

Previous hash: 53439a31b1d09ebb6a7d086599caac0fd468db9478069b1929a53157dc32e10f
Transaction: { from: de22b22f-91d2-4429-9994-d305d2254ed6, to: c55d4393-737a-4367-816f-66ebdb09ada1, description: , amount: 200.0000
Hash: cb72ca00cf4ddfbd9d7709496de4569332e8cbde13c97be54c77fa81f6b43197

--- PASS: TestBlockchain (0.00s)
PASS
ok  _/tmp/demo-blockchain  0.009s
```
