package database

const(
  insertInto = `INSERT INTO items (
    id,
    name,
    img_url,
    stock,
    unit,
    min_stock
  ) VALUES (?, ?, ?, ?, ?, ?)`
  selectByID = `SELECT
    id,
    name,
    img_url,
    stock,
    unit,
    min_stock
    FROM items
    WHERE id = ?`
  updateRecord = `UPDATE items
    Set name = ?,
    img_url = ?,
    stock = ?,
    unit = ?,
    min_stock = ?
    WHERE id = ?`
  selectAllRecords = "SELECT * FROM items"
)

