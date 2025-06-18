CREATE TABLE cats (
                      id UUID PRIMARY KEY,
                      name TEXT NOT NULL,
                      age INT NOT NULL,
                      color TEXT NOT NULL
);

CREATE TABLE actions (
                         id UUID PRIMARY KEY,
                         cat_id UUID REFERENCES cats(id),
                         payload JSONB NOT NULL,
                         action_type TEXT NOT NULL,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);