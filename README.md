# pegnetwinners
shows the winners at a certain factom block height and any other accounts asked for

runs against a local factomd

runs on port 8899

3 available calls

http://localhost:8899/Winners?height=207797&filter=miner1

this pulls all the OPR records from factom block height 207797
filters on winners (shown yellow)
Filters on filter and shows then as light blue
if the filtered miner name is also a winner, it is dark blue.

http://localhost:8899/BlockWinners?height=207797
shows the winner values for the previous block
the values are a substring of the entry hashes
of the OPR records


http://localhost:8899/BlockStats?height=207797
All OPR records from the height block

