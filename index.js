console.println("Start");

setTimeoutCustom(() => {
  console.println("timeout 1");
}, 4000);

setTimeoutCustom(() => {
  console.println("timeout 2");
}, 100);

console.println("End");
