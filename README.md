# 1-by-n_communication-data-distribution-by-channel

This lab is prepared to test the 1-to-n data distribution using the golang channel. 
The main thread creates a number of worker threads giving the common boolean channel and enqueue the 'true' to the channel.

Then the worker thread that firstly fetches the boolean value prints out its own identifier

