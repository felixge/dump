// TODO(felixge) left join
export class NestedLoop {
  constructor(params) {
    this.outer = params.outer;
    this.inner = params.inner;
    this.where = params.where;
    this.select = params.select;
    this.outerRow = null;
  }
  next() {
    while (true) {
      if (!this.outerRow) {
        this.outerRow = this.outer.next();
        if (!this.outerRow) {
          return;
        }
      }

      const innerRow = this.inner.next();
      if (!innerRow) {
        this.inner.rewind();
        this.outerRow = null;
        continue;
      } else if (this.where && !this.where(this.outerRow, innerRow)) {
        continue;
      } else if (this.select) {
        return this.select(this.outerRow, innerRow);
      } else {
        return { ...this.outerRow, ...innerRow };
      }
    }
  }
  rewind() {
    this.index = 0;
  }
}
