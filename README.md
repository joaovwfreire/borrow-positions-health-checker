# Repository with Health Factor monitoring in Go

This repo contains a simple Go application that monitors the health factor of multiple borrowers at a variety of DeFi protocols. The health factor is a metric that measures the health of a borrower's position. It is calculated by dividing the value of the collateral by the value of the debt. The health factor is a ratio, so it is always greater than or equal to 1. If the health factor is less than 1, the borrower is considered to be undercollateralized and is at risk of liquidation.

This repo uses the Multicall contract to make a single call to the protocols' smart contracts to get the health factor of multiple borrowers. This is more efficient than making a separate call to each protocol for each borrower as it enables multiple contract calls in a single view function request. At this repo I've opted to use up to 64 requests at a single multicall.

This repo also contains a list of borrowers generator. It works by concurrently fetching the borrow events and indexing the addresses of borrowers.
 This is useful if you want to monitor the health factor of a large number of borrowers.

## Smart Contracts
Most is interfaces. This repo's scope is to monitor the health factor of borrowers. The smart contracts are just means for such, so I do recommend checking each protocol's documentation according to your own interest.

## Checklist
- [ ] Aave
    - [x] Health factor
    - [x] Borrowers list generator
    - [ ] Cli tool
- [ ] Compound
    - [ ] Health factor
    - [ ] Borrowers list generator
    - [ ] Cli tool
- [ ] MakerDAO
    - [ ] Health factor
    - [ ] Borrowers list generator
    - [ ] Cli tool
- [ ] dYdX
    - [ ] Health factor
    - [ ] Borrowers list generator
    - [ ] Cli tool
