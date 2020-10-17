import { Heap } from "./heap.js";
import { Index } from "./index.js";
const numbers = [4, 3, 5, 9, 2, 7, 1, 10, 8, 6];

const heap = new Heap(numbers.map((num) => ({ num })));
const index = new Index({ heap: heap, expr: (row) => row.num });

test("lookup one row", () => {
  const offset = index.offset((n) => 5 - n);
  expect(offset).toEqual(4);
  const entry = index.get(offset);
  const row = heap.get(entry.heapOffset);
  expect(row.num).toEqual(5);
});

test("iterate all rows", () => {
  let offset = 0;
  const nums = [];
  while (true) {
    const entry = index.get(offset++);
    if (!entry) {
      break;
    }
    const row = heap.get(entry.heapOffset);
    nums.push(row.num);
  }
  expect(nums).toEqual([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);
});

test("find lowest entry", () => {
  const heap = new Heap(
    [1, 1, 1, 1, 2, 2, 3, 3, 3, 3, 3, 3, 3, 4].map((num) => ({ num }))
  );
  const index = new Index({ heap: heap, expr: (row) => row.num });
  expect(index.offset((n) => 1 - n)).toEqual(0);
  expect(index.offset((n) => 2 - n)).toEqual(4);
  expect(index.offset((n) => 3 - n)).toEqual(6);
  expect(index.offset((n) => 4 - n)).toEqual(13);
  expect(index.offset((n) => 5 - n)).toEqual(-1);
});
