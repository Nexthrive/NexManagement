DO $$ BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_enum') THEN
    CREATE TYPE role_enum AS ENUM ('user', 'admin', 'superadmin');
  END IF;
END $$;

CREATE TABLE "user" (
  id VARCHAR PRIMARY KEY,
  Username VARCHAR(255) NOT NULL,
  Passphrase VARCHAR(255) NOT NULL,
  Email VARCHAR(255) NOT NULL,
  No_telp VARCHAR(255),
  Role role_enum DEFAULT 'user'
);

