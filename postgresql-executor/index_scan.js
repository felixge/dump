export class IndexScan {
  constructor(params) {
    if (!params.heap) throw new Error("no heap");
    if (!params.index) throw new Error("no index");

    this.heap = params.heap;
    this.index = params.index;
    this.guide = params.guide;
    this.where = params.where;
    this.select = params.select;
    this.offset = -1;
  }
  next() {
    if (this.offset === -1) {
      this.rewind();
    }

    while (true) {
      const entry = this.index.get(this.offset++);
      if (!entry || (this.guide && this.guide(entry.key) !== 0)) {
        return;
      }

      const row = this.heap.get(entry.heapOffset);
      if (this.where && !this.where(row)) {
        continue;
      }

      if (this.select) {
        return this.select(row);
      }
      return row;
    }
  }
  rewind() {
    this.offset = this.guide
      ? (this.offset = this.index.offset(this.guide))
      : 0;
  }
}
