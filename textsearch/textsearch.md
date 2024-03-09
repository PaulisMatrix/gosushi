# A simple text search in golang

1.  A simple doc search in go
2.  Most of the code taken from this excellent blog: https://artem.krylysov.com/blog/2020/07/28/lets-build-a-full-text-search-engine/
3.  Only change being using bitmaps to find the intersection between the sets

# Steps:

```
Example text: Anarchism is a political philosophy and movement that is skeptical of all justifications for authority and seeks to abolish the institutions it claims maintain unnecessary coercion and hierarchy, typically including nation-states, and capitalism. Anarchism advocates for the replacement of the state with stateless societies and voluntary free associations.
```

1.  Tokenizing: Split the above text into tokens
2.  After splitting, clean the list of tokens:
    * Cleaning includes:

        a.  Normalise each word. Make all small caps<br>
        b.  removing stop words<br>
        c.  stemming: combine similar words together: for ex working, worked, work -> work<br>
3. Build the inverted index:
    `word: <list of docs ids which contains that word>`
4.  Search the given query in the inverted index.
5.  

