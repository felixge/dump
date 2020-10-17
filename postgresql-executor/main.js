import { Heap } from "./heap.js";
import { Index } from "./index.js";
import { users, orders, series } from "./data.js";
import { SeqScan } from "./seq_scan.js";
import { NestedLoop } from "./nested_loop.js";
//import { Sort } from "./sort.js";
//import { Limit } from "./limit.js";

var seriesHeap = new Heap(series);
var seriesIndex = new Index({
  heap: seriesHeap,
  expr: (series) => series.num,
});
let offset = seriesIndex.offset(3);
while (true) {
  const entry = seriesIndex.get(offset++);
  if (!entry || entry.key != 3) {
    break;
  }
  console.log(entry);
}

//const usersScan = new SeqScan({ heap: new Heap(users) });
//const ordersScan = new SeqScan({ heap: new Heap(orders) });
//const nestedLoop = new NestedLoop({
//outer: usersScan,
//inner: ordersScan,
//where: (users, orders) => (users.user_id == orders.user_id),
//});

//console.log(execute(nestedLoop));

function execute(plan) {
  var rows = [];
  while (true) {
    let row = plan.next();
    if (!row) {
      break;
    }
    rows.push(row);
  }
  return rows;
}

