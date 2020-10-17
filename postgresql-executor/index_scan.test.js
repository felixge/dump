import { Heap } from "./heap.js";
import { Index } from "./index.js";
import { IndexScan } from "./index_scan.js";
const numbers = [4, 3, 5, 9, 2, 7, 1, 10, 8, 6];

const heap = new Heap(numbers.map((num) => ({ num })));
const index = new Index({ heap: heap, expr: (row) => row.num });

test("index range scan", () => {
  const guide = (n) => {
    if (n < 3) {
      return 1;
    } else if (n > 5) {
      return -1;
    }
    return 0;
  };
  const nums = execute(guide);
  expect(nums).toEqual([3, 4, 5]);
});

test("single row lookup", () => {
  const guide = (n) => 5 - n;
  const nums = execute(guide);
  expect(nums).toEqual([5]);
});

test("lookup failure", () => {
  const guide = (n) => 20 - n;
  const nums = execute(guide);
  expect(nums).toEqual([]);
});

test("full index scan", () => {
  const nums = execute();
  expect(nums).toEqual([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);
});

function execute(guide) {
  const scan = new IndexScan({
    index: index,
    heap: heap,
    guide: guide,
  });

  const nums = [];
  while (true) {
    const row = scan.next();
    if (!row) {
      break;
    }
    nums.push(row.num);
  }
  return nums;
}
