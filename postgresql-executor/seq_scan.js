export class SeqScan {
  constructor(params) {
    if (!params.heap) throw new Error("no heap");

    this.heap = params.heap;
    this.where = params.where;
    this.select = params.select;
    this.offset = 0;
  }
  next() {
    while (true) {
      const row = this.heap.get(this.offset++);
      if (!row) {
        return;
      } else if (this.where && !this.where(row)) {
        continue;
      } else if (this.select) {
        return this.select(row);
      } else {
        return row;
      }
    }
  }
  rewind() {
    this.offset = 0;
  }
}
