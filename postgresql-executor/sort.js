export class Sort {
  constructor(params) {
    if (!params.order) throw new Error("no order");
    if (!params.node) throw new Error("no node");

    this.node = params.node;
    this.order = params.order;
    this.sorted = null;
    this.index = 0;
  }
  next() {
    if (!this.sorted) {
      this.sorted = [];
      while (true) {
        const row = this.node.next();
        if (!row) {
          break;
        }
        this.sorted.push(row);
      }
      this.sorted.sort(this.order);
    }
    return this.sorted[this.index++];
  }
  rewind() {
    this.index = 0;
  }
}
