// Create DB and User with roles on that DB.
db.createUser(
    {
        user: "dev",
        pwd: "password123",
        roles: [
            {
                role: "readWrite",
                db: "dev"
            }
        ]
    }
);
