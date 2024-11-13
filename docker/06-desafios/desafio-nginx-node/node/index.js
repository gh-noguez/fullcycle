const express = require('express')
const port = 3000
const app = express()
const config = {
    host: 'db',
    user: 'root',
    password: 'root',
    database: 'nodedb'
}
const mysql = require('mysql')
const connection = mysql.createConnection(config)

const createTable = `create table people(id int not null auto_increment, name varchar(255), primary key(id));`
connection.query(createTable)

const sql = `INSERT INTO people(name) VALUES('Felipe Noguez')`
connection.query(sql)


app.get('/', (req, res) => {
    const select = `SELECT * FROM people`;
    connection.query(select, (err, results, fields) => {
        if (err) {
            console.error(err);
            return res.status(500).send('Erro ao executar a query.');
        }
        console.log(results);
        res.send('<h1>Full Cycle Rocks!</h1><ul>' + results.map(result => `<li>${result.name}</li>`).join('') + '</ul>');
    });
});

app.listen(port, () => {
    console.log('Rodando aplicação na porta', port)
})