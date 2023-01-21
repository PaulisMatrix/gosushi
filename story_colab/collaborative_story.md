## Collaborative Story:

- A story is made up of a title and paragraphs.
- Paragraphs are made of sentences.
- A single request can add exactly 1 word to the story.
- The first 2 words added to a story make the title.
- From the 3rd word, the first sentence of the first paragraph begins.
- As people add words, when there are 15 words, a new sentence starts.
- When there are 10 sentences, a new paragraph starts.
- When there are exactly 7 paragraphs, the story ends and a new one is created.
- There is no concept of users in the system. Anyone with the endpoint URL, should be able to add a word to the story.
- The words in /add requests, do not have any order associated with them. For e.g. if hundreds of POST requests come in concurrently then application may process all of those words in any arbitrary order.

## Endpoints

1. POST `/add` to add a new word.
2. GET `/stories` returns list of stories.
3. GET `/stories/:id` will return details of the story.

## Database

Selecting mongodb with each document as a story and sub document(7 documents/story) as a paragraph with paragraph having a sub document(10 documents/paragraph) called sentence.
