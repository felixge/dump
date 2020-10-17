export class Heap {
  constructor(rows) {
    this.rows = rows;
  }
  get(index) {
    return this.rows[index];
  }
  length() {
    return this.rows.length;
  }
}
