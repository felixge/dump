export class Limit {
  constructor(params) {
    if (!params.node) throw new Error("no node");

    this.node = params.node;
    this.limit = params.limit;
    this.index = 0;
  }
  next() {
    if (this.index >= this.limit) {
      return;
    }
    this.index++;
    return this.node.next();
  }
  rewind() {
    this.node.rewind();
    this.index = 0;
  }
}
