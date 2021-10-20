const {Client} = require('pg');

const pgclient = new Client({
    host: process.env.POSTGRES_HOST,
    port: process.env.POSTGRES_PORT,
    user: 'postgres',
    password: 'postgres',
    database: 'test'
});

pgclient.connect();

const table = `
    CREATE TABLE users
    (
        id    SERIAL,
        name  TEXT,
        email VARCHAR UNIQUE
    )
`;

pgclient.query(table, (err, res) => {
    if (err) throw err
    pgclient.end()
});
