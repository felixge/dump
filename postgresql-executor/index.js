// Index implements a sorted list index. This is similar to a one-level B-Tree
// index and has the same access complexities.
export class Index {
  constructor(params) {
    this.index = [];

    // build index
    let i = 0;
    while (true) {
      const row = params.heap.get(i);
      if (!row) {
        break;
      }
      this.index.push({
        key: params.expr(row),
        heapOffset: i,
      });
      i++;
    }
    this.index.sort((a, b) => {
      if (a.key < b.key) {
        return -1;
      } else if (a.key > b.key) {
        return 1;
      }
      return 0;
    });
  }

  // offset evaluates the given guide callback against keys in the index using
  // binary search and returns the lowest index offset that evaluates to true.
  offset(guide) {
    let low = 0;
    let high = this.index.length;
    while (low < high) {
      const mid = Math.floor((low + high) / 2);
      const midKey = this.index[mid].key;
      const hint = guide(midKey);
      if (hint > 0) {
        low = mid + 1;
      } else if (hint < 0) {
        high = mid - 1;
      } else {
        high = mid;
      }
    }
    return low === this.index.length ? -1 : low;
  }

  // get returns the index entry at the given offset.
  get(offset) {
    return this.index[offset];
  }
}
