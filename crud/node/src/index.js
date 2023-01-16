const express = require('express')
const { MongoClient, ObjectId } = require('mongodb')
const app = express()
const port = 3000

let client;
let collection;

async function connect() {

  const uri = "mongodb://127.0.0.1:27017"

  client = await MongoClient.connect(uri)

  collection = client.db("node").collection("products")

  console.log("db connected");
}

connect()

app.use(express.json())

app.get('/api/products', async (req, res) => {

  const products = await collection.find({}).toArray()

  res.status(200).json(products)
})

app.get('/api/products/:id', async (req, res) => {

  const id = req.params.id

  const objectId = new ObjectId(id)

  const product = await collection.findOne({_id: objectId})

  res.status(200).json(product)
})

app.post('/api/products', async (req, res) => {

  const product = req.body;

  const insertResult = await collection.insertOne(product)

  res.status(201).json({id: insertResult.insertedId})
})

app.listen(port, () => {
  console.log(`App listening on port ${port}`)
})