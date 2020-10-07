can you write a program that goes through each line, and determines how many "new" vs "duplicate" coords it sees?

6:48
and make the program accept the amount of decimals we want to consider
6:48
so we can compare the result for 2 or 3 decimals

we could get a bit fancier with our program as well, if we take cache ttl into consideration.
For example, we can assume a cache TTL of 3 minutes. A request that comes in after 3 minutes of the original request that was cached, would be a cache miss (and would go into MPX and then get cached again). (edited) 
6:54
Because we have the request time on the Splunk logs

1. get time object (date and time) and latitude/long object
 - add to struct