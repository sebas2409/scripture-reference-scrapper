# Scripture Study AI (Scripture Reference)

Scripture Study AI is a full-stack web application that leverages Natural Language Processing (NLP) and vector databases
to discover deep theological connections across canonical texts. Instead of relying on traditional keyword matching,
this tool uses a Machine Learning model to understand the semantic meaning of a highlighted phrase or a whole verse.

The architecture consists of a high-performance Go web scraper to ingest the texts, a Python FastAPI backend connected
to a Qdrant vector database for hybrid search, and a modern React frontend built with TanStack Start for a seamless,
app-like reading experience.

## Instructions for Build and Use

Steps to build and/or run the software:

1. Ensure you have **Docker**, **Bun**, **Python 3.10+**, and **Go** installed on your machine.
2. Spin up the local vector database using Docker Compose by running: `docker-compose up -d`.
3. Start the Python ML backend by navigating to the API directory and running: `uvicorn main:app --reload` (ensure your
   virtual environment is active and `requirements.txt` is installed).
4. Start the frontend web application by navigating to the frontend directory and running: `bun install` followed by
   `bun run dev`.

Instructions for using the software:

1. Open your web browser and navigate to the local server address provided by Bun (typically `http://localhost:3000`).
2. From the home page library, select a volume and a book, then choose a chapter from the grid to open the reading
   interface.
3. **Semantic Search:** Highlight any phrase in the text with your mouse. A dark tooltip will appear; click "Search
   phrase" to query the AI for verses with similar meanings.
4. **Clause Deconstruction:** Hover over any verse and click the link icon (🔗) next to the verse number to break the
   verse down by concept and find highly accurate cross-references.

## Development Environment

To recreate the development environment, you need the following software and/or libraries with the specified versions:

* **Go**: Used exclusively for building the rapid web scraper that generates the initial JSON dataset.
* **Python 3.10+**: Used for the backend API. Key libraries include `FastAPI` (for the server) and
  `sentence-transformers` (for generating the ML vector embeddings).
* **Qdrant**: The vector database used to store and query the embeddings (run locally via Docker).
* **Bun**: The fast JavaScript runtime and package manager used to build the frontend.
* **React 18 & TanStack Start / Router**: Used for building the client-side UI and handling server-side data fetching.
* **Tailwind CSS**: Used for styling the user interface.

## Useful Websites to Learn More

I found these websites useful in developing this software:

* [FastAPI Official Documentation](https://fastapi.tiangolo.com/)
* [Qdrant Vector Database Documentation](https://qdrant.tech/documentation/)
* [TanStack Router & Start Docs](https://tanstack.com/router/latest)
* [Sentence-Transformers (Hugging Face)](https://www.sbert.net/)
* [Bun Runtime Documentation](https://bun.sh/docs)

## Future Work

The following items I plan to fix, improve, and/or add to this project in the future:

* [ ] **User Accounts & Bookmarks:** Implement authentication so users can save their favorite AI-generated
  cross-references, create personal study lists, and highlight verses.
* [ ] **Model Fine-Tuning:** Fine-tune the sentence-transformer model specifically on ancient theological texts and
  archaic language (like King James English) to improve the accuracy of the vector embeddings.
* [ ] **Multilingual Support:** Upgrade the ML pipeline to support multilingual vector embeddings, allowing a user to
  search in Spanish and find exact semantic parallels in English or other languages.