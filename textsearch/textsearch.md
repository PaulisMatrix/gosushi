# A simple text search in golang

1.  A simple doc search in go.
2.  Most of the code referenced from this excellent blog: https://artem.krylysov.com/blog/2020/07/28/lets-build-a-full-text-search-engine/
3.  Only change being using a more efficient method, bitmaps to find the intersection between the sets.

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

# TODO:

1.  Use roaring bitmaps to save space cause there's a lot of 0s which dont get set at all in the bitmap.

    Roaring bitmaps are just space efficient bitmaps:<br>
        https://www.vikramoberoi.com/a-primer-on-roaring-bitmaps-what-they-are-and-how-they-work/ <br>
        https://roaringbitmap.org/about/ <br>
        [roaring bitmaps in golang](https://github.com/RoaringBitmap/roaring)<br>


2.  Store the index on disk. Loading and Rebuilding the index everytime on startup takes significant time. 

    Reading the docs: total time in reading the docs: `19.411223 secs` <br>
    Building the index: total time in indexing 676792 docs: `14.139312 secs`

3.  Add more boolean queries like OR, NOT, etc.


# Further Readings: 

1.  Index Data Structures in ES: 

    [ElasticSearch-BottomUp](https://www.elastic.co/blog/found-elasticsearch-from-the-bottom-up) <br>
    Talk on the same : https://youtu.be/PpX7J-G2PEo

2.  Distributed nature of ES: [ElasticSearch-TopDown](https://www.elastic.co/blog/found-elasticsearch-top-down)
3.  [Dissecting Lucene](https://kandepet.com/dissecting-lucene-the-index-format/)
4.  [Lucene-The good parts](https://www.parse.ly/lucene/)
5.  [Blog posts](https://blog.mikemccandless.com/2011/02/visualizing-lucenes-segment-merges.html) of lucene main committer.
6.  Talks on lucene:

    [Algorithms & data-structures that power Lucene & ElasticSearch](https://youtu.be/eQ-rXP-D80U?si=YW-geBKcNfLaVjzd) <br>
    [Adrien Grand, Software Engineer, Elasticsearch](https://youtu.be/T5RmMNDR5XI?si=LWzBt8Mq2z2bVIGl)
7.  
