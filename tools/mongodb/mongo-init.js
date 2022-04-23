db = db.getSiblingDB("service");

db.createCollection("stuff");

db.credential.createIndex(
  {
    email: 1,
  },
  {
    unique: true,
  }
);
