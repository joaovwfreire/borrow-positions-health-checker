# Overview
This script generates a list of borrowers from the protocol queried. It's main purpose is to enable the execution of other code at this repo, such as the health factor script.
In order to make it quick, I've opted to run it in parallel through wait groups. It's important to note that the script will generate a file called `borrowers.csv` in the root of the repo. This file will be used by the other scripts. Bear in mind this will generate thousands of address entries, so make sure to use specific block number filters in order to make the most out of this tool. This tool does not take into account duplicated addresses, nor does it take into account the amount of debt the address has. It's just a list of addresses that have borrowed from the protocol. One should be able to manipulate this data in order to get accurate results.

# Usage
Make sure to place the same .env file as the one used in the protocol repo. Then run the following command:

```bash
cd borrowersGenerator
go run .
```