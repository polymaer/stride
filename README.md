# Stride
Bringing liquid staking to Cosmos

## Making changes to this repository
The easiest way to develop cosmos-sdk applications is by using the ignite cli to scaffold code. Ignite (developed by the core cosmos team at Tendermint) makes it possible to scaffold new chains, run relayers, build cosmos related proto files, add messages/queries, add new data structures and more. The drawback of creating thousands of lines of code using ignite is that it can be difficult to discern which changes were made by the ignite cli and which changes were made by developers. To make it easier to review code written using ignite and to make it easier to retrace our steps if something goes wrong later, add a commit for each ignite command run, directly after running the command.

For example, adding a new message type and updating the logic of that message would be two commits.
```
// add the new data type
>>> ignite scaffold list loan amount fee collateral deadline state borrower lender
>>> git add . && git commit -m 'ignite scaffold list loan amount fee collateral deadline state borrower lender'
// make some updates to the keeper method in your code editor
>>> git add . && git commit -m 'update loan list keeper method'
```

An example of this can be found in [this PR](https://github.com/Stride-Labs/stride/pull/1). Notice, you can easily differentiate changes made by ignite and those made manually by going through commits. For example, in commit fd3e254bc0, it's easy to review the [files changed](https://github.com/Stride-Labs/stride/pull/1/commits/fd3e254bc0844fe65f5e98f12b366feef2a285f9) even though nearly ~300k LOC were scaffolded.

## Code review format
Opening a PR will automatically create a Summary and Test plan field in the description. In the summary, add a high-level summary of what the change entails and the ignite commands run.
**Summary**
Updating some code.
**Test plan**
yolo, we're in testnet


