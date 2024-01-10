CREATE TABLE IF NOT EXISTS public.authors (
                                              author_id SERIAL PRIMARY KEY,
                                              name VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS public.books (
                                            id SERIAL PRIMARY KEY,
                                            title VARCHAR NOT NULL,
                                            author_id INT NOT NULL,
                                            quantity INT NOT NULL,
                                            status BOOL NOT NULL,
                                            CONSTRAINT books_fk FOREIGN KEY (author_id) REFERENCES authors(author_id)
    );