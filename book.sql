CREATE TABLE book
(
    id serial NOT NULL,
    title character varying NOT NULL,
    author character varying,
    published_at date,
    CONSTRAINT pk_book PRIMARY KEY (id )
)
    WITH (
        OIDS=FALSE
        );
ALTER TABLE book
    OWNER TO postgres;