export const users = [
  { user_id: 1, email: "bob@example.org" },
  { user_id: 2, email: "alice@example.org" },
  { user_id: 3, email: "john@example.org" },
];

export const orders = [
  { order_id: 1, user_id: 1, amount: 5 },
  { order_id: 2, user_id: 1, amount: 13 },
  { order_id: 3, user_id: 3, amount: 7 },
  { order_id: 4, user_id: 3, amount: 9 },
];

export const series = [5, 2, 3, 4, 1, 6, 7, 9, 3, 8, 10, 3, 3, 3, 3].map((num) => ({
  num: num,
}));
