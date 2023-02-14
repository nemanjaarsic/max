CREATE TABLE IF NOT EXISTS "Users"(
    "Id" uuid PRIMARY KEY NOT NULL,
    "Username" text UNIQUE NOT NULL,
    "Password" text NOT NULL,
    "Name" text NOT NULL
);

CREATE TABLE IF NOT EXISTS "Financies"(
    "Id" uuid PRIMARY KEY NOT NULL,
    "UserId" uuid NOT NULL,
    "Balance" double precision NOT NULL,
    CONSTRAINT "UserId" FOREIGN KEY ("UserId")
        REFERENCES "Users" ("Id") MATCH SIMPLE
        ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "Transactions"(
    "Id" uuid PRIMARY KEY NOT NULL,
    "UserId" uuid NOT NULL,
    "OperationType" int NOT NULL,
    "Amount" double precision NOT NULL,
    "Timestamp" timestamp NOT NULL DEFAULT NOW(),
    CONSTRAINT "UserId" FOREIGN KEY ("UserId")
        REFERENCES "Users" ("Id") MATCH SIMPLE
);

CREATE TABLE IF NOT EXISTS "Auth"(
    "Id" uuid PRIMARY KEY NOT NULL,
    "UserId" uuid NOT NULL,
    "Token" text UNIQUE NOT NULL,
    "Active" bool NOT NULL,
    "Created_at" timestamp NOT NULL DEFAULT NOW()
);

INSERT INTO "Users" ("Id", "Username", "Password", "Name")
VALUES 
    ('04267f13-0eef-4594-82a6-4a80ebc22bd7', 'username1', '$2a$10$WvYJAJxw43ZRZfzad0RENumS8a4BWNZ/xgo57BRR/gI8PzYE6YlRO', 'User1'),
    ('a0c26e6e-cde9-4e5a-a1a9-d58f89ffc8bd', 'username2', '$2a$10$vOwIbLLQ5D1HNFs52iJGLOQRMqTXZ6jaIdqMWjsrfVUIlbHE5z2PO', 'User2'),
    ('2ad3d75f-031e-4147-b509-41832b5422b9', 'username3', '$2a$10$OhtDOaI4oaGcQgKSADz51uKRd851OFlVQfJojI.Ybxzy0hErmtaZS', 'User3'),
    ('f05636e6-d552-43ef-addf-d655d5f800ed', 'username4', '$2a$10$Gf9XI3yhAI/.9OP4jrrimeYLgV.8ynNbQHFGrBUkwWUCN4PX.2Equ', 'User4'),
    ('087c568b-b2c5-49ef-873f-e8c187fe2ded', 'username5', '$2a$10$dF6xnBcnsnOByof811KEROvQidHt57IXkoKsUIL5/l2D9JmUGKjRG', 'User5');