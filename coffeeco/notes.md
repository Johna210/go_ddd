### Internal
The internal folder is a special folder in golang. Anything  within the internal directory 
cannot be imported by any other package outside the parent directory. This is a great fit for DDD as we do not want our domain code to be part of our public API. To automatically enforce this, we will be putting all our domain code inside this internal folder.

##### If you can answer yes to all three questions it should be a value object not an entity / upgrade it later
- Is it possible for me to treat this object as immutable?
- Does it measure, quantify, or describe a domain concept?
- Can it be compared to other objects of the same type just by its values?

